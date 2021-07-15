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
	"encoding/json"
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// JSON represents JSON format.
type JSON struct {
	settings *print.Settings
}

// NewJSON returns new instance of JSON.
func NewJSON(settings *print.Settings) print.Engine {
	return &JSON{
		settings: settings,
	}
}

// Generate a Terraform module as json.
func (j *JSON) Generate(module *terraform.Module) (*print.Generator, error) {
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

	print.CopySections(j.settings, module, copy)

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(j.settings.EscapeCharacters)

	err := encoder.Encode(copy)
	if err != nil {
		return nil, err
	}

	return print.NewGenerator(
		"json",
		print.WithContent(strings.TrimSuffix(buffer.String(), "\n")),
	), nil
}

func init() {
	register(map[string]initializerFn{
		"json": NewJSON,
	})
}
