package format

import (
	"encoding/xml"
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/pkg/print"
)

// XML represents XML format.
type XML struct{}

// NewXML returns new instance of XML.
func NewXML(settings *print.Settings) *XML {
	return &XML{}
}

// Print prints a Terraform module as xml.
func (x *XML) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	copy := &terraform.Module{
		Header:       "",
		Inputs:       make([]*terraform.Input, 0),
		Outputs:      make([]*terraform.Output, 0),
		Providers:    make([]*terraform.Provider, 0),
		Requirements: make([]*terraform.Requirement, 0),
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

	out, err := xml.MarshalIndent(copy, "", "  ")
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(out), "\n"), nil
}
