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

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/template"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

//go:embed templates/tfvars_hcl.tmpl
var tfvarsHCLTpl []byte

// TfvarsHCL represents Terraform tfvars HCL format.
type TfvarsHCL struct {
	template *template.Template
}

var padding []int

// NewTfvarsHCL returns new instance of TfvarsHCL.
func NewTfvarsHCL(settings *print.Settings) print.Engine {
	tt := template.New(settings, &template.Item{
		Name: "tfvars",
		Text: string(tfvarsHCLTpl),
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
	})
	return &TfvarsHCL{
		template: tt,
	}
}

// Print a Terraform module as Terraform tfvars HCL.
func (h *TfvarsHCL) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	alignments(module.Inputs)
	rendered, err := h.template.Render(module)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(sanitize(rendered), "\n"), nil
}

func alignments(inputs []*terraform.Input) {
	padding = make([]int, len(inputs))
	maxlen := 0
	index := 0
	for i, input := range inputs {
		isList := input.Type == "list" || reflect.TypeOf(input.Default).Name() == "List"
		isMap := input.Type == "map" || reflect.TypeOf(input.Default).Name() == "Map"
		l := len(input.Name)
		if (isList || isMap) && input.Default.Length() > 0 {
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
