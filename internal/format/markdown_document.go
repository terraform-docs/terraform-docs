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

//go:embed templates/markdown_document.tmpl
var markdownDocumentTpl []byte

// MarkdownDocument represents Markdown Document format.
type MarkdownDocument struct {
	template *template.Template
}

// NewMarkdownDocument returns new instance of Document.
func NewMarkdownDocument(settings *print.Settings) print.Engine {
	tt := template.New(settings, &template.Item{
		Name: "document",
		Text: string(markdownDocumentTpl),
	})
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
	}
}

// Print a Terraform module as Markdown document.
func (d *MarkdownDocument) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	rendered, err := d.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
}

func init() {
	register(map[string]initializerFn{
		"markdown document": NewMarkdownDocument,
		"markdown doc":      NewMarkdownDocument,
		"md document":       NewMarkdownDocument,
		"md doc":            NewMarkdownDocument,
	})
}
