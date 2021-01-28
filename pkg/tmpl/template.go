/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package tmpl

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/pkg/print"
)

// Item represents a named templated which can reference
// other named templated too
type Item struct {
	Name string
	Text string
}

// Template represents a new Template with given name and content
// to be rendered with provided settings with use of built-in and
// custom functions
type Template struct {
	Items []*Item

	settings *print.Settings
	funcMap  template.FuncMap
}

// NewTemplate returns new instance of Template
func NewTemplate(items ...*Item) *Template {
	settings := print.NewSettings()
	return &Template{
		Items:    items,
		settings: settings,
		funcMap:  builtinFuncs(settings),
	}
}

// CustomFunc adds new custom functions to the template
// if functions with the same names didn't exist
func (t *Template) CustomFunc(funcs template.FuncMap) {
	for name, fn := range funcs {
		if t.funcMap[name] == nil {
			t.funcMap[name] = fn
		}
	}
}

// Settings adds current user-selected settings to the Template
func (t *Template) Settings(settings *print.Settings) {
	t.settings = settings
	if settings != nil {
		t.funcMap = builtinFuncs(settings)
	}
}

// Render renders the Template with given Module struct
func (t *Template) Render(module *terraform.Module) (string, error) {
	if len(t.Items) < 1 {
		return "", fmt.Errorf("base template not found")
	}
	var buffer bytes.Buffer
	tmpl := template.New(t.Items[0].Name)
	tmpl.Funcs(t.funcMap)
	template.Must(tmpl.Parse(normalize(t.Items[0].Text)))
	for i, item := range t.Items {
		if i == 0 {
			continue
		}
		tt := tmpl.New(item.Name)
		tt.Funcs(t.funcMap)
		template.Must(tt.Parse(normalize(item.Text)))
	}
	err := tmpl.ExecuteTemplate(&buffer, t.Items[0].Name, struct {
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

func builtinFuncs(settings *print.Settings) template.FuncMap {
	return template.FuncMap{
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
		"tostring": func(s types.String) string {
			return string(s)
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
		"sanitizeHeader": func(s string) string {
			settings.EscapePipe = false
			s = sanitizeItemForDocument(s, settings)
			settings.EscapePipe = true
			return s
		},
		"sanitizeDoc": func(s string) string {
			return sanitizeItemForDocument(s, settings)
		},
		"sanitizeTbl": func(s string) string {
			settings.EscapePipe = true
			s = sanitizeItemForTable(s, settings)
			settings.EscapePipe = false
			return s
		},
		"sanitizeAsciidocTbl": func(s string) string {
			settings.EscapePipe = true
			s = sanitizeItemForAsciidocTable(s, settings)
			settings.EscapePipe = false
			return s
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
