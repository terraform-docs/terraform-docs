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

//go:embed templates/asciidoc_document*.tmpl
var asciidocsDocumentFS embed.FS

// asciidocDocument represents AsciiDoc Document format.
type asciidocDocument struct {
	*print.Generator

	config   *print.Config
	template *template.Template
	settings *print.Settings
}

// NewAsciidocDocument returns new instance of Asciidoc Document.
func NewAsciidocDocument(config *print.Config) Type {
	settings, _ := config.Extract()
	items := readTemplateItems(asciidocsDocumentFS, "asciidoc_document")

	settings.EscapeCharacters = false

	tt := template.New(settings, items...)
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
			return settings.ShowRequired
		},
	})

	return &asciidocDocument{
		Generator: print.NewGenerator("json", config.ModuleRoot),
		config:    config,
		template:  tt,
		settings:  settings,
	}
}

// Generate a Terraform module as AsciiDoc document.
func (d *asciidocDocument) Generate(module *terraform.Module) error {
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
		"asciidoc document": NewAsciidocDocument,
		"asciidoc doc":      NewAsciidocDocument,
		"adoc document":     NewAsciidocDocument,
		"adoc doc":          NewAsciidocDocument,
	})
}
