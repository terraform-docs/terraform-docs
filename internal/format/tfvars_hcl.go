package format

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	tfvarsHCLTpl = `
	{{- if .Module.Inputs -}}
		{{- range $i, $k := .Module.Inputs -}}
			{{ align $k.Name $i }} = {{ default "\"\"" $k.GetValue }}
		{{ end -}}
	{{- end -}}
	`
)

// TfvarsHCL represents Terraform tfvars HCL format.
type TfvarsHCL struct {
	template *tmpl.Template
}

var padding []int

// NewTfvarsHCL returns new instance of TfvarsHCL.
func NewTfvarsHCL(settings *print.Settings) *TfvarsHCL {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "tfvars",
		Text: tfvarsHCLTpl,
	})
	tt.Settings(settings)
	tt.CustomFunc(template.FuncMap{
		"align": func(s string, i int) string {
			return fmt.Sprintf("%-*s", padding[i], s)
		},
	})
	return &TfvarsHCL{
		template: tt,
	}
}

// Print prints a Terraform module as Terraform tfvars HCL document.
func (h *TfvarsHCL) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	align(module.Inputs)
	rendered, err := h.template.Render(module)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(sanitize(rendered), "\n"), nil
}

func align(inputs []*tfconf.Input) {
	padding = make([]int, len(inputs))
	maxlen := 0
	index := 0
	for i, input := range inputs {
		isList := input.Type == "list" || reflect.TypeOf(input.Default).Name() == "List"
		isMap := input.Type == "map" || reflect.TypeOf(input.Default).Name() == "Map"
		l := len(input.Name)
		if (isList || isMap) && input.Default.Length() > 0 {
			for j := index; j < i; j++ {
				padding[j] = maxlen
			}
			padding[i] = l
			maxlen = 0
			index = i + 1
		} else {
			if l > maxlen {
				maxlen = l
			}
		}
	}
	for i := index; i < len(inputs); i++ {
		padding[i] = maxlen
	}
}
