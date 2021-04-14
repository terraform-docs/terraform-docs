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
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// JSON represents JSON format.
type JSON struct{}

// NewJSON returns new instance of JSON.
func NewJSON(settings *print.Settings) print.Engine {
	return &JSON{}
}

// Print a Terraform module as json.
func (j *JSON) Print(module *terraform.Module, settings *print.Settings) (string, error) {
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

	print.CopySections(settings, module, copy)

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(settings.EscapeCharacters)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}

func init() {
	register(map[string]initializerFn{
		"json": NewJSON,
	})
}
