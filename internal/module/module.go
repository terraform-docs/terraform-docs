package module

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/reader"
	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// LoadWithOptions returns new instance of Module with all the inputs and
// outputs dircoverd from provided 'path' containing Terraform config
func LoadWithOptions(options *Options) (*tfconf.Module, error) {
	tfmodule := loadModule(options.Path)

	header := loadHeader(options.Path)
	inputs, required, optional := loadInputs(tfmodule)
	outputs := loadOutputs(tfmodule)
	providers := loadProviders(tfmodule)

	module := &tfconf.Module{
		Header:    header,
		Inputs:    inputs,
		Outputs:   outputs,
		Providers: providers,

		RequiredInputs: required,
		OptionalInputs: optional,
	}
	sortItems(module, options.SortBy)
	return module, nil
}

func loadModule(path string) *tfconfig.Module {
	module, diag := tfconfig.LoadModule(path)
	if diag != nil && diag.HasErrors() {
		log.Fatal(diag)
	}
	return module
}

func loadHeader(path string) string {
	filename := filepath.Join(path, "main.tf")
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
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
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(header, "\n")
}

func loadInputs(tfmodule *tfconfig.Module) ([]*tfconf.Input, []*tfconf.Input, []*tfconf.Input) {
	var inputs = make([]*tfconf.Input, 0, len(tfmodule.Variables))
	var required = make([]*tfconf.Input, 0, len(tfmodule.Variables))
	var optional = make([]*tfconf.Input, 0, len(tfmodule.Variables))

	for _, input := range tfmodule.Variables {
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
			inputDescription = loadComments(input.Pos.Filename, input.Pos.Line-1)
		}

		i := &tfconf.Input{
			Name:        input.Name,
			Type:        types.TFString(inputType),
			Description: types.TFString(inputDescription),
			Default:     input.Default,
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

func loadOutputs(tfmodule *tfconfig.Module) []*tfconf.Output {
	outputs := make([]*tfconf.Output, 0, len(tfmodule.Outputs))
	for _, output := range tfmodule.Outputs {
		description := output.Description
		if description == "" {
			description = loadComments(output.Pos.Filename, output.Pos.Line-1)
		}
		outputs = append(outputs, &tfconf.Output{
			Name:        output.Name,
			Description: types.TFString(description),
			Position: tfconf.Position{
				Filename: output.Pos.Filename,
				Line:     output.Pos.Line,
			},
		})
	}
	return outputs
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
				Alias:   types.TFString(r.Provider.Alias),
				Version: types.TFString(version),
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
	if sortby.Name {
		sort.Sort(providersSortedByName(tfmodule.Providers))
	} else {
		sort.Sort(providersSortedByPosition(tfmodule.Providers))
	}

	if sortby.Name {
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

	if sortby.Name {
		sort.Sort(outputsSortedByName(tfmodule.Outputs))
	} else {
		sort.Sort(outputsSortedByPosition(tfmodule.Outputs))
	}
}
