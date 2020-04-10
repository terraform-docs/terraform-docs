package format

import (
	"text/template"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	documentHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ sanitizeHeader . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`

	documentRequirementsTpl = `
	{{- if .Settings.ShowRequirements -}}
		{{ indent 0 }} Requirements
		{{ if not .Module.Requirements }}
			No requirements.
		{{ else }}
			The following requirements are needed by this module:
			{{- range .Module.Requirements }}
				{{ $version := ternary (tostring .Version) (printf " (%s)" .Version) "" }}
				- {{ name .Name }}{{ $version }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	documentProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{ indent 0 }} Providers
		{{ if not .Module.Providers }}
			No provider.
		{{ else }}
			The following providers are used by this module:
			{{- range .Module.Providers }}
				{{ $version := ternary (tostring .Version) (printf " (%s)" .Version) "" }}
				- {{ name .FullName }}{{ $version }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	documentInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{- if .Settings.ShowRequired -}}
			{{ indent 0 }} Required Inputs
			{{ if not .Module.RequiredInputs }}
				No required input.
			{{ else }}
				The following input variables are required:
				{{- range .Module.RequiredInputs }}
					{{ template "input" . }}
				{{- end }}
			{{- end }}
			{{ indent 0 }} Optional Inputs
			{{ if not .Module.OptionalInputs }}
				No optional input.
			{{ else }}
				The following input variables are optional (have default values):
				{{- range .Module.OptionalInputs }}
					{{ template "input" . }}
				{{- end }}
			{{ end }}
		{{ else -}}
			{{ indent 0 }} Inputs
			{{ if not .Module.Inputs }}
				No input.
			{{ else }}
				The following input variables are supported:
				{{- range .Module.Inputs }}
					{{ template "input" . }}
				{{- end }}
			{{ end }}
		{{- end }}
	{{ end -}}
	`

	documentInputTpl = `
	{{ printf "\n" }}
	{{ indent 1 }} {{ name .Name }}

	Description: {{ tostring .Description | sanitizeDoc }}

	Type: {{ tostring .Type | type }}

	{{ if or .HasDefault (not isRequired) }}
		Default: {{ default "n/a" .GetValue | value }}
	{{- end }}
	`

	documentOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			The following outputs are exported:
			{{- range .Module.Outputs }}

				{{ indent 1 }} {{ name .Name }}

				Description: {{ tostring .Description | sanitizeDoc }}

				{{ if $.Settings.OutputValues }}
					{{- $sensitive := ternary .Sensitive "<sensitive>" .GetValue -}}
					Value: {{ value $sensitive | sanitizeDoc }}

					{{ if $.Settings.ShowSensitivity -}}
						Sensitive: {{ ternary (.Sensitive) "yes" "no" }}
					{{- end }}
				{{ end }}
			{{ end }}
		{{ end }}
	{{ end -}}
	`

	documentTpl = `
	{{- template "header" . -}}
	{{- template "requirements" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Document represents Markdown Document format.
type Document struct {
	template *tmpl.Template
}

// NewDocument returns new instance of Document.
func NewDocument(settings *print.Settings) *Document {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "document",
		Text: documentTpl,
	}, &tmpl.Item{
		Name: "header",
		Text: documentHeaderTpl,
	}, &tmpl.Item{
		Name: "requirements",
		Text: documentRequirementsTpl,
	}, &tmpl.Item{
		Name: "providers",
		Text: documentProvidersTpl,
	}, &tmpl.Item{
		Name: "inputs",
		Text: documentInputsTpl,
	}, &tmpl.Item{
		Name: "input",
		Text: documentInputTpl,
	}, &tmpl.Item{
		Name: "outputs",
		Text: documentOutputsTpl,
	})
	tt.Settings(settings)
	tt.CustomFunc(template.FuncMap{
		"type": func(t string) string {
			result, extraline := printFencedCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := printFencedCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})
	return &Document{
		template: tt,
	}
}

// Print prints a Terraform module as Markdown document.
func (d *Document) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	rendered, err := d.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
}
