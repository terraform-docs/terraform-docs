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

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/template"
	"github.com/rquadling/terraform-docs/terraform"
)

//go:embed templates/markdown_table*.tmpl
var markdownTableFS embed.FS

// markdownTable represents Markdown Table format.
type markdownTable struct {
	*generator

	config   *print.Config
	template *template.Template
}

// NewMarkdownTable returns new instance of Markdown Table.
func NewMarkdownTable(config *print.Config) Type {
	items := readTemplateItems(markdownTableFS, "markdown_table")

	tt := template.New(config, items...)
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

	return &markdownTable{
		generator: newGenerator(config, true),
		config:    config,
		template:  tt,
	}
}

// Generate a Terraform module as Markdown tables.
func (t *markdownTable) Generate(module *terraform.Module) error {
	err := t.forEach(func(name string) (string, error) {
		rendered, err := t.template.Render(name, module)
		if err != nil {
			return "", err
		}
		return sanitize(rendered), nil
	})

	t.funcs(withModule(module))

	return err
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
