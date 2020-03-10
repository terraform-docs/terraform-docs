package tmpl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"text/template"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestTemplateRender(t *testing.T) {
	sectionTpl := `
	{{- with .Module.Header -}}
		{{ custom . }}
	{{- end -}}
	`
	customFuncs := template.FuncMap{
		"custom": func(s string) string {
			return fmt.Sprintf("customized <<%s>>", s)
		},
	}
	module := &tfconf.Module{
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
				&Item{
					Name: "all",
					Text: `{{- template "section" . -}}`,
				}, &Item{
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
			tpl := NewTemplate(tt.items...)
			tpl.Settings(print.NewSettings())
			tpl.CustomFunc(customFuncs)
			rendered, err := tpl.Render(module)
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
		name       string
		funcName   string
		funcArgs   []string
		escapeChar bool
		escapePipe bool
		expected   string
	}{
		// default
		{
			name:       "template builtin functions default",
			funcName:   "default",
			funcArgs:   []string{`"a"`, `"b"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "b",
		},
		{
			name:       "template builtin functions default",
			funcName:   "default",
			funcArgs:   []string{`"a"`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "a",
		},
		{
			name:       "template builtin functions default",
			funcName:   "default",
			funcArgs:   []string{`""`, `"b"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "b",
		},
		{
			name:       "template builtin functions default",
			funcName:   "default",
			funcArgs:   []string{`""`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// tostring
		{
			name:       "template builtin functions tostring",
			funcName:   "tostring",
			funcArgs:   []string{`"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions tostring",
			funcName:   "tostring",
			funcArgs:   []string{`""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// trim
		{
			name:       "template builtin functions trim",
			funcName:   "trim",
			funcArgs:   []string{`" "`, `"   foo   "`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trim",
			funcName:   "trim",
			funcArgs:   []string{`" "`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trim",
			funcName:   "trim",
			funcArgs:   []string{`""`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trim",
			funcName:   "trim",
			funcArgs:   []string{`" "`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// trimLeft
		{
			name:       "template builtin functions trimLeft",
			funcName:   "trimLeft",
			funcArgs:   []string{`" "`, `"   foo   "`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo   ",
		},
		{
			name:       "template builtin functions trimLeft",
			funcName:   "trimLeft",
			funcArgs:   []string{`" "`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimLeft",
			funcName:   "trimLeft",
			funcArgs:   []string{`""`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimLeft",
			funcName:   "trimLeft",
			funcArgs:   []string{`" "`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// trimRight
		{
			name:       "template builtin functions trimRight",
			funcName:   "trimRight",
			funcArgs:   []string{`" "`, `"   foo   "`},
			escapeChar: true,
			escapePipe: true,
			expected:   "   foo",
		},
		{
			name:       "template builtin functions trimRight",
			funcName:   "trimRight",
			funcArgs:   []string{`" "`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimRight",
			funcName:   "trimRight",
			funcArgs:   []string{`""`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimRight",
			funcName:   "trimRight",
			funcArgs:   []string{`" "`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// trimPrefix
		{
			name:       "template builtin functions trimPrefix",
			funcName:   "trimPrefix",
			funcArgs:   []string{`" "`, `"   foo   "`},
			escapeChar: true,
			escapePipe: true,
			expected:   "  foo   ",
		},
		{
			name:       "template builtin functions trimPrefix",
			funcName:   "trimPrefix",
			funcArgs:   []string{`" "`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimPrefix",
			funcName:   "trimPrefix",
			funcArgs:   []string{`""`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimPrefix",
			funcName:   "trimPrefix",
			funcArgs:   []string{`" "`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// trimSuffix
		{
			name:       "template builtin functions trimSuffix",
			funcName:   "trimSuffix",
			funcArgs:   []string{`" "`, `"   foo   "`},
			escapeChar: true,
			escapePipe: true,
			expected:   "   foo  ",
		},
		{
			name:       "template builtin functions trimSuffix",
			funcName:   "trimSuffix",
			funcArgs:   []string{`" "`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimSuffix",
			funcName:   "trimSuffix",
			funcArgs:   []string{`""`, `"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions trimSuffix",
			funcName:   "trimSuffix",
			funcArgs:   []string{`" "`, `""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// indent
		{
			name:       "template builtin functions indent",
			funcName:   "indent",
			funcArgs:   []string{`0`},
			escapeChar: true,
			escapePipe: true,
			expected:   "##",
		},
		{
			name:       "template builtin functions indent",
			funcName:   "indent",
			funcArgs:   []string{`1`},
			escapeChar: true,
			escapePipe: true,
			expected:   "###",
		},
		{
			name:       "template builtin functions indent",
			funcName:   "indent",
			funcArgs:   []string{`2`},
			escapeChar: true,
			escapePipe: true,
			expected:   "####",
		},
		{
			name:       "template builtin functions indent",
			funcName:   "indent",
			funcArgs:   []string{`3`},
			escapeChar: true,
			escapePipe: true,
			expected:   "#####",
		},

		// name
		{
			name:       "template builtin functions name",
			funcName:   "name",
			funcArgs:   []string{`"foo"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo",
		},
		{
			name:       "template builtin functions name",
			funcName:   "name",
			funcArgs:   []string{`"foo_bar"`},
			escapeChar: true,
			escapePipe: true,
			expected:   "foo\\_bar",
		},
		{
			name:       "template builtin functions name",
			funcName:   "name",
			funcArgs:   []string{`"foo_bar"`},
			escapeChar: false,
			escapePipe: true,
			expected:   "foo_bar",
		},
		{
			name:       "template builtin functions name",
			funcName:   "name",
			funcArgs:   []string{`""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "",
		},

		// sanitizeHeader
		{
			name:       "template builtin functions sanitizeHeader",
			funcName:   "sanitizeHeader",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\n|-----|-----|\n| foo | bar |\""},
			escapeChar: true,
			escapePipe: true,
			expected:   "Example of 'foo\\_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\n|-----|-----|\n| foo | bar |",
		},
		{
			name:       "template builtin functions sanitizeHeader",
			funcName:   "sanitizeHeader",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: true,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:       "template builtin functions sanitizeHeader",
			funcName:   "sanitizeHeader",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: false,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:       "template builtin functions sanitizeHeader",
			funcName:   "sanitizeHeader",
			funcArgs:   []string{`""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "n/a",
		},

		// sanitizeDoc
		{
			name:       "template builtin functions sanitizeDoc",
			funcName:   "sanitizeDoc",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: true,
			escapePipe: true,
			expected:   "Example of 'foo\\_bar' module in `foo_bar.tf`.\n\n\\| Foo \\| Bar \\|",
		},
		{
			name:       "template builtin functions sanitizeDoc",
			funcName:   "sanitizeDoc",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: true,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.\n\n\\| Foo \\| Bar \\|",
		},
		{
			name:       "template builtin functions sanitizeDoc",
			funcName:   "sanitizeDoc",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: false,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |",
		},
		{
			name:       "template builtin functions sanitizeDoc",
			funcName:   "sanitizeDoc",
			funcArgs:   []string{`""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "n/a",
		},

		// sanitizeTbl
		{
			name:       "template builtin functions sanitizeTbl",
			funcName:   "sanitizeTbl",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: true,
			escapePipe: true,
			expected:   "Example of 'foo\\_bar' module in `foo_bar.tf`.<br><br>\\| Foo \\| Bar \\|",
		},
		{
			name:       "template builtin functions sanitizeTbl",
			funcName:   "sanitizeTbl",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: true,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.<br><br>\\| Foo \\| Bar \\|",
		},
		{
			name:       "template builtin functions sanitizeTbl",
			funcName:   "sanitizeTbl",
			funcArgs:   []string{"\"Example of 'foo_bar' module in `foo_bar.tf`.\n\n| Foo | Bar |\""},
			escapeChar: false,
			escapePipe: false,
			expected:   "Example of 'foo_bar' module in `foo_bar.tf`.<br><br>\\| Foo \\| Bar \\|",
		},
		{
			name:       "template builtin functions sanitizeTbl",
			funcName:   "sanitizeTbl",
			funcArgs:   []string{`""`},
			escapeChar: true,
			escapePipe: true,
			expected:   "n/a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := print.NewSettings()
			settings.EscapeCharacters = tt.escapeChar
			settings.EscapePipe = tt.escapePipe
			funcs := builtinFuncs(settings)

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
