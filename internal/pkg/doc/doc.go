package doc

import (
	"encoding/json"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

// Doc represents a Terraform module.
type Doc struct {
	Inputs         []Input  `json:"inputs"`
	Outputs        []Output `json:"outputs"`
	RequiredInputs []Input  `json:"-"`
	OptionalInputs []Input  `json:"-"`
}

// HasInputs indicates if the document has inputs.
func (d *Doc) HasInputs() bool {
	return len(d.Inputs) > 0
}

// HasOutputs indicates if the document has outputs.
func (d *Doc) HasOutputs() bool {
	return len(d.Outputs) > 0
}

// Create TODO
func Create(module *tfconfig.Module) (*Doc, error) {
	var inputs = make([]Input, 0, len(module.Variables))
	var requiredInputs = make([]Input, 0, len(module.Variables))
	var optionalInputs = make([]Input, 0, len(module.Variables))

	for _, input := range module.Variables {
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
		}

		i := Input{
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

	var outputs = make([]Output, 0, len(module.Outputs))
	for _, output := range module.Outputs {
		outputs = append(outputs, Output{
			Name:        output.Name,
			Description: output.Description,
		})
	}

	doc := &Doc{
		Inputs:         inputs,
		Outputs:        outputs,
		RequiredInputs: requiredInputs,
		OptionalInputs: optionalInputs,
	}
	return doc, nil
}
