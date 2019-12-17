package doc

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

type Doc struct {
	Inputs    []Input    `json:"variables"`
	Outputs   []Output   `json:"outputs"`
	Providers []Provider `json:"providers"`
}

// TODO: verify that the side effects to tracker stick
func discoverAliases(tracker map[string]Provider, versionLookup map[string][]string, resources map[string]*tfconfig.Resource) {
	for _, resource := range resources {
		key := fmt.Sprintf("%s.%s", resource.Provider.Name, resource.Provider.Alias)
		var version = ""
		if requiredVersion, ok := versionLookup[resource.Provider.Name]; ok {
			version = strings.Join(requiredVersion, " ")
		}
		tracker[key] = Provider{
			Name:    resource.Provider.Name,
			Alias:   resource.Provider.Alias,
			Version: version,
		}
	}
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

	var providerSet = make(map[string]Provider)
	discoverAliases(providerSet, module.RequiredProviders, module.DataResources)
	discoverAliases(providerSet, module.RequiredProviders, module.ManagedResources)
	var providers = make([]Provider, 0, len(providerSet))
	for _, provider := range providerSet {
		providers = append(providers, provider)
	}

	if settings.SortInputsByRequired {
		sort.Sort(variablesSortedByRequired(inputs))
	} else {
		sort.Sort(variablesSortedByName(inputs))
	}
	sort.Sort(outputsSortedByName(outputs))
	sort.Sort(providersSortedByRequired(providers))

	doc := &Doc{
		Inputs:    inputs,
		Outputs:   outputs,
		Providers: providers,
	}
	return doc, nil

}
