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

	"github.com/iancoleman/orderedmap"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// TfvarsJSON represents Terraform tfvars JSON format.
type TfvarsJSON struct {
	settings *print.Settings
}

// NewTfvarsJSON returns new instance of TfvarsJSON.
func NewTfvarsJSON(settings *print.Settings) print.Engine {
	return &TfvarsJSON{
		settings: settings,
	}
}

// Generate a Terraform module as Terraform tfvars JSON.
func (j *TfvarsJSON) Generate(module *terraform.Module) (*print.Generator, error) {
	copy := orderedmap.New()
	copy.SetEscapeHTML(false)
	for _, i := range module.Inputs {
		copy.Set(i.Name, i.Default)
	}

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(copy)
	if err != nil {
		return nil, err
	}

	return print.NewGenerator(
		"tfvars json",
		print.WithContent(strings.TrimSuffix(buffer.String(), "\n")),
	), nil

}

func init() {
	register(map[string]initializerFn{
		"tfvars json": NewTfvarsJSON,
	})
}
