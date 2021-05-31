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

//go:embed templates/asciidoc_document*.tmpl
var asciidocsDocumentFS embed.FS

// AsciidocDocument represents AsciiDoc Document format.
type AsciidocDocument struct {
	template *template.Template
	settings *print.Settings
}

// NewAsciidocDocument returns new instance of AsciidocDocument.
func NewAsciidocDocument(settings *print.Settings) print.Engine {
	items := readTemplateItems(asciidocsDocumentFS, "asciidoc_document")

	settings.EscapeCharacters = false

	tt := template.New(settings, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			result, extraline := printFencedAsciidocCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := printFencedAsciidocCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})
	return &AsciidocDocument{
		template: tt,
		settings: settings,
	}
}

// Generate a Terraform module as AsciiDoc document.
func (d *AsciidocDocument) Generate(module *terraform.Module) (*print.Generator, error) {
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

	return print.NewGenerator("asciidoc document", funcs...), nil
}

func init() {
	register(map[string]initializerFn{
		"asciidoc document": NewAsciidocDocument,
		"asciidoc doc":      NewAsciidocDocument,
		"adoc document":     NewAsciidocDocument,
		"adoc doc":          NewAsciidocDocument,
	})
}
