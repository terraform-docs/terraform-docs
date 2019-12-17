package doc

import (
	"encoding/json"
	"sort"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

type Doc struct {
	Inputs  []Input  `json:"variables"`
	Outputs []Output `json:"outputs"`
}

func Create(module *tfconfig.Module, settings *print.Settings) (*Doc, error) {
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

	if settings.SortInputsByRequired {
		sort.Sort(variablesSortedByRequired(inputs))
	} else {
		sort.Sort(variablesSortedByName(inputs))
	}
	sort.Sort(outputsSortedByName(outputs))

	doc := &Doc{
		Inputs:  inputs,
		Outputs: outputs,
	}
	return doc, nil

}
