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
		g.header = header
	}
}

// WithFooter specifies how the Generator should add Footer.
func WithFooter(footer string) GenerateFunc {
	return func(g *Generator) {
		g.footer = footer
	}
}

// WithInputs specifies how the Generator should add Inputs.
func WithInputs(inputs string) GenerateFunc {
	return func(g *Generator) {
		g.inputs = inputs
	}
}

// WithModules specifies how the Generator should add Modules.
func WithModules(modules string) GenerateFunc {
	return func(g *Generator) {
		g.modules = modules
	}
}

// WithOutputs specifies how the Generator should add Outputs.
func WithOutputs(outputs string) GenerateFunc {
	return func(g *Generator) {
		g.outputs = outputs
	}
}

// WithProviders specifies how the Generator should add Providers.
func WithProviders(providers string) GenerateFunc {
	return func(g *Generator) {
		g.providers = providers
	}
}

// WithRequirements specifies how the Generator should add Requirements.
func WithRequirements(requirements string) GenerateFunc {
	return func(g *Generator) {
		g.requirements = requirements
	}
}

// WithResources specifies how the Generator should add Resources.
func WithResources(resources string) GenerateFunc {
	return func(g *Generator) {
		g.resources = resources
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
	// all the content combined
	content string

	// individual sections
	header       string
	footer       string
	inputs       string
	modules      string
	outputs      string
	providers    string
	requirements string
	resources    string

	path      string // module's path
	formatter string // formatter name

	funcs []GenerateFunc
}

// NewGenerator returns a Generator for specific formatter name and with
// provided sets of GeneratorFunc functions to build and add individual
// sections.
func NewGenerator(name string, root string, fns ...GenerateFunc) *Generator {
	g := &Generator{
		path:      root,
		formatter: name,
		funcs:     []GenerateFunc{},
	}

	g.Funcs(fns...)

	return g
}

// Content returns generted all the sections combined based on the underlying format.
func (g *Generator) Content() string { return g.content }

// Header returns generted header section based on the underlying format.
func (g *Generator) Header() string { return g.header }

// Footer returns generted footer section based on the underlying format.
func (g *Generator) Footer() string { return g.footer }

// Inputs returns generted inputs section based on the underlying format.
func (g *Generator) Inputs() string { return g.inputs }

// Modules returns generted modules section based on the underlying format.
func (g *Generator) Modules() string { return g.modules }

// Outputs returns generted outputs section based on the underlying format.
func (g *Generator) Outputs() string { return g.outputs }

// Providers returns generted providers section based on the underlying format.
func (g *Generator) Providers() string { return g.providers }

// Requirements returns generted resources section based on the underlying format.
func (g *Generator) Requirements() string { return g.requirements }

// Resources returns generted requirements section based on the underlying format.
func (g *Generator) Resources() string { return g.resources }

// Funcs adds GenerateFunc to the list of available functions, for further use
// if need be, and then runs them.
func (g *Generator) Funcs(fns ...GenerateFunc) {
	for _, fn := range fns {
		g.funcs = append(g.funcs, fn)
		fn(g)
	}
}

// Path set path of module's root directory.
func (g *Generator) Path(root string) {
	g.path = root
}

// ExecuteTemplate applies the template with Renderer known items. If template
// is empty Renderer.content is returned as is. If template is not empty this
// still returns Renderer.content for incompatible formatters.
// func (g *Renderer) Render(contentTmpl string) (string, error) {
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

// generatorCallback renders a Terraform module and creates a GenerateFunc.
type generatorCallback func(string) GenerateFunc

// ForEach section executes generatorCallback to render the content for that
// section and create corresponding GeneratorFunc. If there is any error in
// the executing the template for the section ForEach function immediately
// returns it and exit.
func (g *Generator) ForEach(render func(string) (string, error)) error {
	mappings := map[string]generatorCallback{
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
	for name, callback := range mappings {
		result, err := render(name)
		if err != nil {
			return err
		}
		fn := callback(result)
		g.funcs = append(g.funcs, fn)
		fn(g)
	}
	return nil
}
