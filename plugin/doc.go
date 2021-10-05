/*
Copyright 2021 The terraform-docs Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package plugin contains the implementations needed to make
// the built binary act as a plugin.
//
// A plugin is implemented as an RPC server and the host acts
// as the client, sending analysis requests to the plugin.
// Note that the server-client relationship here is the opposite of
// the communication that takes place during the checking phase.
//
// Implementation details are hidden in go-plugin. This package is
// essentially a wrapper for go-plugin.
//
// Usage
//
// A simple plugin can look like this:
//
//     package main
//
//     import (
//         _ "embed" //nolint
//
//         "github.com/terraform-docs/terraform-docs/plugin"
//         "github.com/terraform-docs/terraform-docs/print"
//         "github.com/terraform-docs/terraform-docs/template"
//         "github.com/terraform-docs/terraform-docs/terraform"
//     )
//
//     func main() {
//         plugin.Serve(&plugin.ServeOpts{
//             Name:    "template",
//             Version: "0.1.0",
//             Printer: printerFunc,
//         })
//     }
//
//     //go:embed sections.tmpl
//     var tplCustom []byte
//
//     // printerFunc the function being executed by the plugin client.
//     func printerFunc(config *print.Config, module *terraform.Module) (string, error) {
//         tpl := template.New(config,
//             &template.Item{Name: "custom", Text: string(tplCustom)},
//         )
//
//         rendered, err := tpl.Render("custom", module)
//         if err != nil {
//             return "", err
//         }
//
//         return rendered, nil
//     }
//
package plugin
