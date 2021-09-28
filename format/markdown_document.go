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

//go:embed templates/markdown_document*.tmpl
var markdownDocumentFS embed.FS

// markdownDocument represents Markdown Document format.
type markdownDocument struct {
	*print.Generator

	config   *print.Config
	template *template.Template
	settings *print.Settings
}

// NewMarkdownDocument returns new instance of Markdown Document.
func NewMarkdownDocument(config *print.Config) Type {
	settings, _ := config.Extract()
	items := readTemplateItems(markdownDocumentFS, "markdown_document")

	tt := template.New(settings, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			result, extraline := PrintFencedCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := PrintFencedCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})

	return &markdownDocument{
		Generator: print.NewGenerator("json", config.ModuleRoot),
		config:    config,
		template:  tt,
		settings:  settings,
	}
}

// Generate a Terraform module as Markdown document.
func (d *markdownDocument) Generate(module *terraform.Module) error {
	err := d.Generator.ForEach(func(name string) (string, error) {
		rendered, err := d.template.Render(name, module)
		if err != nil {
			return "", err
		}
		return sanitize(rendered), nil
	})

	return err
}

func init() {
	register(map[string]initializerFn{
		"markdown document": NewMarkdownDocument,
		"markdown doc":      NewMarkdownDocument,
		"md document":       NewMarkdownDocument,
		"md doc":            NewMarkdownDocument,
	})
}
