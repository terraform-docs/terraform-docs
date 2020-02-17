package format

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// TOML represents TOML format.
type TOML struct{}

// NewTOML returns new instance of TOML.
func NewTOML(settings *print.Settings) *TOML {
	return &TOML{}
}

// Print prints a Terraform module as toml.
func (t *TOML) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	copy := tfconf.Module{
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

	buffer := new(bytes.Buffer)
	encoder := toml.NewEncoder(buffer)
	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
