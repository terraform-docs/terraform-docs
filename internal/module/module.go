package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/segmentio/terraform-docs/internal/reader"
	"github.com/segmentio/terraform-docs/internal/tfconfig"
	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// LoadWithOptions returns new instance of Module with all the inputs and
// outputs discovered from provided 'path' containing Terraform config
func LoadWithOptions(options *Options) (*tfconf.Module, error) {
	tfmodule, err := loadModule(options.Path)
	if err != nil {
		return nil, err
	}
	module, err := loadModuleItems(tfmodule, options)
	if err != nil {
		return nil, err
	}
	sortItems(module, options.SortBy)
	return module, nil
}

func loadModule(path string) (*tfconfig.Module, error) {
	module, diag := tfconfig.LoadModule(path)
	if diag != nil && diag.HasErrors() {
		return nil, diag
	}
	return module, nil
}

func loadModuleItems(tfmodule *tfconfig.Module, options *Options) (*tfconf.Module, error) {
	header, err := loadHeader(options)
	if err != nil {
		return nil, err
	}

	inputs, required, optional := loadInputs(tfmodule)
	outputs, err := loadOutputs(tfmodule, options)
	if err != nil {
		return nil, err
	}
	providers := loadProviders(tfmodule)
	requirements := loadRequirements(tfmodule)

	return &tfconf.Module{
		Header:       header,
		Inputs:       inputs,
		Outputs:      outputs,
		Providers:    providers,
		Requirements: requirements,

		RequiredInputs: required,
		OptionalInputs: optional,
	}, nil
}

func getFileFormat(filename string) string {
	if filename == "" {
		return ""
	}
	last := strings.LastIndex(filename, ".")
	if last == -1 {
		return ""
	}
	return filename[last:]
}
func isFileFormatSupported(filename string) (bool, error) {
	if filename == "" {
		return false, fmt.Errorf("--header-from value is missing")
	}
	switch getFileFormat(filename) {
	case ".adoc", ".md", ".tf", ".txt":
		return true, nil
	}
	return false, fmt.Errorf("only .adoc, .md, .tf and .txt formats are supported to read header from")
}

func loadHeader(options *Options) (string, error) {
	if !options.ShowHeader {
		return "", nil
	}
	if ok, err := isFileFormatSupported(options.HeaderFromFile); !ok {
		return "", err
	}
	filename := filepath.Join(options.Path, options.HeaderFromFile)
	if info, err := os.Stat(filename); os.IsNotExist(err) || info.IsDir() {
		if options.HeaderFromFile != "main.tf" {
			return "", err // user explicitly asked for a file which doesn't exist
		}
		return "", nil // absorb the error to not break workflow of users who don't have 'main.tf at all
	}
	if getFileFormat(options.HeaderFromFile) != ".tf" {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}
	lines := reader.Lines{
		FileName: filename,
		LineNum:  -1,
		Condition: func(line string) bool {
			line = strings.TrimSpace(line)
			return strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") || strings.HasPrefix(line, "*/")
		},
		Parser: func(line string) (string, bool) {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*/") {
				return "", false
			}
			if line == "*" {
				return "", true
			}
			line = strings.TrimPrefix(line, "* ")
			return line, true
		},
	}
	header, err := lines.Extract()
	if err != nil {
		return "", err
	}
	return strings.Join(header, "\n"), nil
}

func loadInputs(tfmodule *tfconfig.Module) ([]*tfconf.Input, []*tfconf.Input, []*tfconf.Input) {
	var inputs = make([]*tfconf.Input, 0, len(tfmodule.Variables))
	var required = make([]*tfconf.Input, 0, len(tfmodule.Variables))
	var optional = make([]*tfconf.Input, 0, len(tfmodule.Variables))

	for _, input := range tfmodule.Variables {
		inputDescription := input.Description
		if inputDescription == "" {
			inputDescription = loadComments(input.Pos.Filename, input.Pos.Line)
		}

		i := &tfconf.Input{
			Name:        input.Name,
			Type:        types.TypeOf(input.Type, input.Default),
			Description: types.String(inputDescription),
			Default:     types.ValueOf(input.Default),
			Required:    input.Required,
			Position: tfconf.Position{
				Filename: input.Pos.Filename,
				Line:     input.Pos.Line,
			},
		}

		inputs = append(inputs, i)
		if i.HasDefault() {
			optional = append(optional, i)
		} else {
			required = append(required, i)
		}
	}
	return inputs, required, optional
}

func loadOutputs(tfmodule *tfconfig.Module, options *Options) ([]*tfconf.Output, error) {
	outputs := make([]*tfconf.Output, 0, len(tfmodule.Outputs))
	values := make(map[string]*TerraformOutput)
	if options.OutputValues {
		var err error
		values, err = loadOutputValues(options)
		if err != nil {
			return nil, err
		}
	}
	for _, o := range tfmodule.Outputs {
		description := o.Description
		if description == "" {
			description = loadComments(o.Pos.Filename, o.Pos.Line)
		}
		output := &tfconf.Output{
			Name:        o.Name,
			Description: types.String(description),
			Position: tfconf.Position{
				Filename: o.Pos.Filename,
				Line:     o.Pos.Line,
			},
			ShowValue: options.OutputValues,
		}
		if options.OutputValues {
			output.Sensitive = values[output.Name].Sensitive
			if values[output.Name].Sensitive {
				output.Value = types.ValueOf(`<sensitive>`)
			} else {
				output.Value = types.ValueOf(values[output.Name].Value)
			}
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func loadOutputValues(options *Options) (map[string]*TerraformOutput, error) {
	var out []byte
	var err error
	if options.OutputValuesPath == "" {
		cmd := exec.Command("terraform", "output", "-json")
		cmd.Dir = options.Path
		if out, err = cmd.Output(); err != nil {
			return nil, fmt.Errorf("caught error while reading the terraform outputs: %v", err)
		}
	} else {
		if out, err = ioutil.ReadFile(options.OutputValuesPath); err != nil {
			return nil, fmt.Errorf("caught error while reading the terraform outputs file at %s: %v", options.OutputValuesPath, err)
		}
	}
	var terraformOutputs map[string]*TerraformOutput
	err = json.Unmarshal(out, &terraformOutputs)
	if err != nil {
		return nil, err
	}
	return terraformOutputs, err
}

func loadProviders(tfmodule *tfconfig.Module) []*tfconf.Provider {
	resources := []map[string]*tfconfig.Resource{tfmodule.ManagedResources, tfmodule.DataResources}
	discovered := make(map[string]*tfconf.Provider)
	for _, resource := range resources {
		for _, r := range resource {
			var version = ""
			if rv, ok := tfmodule.RequiredProviders[r.Provider.Name]; ok && len(rv.VersionConstraints) > 0 {
				version = strings.Join(rv.VersionConstraints, " ")
			}
			key := fmt.Sprintf("%s.%s", r.Provider.Name, r.Provider.Alias)
			discovered[key] = &tfconf.Provider{
				Name:    r.Provider.Name,
				Alias:   types.String(r.Provider.Alias),
				Version: types.String(version),
				Position: tfconf.Position{
					Filename: r.Pos.Filename,
					Line:     r.Pos.Line,
				},
			}
		}
	}
	providers := make([]*tfconf.Provider, 0, len(discovered))
	for _, provider := range discovered {
		providers = append(providers, provider)
	}
	return providers
}

func loadRequirements(tfmodule *tfconfig.Module) []*tfconf.Requirement {
	var requirements = make([]*tfconf.Requirement, 0)
	for _, core := range tfmodule.RequiredCore {
		requirements = append(requirements, &tfconf.Requirement{
			Name:    "terraform",
			Version: types.String(core),
		})
	}
	names := make([]string, 0, len(tfmodule.RequiredProviders))
	for n := range tfmodule.RequiredProviders {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, name := range names {
		for _, version := range tfmodule.RequiredProviders[name].VersionConstraints {
			requirements = append(requirements, &tfconf.Requirement{
				Name:    name,
				Version: types.String(version),
			})
		}
	}
	return requirements
}

func loadComments(filename string, lineNum int) string {
	lines := reader.Lines{
		FileName: filename,
		LineNum:  lineNum,
		Condition: func(line string) bool {
			return strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//")
		},
		Parser: func(line string) (string, bool) {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "#")
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			return line, true
		},
	}
	comment, err := lines.Extract()
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(comment, " ")
}

func sortItems(tfmodule *tfconf.Module, sortby *SortBy) {
	if sortby.Type {
		sort.Sort(inputsSortedByType(tfmodule.Inputs))
		sort.Sort(inputsSortedByType(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByType(tfmodule.OptionalInputs))
	} else if sortby.Name {
		if sortby.Required {
			sort.Sort(inputsSortedByRequired(tfmodule.Inputs))
			sort.Sort(inputsSortedByRequired(tfmodule.RequiredInputs))
			sort.Sort(inputsSortedByRequired(tfmodule.OptionalInputs))
		} else {
			sort.Sort(inputsSortedByName(tfmodule.Inputs))
			sort.Sort(inputsSortedByName(tfmodule.RequiredInputs))
			sort.Sort(inputsSortedByName(tfmodule.OptionalInputs))
		}
	} else {
		sort.Sort(inputsSortedByPosition(tfmodule.Inputs))
		sort.Sort(inputsSortedByPosition(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByPosition(tfmodule.OptionalInputs))
	}

	if sortby.Name || sortby.Type {
		sort.Sort(outputsSortedByName(tfmodule.Outputs))
	} else {
		sort.Sort(outputsSortedByPosition(tfmodule.Outputs))
	}

	if sortby.Name || sortby.Type {
		sort.Sort(providersSortedByName(tfmodule.Providers))
	} else {
		sort.Sort(providersSortedByPosition(tfmodule.Providers))
	}
}
