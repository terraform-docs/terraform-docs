package json

import (
	"encoding/json"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

const (
	indent string = "  "
	prefix string = ""
)

// Print prints a document as json.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	copy := &tfconf.Module{
		Providers: make([]*tfconf.Provider, 0),
		Inputs:    make([]*tfconf.Input, 0),
		Outputs:   make([]*tfconf.Output, 0),
	}

	if settings.ShowProviders {
		copy.Providers = module.Providers
	}
	if settings.ShowInputs {
		copy.Inputs = module.Inputs
	}
	if settings.ShowOutputs {
		copy.Outputs = module.Outputs
	}

	buffer, err := json.MarshalIndent(copy, prefix, indent)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
