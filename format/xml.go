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

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// xml represents XML format.
type xml struct {
	*print.Generator

	config   *print.Config
	settings *print.Settings
}

// NewXML returns new instance of XML.
func NewXML(config *print.Config) Type {
	settings, _ := config.Extract()

	return &xml{
		Generator: print.NewGenerator("xml", config.ModuleRoot),
		config:    config,
		settings:  settings,
	}
}

// Generate a Terraform module as xml.
func (x *xml) Generate(module *terraform.Module) error {
	copy := copySections(x.settings, module)

	out, err := xmlsdk.MarshalIndent(copy, "", "  ")
	if err != nil {
		return err
	}

	x.Generator.Funcs(print.WithContent(strings.TrimSuffix(string(out), "\n")))

	return nil
}

func init() {
	register(map[string]initializerFn{
		"xml": NewXML,
	})
}
