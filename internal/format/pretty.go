/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	_ "embed" //nolint
	"fmt"
	"regexp"
	gotemplate "text/template"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/template"
	"github.com/terraform-docs/terraform-docs/terraform"
)

//go:embed templates/pretty.tmpl
var prettyTpl []byte

// Pretty represents colorized pretty format.
type Pretty struct {
	template *template.Template
	settings *print.Settings
}

// NewPretty returns new instance of Pretty.
func NewPretty(settings *print.Settings) print.Engine {
	tt := template.New(settings, &template.Item{
		Name: "pretty",
		Text: string(prettyTpl),
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
		settings: settings,
	}
}

// Generate a Terraform module document.
func (p *Pretty) Generate(module *terraform.Module) (*print.Generator, error) {
	rendered, err := p.template.Render("pretty", module)
	if err != nil {
		return nil, err
	}

	return print.NewGenerator(
		"pretty",
		print.WithContent(regexp.MustCompile(`(\r?\n)*$`).ReplaceAllString(rendered, "")),
	), nil
}

func init() {
	register(map[string]initializerFn{
		"pretty": NewPretty,
	})
}
