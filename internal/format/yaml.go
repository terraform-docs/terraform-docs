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

	"gopkg.in/yaml.v3"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// YAML represents YAML format.
type YAML struct{}

// NewYAML returns new instance of YAML.
func NewYAML(settings *print.Settings) print.Engine {
	return &YAML{}
}

// Print a Terraform module as yaml.
func (y *YAML) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	copy := &terraform.Module{
		Header:       "",
		Inputs:       make([]*terraform.Input, 0),
		ModuleCalls:  make([]*terraform.ModuleCall, 0),
		Outputs:      make([]*terraform.Output, 0),
		Providers:    make([]*terraform.Provider, 0),
		Requirements: make([]*terraform.Requirement, 0),
		Resources:    make([]*terraform.Resource, 0),
	}

	if settings.ShowHeader {
		copy.Header = module.Header
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

	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(2)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}

func init() {
	register(map[string]initializerFn{
		"yaml": NewYAML,
	})
}
