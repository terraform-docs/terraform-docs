package format

import (
	"bytes"
	"strings"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"gopkg.in/yaml.v3"
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

	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(2)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}
