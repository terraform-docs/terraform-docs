/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateFunc configures Generator.
type GenerateFunc func(*Generator)

// WithContent specifies how the Generator should add content.
func WithContent(content string) GenerateFunc {
	return func(g *Generator) {
		g.content = content
	}
}

// WithHeader specifies how the Generator should add Header.
func WithHeader(header string) GenerateFunc {
	return func(g *Generator) {
		g.Header = header
	}
}

// WithFooter specifies how the Generator should add Footer.
func WithFooter(footer string) GenerateFunc {
	return func(g *Generator) {
		g.Footer = footer
	}
}

// WithInputs specifies how the Generator should add Inputs.
func WithInputs(inputs string) GenerateFunc {
	return func(g *Generator) {
		g.Inputs = inputs
	}
}

// WithModules specifies how the Generator should add Modules.
func WithModules(modules string) GenerateFunc {
	return func(g *Generator) {
		g.Modules = modules
	}
}

// WithOutputs specifies how the Generator should add Outputs.
func WithOutputs(outputs string) GenerateFunc {
	return func(g *Generator) {
		g.Outputs = outputs
	}
}

// WithProviders specifies how the Generator should add Providers.
func WithProviders(providers string) GenerateFunc {
	return func(g *Generator) {
		g.Providers = providers
	}
}

// WithRequirements specifies how the Generator should add Requirements.
func WithRequirements(requirements string) GenerateFunc {
	return func(g *Generator) {
		g.Requirements = requirements
	}
}

// WithResources specifies how the Generator should add Resources.
func WithResources(resources string) GenerateFunc {
	return func(g *Generator) {
		g.Resources = resources
	}
}

// Generator represents all the sections that can be generated for a Terraform
// modules (e.g. header, footer, inputs, etc). All these sections are being
// generated individually and if no content template was passed they will be
// combined together with a predefined order.
//
// On the other hand these sections can individually be used in content template
// to form a custom format (and order).
//
// Note that the notion of custom content template will be ignored for incompatible
// formatters and custom plugins. Compatible formatters are:
//
// - asciidoc document
// - asciidoc table
// - markdown document
// - markdown table
type Generator struct {
	Header       string
	Footer       string
	Inputs       string
	Modules      string
	Outputs      string
	Providers    string
	Requirements string
	Resources    string

	path      string // module's path
	content   string // all the content combined
	formatter string // name of the formatter
}

// NewGenerator returns a Generator for specific formatter name and with
// provided sets of GeneratorFunc functions to build and add individual
// sections.
func NewGenerator(name string, fns ...GenerateFunc) *Generator {
	g := &Generator{
		formatter: name,
	}

	for _, fn := range fns {
		fn(g)
	}

	return g
}

// Path of module's directory.
func (g *Generator) Path(root string) {
	g.path = root
}

// ExecuteTemplate applies the template with Generator known items. If template
// is empty Generator.content is returned as is. If template is not empty this
// still returns Generator.content for incompatible formatters.
func (g *Generator) ExecuteTemplate(contentTmpl string) (string, error) {
	if !g.isCompatible() {
		return g.content, nil
	}

	if contentTmpl == "" {
		return g.content, nil
	}

	var buf bytes.Buffer

	tmpl := template.New("content")
	tmpl.Funcs(template.FuncMap{
		"include": func(s string) string {
			content, err := os.ReadFile(filepath.Join(g.path, s))
			if err != nil {
				panic(err)
			}
			return string(content)
		},
	})
	template.Must(tmpl.Parse(contentTmpl))

	if err := tmpl.ExecuteTemplate(&buf, "content", g); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *Generator) isCompatible() bool {
	switch g.formatter {
	case "asciidoc document", "asciidoc table", "markdown document", "markdown table":
		return true
	}
	return false
}

// GeneratorCallback renders a Terraform module and creates a GenerateFunc.
type GeneratorCallback func(string) GenerateFunc

// ForEach section executes GeneratorCallback to render the content for that
// section and create corresponding GeneratorFunc. If there is any error in
// the executing the template for the section ForEach function immediately
// returns it and exit.
func ForEach(callback func(string, GeneratorCallback) error) error {
	mappings := map[string]GeneratorCallback{
		"all":          WithContent,
		"header":       WithHeader,
		"footer":       WithFooter,
		"inputs":       WithInputs,
		"modules":      WithModules,
		"outputs":      WithOutputs,
		"providers":    WithProviders,
		"requirements": WithRequirements,
		"resources":    WithResources,
	}
	for name, fn := range mappings {
		if err := callback(name, fn); err != nil {
			return err
		}
	}
	return nil
}
