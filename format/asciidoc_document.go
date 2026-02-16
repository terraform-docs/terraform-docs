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

//go:embed templates/asciidoc_document*.tmpl
var asciidocsDocumentFS embed.FS

// asciidocDocument represents AsciiDoc Document format.
type asciidocDocument struct {
	*generator

	config   *print.Config
	template *template.Template
}

// NewAsciidocDocument returns new instance of Asciidoc Document.
func NewAsciidocDocument(config *print.Config) Type {
	items := readTemplateItems(asciidocsDocumentFS, "asciidoc_document")

	config.Settings.Escape = false

	tt := template.New(config, items...)
	tt.CustomFunc(gotemplate.FuncMap{
		"type": func(t string) string {
			result, extraline := PrintFencedAsciidocCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := PrintFencedAsciidocCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return config.Settings.Required
		},
	})

	return &asciidocDocument{
		generator: newGenerator(config, true),
		config:    config,
		template:  tt,
	}
}

// Generate a Terraform module as AsciiDoc document.
func (d *asciidocDocument) Generate(module *terraform.Module) error {
	err := d.forEach(func(name string) (string, error) {
		rendered, err := d.template.Render(name, module)
		if err != nil {
			return "", err
		}
		return sanitize(rendered), nil
	})

	d.funcs(withModule(module))

	return err
}

func init() {
	register(map[string]initializerFn{
		"asciidoc document": NewAsciidocDocument,
		"asciidoc doc":      NewAsciidocDocument,
		"adoc document":     NewAsciidocDocument,
		"adoc doc":          NewAsciidocDocument,
	})
}
