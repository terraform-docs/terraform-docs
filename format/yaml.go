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

	yamlv3 "gopkg.in/yaml.v3"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// yaml represents YAML format.
type yaml struct {
	*generator

	config *print.Config
}

// NewYAML returns new instance of YAML.
func NewYAML(config *print.Config) Type {
	return &yaml{
		generator: newGenerator(config, false),
		config:    config,
	}
}

// Generate a Terraform module as YAML.
func (y *yaml) Generate(module *terraform.Module) error {
	copy := copySections(y.config, module)

	buffer := new(bytes.Buffer)
	encoder := yamlv3.NewEncoder(buffer)
	encoder.SetIndent(2)

	if err := encoder.Encode(copy); err != nil {
		return err
	}

	y.funcs(withContent(strings.TrimSuffix(buffer.String(), "\n")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"yaml": NewYAML,
	})
}
