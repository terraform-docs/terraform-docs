package table

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
			| Name | Version |
			|------|---------|
			{{- range .Module.Providers }}
				| {{ name .FullName }} | {{ tostring .Version | default "n/a" }} |
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	inputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{ indent 0 }} Inputs
		{{ if not .Module.Inputs }}
			No input.
		{{ else }}
			{{ if not .Settings.ShowRequired }}
				| Name | Description | Type | Default |
				|------|-------------|------|---------|
			{{- else }}
				| Name | Description | Type | Default | Required |
				|------|-------------|------|---------|:--------:|
			{{- end }}
			{{- range .Module.Inputs }}
				{{- if not $.Settings.ShowRequired }}
					| {{ name .Name }} | {{ tostring .Description | sanitizeTbl }} | {{ tostring .Type | type | sanitizeTbl }} | {{ value .Value | sanitizeTbl }} |
				{{- else }}
					| {{ name .Name }} | {{ tostring .Description | sanitizeTbl }} | {{ tostring .Type | type | sanitizeTbl }} | {{ value .Value | sanitizeTbl }} | {{ ternary (.Value) "no" "yes" }} |
				{{- end }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	outputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			| Name | Description | {{if $.Settings.OutputValues}}Value |{{end}}
			|------|-------------|{{if $.Settings.OutputValues}}-------|{{end}}
			{{- range .Module.Outputs }}
				| {{ name .Name }} | {{ tostring .Description | sanitizeTbl }} | {{if $.Settings.OutputValues}}{{ .Value | sanitizeTbl }} |{{end}}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tableTpl = `
	{{- template "header" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Print prints a document as Markdown tables.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	t := tmpl.NewTemplate(&tmpl.Item{
		Name: "table",
		Text: tableTpl,
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
		Name: "outputs",
		Text: outputsTpl,
	})
	t.Settings(settings)
	t.CustomFunc(template.FuncMap{
		"type": func(t string) string {
			inputType, _ := markdown.PrintFencedCodeBlock(t, "")
			return inputType
		},
		"value": func(v string) string {
			var result = "n/a"
			if v != "" {
				result, _ = markdown.PrintFencedCodeBlock(v, "")
			}
			return result
		},
	})
	rendered, err := t.Render(module)
	if err != nil {
		return "", err
	}

	return markdown.Sanitize(rendered), nil
}
