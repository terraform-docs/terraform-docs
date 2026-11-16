/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	xmlsdk "encoding/xml"
	"strings"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// xml represents XML format.
type xml struct {
	*generator

	config *print.Config
}

// NewXML returns new instance of XML.
func NewXML(config *print.Config) Type {
	return &xml{
		generator: newGenerator(config, false),
		config:    config,
	}
}

// Generate a Terraform module as xml.
func (x *xml) Generate(module *terraform.Module) error {
	copy := copySections(x.config, module)

	out, err := xmlsdk.MarshalIndent(copy, "", "  ")
	if err != nil {
		return err
	}

	x.funcs(withContent(strings.TrimSuffix(string(out), "\n")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"xml": NewXML,
	})
}
