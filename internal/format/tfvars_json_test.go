/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
	"github.com/terraform-docs/terraform-docs/terraform"
)

func TestTfvarsJson(t *testing.T) {
	tests := map[string]struct {
		settings print.Settings
		options  terraform.Options
	}{
		// Base
		"Base": {
			settings: testutil.WithSections(),
			options:  terraform.Options{},
		},
		"Empty": {
			settings: testutil.WithSections(),
			options: terraform.Options{
				Path: "empty",
			},
		},

		// Settings
		"EscapeCharacters": {
			settings: print.Settings{EscapeCharacters: true},
			options:  terraform.Options{},
		},
		"SortByName": {
			settings: testutil.WithSections(),
			options: terraform.Options{
				SortBy: &terraform.SortBy{
					Name: true,
				},
			},
		},
		"SortByRequired": {
			settings: testutil.WithSections(),
			options: terraform.Options{
				SortBy: &terraform.SortBy{
					Name:     true,
					Required: true,
				},
			},
		},
		"SortByType": {
			settings: testutil.WithSections(),
			options: terraform.Options{
				SortBy: &terraform.SortBy{
					Type: true,
				},
			},
		},

		// No section
		"NoInputs": {
			settings: testutil.WithSections(
				print.Settings{
					ShowInputs: false,
				},
			),
			options: terraform.Options{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("tfvars", "json-"+name)
			assert.Nil(err)

			options, err := terraform.NewOptions().With(&tt.options)
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			formatter := NewTfvarsJSON(&tt.settings)

			generator, err := formatter.Generate(module)
			assert.Nil(err)

			actual, err := generator.ExecuteTemplate("")

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}
