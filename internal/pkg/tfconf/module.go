package tfconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

// Module represents a Terraform mod. It consists of
// - Header    ('header' json key):    Module header found in shape of multi line comments at the beginning of 'main.tf'
// - Inputs    ('inputs' json key):    List of input 'variables' extracted from the Terraform module .tf files
// - Outputs   ('outputs' json key):   List of 'outputs' extracted from Terraform module .tf files
// - Providers ('providers' json key): List of 'providers' extracted from resources used in Terraform module
type Module struct {
	Header         string      `json:"header" yaml:"header"`
	Inputs         []*Input    `json:"inputs" yaml:"inputs"`
	Outputs        []*Output   `json:"outputs" yaml:"outputs"`
	Providers      []*Provider `json:"providers" yaml:"providers"`
	RequiredInputs []*Input    `json:"-" yaml:"-"`
	OptionalInputs []*Input    `json:"-" yaml:"-"`
	Options        *Options    `json:"-" yaml:"-"`
}

// HasInputs indicates if the document has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Inputs) > 0
}

// HasOutputs indicates if the document has outputs.
func (m *Module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

// Sort sorts list of inputs and outputs based on provided flags (name, required, etc)
func (m *Module) Sort(settings *print.Settings) {
	if settings.SortByName {
		sort.Sort(providersSortedByName(m.Providers))
	} else {
		sort.Sort(providersSortedByPosition(m.Providers))
	}

	if settings.SortByName {
		if settings.SortByRequired {
			sort.Sort(inputsSortedByRequired(m.Inputs))
			sort.Sort(inputsSortedByRequired(m.RequiredInputs))
			sort.Sort(inputsSortedByRequired(m.OptionalInputs))
		} else {
			sort.Sort(inputsSortedByName(m.Inputs))
			sort.Sort(inputsSortedByName(m.RequiredInputs))
			sort.Sort(inputsSortedByName(m.OptionalInputs))
		}
	} else {
		sort.Sort(inputsSortedByPosition(m.Inputs))
		sort.Sort(inputsSortedByPosition(m.RequiredInputs))
		sort.Sort(inputsSortedByPosition(m.OptionalInputs))
	}

	if settings.SortByName {
		sort.Sort(outputsSortedByName(m.Outputs))
	} else {
		sort.Sort(outputsSortedByPosition(m.Outputs))
	}
}

// CreateModule returns new instance of Module with all the inputs and
// outputs discovered from provided 'path' containing Terraform config
func CreateModule(options *Options) (*Module, error) {
	mod := loadModule(options.Path)

	header := readHeader(options.Path)

	var inputs = make([]*Input, 0, len(mod.Variables))
	var requiredInputs = make([]*Input, 0, len(mod.Variables))
	var optionalInputs = make([]*Input, 0, len(mod.Variables))

	for _, input := range mod.Variables {
		inputType := input.Type
		if input.Type == "" {
			inputType = "any"
			if input.Default != nil {
				switch xType := fmt.Sprintf("%T", input.Default); xType {
				case "string":
					inputType = "string"
				case "int", "int8", "int16", "int32", "int64", "float32", "float64":
					inputType = "number"
				case "bool":
					inputType = "bool"
				case "[]interface {}":
					inputType = "list"
				case "map[string]interface {}":
					inputType = "map"
				}
			}
		}

		inputDescription := input.Description
		if inputDescription == "" {
			inputDescription = readComment(input.Pos.Filename, input.Pos.Line-1)
		}

		i := &Input{
			Name:        input.Name,
			Type:        String(inputType),
			Description: String(inputDescription),
			Default:     input.Default,
			Position: Position{
				Filename: input.Pos.Filename,
				Line:     input.Pos.Line,
			},
		}

		inputs = append(inputs, i)
		if i.HasDefault() {
			optionalInputs = append(optionalInputs, i)
		} else {
			requiredInputs = append(requiredInputs, i)
		}
	}

	// output module
	var outputs = make([]*Output, 0, len(mod.Outputs))
	for _, o := range mod.Outputs {
		outputDescription := o.Description
		if outputDescription == "" {
			outputDescription = readComment(o.Pos.Filename, o.Pos.Line-1)
		}
		output := &Output{
			Name:        o.Name,
			Description: String(outputDescription),
			Position: Position{
				Filename: o.Pos.Filename,
				Line:     o.Pos.Line,
			},
		}
		if options.OutputValues {
			terraformOutputs, err := loadOutputValues(options)
			if err != nil {
				log.Fatal(err)
			}
			output.Value = terraformOutputs[output.Name].Value
		}
		outputs = append(outputs, output)
	}

	var providerSet = loadProviders(mod.RequiredProviders, mod.ManagedResources, mod.DataResources)
	var providers = make([]*Provider, 0, len(providerSet))
	for _, provider := range providerSet {
		providers = append(providers, provider)
	}

	module := &Module{
		Header:         header,
		Inputs:         inputs,
		Outputs:        outputs,
		Providers:      providers,
		RequiredInputs: requiredInputs,
		OptionalInputs: optionalInputs,

		Options: options,
	}
	return module, nil
}

func loadModule(path string) *tfconfig.Module {
	module, diag := tfconfig.LoadModule(path)
	if diag != nil && diag.HasErrors() {
		log.Fatal(diag)
	}
	return module
}

func loadProviders(requiredProviders map[string]*tfconfig.ProviderRequirement, resources ...map[string]*tfconfig.Resource) map[string]*Provider {
	var providers = make(map[string]*Provider)
	for _, resource := range resources {
		for _, r := range resource {
			var version = ""
			if requiredVersion, ok := requiredProviders[r.Provider.Name]; ok && len(requiredVersion.VersionConstraints) > 0 {
				version = strings.Join(requiredVersion.VersionConstraints, " ")
			}
			key := fmt.Sprintf("%s.%s", r.Provider.Name, r.Provider.Alias)
			providers[key] = &Provider{
				Name:    r.Provider.Name,
				Alias:   String(r.Provider.Alias),
				Version: String(version),
				Position: Position{
					Filename: r.Pos.Filename,
					Line:     r.Pos.Line,
				},
			}
		}
	}
	return providers
}

func loadOutputValues(options *Options) (map[string]*TerraformOutput, error) {
	if options.OutputValuesPath == "terraform-outputs.json" {
		_, err := exec.Command("cd", options.Path, "&&", "terraform", "output", "--json", ">>", options.OutputValuesPath).Output()
		if err != nil {
			log.Fatal(err)
		}
	}

	var terraformOutputs map[string]*TerraformOutput

	byteValue, err := ioutil.ReadFile(options.OutputValuesPath)
	if err != nil {
		return nil, fmt.Errorf("caught error while reading the terraform outputs file at %s: %v", options.OutputValuesPath, err)
	}
	err = json.Unmarshal(byteValue, &terraformOutputs)
	if err != nil {
		return nil, err
	}
	return terraformOutputs, err
}
