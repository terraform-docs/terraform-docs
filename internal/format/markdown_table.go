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
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/template"
)

//go:embed templates/markdown_table*.tmpl
var markdownTableFS embed.FS

// MarkdownTable represents Markdown Table format.
type MarkdownTable struct {
	template *template.Template
	settings *print.Settings
}

// NewMarkdownTable returns new instance of Table.
func NewMarkdownTable(settings *print.Settings) print.Engine {
	items := readTemplateItems(markdownTableFS, "markdown_table")

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
	return &MarkdownTable{
		template: tt,
		settings: settings,
	}
}

// Generate a Terraform module as Markdown tables.
func (t *MarkdownTable) Generate(module *terraform.Module) (*print.Generator, error) {
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

	return print.NewGenerator("markdown table", funcs...), nil
}

func init() {
	register(map[string]initializerFn{
		"markdown":       NewMarkdownTable,
		"markdown table": NewMarkdownTable,
		"markdown tbl":   NewMarkdownTable,
		"md":             NewMarkdownTable,
		"md table":       NewMarkdownTable,
		"md tbl":         NewMarkdownTable,
	})
}
