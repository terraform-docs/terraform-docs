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
	items    []*Item
	settings *print.Settings

	funcMap    gotemplate.FuncMap
	customFunc gotemplate.FuncMap
}

// New returns new instance of Template.
func New(settings *print.Settings, items ...*Item) *Template {
	t := &Template{
		items:      items,
		settings:   settings,
		funcMap:    builtinFuncs(settings),
		customFunc: make(gotemplate.FuncMap),
	}
	t.CustomFunc(gotemplate.FuncMap{
		"tostring": func(s types.String) string {
			return string(s)
		},
		"sanitizeHeader": func(s string) string {
			copy := *settings
			copy.EscapePipe = false
			s = sanitizeItemForDocument(s, &copy)
			return s
		},
		"sanitizeDoc": func(s string) string {
			return sanitizeItemForDocument(s, settings)
		},
		"sanitizeTbl": func(s string) string {
			copy := *settings
			copy.EscapePipe = true
			s = sanitizeItemForTable(s, &copy)
			return s
		},
		"sanitizeAsciidocTbl": func(s string) string {
			copy := *settings
			copy.EscapePipe = true
			s = sanitizeItemForAsciidocTable(s, &copy)
			return s
		},
	})

	return t
}

// Funcs return available template out of the box and custom functions.
func (t Template) Funcs() gotemplate.FuncMap {
	return t.funcMap
}

// CustomFunc adds new custom functions to the template
// if functions with the same names didn't exist.
func (t Template) CustomFunc(funcs gotemplate.FuncMap) {
	for name, fn := range funcs {
		if _, found := t.customFunc[name]; !found {
			t.customFunc[name] = fn
		}
	}
	t.applyCustomFunc()
}

// applyCustomFunc is re-adding the custom functions to list
// of avialable functions.
func (t *Template) applyCustomFunc() {
	for name, fn := range t.customFunc {
		if _, found := t.funcMap[name]; !found {
			t.funcMap[name] = fn
		}
	}
}

// Render template with given Module struct.
func (t Template) Render(module *terraform.Module) (string, error) {
	if len(t.items) < 1 {
		return "", fmt.Errorf("base template not found")
	}
	var buffer bytes.Buffer
	tmpl := gotemplate.New(t.items[0].Name)
	tmpl.Funcs(t.funcMap)
	gotemplate.Must(tmpl.Parse(normalize(t.items[0].Text)))
	for i, item := range t.items {
		if i == 0 {
			continue
		}
		tt := tmpl.New(item.Name)
		tt.Funcs(t.funcMap)
		gotemplate.Must(tt.Parse(normalize(item.Text)))
	}
	err := tmpl.ExecuteTemplate(&buffer, t.items[0].Name, struct {
		Module   *terraform.Module
		Settings *print.Settings
	}{
		Module:   module,
		Settings: t.settings,
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
func builtinFuncs(settings *print.Settings) gotemplate.FuncMap {
	return gotemplate.FuncMap{
		"default": func(d string, s string) string {
			if s != "" {
				return s
			}
			return d
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
		"indent": func(l int, char string) string {
			return generateIndentation(l, char, settings)
		},
		"name": func(n string) string {
			return sanitizeName(n, settings)
		},
	}
}

// Normalizes the template and remove any space from all the lines.
// This makes it possible to have a indented, human-readable template
// which doesn't affect the rendering of them.
func normalize(s string) string {
	segments := strings.Split(s, "\n")
	buffer := bytes.NewBufferString("")
	for _, segment := range segments {
		buffer.WriteString(strings.TrimSpace(segment))
		buffer.WriteString("\n")
	}
	return buffer.String()
}
