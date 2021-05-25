/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	gotemplate "text/template"

	templatesdk "github.com/terraform-docs/plugin-sdk/template"
	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Item represents a named templated which can reference
// other named templated too.
type Item struct {
	Name string
	Text string
}

// Template represents a new Template with given name and content
// to be rendered with provided settings with use of built-in and
// custom functions.
type Template struct {
	engine   *templatesdk.Template
	settings *print.Settings
}

// New returns new instance of Template.
func New(settings *print.Settings, items ...*Item) *Template {
	ii := []*templatesdk.Item{}
	for _, v := range items {
		ii = append(ii, &templatesdk.Item{Name: v.Name, Text: v.Text})
	}

	engine := templatesdk.New(settings.Convert(), ii...)
	engine.CustomFunc(gotemplate.FuncMap{
		"tostring": func(s types.String) string {
			return string(s)
		},
		"sanitizeSection": func(s string) string {
			return sanitizeSection(s, settings)
		},
		"sanitizeDoc": func(s string) string {
			return sanitizeDocument(s, settings)
		},
		"sanitizeMarkdownTbl": func(s string) string {
			return sanitizeMarkdownTable(s, settings)
		},
		"sanitizeAsciidocTbl": func(s string) string {
			return sanitizeAsciidocTable(s, settings)
		},
		"anchorNameMarkdown": func(s string, t string) string {
			return createAnchorMarkdown(s, t, settings)
		},
		"anchorNameAsciidoc": func(s string, t string) string {
			return createAnchorAsciidoc(s, t, settings)
		},
	})

	return &Template{
		engine:   engine,
		settings: settings,
	}
}

// Funcs return available template out of the box and custom functions.
func (t Template) Funcs() gotemplate.FuncMap {
	return t.engine.Funcs()
}

// CustomFunc adds new custom functions to the template
// if functions with the same names didn't exist.
func (t Template) CustomFunc(funcs gotemplate.FuncMap) {
	t.engine.CustomFunc(funcs)
}

// Render template with given Module struct.
func (t Template) Render(name string, module *terraform.Module) (string, error) {
	return t.engine.Render(name, module)
}
