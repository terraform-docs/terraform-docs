package format

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// JSON represents JSON format.
type JSON struct{}

// NewJSON returns new instance of JSON.
func NewJSON(settings *print.Settings) *JSON {
	return &JSON{}
}

// Print prints a Terraform module as json.
func (j *JSON) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	copy := &tfconf.Module{
		Header:       "",
		Inputs:       make([]*tfconf.Input, 0),
		Outputs:      make([]*tfconf.Output, 0),
		Providers:    make([]*tfconf.Provider, 0),
		Requirements: make([]*tfconf.Requirement, 0),
	}

	if settings.ShowHeader {
		copy.Header = module.Header
	}
	if settings.ShowInputs {
		copy.Inputs = module.Inputs
	}
	if settings.ShowOutputs {
		copy.Outputs = module.Outputs
	}
	if settings.ShowProviders {
		copy.Providers = module.Providers
	}
	if settings.ShowRequirements {
		copy.Requirements = module.Requirements
	}

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(settings.EscapeCharacters)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}
