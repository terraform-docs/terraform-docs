package format

import (
	"strings"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"gopkg.in/yaml.v2"
)

// YAML represents YAML format.
type YAML struct{}

// NewYAML returns new instance of YAML.
func NewYAML(settings *print.Settings) *YAML {
	return &YAML{}
}

// Print prints a Terraform module as yaml.
func (y *YAML) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	copy := &tfconf.Module{
		Requirements: make([]*tfconf.Requirement, 0),
		Header:       "",
		Providers:    make([]*tfconf.Provider, 0),
		Inputs:       make([]*tfconf.Input, 0),
		Outputs:      make([]*tfconf.Output, 0),
	}

	if settings.ShowRequirements {
		copy.Requirements = module.Requirements
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
