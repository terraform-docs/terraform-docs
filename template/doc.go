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
//     import (
//         "fmt"
//         gotemplate "text/template"
//
//         "github.com/terraform-docs/terraform-docs/internal/print"
//         "github.com/terraform-docs/terraform-docs/internal/terraform"
//         "github.com/terraform-docs/terraform-docs/template"
//     )
//
//     const mainTpl =`
//     {{- if .Settings.ShowHeader -}}
//         {{- with .Module.Header -}}
//             {{ colorize "\033[90m" . }}
//         {{ end -}}
//         {{- printf "\n\n" -}}
//     {{ end -}}`
//
//     func render(settings *print.Settings, module *terraform.Module) (string, error) {
//         tt := template.New(settings, &template.Item{
//             Name: "main",
//             Text: mainTpl,
//         })
//
//         tt := template.New(settings, items...)
//         tt.CustomFunc(gotemplate.FuncMap{
//             "colorize": func(c string, s string) string {
//                 r := "\033[0m"
//                 if !settings.ShowColor {
//                     c = ""
//                     r = ""
//                 }
//                 return fmt.Sprintf("%s%s%s", c, s, r)
//             },
//         })
//
//         return tt.Render("main", module)
//     }
//
package template
