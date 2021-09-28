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
	"github.com/terraform-docs/terraform-docs/terraform"
)

// YAML represents YAML format.
type YAML struct {
	settings *print.Settings
}

// NewYAML returns new instance of YAML.
func NewYAML(settings *print.Settings) print.Engine {
	return &YAML{
		settings: settings,
	}
}

// Generate a Terraform module as yaml.
func (y *YAML) Generate(module *terraform.Module) (*print.Generator, error) {
	copy := &terraform.Module{
		Header:       "",
		Footer:       "",
		Inputs:       make([]*terraform.Input, 0),
		ModuleCalls:  make([]*terraform.ModuleCall, 0),
		Outputs:      make([]*terraform.Output, 0),
		Providers:    make([]*terraform.Provider, 0),
		Requirements: make([]*terraform.Requirement, 0),
		Resources:    make([]*terraform.Resource, 0),
	}

	print.CopySections(y.settings, module, copy)

	buffer := new(bytes.Buffer)

	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(2)

	err := encoder.Encode(copy)
	if err != nil {
		return nil, err
	}

	return print.NewGenerator(
		"yaml",
		print.WithContent(strings.TrimSuffix(buffer.String(), "\n")),
	), nil
}

func init() {
	register(map[string]initializerFn{
		"yaml": NewYAML,
	})
}
