/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"fmt"
	"regexp"
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/template"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
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

	prettyResourcesTpl = `
	{{- if .Settings.ShowResources -}}
		{{- with .Module.Resources }}
			{{- range . }}
				{{- if eq (len .URL) 0 }}
					{{- printf "resource.%s" .FullType | colorize "\033[36m" }}
				{{- else -}}
					{{- printf "resource.%s" .FullType | colorize "\033[36m" }} ({{ .URL}})
				{{- end }}
			{{ end -}}
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
	{{- template "resources" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// Pretty represents colorized pretty format.
type Pretty struct {
	template *template.Template
}

// NewPretty returns new instance of Pretty.
func NewPretty(settings *print.Settings) print.Engine {
	tt := template.New(settings, &template.Item{
		Name: "pretty",
		Text: prettyTpl,
	}, &template.Item{
		Name: "header",
		Text: prettyHeaderTpl,
	}, &template.Item{
		Name: "requirements",
		Text: prettyRequirementsTpl,
	}, &template.Item{
		Name: "providers",
		Text: prettyProvidersTpl,
	}, &template.Item{
		Name: "resources",
		Text: prettyResourcesTpl,
	}, &template.Item{
		Name: "inputs",
		Text: prettyInputsTpl,
	}, &template.Item{
		Name: "outputs",
		Text: prettyOutputsTpl,
	})
	tt.CustomFunc(gotemplate.FuncMap{
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

// Print a Terraform module document.
func (p *Pretty) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	rendered, err := p.template.Render(module)
	if err != nil {
		return "", err
	}
	return regexp.MustCompile(`(\r?\n)*$`).ReplaceAllString(rendered, ""), nil
}

func init() {
	register(map[string]initializerFn{
		"pretty": NewPretty,
	})
}
