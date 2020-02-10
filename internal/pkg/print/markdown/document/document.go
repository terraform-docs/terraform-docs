package document

import (
	"text/template"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/segmentio/terraform-docs/internal/pkg/tmpl"
)

const (
	headerTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ sanitizeHeader . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`

	providersTpl = `
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

	inputsTpl = `
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

	inputTpl = `
	{{ printf "\n" }}
	{{ indent 1 }} {{ name .Name }}

	Description: {{ tostring .Description | sanitizeDoc }}

	Type: {{ tostring .Type | type }}

	{{ if or .HasDefault (not isRequired) }}
		Default: {{ default "n/a" .Value | value }}
	{{- end }}
	`

	outputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			The following outputs are exported:
			{{- range .Module.Outputs }}

				{{ indent 1 }} {{ name .Name }}

				Description: {{ tostring .Description | sanitizeDoc }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	documentTpl = `
	{{- template "header" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Print prints a document as Markdown document.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	t := tmpl.NewTemplate(&tmpl.Item{
		Name: "document",
		Text: documentTpl,
	}, &tmpl.Item{
		Name: "header",
		Text: headerTpl,
	}, &tmpl.Item{
		Name: "providers",
		Text: providersTpl,
	}, &tmpl.Item{
		Name: "inputs",
		Text: inputsTpl,
	}, &tmpl.Item{
		Name: "input",
		Text: inputTpl,
	}, &tmpl.Item{
		Name: "outputs",
		Text: outputsTpl,
	})
	t.Settings(settings)
	t.CustomFunc(template.FuncMap{
		"type": func(t string) string {
			result, extraline := markdown.PrintFencedCodeBlock(t, "hcl")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"value": func(v string) string {
			if v == "n/a" {
				return v
			}
			result, extraline := markdown.PrintFencedCodeBlock(v, "json")
			if !extraline {
				result += "\n"
			}
			return result
		},
		"isRequired": func() bool {
			return settings.ShowRequired
		},
	})
	rendered, err := t.Render(module)
	if err != nil {
		return "", err
	}

	return markdown.Sanitize(rendered), nil
}
