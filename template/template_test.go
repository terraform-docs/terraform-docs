/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	gotemplate "text/template"

	"github.com/stretchr/testify/assert"

	"github.com/rquadling/terraform-docs/internal/types"
	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

func TestTemplateRender(t *testing.T) {
	sectionTpl := `
	{{- with .Module.Header -}}
		{{ custom . }}
	{{- end -}}
	`
	customFuncs := gotemplate.FuncMap{
		"custom": func(s string) string {
			return fmt.Sprintf("customized <<%s>>", s)
		},
	}
	module := &terraform.Module{
		Header: "sample header",
	}
	tests := []struct {
		name     string
		items    []*Item
		expected string
		wantErr  bool
	}{
		{
			name: "template render with custom functions",
			items: []*Item{
				{
					Name: "all",
					Text: `{{- template "section" . -}}`,
				}, {
					Name: "section",
					Text: sectionTpl,
				},
			},
			expected: "customized <<sample header>>",
			wantErr:  false,
		},
		{
			name:     "template render with custom functions",
			items:    []*Item{},
			expected: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			tpl := New(print.DefaultConfig(), tt.items...)
			tpl.CustomFunc(customFuncs)
			rendered, err := tpl.Render("", module)
			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, rendered)
			}
		})
	}
}

func TestBuiltinFunc(t *testing.T) {
	tests := []struct {
		name     string
		funcName string
		funcArgs []string
		escape   bool
		expected string
	}{
		// default
		{
			name:     "template builtin functions default",
			funcName: "default",
			funcArgs: []string{`"a"`, `"b"`},
			escape:   true,
			expected: "b",
		},
		{
			name:     "template builtin functions default",
			funcName: "default",
			funcArgs: []string{`"a"`, `""`},
			escape:   true,
			expected: "a",
		},
		{
			name:     "template builtin functions default",
			funcName: "default",
			funcArgs: []string{`""`, `"b"`},
			escape:   true,
			expected: "b",
		},
		{
			name:     "template builtin functions default",
			funcName: "default",
			funcArgs: []string{`""`, `""`},
			escape:   true,
			expected: "",
		},

		// tostring
		{
			name:     "template builtin functions tostring",
			funcName: "tostring",
			funcArgs: []string{`"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions tostring",
			funcName: "tostring",
			funcArgs: []string{`""`},
			escape:   true,
			expected: "",
		},

		// trim
		{
			name:     "template builtin functions trim",
			funcName: "trim",
			funcArgs: []string{`" "`, `"   foo   "`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trim",
			funcName: "trim",
			funcArgs: []string{`" "`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trim",
			funcName: "trim",
			funcArgs: []string{`""`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trim",
			funcName: "trim",
			funcArgs: []string{`" "`, `""`},
			escape:   true,
			expected: "",
		},

		// trimLeft
		{
			name:     "template builtin functions trimLeft",
			funcName: "trimLeft",
			funcArgs: []string{`" "`, `"   foo   "`},
			escape:   true,
			expected: "foo   ",
		},
		{
			name:     "template builtin functions trimLeft",
			funcName: "trimLeft",
			funcArgs: []string{`" "`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimLeft",
			funcName: "trimLeft",
			funcArgs: []string{`""`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimLeft",
			funcName: "trimLeft",
			funcArgs: []string{`" "`, `""`},
			escape:   true,
			expected: "",
		},

		// trimRight
		{
			name:     "template builtin functions trimRight",
			funcName: "trimRight",
			funcArgs: []string{`" "`, `"   foo   "`},
			escape:   true,
			expected: "   foo",
		},
		{
			name:     "template builtin functions trimRight",
			funcName: "trimRight",
			funcArgs: []string{`" "`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimRight",
			funcName: "trimRight",
			funcArgs: []string{`""`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimRight",
			funcName: "trimRight",
			funcArgs: []string{`" "`, `""`},
			escape:   true,
			expected: "",
		},

		// trimPrefix
		{
			name:     "template builtin functions trimPrefix",
			funcName: "trimPrefix",
			funcArgs: []string{`" "`, `"   foo   "`},
			escape:   true,
			expected: "  foo   ",
		},
		{
			name:     "template builtin functions trimPrefix",
			funcName: "trimPrefix",
			funcArgs: []string{`" "`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimPrefix",
			funcName: "trimPrefix",
			funcArgs: []string{`""`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimPrefix",
			funcName: "trimPrefix",
			funcArgs: []string{`" "`, `""`},
			escape:   true,
			expected: "",
		},

		// trimSuffix
		{
			name:     "template builtin functions trimSuffix",
			funcName: "trimSuffix",
			funcArgs: []string{`" "`, `"   foo   "`},
			escape:   true,
			expected: "   foo  ",
		},
		{
			name:     "template builtin functions trimSuffix",
			funcName: "trimSuffix",
			funcArgs: []string{`" "`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimSuffix",
			funcName: "trimSuffix",
			funcArgs: []string{`""`, `"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions trimSuffix",
			funcName: "trimSuffix",
			funcArgs: []string{`" "`, `""`},
			escape:   true,
			expected: "",
		},

		// indent
		{
			name:     "template builtin functions indent",
			funcName: "indent",
			funcArgs: []string{`0`, `"#"`},
			escape:   true,
			expected: "##",
		},
		{
			name:     "template builtin functions indent",
			funcName: "indent",
			funcArgs: []string{`1`, `"#"`},
			escape:   true,
			expected: "###",
		},
		{
			name:     "template builtin functions indent",
			funcName: "indent",
			funcArgs: []string{`2`, `"#"`},
			escape:   true,
			expected: "####",
		},
		{
			name:     "template builtin functions indent",
			funcName: "indent",
			funcArgs: []string{`3`, `"#"`},
			escape:   true,
			expected: "#####",
		},

		// name
		{
			name:     "template builtin functions name",
			funcName: "name",
			funcArgs: []string{`"foo"`},
			escape:   true,
			expected: "foo",
		},
		{
			name:     "template builtin functions name",
			funcName: "name",
			funcArgs: []string{`"foo_bar"`},
			escape:   true,
			expected: "foo\\_bar",
		},
		{
			name:     "template builtin functions name",
			funcName: "name",
			funcArgs: []string{`"foo_bar"`},
			escape:   false,
			expected: "foo_bar",
		},
		{
			name:     "template builtin functions name",
			funcName: "name",
			funcArgs: []string{`""`},
			escape:   true,
			expected: "",
		},

		// sanitizeSection
		{
			name:     "template builtin functions sanitizeSection",
			funcName: "sanitizeSection",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\n|-----|-----|\n| foo | bar |\""},
			escape:   true,
			expected: "Example of 'foo\\_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\n|-----|-----|\n| foo | bar |",
		},
		{
			name:     "template builtin functions sanitizeSection",
			funcName: "sanitizeSection",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escape:   false,
			expected: "Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:     "template builtin functions sanitizeSection",
			funcName: "sanitizeSection",
			funcArgs: []string{`""`},
			escape:   true,
			expected: "n/a",
		},

		// sanitizeDoc
		{
			name:     "template builtin functions sanitizeDoc",
			funcName: "sanitizeDoc",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escape:   true,
			expected: "Example of 'foo\\_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:     "template builtin functions sanitizeDoc",
			funcName: "sanitizeDoc",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escape:   false,
			expected: "Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:     "template builtin functions sanitizeDoc",
			funcName: "sanitizeDoc",
			funcArgs: []string{`""`},
			escape:   true,
			expected: "n/a",
		},

		// sanitizeMarkdownTbl
		{
			name:     "template builtin functions sanitizeMarkdownTbl",
			funcName: "sanitizeMarkdownTbl",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escape:   true,
			expected: "Example of 'foo\\_bar' module in `foo_bar.tf`.<br/><br/>\\| Foo \\| Bar \\|",
		},
		{
			name:     "template builtin functions sanitizeMarkdownTbl",
			funcName: "sanitizeMarkdownTbl",
			funcArgs: []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escape:   false,
			expected: "Example of 'foo_bar' module in `foo_bar.tf`.<br/><br/>\\| Foo \\| Bar \\|",
		},
		{
			name:     "template builtin functions sanitizeMarkdownTbl",
			funcName: "sanitizeMarkdownTbl",
			funcArgs: []string{`""`},
			escape:   true,
			expected: "n/a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			config := print.DefaultConfig()
			config.Settings.Escape = tt.escape
			funcs := builtinFuncs(config)

			fn, ok := funcs[tt.funcName]
			assert.Truef(ok, "function %s is not defined", tt.funcName)

			v := reflect.ValueOf(fn)
			tp := v.Type()
			assert.Equalf(len(tt.funcArgs), tp.NumIn(), "invalid number of arguments. got: %v, want: %v", len(tt.funcArgs), tp.NumIn())

			argv := make([]reflect.Value, len(tt.funcArgs))

			for i := range argv {
				var argType reflect.Kind
				if strings.HasPrefix(tt.funcArgs[i], "\"") {
					if tt.funcName == "tostring" {
						argType = reflect.TypeOf(types.String("")).Kind()
						argv[i] = reflect.ValueOf(types.String(strings.Trim(tt.funcArgs[i], "\"")))
					} else {
						argType = reflect.String
						argv[i] = reflect.ValueOf(strings.Trim(tt.funcArgs[i], "\""))
					}
				} else {
					argType = reflect.Int
					num, _ := strconv.Atoi(tt.funcArgs[i])
					argv[i] = reflect.ValueOf(num)
				}
				if tp.In(i).Kind() != argType {
					assert.Fail("Invalid argument. got: %v, want: %v", argType, tp.In(i).Kind())
				}
			}

			result := v.Call(argv)

			if len(result) != 1 || result[0].Kind() != reflect.String {
				assert.Fail("function %s must return a one string value", tt.funcName)
			}

			assert.Equal(tt.expected, result[0].String())
		})
	}
}

func TestGenerateIndentation(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		extra    int
		expected string
	}{
		{
			name:     "generate indentation",
			base:     2,
			extra:    1,
			expected: "###",
		},
		{
			name:     "generate indentation",
			extra:    2,
			expected: "####",
		},
		{
			name:     "generate indentation",
			base:     4,
			extra:    3,
			expected: "#######",
		},
		{
			name:     "generate indentation",
			base:     0,
			extra:    0,
			expected: "##",
		},
		{
			name:     "generate indentation",
			base:     6,
			extra:    1,
			expected: "###",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := GenerateIndentation(tt.base, tt.extra, "#")

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		trim     bool
		expected string
	}{
		{
			name:     "normalize with trim space",
			text:     "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit",
			trim:     true,
			expected: "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit",
		},
		{
			name:     "normalize with trim space",
			text:     "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit\n",
			trim:     true,
			expected: "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit\n",
		},
		{
			name:     "normalize with trim space",
			text:     "Lorem ipsum\ndolor sit amet,\n  consectetur\nadipiscing\nelit",
			trim:     true,
			expected: "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit",
		},
		{
			name:     "normalize without trim space",
			text:     "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit",
			trim:     false,
			expected: "Lorem ipsum\ndolor sit amet,\nconsectetur\nadipiscing\nelit",
		},
		{
			name:     "normalize without trim space",
			text:     "Lorem ipsum\ndolor sit amet,\n  consectetur\nadipiscing\nelit",
			trim:     false,
			expected: "Lorem ipsum\ndolor sit amet,\n  consectetur\nadipiscing\nelit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := normalize(tt.text, tt.trim)

			assert.Equal(tt.expected, actual)
		})
	}
}
