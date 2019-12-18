package doc

import (
	"encoding/json"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

// Doc represents a Terraform module.
type Doc struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
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
	for _, input := range module.Variables {
		var defaultValue string
		if input.Default != nil {
			marshaled, err := json.MarshalIndent(input.Default, "", "  ")
			if err != nil {
				return nil, err
			}
			defaultValue = string(marshaled)
		}
		inputs = append(inputs, Input{
			Name:        input.Name,
			Type:        input.Type,
			Description: input.Description,
			Default:     defaultValue,
		})
	}

	var outputs = make([]Output, 0, len(module.Outputs))
	for _, output := range module.Outputs {
		outputs = append(outputs, Output{
			Name:        output.Name,
			Description: output.Description,
		})
	}

	doc := &Doc{
		Inputs:  inputs,
		Outputs: outputs,
	}
	return doc, nil
}
