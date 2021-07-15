/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"encoding/xml"
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// XML represents XML format.
type XML struct {
	settings *print.Settings
}

// NewXML returns new instance of XML.
func NewXML(settings *print.Settings) print.Engine {
	return &XML{
		settings: settings,
	}
}

// Generate a Terraform module as xml.
func (x *XML) Generate(module *terraform.Module) (*print.Generator, error) {
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

	print.CopySections(x.settings, module, copy)

	out, err := xml.MarshalIndent(copy, "", "  ")
	if err != nil {
		return nil, err
	}

	return print.NewGenerator(
		"xml",
		print.WithContent(strings.TrimSuffix(string(out), "\n")),
	), nil
}

func init() {
	register(map[string]initializerFn{
		"xml": NewXML,
	})
}
