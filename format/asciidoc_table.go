/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"embed"
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/template"
	"github.com/terraform-docs/terraform-docs/terraform"
)

//go:embed templates/asciidoc_table*.tmpl
var asciidocTableFS embed.FS

// asciidocTable represents AsciiDoc Table format.
type asciidocTable struct {
	*print.Generator

	config   *print.Config
	template *template.Template
	settings *print.Settings
}

// NewAsciidocTable returns new instance of Asciidoc Table.
func NewAsciidocTable(config *print.Config) Type {
	settings, _ := config.Extract()
	items := readTemplateItems(asciidocTableFS, "asciidoc_table")

	settings.EscapeCharacters = false

	tt := template.New(settings, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			inputType, _ := PrintFencedCodeBlock(t, "")
			return inputType
		},
		"value": func(v string) string {
			var result = "n/a"
			if v != "" {
				result, _ = PrintFencedCodeBlock(v, "")
			}
			return result
		},
	})

	return &asciidocTable{
		Generator: print.NewGenerator("json", config.ModuleRoot),
		config:    config,
		template:  tt,
		settings:  settings,
	}
}

// Generate a Terraform module as AsciiDoc tables.
func (t *asciidocTable) Generate(module *terraform.Module) error {
	err := t.Generator.ForEach(func(name string) (string, error) {
		rendered, err := t.template.Render(name, module)
		if err != nil {
			return "", err
		}
		return sanitize(rendered), nil
	})

	return err
}

func init() {
	register(map[string]initializerFn{
		"asciidoc":       NewAsciidocTable,
		"asciidoc table": NewAsciidocTable,
		"asciidoc tbl":   NewAsciidocTable,
		"adoc":           NewAsciidocTable,
		"adoc table":     NewAsciidocTable,
		"adoc tbl":       NewAsciidocTable,
	})
}
