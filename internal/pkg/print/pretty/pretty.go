package pretty

import (
	"fmt"
	"text/template"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/segmentio/terraform-docs/internal/pkg/tmpl"
)

const (
	headerTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header }}
			{{- printf "\n" }}
			{{ colorize "\033[90m" . }}
			{{- printf "\n" }}
		{{ end -}}
	{{ end -}}
	`

	providersTpl = `
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

	inputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{- with .Module.Inputs }}
			{{- printf "\n" -}}
			{{- range . }}
				{{ printf "input.%s" .Name | colorize "\033[36m" }} ({{ default "required" .Value }})
				{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
			{{ end }}
			{{- printf "\n" -}}
		{{ end -}}
	{{ end -}}
	`

	outputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{- with .Module.Outputs }}
			{{- printf "\n" -}}
			{{- range . }}
				{{ printf "output.%s" .Name | colorize "\033[36m" }}
				{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
				{{ if $.Settings.OutputValues }}{{ .Value | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}{{ end}}
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

// Print prints a pretty document.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	t := tmpl.NewTemplate(&tmpl.Item{
		Name: "pretty",
		Text: prettyTpl,
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
		"colorize": func(c string, s string) string {
			r := "\033[0m"
			if !settings.ShowColor {
				c = ""
				r = ""
			}
			return fmt.Sprintf("%s%s%s", c, s, r)
		},
	})
	rendered, err := t.Render(module)
	if err != nil {
		return "", err
	}

	return rendered, nil
}
