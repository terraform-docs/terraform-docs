package format

import (
	"text/template"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	asciidocDocumentHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ sanitizeHeader . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`

	asciidocDocumentRequirementsTpl = `
	{{- if .Settings.ShowRequirements -}}
		{{ indentAsciidoc 0 }} Requirements
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

	asciidocDocumentProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{ indentAsciidoc 0 }} Providers
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

	asciidocDocumentInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{- if .Settings.ShowRequired -}}
			{{ indentAsciidoc 0 }} Required Inputs
			{{ if not .Module.RequiredInputs }}
				No required input.
			{{ else }}
				The following input variables are required:
				{{- range .Module.RequiredInputs }}
					{{ template "input" . }}
				{{- end }}
			{{- end }}
			{{ indentAsciidoc 0 }} Optional Inputs
			{{ if not .Module.OptionalInputs }}
				No optional input.
			{{ else }}
				The following input variables are optional (have default values):
				{{- range .Module.OptionalInputs }}
					{{ template "input" . }}
				{{- end }}
			{{ end }}
		{{ else -}}
			{{ indentAsciidoc 0 }} Inputs
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

	asciidocDocumentInputTpl = `
	{{ printf "\n" }}
	{{ indentAsciidoc 1 }} {{ name .Name }}

	Description: {{ tostring .Description | sanitizeDoc }}

	Type: {{ tostring .Type | type }}

	{{ if or .HasDefault (not isRequired) }}
		Default: {{ default "n/a" .GetValue | value }}
	{{- end }}
	`

	asciidocDocumentOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indentAsciidoc 0 }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			The following outputs are exported:
			{{- range .Module.Outputs }}

				{{ indentAsciidoc 1 }} {{ name .Name }}

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

	asciidocDocumentTpl = `
	{{- template "header" . -}}
	{{- template "requirements" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// AsciidocDocument represents AsciiDoc Document format.
type AsciidocDocument struct {
	template *tmpl.Template
}

// NewAsciidocDocument returns new instance of AsciidocDocument.
func NewAsciidocDocument(settings *print.Settings) *AsciidocDocument {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "document",
		Text: asciidocDocumentTpl,
	}, &tmpl.Item{
		Name: "header",
		Text: asciidocDocumentHeaderTpl,
	}, &tmpl.Item{
		Name: "requirements",
		Text: asciidocDocumentRequirementsTpl,
	}, &tmpl.Item{
		Name: "providers",
		Text: asciidocDocumentProvidersTpl,
	}, &tmpl.Item{
		Name: "inputs",
		Text: asciidocDocumentInputsTpl,
	}, &tmpl.Item{
		Name: "input",
		Text: asciidocDocumentInputTpl,
	}, &tmpl.Item{
		Name: "outputs",
		Text: asciidocDocumentOutputsTpl,
	})
	settings.EscapeCharacters = false
	tt.Settings(settings)
	tt.CustomFunc(template.FuncMap{
		"type": func(t string) string {
			result, extraline := printFencedAsciidocCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := printFencedAsciidocCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})
	return &AsciidocDocument{
		template: tt,
	}
}

// Print prints a Terraform module as AsciiDoc document.
func (d *AsciidocDocument) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	rendered, err := d.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
}
