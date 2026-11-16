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

	"github.com/iancoleman/orderedmap"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// tfvarsJSON represents Terraform tfvars JSON format.
type tfvarsJSON struct {
	*generator

	config *print.Config
}

// NewTfvarsJSON returns new instance of TfvarsJSON.
func NewTfvarsJSON(config *print.Config) Type {
	return &tfvarsJSON{
		generator: newGenerator(config, false),
		config:    config,
	}
}

// Generate a Terraform module as Terraform tfvars JSON.
func (j *tfvarsJSON) Generate(module *terraform.Module) error {
	copy := orderedmap.New()
	copy.SetEscapeHTML(false)
	for _, i := range module.Inputs {
		copy.Set(i.Name, i.Default)
	}

	buffer := new(bytes.Buffer)
	encoder := jsonsdk.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(copy); err != nil {
		return err
	}

	j.funcs(withContent(strings.TrimSuffix(buffer.String(), "\n")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"tfvars json": NewTfvarsJSON,
	})
}
