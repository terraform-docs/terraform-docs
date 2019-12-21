package tfconf

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

// Module represents a Terraform mod.
type Module struct {
	Inputs         []*Input  `json:"inputs"`
	Outputs        []*Output `json:"outputs"`
	RequiredInputs []*Input  `json:"-"`
	OptionalInputs []*Input  `json:"-"`
}

// HasInputs indicates if the document has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Inputs) > 0
}

// HasOutputs indicates if the document has outputs.
func (m *Module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

func (m *Module) sortInputsByName() {
	sort.Sort(inputsSortedByName(m.Inputs))
	sort.Sort(inputsSortedByName(m.RequiredInputs))
	sort.Sort(inputsSortedByName(m.OptionalInputs))
}

func (m *Module) sortInputsByRequired() {
	sort.Sort(inputsSortedByRequired(m.Inputs))
	sort.Sort(inputsSortedByRequired(m.RequiredInputs))
	sort.Sort(inputsSortedByRequired(m.OptionalInputs))
}

func (m *Module) sortOutputsByName() {
	sort.Sort(outputsSortedByName(m.Outputs))
}

// Sort sorts list of inputs and outputs based on provided flags (name, required, etc)
func (m *Module) Sort(settings *print.Settings) {
	if settings.SortByName {
		if settings.SortInputsByRequired {
			m.sortInputsByRequired()
		} else {
			m.sortInputsByName()
		}
	}

	if settings.SortByName {
		m.sortOutputsByName()
	}

}

// CreateModule returns new instance of Module with all the inputs and
// outputs dircoverd from provided 'path' containing Terraform config
func CreateModule(path string) (*Module, error) {
	mod := loadModule(path)

	var inputs = make([]*Input, 0, len(mod.Variables))
	var requiredInputs = make([]*Input, 0, len(mod.Variables))
	var optionalInputs = make([]*Input, 0, len(mod.Variables))

	for _, input := range mod.Variables {
		inputType := input.Type
		if input.Type == "" {
			inputType = "any"
		}

		var defaultValue string
		if input.Default != nil {
			marshaled, err := json.MarshalIndent(input.Default, "", "  ")
			if err != nil {
				return nil, err
			}
			defaultValue = string(marshaled)

			if inputType == "any" {
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

		i := &Input{
			Name:        input.Name,
			Type:        inputType,
			Description: input.Description,
			Default:     defaultValue,
		}

		inputs = append(inputs, i)
		if i.HasDefault() {
			optionalInputs = append(optionalInputs, i)
		} else {
			requiredInputs = append(requiredInputs, i)
		}
	}

	var outputs = make([]*Output, 0, len(mod.Outputs))
	for _, output := range mod.Outputs {
		outputs = append(outputs, &Output{
			Name:        output.Name,
			Description: output.Description,
		})
	}

	module := &Module{
		Inputs:         inputs,
		Outputs:        outputs,
		RequiredInputs: requiredInputs,
		OptionalInputs: optionalInputs,
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
