package tmpl

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
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
func (t *Template) Render(module *tfconf.Module) (string, error) {
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
		Module   *tfconf.Module
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
		"ternary": func(c interface{}, t string, f string) string {
			var condition bool
			switch x := fmt.Sprintf("%T", c); x {
			case "string":
				condition = c.(string) != ""
			case "int":
				condition = c.(int) != 0
			case "bool":
				condition = c.(bool)
			}
			if condition {
				return t
			}
			return f
		},
		"tostring": func(s types.String) string {
			return string(s)
		},
		"trim": func(t string, s string) string {
			if s != "" {
				return strings.Trim(s, t)
			}
			return s
		},
		"trimLeft": func(t string, s string) string {
			if s != "" {
				return strings.TrimLeft(s, t)
			}
			return s
		},
		"trimRight": func(t string, s string) string {
			if s != "" {
				return strings.TrimRight(s, t)
			}
			return s
		},
		"trimPrefix": func(t string, s string) string {
			if s != "" {
				return strings.TrimPrefix(s, t)
			}
			return s
		},
		"trimSuffix": func(t string, s string) string {
			if s != "" {
				return strings.TrimSuffix(s, t)
			}
			return s
		},

		"indent": func(l int) string {
			return generateIndentation(l, settings)
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
			return sanitizeItemForTable(s, settings)
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
