package format

import (
	"fmt"
	"regexp"
	"text/template"

	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/pkg/print"
	"github.com/terraform-docs/terraform-docs/pkg/tmpl"
)

const (
	prettyHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ colorize "\033[90m" . }}
		{{ end -}}
		{{- printf "\n\n" -}}
	{{ end -}}
	`

	prettyRequirementsTpl = `
	{{- if .Settings.ShowRequirements -}}
		{{- with .Module.Requirements }}
			{{- range . }}
				{{- $version := ternary (tostring .Version) (printf " (%s)" .Version) "" }}
				{{- printf "requirement.%s" .Name | colorize "\033[36m" }}{{ $version }}
			{{ end -}}
		{{ end -}}
		{{- printf "\n\n" -}}
	{{ end -}}
	`

	prettyProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{- with .Module.Providers }}
			{{- range . }}
				{{- $version := ternary (tostring .Version) (printf " (%s)" .Version) "" }}
				{{- printf "provider.%s" .FullName | colorize "\033[36m" }}{{ $version }}
			{{ end -}}
		{{ end -}}
		{{- printf "\n\n" -}}
	{{ end -}}
	`

	prettyInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{- with .Module.Inputs }}
			{{- range . }}
				{{- printf "input.%s" .Name | colorize "\033[36m" }} ({{ default "required" .GetValue }})
				{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
				{{- printf "\n\n" -}}
			{{ end -}}
		{{ end -}}
		{{- printf "\n" -}}
	{{ end -}}
	`

	prettyOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{- with .Module.Outputs }}
			{{- range . }}
				{{- printf "output.%s" .Name | colorize "\033[36m" }}
				{{- if $.Settings.OutputValues -}}
					{{- printf " " -}}
					({{ ternary .Sensitive "<sensitive>" .GetValue }})
				{{- end }}
				{{ tostring .Description | trimSuffix "\n" | default "n/a" | colorize "\033[90m" }}
				{{- printf "\n\n" -}}
			{{ end -}}
		{{ end -}}
	{{ end -}}
	`

	prettyTpl = `
	{{- template "header" . -}}
	{{- template "requirements" . -}}
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
		Name: "requirements",
		Text: prettyRequirementsTpl,
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
func (p *Pretty) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	rendered, err := p.template.Render(module)
	if err != nil {
		return "", err
	}
	return regexp.MustCompile(`(\r?\n)*$`).ReplaceAllString(rendered, ""), nil
}
