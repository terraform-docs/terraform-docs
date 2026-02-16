/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"bytes"
	"fmt"
	"strings"
	gotemplate "text/template"

	sprig "github.com/Masterminds/sprig/v3"

	"github.com/rquadling/terraform-docs/internal/types"
	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// Item represents a named templated which can reference other named templated too.
type Item struct {
	Name      string
	Text      string
	TrimSpace bool
}

// Template represents a new Template with given name and content to be rendered
// with provided settings with use of built-in and custom functions.
type Template struct {
	items  []*Item
	config *print.Config

	funcMap    gotemplate.FuncMap
	customFunc gotemplate.FuncMap
}

// New returns new instance of Template.
func New(config *print.Config, items ...*Item) *Template {
	return &Template{
		items:      items,
		config:     config,
		funcMap:    builtinFuncs(config),
		customFunc: make(gotemplate.FuncMap),
	}
}

// Funcs return available template out of the box and custom functions.
func (t *Template) Funcs() gotemplate.FuncMap {
	return t.funcMap
}

// CustomFunc adds new custom functions to the template if functions with the same
// names didn't exist.
func (t *Template) CustomFunc(funcs gotemplate.FuncMap) {
	for name, fn := range funcs {
		if _, found := t.customFunc[name]; !found {
			t.customFunc[name] = fn
		}
	}
	t.applyCustomFunc()
}

// applyCustomFunc is re-adding the custom functions to list of available functions.
func (t *Template) applyCustomFunc() {
	for name, fn := range t.customFunc {
		if _, found := t.funcMap[name]; !found {
			t.funcMap[name] = fn
		}
	}
}

// Render template with given Module struct.
func (t *Template) Render(name string, module *terraform.Module) (string, error) {
	data := struct {
		Config *print.Config
		Module *terraform.Module
	}{
		Config: t.config,
		Module: module,
	}
	return t.RenderContent(name, data)
}

// RenderContent template with given data. It can contain anything but most
// probably it will only contain terraform.Module and print.generator.
func (t *Template) RenderContent(name string, data interface{}) (string, error) {
	if len(t.items) < 1 {
		return "", fmt.Errorf("base template not found")
	}

	item := t.findByName(name)
	if item == nil {
		return "", fmt.Errorf("%s template not found", name)
	}

	var buffer bytes.Buffer

	tmpl := gotemplate.New(item.Name)
	tmpl.Funcs(t.funcMap)
	gotemplate.Must(tmpl.Parse(normalize(item.Text, item.TrimSpace)))

	for _, ii := range t.items {
		tt := tmpl.New(ii.Name)
		tt.Funcs(t.funcMap)
		gotemplate.Must(tt.Parse(normalize(ii.Text, ii.TrimSpace)))
	}

	if err := tmpl.ExecuteTemplate(&buffer, item.Name, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (t *Template) findByName(name string) *Item {
	if name == "" {
		if len(t.items) > 0 {
			return t.items[0]
		}
		return nil
	}
	for _, i := range t.items {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func builtinFuncs(config *print.Config) gotemplate.FuncMap { // nolint:gocyclo
	fns := gotemplate.FuncMap{
		"default": func(_default string, value string) string {
			if value != "" {
				return value
			}
			return _default
		},
		"indent": func(extra int, char string) string {
			return GenerateIndentation(config.Settings.Indent, extra, char)
		},
		"name": func(name string) string {
			return SanitizeName(name, config.Settings.Escape)
		},
		"ternary": func(condition interface{}, trueValue string, falseValue string) string {
			var c bool
			switch x := fmt.Sprintf("%T", condition); x {
			case "string":
				c = condition.(string) != ""
			case "int":
				c = condition.(int) != 0
			case "bool":
				c = condition.(bool)
			}
			if c {
				return trueValue
			}
			return falseValue
		},
		"tostring": func(s types.String) string {
			return string(s)
		},

		// trim
		"trim": func(cut string, s string) string {
			if s != "" {
				return strings.Trim(s, cut)
			}
			return s
		},
		"trimLeft": func(cut string, s string) string {
			if s != "" {
				return strings.TrimLeft(s, cut)
			}
			return s
		},
		"trimRight": func(cut string, s string) string {
			if s != "" {
				return strings.TrimRight(s, cut)
			}
			return s
		},
		"trimPrefix": func(prefix string, s string) string {
			if s != "" {
				return strings.TrimPrefix(s, prefix)
			}
			return s
		},
		"trimSuffix": func(suffix string, s string) string {
			if s != "" {
				return strings.TrimSuffix(s, suffix)
			}
			return s
		},

		// sanitize
		"sanitizeSection": func(s string) string {
			return SanitizeSection(s, config.Settings.Escape, config.Settings.HTML)
		},
		"sanitizeDoc": func(s string) string {
			return SanitizeDocument(s, config.Settings.Escape, config.Settings.HTML)
		},
		"sanitizeMarkdownTbl": func(s string) string {
			return SanitizeMarkdownTable(s, config.Settings.Escape, config.Settings.HTML)
		},
		"sanitizeAsciidocTbl": func(s string) string {
			return SanitizeAsciidocTable(s, config.Settings.Escape, config.Settings.HTML)
		},

		// anchors
		"anchorNameMarkdown": func(prefix string, value string) string {
			return CreateAnchorMarkdown(prefix, value, config.Settings.Anchor, config.Settings.Escape)
		},
		"anchorNameAsciidoc": func(prefix string, value string) string {
			return CreateAnchorAsciidoc(prefix, value, config.Settings.Anchor, config.Settings.Escape)
		},
	}

	for name, fn := range sprig.FuncMap() {
		if _, found := fns[name]; !found {
			fns[name] = fn
		}
	}

	return fns
}

// normalize the template and remove any space from all the lines. This makes
// it possible to have a indented, human-readable template which doesn't affect
// the rendering of them.
func normalize(s string, trimSpace bool) string {
	if !trimSpace {
		return s
	}
	splitted := strings.Split(s, "\n")
	for i, v := range splitted {
		splitted[i] = strings.TrimSpace(v)
	}
	return strings.Join(splitted, "\n")
}

// GenerateIndentation generates indentation of Markdown and AsciiDoc headers
// with base level of provided 'settings.IndentLevel' plus any extra level needed
// for subsection (e.g. 'Required Inputs' which is a subsection of 'Inputs' section)
func GenerateIndentation(base int, extra int, char string) string {
	if char == "" {
		return ""
	}
	if base < 1 || base > 5 {
		base = 2
	}
	var indent string
	for i := 0; i < base+extra; i++ {
		indent += char
	}
	return indent
}
