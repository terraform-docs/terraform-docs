package format

import (
	"fmt"
	"text/template"

	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/segmentio/terraform-docs/pkg/tmpl"
)

const (
	prettyHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header }}
			{{- printf "\n" }}
			{{ colorize "\033[90m" . }}
			{{- printf "\n" }}
		{{ end -}}
	{{ end -}}
	`

	prettyProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{- with .Module.Providers }}
			{{- printf "\n" -}}
			{{- range . }}
				{{- $version := ternary (tostring .Version) (printf " (%s)" .Version) "" }}
				{{ printf "provider.%s" .FullName | colorize "\033[36m" }}{{ $version }}
			{{ end }}
			{{- printf "\n" -}}
		{{ end -}}
	{{ end -}}
	`

	prettyInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{- with .Module.Inputs }}
			{{- printf "\n" -}}
			{{- range . }}
				{{ printf "input.%s" .Name | colorize "\033[36m" }} ({{ default "required" .GetValue }})
				{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
			{{ end }}
			{{- printf "\n" -}}
		{{ end -}}
	{{ end -}}
	`

	prettyOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{- with .Module.Outputs }}
			{{- printf "\n" -}}
			{{- range . }}
				{{ printf "output.%s" .Name | colorize "\033[36m" }}
				{{- if $.Settings.OutputValues -}}
					{{- printf " " -}}
					({{ ternary .Sensitive "<sensitive>" .GetValue }})
			{{- end }}
			{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
			{{ end }}
		{{ end -}}
	{{ end -}}
	`

	prettyTpl = `
	{{- template "header" . -}}
	{{- template "providers" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Pretty represents colorized pretty format.
type Pretty struct {
	template *tmpl.Template
}

// NewPretty returns new instance of Pretty.
func NewPretty(settings *print.Settings) *Pretty {
	tt := tmpl.NewTemplate(&tmpl.Item{
		Name: "pretty",
		Text: prettyTpl,
	}, &tmpl.Item{
		Name: "header",
		Text: prettyHeaderTpl,
	}, &tmpl.Item{
		Name: "providers",
		Text: prettyProvidersTpl,
	}, &tmpl.Item{
		Name: "inputs",
		Text: prettyInputsTpl,
	}, &tmpl.Item{
		Name: "outputs",
		Text: prettyOutputsTpl,
	})
	tt.Settings(settings)
	tt.CustomFunc(template.FuncMap{
		"colorize": func(c string, s string) string {
			r := "\033[0m"
			if !settings.ShowColor {
				c = ""
				r = ""
			}
			return fmt.Sprintf("%s%s%s", c, s, r)
		},
	})
	return &Pretty{
		template: tt,
	}
}

// Print prints a Terraform module document.
func (p *Pretty) Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	rendered, err := p.template.Render(module)
	if err != nil {
		return "", err
	}
	return rendered, nil
}
