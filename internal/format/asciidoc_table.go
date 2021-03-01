/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	_ "embed" //nolint
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/template"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

//go:embed templates/asciidoc_table.tmpl
var asciidocTableTpl []byte

// AsciidocTable represents AsciiDoc Table format.
type AsciidocTable struct {
	template *template.Template
}

// NewAsciidocTable returns new instance of AsciidocTable.
func NewAsciidocTable(settings *print.Settings) print.Engine {
	settings.EscapeCharacters = false
	tt := template.New(settings, &template.Item{
		Name: "table",
		Text: string(asciidocTableTpl),
	})
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
	}
}

// Print a Terraform module as AsciiDoc tables.
func (t *AsciidocTable) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	rendered, err := t.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
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
