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
	jsonsdk "encoding/json"
	"strings"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// json represents JSON format.
type json struct {
	*print.Generator

	config   *print.Config
	settings *print.Settings
}

// NewJSON returns new instance of JSON.
func NewJSON(config *print.Config) Type {
	settings, _ := config.Extract()

	return &json{
		Generator: print.NewGenerator("json", config.ModuleRoot),
		config:    config,
		settings:  settings,
	}
}

// Generate a Terraform module as json.
func (j *json) Generate(module *terraform.Module) error {
	copy := copySections(j.settings, module)

	buffer := new(bytes.Buffer)
	encoder := jsonsdk.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(j.settings.EscapeCharacters)

	if err := encoder.Encode(copy); err != nil {
		return err
	}

	j.Generator.Funcs(print.WithContent(strings.TrimSuffix(buffer.String(), "\n")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"json": NewJSON,
	})
}
