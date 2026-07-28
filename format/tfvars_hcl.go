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
	"reflect"
	"strings"
	gotemplate "text/template"

	"github.com/rquadling/terraform-docs/internal/types"
	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/template"
	"github.com/rquadling/terraform-docs/terraform"
)

//go:embed templates/tfvars_hcl.tmpl
var tfvarsHCLTpl []byte

// tfvarsHCL represents Terraform tfvars HCL format.
type tfvarsHCL struct {
	*generator

	config   *print.Config
	template *template.Template
}

var padding []int

// NewTfvarsHCL returns new instance of TfvarsHCL.
func NewTfvarsHCL(config *print.Config) Type {
	tt := template.New(config, &template.Item{
		Name:      "tfvars",
		Text:      string(tfvarsHCLTpl),
		TrimSpace: true,
	})
	tt.CustomFunc(gotemplate.FuncMap{
		"align": func(s string, i int) string {
			return fmt.Sprintf("%-*s", padding[i], s)
		},
		"value": func(s string) string {
			if s == "" || s == "null" {
				return "\"\""
			}
			return s
		},
		"convertToComment": func(s types.String) string {
			return "\n# " + strings.ReplaceAll(string(s), "\n", "\n# ")
		},
		"showDescription": func() bool {
			return config.Settings.Description
		},
		"showValidation": func() bool {
			return config.Settings.Validation
		},
	})

	return &tfvarsHCL{
		generator: newGenerator(config, false),
		config:    config,
		template:  tt,
	}
}

// Generate a Terraform module as Terraform tfvars HCL.
func (h *tfvarsHCL) Generate(module *terraform.Module) error {
	alignments(module.Inputs, h.config)

	rendered, err := h.template.Render("tfvars", module)
	if err != nil {
		return err
	}

	h.funcs(withContent(strings.TrimSuffix(sanitize(rendered), "\n")))

	return nil
}

func isMultilineFormat(input *terraform.Input) bool {
	isList := input.Type == "list" || reflect.TypeOf(input.Default).Name() == "List"
	isMap := input.Type == "map" || reflect.TypeOf(input.Default).Name() == "Map"
	return (isList || isMap) && input.Default.Length() > 0
}

func alignments(inputs []*terraform.Input, config *print.Config) {
	padding = make([]int, len(inputs))
	maxlen := 0
	index := 0
	for i, input := range inputs {
		isDescribed := config.Settings.Description && input.Description.Length() > 0
		isValidated := config.Settings.Validation && input.Validation.Length() > 0
		l := len(input.Name)
		if isMultilineFormat(input) || isDescribed || isValidated {
			for j := index; j < i; j++ {
				padding[j] = maxlen
			}
			padding[i] = l
			maxlen = 0
			index = i + 1
		} else if l > maxlen {
			maxlen = l
		}
	}
	for i := index; i < len(inputs); i++ {
		padding[i] = maxlen
	}
}

func init() {
	register(map[string]initializerFn{
		"tfvars hcl": NewTfvarsHCL,
	})
}
