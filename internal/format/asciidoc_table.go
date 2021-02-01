/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/template"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

const (
	asciidocTableHeaderTpl = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ sanitizeHeader . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`
	asciidocTableResourcesTpl = `
	{{- if .Settings.ShowResources -}}
		{{ indent 0 "=" }} Resources
		{{ if not .Module.Resources }}
			No resources.
		{{ else }}
			[cols="a",options="header,autowidth"]
			|===
			|Name
			{{- range .Module.Resources }}
				{{ if eq (len .URL) 0 }}
				|{{ .FullType }}
				{{- else -}}
				|{{ .URL }}[{{ .FullType }}]
				{{- end }}
			{{- end }}
			|===
		{{ end }}
	{{ end -}}
	`

	asciidocTableRequirementsTpl = `
	{{- if .Settings.ShowRequirements -}}
		{{ indent 0 "=" }} Requirements
		{{ if not .Module.Requirements }}
			No requirements.
		{{ else }}
			[cols="a,a",options="header,autowidth"]
			|===
			|Name |Version
			{{- range .Module.Requirements }}
				|{{ .Name }} |{{ tostring .Version | default "n/a" }}
			{{- end }}
			|===
		{{ end }}
	{{ end -}}
	`

	asciidocTableProvidersTpl = `
	{{- if .Settings.ShowProviders -}}
		{{ indent 0 "=" }} Providers
		{{ if not .Module.Providers }}
			No provider.
		{{ else }}
			[cols="a,a",options="header,autowidth"]
			|===
			|Name |Version
			{{- range .Module.Providers }}
				|{{ .FullName }} |{{ tostring .Version | default "n/a" }}
			{{- end }}
			|===
		{{ end }}
	{{ end -}}
	`

	asciidocTableInputsTpl = `
	{{- if .Settings.ShowInputs -}}
		{{ indent 0 "=" }} Inputs
		{{ if not .Module.Inputs }}
			No input.
		{{ else }}
			[cols="a,a,a,a{{ if .Settings.ShowRequired }},a{{ end }}",options="header,autowidth"]
			|===
			|Name |Description |Type |Default{{ if .Settings.ShowRequired }} |Required{{ end }}
			{{- range .Module.Inputs }}
				|{{ .Name }}
				|{{ tostring .Description | sanitizeAsciidocTbl }}
				|{{ tostring .Type | type | sanitizeAsciidocTbl }}
				|{{ value .GetValue | sanitizeAsciidocTbl }}
				{{ if $.Settings.ShowRequired }}|{{ ternary .Required "yes" "no" }}{{ end }}
			{{ end }}
			|===
		{{ end }}
	{{ end -}}
	`

	asciidocTableOutputsTpl = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 "=" }} Outputs
		{{ if not .Module.Outputs }}
			No output.
		{{ else }}
			[cols="a,a{{ if .Settings.OutputValues }},a{{ if $.Settings.ShowSensitivity }},a{{ end }}{{ end }}",options="header,autowidth"]
			|===
			|Name |Description{{ if .Settings.OutputValues }} |Value{{ if $.Settings.ShowSensitivity }} |Sensitive{{ end }}{{ end }}
			{{- range .Module.Outputs }}
				|{{ .Name }} |{{ tostring .Description | sanitizeAsciidocTbl }}
				{{- if $.Settings.OutputValues -}}
					{{- $sensitive := ternary .Sensitive "<sensitive>" .GetValue -}}
					{{ printf " " }}|{{ value $sensitive }}
					{{- if $.Settings.ShowSensitivity -}}
						{{ printf " " }}|{{ ternary .Sensitive "yes" "no" }}
					{{- end -}}
				{{- end -}}
			{{- end }}
			|===
		{{ end }}
	{{ end -}}
	`

	asciidocTableTpl = `
	{{- template "header" . -}}
	{{- template "requirements" . -}}
	{{- template "providers" . -}}
	{{- template "resources" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	`
)

// AsciidocTable represents AsciiDoc Table format.
type AsciidocTable struct {
	template *template.Template
}

// NewAsciidocTable returns new instance of AsciidocTable.
func NewAsciidocTable(settings *print.Settings) print.Engine {
	settings.EscapeCharacters = false
	tt := template.New(settings, &template.Item{
		Name: "table",
		Text: asciidocTableTpl,
	}, &template.Item{
		Name: "header",
		Text: asciidocTableHeaderTpl,
	}, &template.Item{
		Name: "resources",
		Text: asciidocTableResourcesTpl,
	}, &template.Item{
		Name: "requirements",
		Text: asciidocTableRequirementsTpl,
	}, &template.Item{
		Name: "providers",
		Text: asciidocTableProvidersTpl,
	}, &template.Item{
		Name: "inputs",
		Text: asciidocTableInputsTpl,
	}, &template.Item{
		Name: "outputs",
		Text: asciidocTableOutputsTpl,
	})
	tt.CustomFunc(gotemplate.FuncMap{
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
	return &AsciidocTable{
		template: tt,
	}
}

// Print a Terraform module as AsciiDoc tables.
func (t *AsciidocTable) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	rendered, err := t.template.Render(module)
	if err != nil {
		return "", err
	}
	return sanitize(rendered), nil
}

func init() {
	register(map[string]initializerFn{
		"asciidoc":       NewAsciidocTable,
		"asciidoc table": NewAsciidocTable,
		"asciidoc tbl":   NewAsciidocTable,
		"adoc":           NewAsciidocTable,
		"adoc table":     NewAsciidocTable,
		"adoc tbl":       NewAsciidocTable,
	})
}
