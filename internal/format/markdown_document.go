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

//go:embed templates/markdown_document*.tmpl
var markdownDocumentFS embed.FS

// MarkdownDocument represents Markdown Document format.
type MarkdownDocument struct {
	template *template.Template
	settings *print.Settings
}

// NewMarkdownDocument returns new instance of Document.
func NewMarkdownDocument(settings *print.Settings) print.Engine {
	items := readTemplateItems(markdownDocumentFS, "markdown_document")

	tt := template.New(settings, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			result, extraline := printFencedCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := printFencedCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})
	return &MarkdownDocument{
		template: tt,
		settings: settings,
	}
}

// Generate a Terraform module as Markdown document.
func (d *MarkdownDocument) Generate(module *terraform.Module) (*print.Generator, error) {
	funcs := []print.GenerateFunc{}

	err := print.ForEach(func(name string, fn print.GeneratorCallback) error {
		rendered, err := d.template.Render(name, module)
		if err != nil {
			return err
		}

		funcs = append(funcs, fn(sanitize(rendered)))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return print.NewGenerator("markdown document", funcs...), nil
}

func init() {
	register(map[string]initializerFn{
		"markdown document": NewMarkdownDocument,
		"markdown doc":      NewMarkdownDocument,
		"md document":       NewMarkdownDocument,
		"md doc":            NewMarkdownDocument,
	})
}
