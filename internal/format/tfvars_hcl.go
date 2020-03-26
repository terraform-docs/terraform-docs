package format

import (
	"strings"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	tfvarsHCLTpl = `
	{{- if .Module.Inputs -}}
		{{- range .Module.Inputs -}}
			{{ .Name }} = {{ default "\"\"" .GetValue }}
		{{ end -}}
	{{- end -}}
	`
)

// TfvarsHCL represents Terraform tfvars HCL format.
type TfvarsHCL struct {
	template *tmpl.Template
}

// NewTfvarsHCL returns new instance of TfvarsHCL.
func NewTfvarsHCL(settings *print.Settings) *TfvarsHCL {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "tfvars",
		Text: tfvarsHCLTpl,
	})
	tt.Settings(settings)
	return &TfvarsHCL{
		template: tt,
	}
}

// Print prints a Terraform module as Terraform tfvars HCL document.
func (h *TfvarsHCL) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	rendered, err := h.template.Render(module)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(sanitize(rendered), "\n"), nil
}
