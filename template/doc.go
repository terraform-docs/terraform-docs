/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

// Package template provides templating functionality.
//
// Usage
//
//	import (
//	    "fmt"
//	    gotemplate "text/template"
//
//	    "github.com/rquadling/terraform-docs/print"
//	    "github.com/rquadling/terraform-docs/template"
//	    "github.com/rquadling/terraform-docs/terraform"
//	)
//
//	const mainTpl =`
//	{{- if .Config.Sections.Header -}}
//	    {{- with .Module.Header -}}
//	        {{ colorize "\033[90m" . }}
//	    {{ end -}}
//	    {{- printf "\n\n" -}}
//	{{ end -}}`
//
//	func render(config *print.Config, module *terraform.Module) (string, error) {
//	    tt := template.New(config, &template.Item{
//	        Name:      "main",
//	        Text:      mainTpl,
//	        TrimSpace: true,
//	    })
//
//	    tt := template.New(config, items...)
//	    tt.CustomFunc(gotemplate.FuncMap{
//	        "colorize": func(color string, s string) string {
//	            reset := "\033[0m"
//	            if !config.Settings.Color {
//	                color = ""
//	                reset = ""
//	            }
//	            return fmt.Sprintf("%s%s%s", color, s, reset)
//	        },
//	    })
//
//	    return tt.Render("main", module)
//	}
package template
