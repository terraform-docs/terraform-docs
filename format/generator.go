/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"os"
	"path/filepath"
	"strings"
	gotemplate "text/template"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/template"
	"github.com/rquadling/terraform-docs/terraform"
)

// generateFunc configures generator.
type generateFunc func(*generator)

// withContent specifies how the generator should add content.
func withContent(content string) generateFunc {
	return func(g *generator) {
		g.content = content
	}
}

// withHeader specifies how the generator should add Header.
func withHeader(header string) generateFunc {
	return func(g *generator) {
		g.header = header
	}
}

// withFooter specifies how the generator should add Footer.
func withFooter(footer string) generateFunc {
	return func(g *generator) {
		g.footer = footer
	}
}

// withInputs specifies how the generator should add Inputs.
func withInputs(inputs string) generateFunc {
	return func(g *generator) {
		g.inputs = inputs
	}
}

// withModules specifies how the generator should add Modules.
func withModules(modules string) generateFunc {
	return func(g *generator) {
		g.modules = modules
	}
}

// withOutputs specifies how the generator should add Outputs.
func withOutputs(outputs string) generateFunc {
	return func(g *generator) {
		g.outputs = outputs
	}
}

// withProviders specifies how the generator should add Providers.
func withProviders(providers string) generateFunc {
	return func(g *generator) {
		g.providers = providers
	}
}

// withRequirements specifies how the generator should add Requirements.
func withRequirements(requirements string) generateFunc {
	return func(g *generator) {
		g.requirements = requirements
	}
}

// withResources specifies how the generator should add Resources.
func withResources(resources string) generateFunc {
	return func(g *generator) {
		g.resources = resources
	}
}

// withModule specifies how the generator should add Resources.
func withModule(module *terraform.Module) generateFunc {
	return func(g *generator) {
		g.module = module
	}
}

// generator represents all the sections that can be generated for a Terraform
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
type generator struct {
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

	config *print.Config
	module *terraform.Module

	path string         // module's path
	fns  []generateFunc // generator helper functions

	canRender bool // indicates if the generator can render with custom template
}

// newGenerator returns a generator for specific formatter name and with
// provided sets of GeneratorFunc functions to build and add individual
// sections.
//
//nolint:unparam
func newGenerator(config *print.Config, canRender bool, fns ...generateFunc) *generator {
	g := &generator{
		config: config,

		path: config.ModuleRoot,
		fns:  []generateFunc{},

		canRender: canRender,
	}

	g.funcs(fns...)

	return g
}

// Content returns generted all the sections combined based on the underlying format.
func (g *generator) Content() string { return g.content }

// Header returns generted header section based on the underlying format.
func (g *generator) Header() string { return g.header }

// Footer returns generted footer section based on the underlying format.
func (g *generator) Footer() string { return g.footer }

// Inputs returns generted inputs section based on the underlying format.
func (g *generator) Inputs() string { return g.inputs }

// Modules returns generted modules section based on the underlying format.
func (g *generator) Modules() string { return g.modules }

// Outputs returns generted outputs section based on the underlying format.
func (g *generator) Outputs() string { return g.outputs }

// Providers returns generted providers section based on the underlying format.
func (g *generator) Providers() string { return g.providers }

// Requirements returns generted resources section based on the underlying format.
func (g *generator) Requirements() string { return g.requirements }

// Resources returns generted requirements section based on the underlying format.
func (g *generator) Resources() string { return g.resources }

// Module returns generted requirements section based on the underlying format.
func (g *generator) Module() *terraform.Module { return g.module }

// funcs adds GenerateFunc to the list of available functions, for further use
// if need be, and then runs them.
func (g *generator) funcs(fns ...generateFunc) {
	for _, fn := range fns {
		g.fns = append(g.fns, fn)
		fn(g)
	}
}

// Path set path of module's root directory.
func (g *generator) Path(root string) {
	g.path = root
}

func (g *generator) Render(tpl string) (string, error) {
	if !g.canRender {
		return g.content, nil
	}

	if tpl == "" {
		return g.content, nil
	}

	tt := template.New(g.config, &template.Item{
		Name: "content",
		Text: tpl,
	})
	tt.CustomFunc(gotemplate.FuncMap{
		"include": func(s string) string {
			content, err := os.ReadFile(filepath.Join(g.path, filepath.Clean(s)))
			if err != nil {
				panic(err)
			}
			return strings.TrimSuffix(string(content), "\n")
		},
	})

	data := struct {
		*generator
		Config *print.Config
		Module *terraform.Module
	}{
		generator: g,
		Config:    g.config,
		Module:    g.module,
	}

	rendered, err := tt.RenderContent("content", data)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(rendered, "\n"), nil
}

// generatorCallback renders a Terraform module and creates a GenerateFunc.
type generatorCallback func(string) generateFunc

// forEach section executes generatorCallback to render the content for that
// section and create corresponding GeneratorFunc. If there is any error in
// executing the template for the section forEach function immediately returns
// it and exits.
func (g *generator) forEach(render func(string) (string, error)) error {
	mappings := map[string]generatorCallback{
		"all":          withContent,
		"header":       withHeader,
		"footer":       withFooter,
		"inputs":       withInputs,
		"modules":      withModules,
		"outputs":      withOutputs,
		"providers":    withProviders,
		"requirements": withRequirements,
		"resources":    withResources,
	}
	for name, callback := range mappings {
		result, err := render(name)
		if err != nil {
			return err
		}
		fn := callback(result)
		g.fns = append(g.fns, fn)
		fn(g)
	}
	return nil
}
