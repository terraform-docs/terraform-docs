package yaml

import (
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"gopkg.in/yaml.v2"
)

// Print prints a document as yaml.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	copy := &tfconf.Module{
		Header:    "",
		Providers: make([]*tfconf.Provider, 0),
		Inputs:    make([]*tfconf.Input, 0),
		Outputs:   make([]*tfconf.Output, 0),
	}

	if settings.ShowHeader {
		copy.Header = module.Header
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

	out, err := yaml.Marshal(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(out), "\n"), nil
}
