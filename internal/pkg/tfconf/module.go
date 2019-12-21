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
	Variables         []*Variable `json:"variables"`
	Outputs           []*Output   `json:"outputs"`
	RequiredVariables []*Variable `json:"-"`
	OptionalVariables []*Variable `json:"-"`
}

// HasInputs indicates if the document has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Variables) > 0
}

// HasOutputs indicates if the document has outputs.
func (m *Module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

func (m *Module) sortVariablesByName() {
	sort.Sort(variablesSortedByName(m.Variables))
	sort.Sort(variablesSortedByName(m.RequiredVariables))
	sort.Sort(variablesSortedByName(m.OptionalVariables))
}

func (m *Module) sortVariablesByRequired() {
	sort.Sort(variablesSortedByRequired(m.Variables))
	sort.Sort(variablesSortedByRequired(m.RequiredVariables))
	sort.Sort(variablesSortedByRequired(m.OptionalVariables))
}

func (m *Module) sortOutputsByName() {
	sort.Sort(outputsSortedByName(m.Outputs))
}

// Sort sorts list of inputs and outputs based on provided flags (name, required, etc)
func (m *Module) Sort(settings *print.Settings) {
	if settings.SortByName {
		if settings.SortVariablesByRequired {
			m.sortVariablesByRequired()
		} else {
			m.sortVariablesByName()
		}
	}

	if settings.SortByName {
		m.sortOutputsByName()
	}

}

// CreateModule returns new instance of Module with all the variables and
// outputs dircoverd from provided 'path' containing Terraform config
func CreateModule(path string) (*Module, error) {
	mod := loadModule(path)

	var variables = make([]*Variable, 0, len(mod.Variables))
	var requiredVariables = make([]*Variable, 0, len(mod.Variables))
	var optionalVariables = make([]*Variable, 0, len(mod.Variables))

	for _, variable := range mod.Variables {
		variableType := variable.Type
		if variable.Type == "" {
			variableType = "any"
		}

		var defaultValue string
		if variable.Default != nil {
			marshaled, err := json.MarshalIndent(variable.Default, "", "  ")
			if err != nil {
				return nil, err
			}
			defaultValue = string(marshaled)

			if variableType == "any" {
				switch xType := fmt.Sprintf("%T", variable.Default); xType {
				case "string":
					variableType = "string"
				case "int", "int8", "int16", "int32", "int64", "float32", "float64":
					variableType = "number"
				case "bool":
					variableType = "bool"
				case "[]interface {}":
					variableType = "list"
				case "map[string]interface {}":
					variableType = "map"
				}
			}
		}

		v := &Variable{
			Name:        variable.Name,
			Type:        variableType,
			Description: variable.Description,
			Default:     defaultValue,
		}

		variables = append(variables, v)
		if v.HasDefault() {
			optionalVariables = append(optionalVariables, v)
		} else {
			requiredVariables = append(requiredVariables, v)
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
		Variables:         variables,
		Outputs:           outputs,
		RequiredVariables: requiredVariables,
		OptionalVariables: optionalVariables,
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
