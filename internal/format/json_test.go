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
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
)

func TestJson(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("json", "json")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-SortByName")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-SortByRequired")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name:     true,
			Required: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-SortByType")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoModulecalls(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoReources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-NoResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyModulecalls(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OnlyResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		EscapeCharacters: true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-EscapeCharacters")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-OutputValues")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonHeaderFromFile(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		file   string
	}{
		{
			name:   "load module header from .adoc",
			golden: "json-HeaderFromADOCFile",
			file:   "doc.adoc",
		},
		{
			name:   "load module header from .md",
			golden: "json-HeaderFromMDFile",
			file:   "doc.md",
		},
		{
			name:   "load module header from .tf",
			golden: "json-HeaderFromTFFile",
			file:   "doc.tf",
		},
		{
			name:   "load module header from .txt",
			golden: "json-HeaderFromTXTFile",
			file:   "doc.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := testutil.Settings().WithSections().Build()

			expected, err := testutil.GetExpected("json", tt.golden)
			assert.Nil(err)

			options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
				HeaderFromFile: tt.file,
			})
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewJSON(settings)
			actual, err := printer.Print(module, settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}

func TestJsonEmpty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("json", "json-Empty")
	assert.Nil(err)

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewJSON(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
