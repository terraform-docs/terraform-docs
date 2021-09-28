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

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/template"
	"github.com/terraform-docs/terraform-docs/terraform"
)

//go:embed templates/asciidoc_table*.tmpl
var asciidocTableFS embed.FS

// AsciidocTable represents AsciiDoc Table format.
type AsciidocTable struct {
	template *template.Template
	settings *print.Settings
}

// NewAsciidocTable returns new instance of AsciidocTable.
func NewAsciidocTable(settings *print.Settings) print.Engine {
	items := readTemplateItems(asciidocTableFS, "asciidoc_table")

	settings.EscapeCharacters = false

	tt := template.New(settings, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			inputType, _ := printFencedCodeBlock(t, "")
			return inputType
		},
		"value": func(v string) string {
			var result = "n/a"
			if v != "" {
				result, _ = printFencedCodeBlock(v, "")
			}
			return result
		},
	})
	return &AsciidocTable{
		template: tt,
		settings: settings,
	}
}

// Generate a Terraform module as AsciiDoc tables.
func (t *AsciidocTable) Generate(module *terraform.Module) (*print.Generator, error) {
	funcs := []print.GenerateFunc{}

	err := print.ForEach(func(name string, fn print.GeneratorCallback) error {
		rendered, err := t.template.Render(name, module)
		if err != nil {
			return err
		}

		funcs = append(funcs, fn(sanitize(rendered)))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return print.NewGenerator("asciidoc table", funcs...), nil
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
