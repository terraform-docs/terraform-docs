package format

import (
	"text/template"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	tableHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ sanitizeHeader . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`

	tableRequirementsTpl = `
	{{- if .Settings.ShowRequirements -}}
		{{ indent 0 }} Requirements
		{{ if not .Module.Requirements }}
			No requirements.
		{{ else }}
			| Name | Version |
			|------|---------|
			{{- range .Module.Requirements }}
				| {{ name .Name }} | {{ tostring .Version | default "n/a" }} |
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tableProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{ indent 0 }} Providers
		{{ if not .Module.Providers }}
			No provider.
		{{ else }}
			| Name | Version |
			|------|---------|
			{{- range .Module.Providers }}
				| {{ name .FullName }} | {{ tostring .Version | default "n/a" }} |
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tableInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{ indent 0 }} Inputs
		{{ if not .Module.Inputs }}
			No input.
		{{ else }}
			| Name | Description | Type | Default |{{ if .Settings.ShowRequired }} Required |{{ end }}
			|------|-------------|------|---------|{{ if .Settings.ShowRequired }}:--------:|{{ end }}
			{{- range .Module.Inputs }}
				| {{ name .Name }} | {{ tostring .Description | sanitizeTbl }} | {{ tostring .Type | type | sanitizeTbl }} | {{ value .GetValue | sanitizeTbl }} |
				{{- if $.Settings.ShowRequired -}}
					{{ printf " " }}{{ ternary (.GetValue) "no" "yes" }} |
				{{- end -}}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tableOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			| Name | Description |{{ if .Settings.OutputValues }} Value | Sensitive |{{ end }}
			|------|-------------|{{ if .Settings.OutputValues }}-------|:---------:|{{ end }}
			{{- range .Module.Outputs }}
				| {{ name .Name }} | {{ tostring .Description | sanitizeTbl }} |
				{{- if $.Settings.OutputValues -}}
					{{- $sensitive := ternary .Sensitive "<sensitive>" .GetValue -}}
					{{ printf " " }}{{ value $sensitive | sanitizeTbl }} | {{ ternary (.Sensitive) "yes" "no" }} |
				{{- end -}}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tableTpl = `
	{{- template "header" . -}}
	{{- template "requirements" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Table represents Markdown Table format.
type Table struct {
	template *tmpl.Template
}

// NewTable returns new instance of Table.
func NewTable(settings *print.Settings) *Table {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "table",
		Text: tableTpl,
	}, &tmpl.Item{
		Name: "requirements",
		Text: tableRequirementsTpl,
	}, &tmpl.Item{
		Name: "header",
		Text: tableHeaderTpl,
	}, &tmpl.Item{
		Name: "providers",
		Text: tableProvidersTpl,
	}, &tmpl.Item{
		Name: "inputs",
		Text: tableInputsTpl,
	}, &tmpl.Item{
		Name: "outputs",
		Text: tableOutputsTpl,
	})
	tt.Settings(settings)
	tt.CustomFunc(template.FuncMap{
		"type": func(t string) string {
			inputType, _ := printFencedCodeBlock(t, "")
			return inputType
		},
		"value": func(v string) string {
			var result = "n/a"
			if v != "" {
				result, _ = printFencedCodeBlock(v, "")
			}
			return result
		},
	})
	return &Table{
		template: tt,
	}
}

// Print prints a Terraform module as Markdown tables.
func (t *Table) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	rendered, err := t.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
}
