/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"bytes"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// TOML represents TOML format.
type TOML struct{}

// NewTOML returns new instance of TOML.
func NewTOML(settings *print.Settings) print.Engine {
	return &TOML{}
}

// Print a Terraform module as toml.
func (t *TOML) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	copy := terraform.Module{
		Header:       "",
		Footer:       "",
		Providers:    make([]*terraform.Provider, 0),
		Inputs:       make([]*terraform.Input, 0),
		ModuleCalls:  make([]*terraform.ModuleCall, 0),
		Outputs:      make([]*terraform.Output, 0),
		Requirements: make([]*terraform.Requirement, 0),
		Resources:    make([]*terraform.Resource, 0),
	}

	if settings.ShowHeader {
		copy.Header = module.Header
	}
	if settings.ShowFooter {
		copy.Footer = module.Footer
	}
	if settings.ShowInputs {
		copy.Inputs = module.Inputs
	}
	if settings.ShowModuleCalls {
		copy.ModuleCalls = module.ModuleCalls
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
	if settings.ShowResources {
		copy.Resources = module.Resources
	}

	buffer := new(bytes.Buffer)
	encoder := toml.NewEncoder(buffer)
	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}

func init() {
	register(map[string]initializerFn{
		"toml": NewTOML,
	})
}
