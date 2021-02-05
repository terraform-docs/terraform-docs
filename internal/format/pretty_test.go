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

func TestPretty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().Build()

	expected, err := testutil.GetExpected("pretty", "pretty")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-SortByName")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-SortByRequired")
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

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-SortByType")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoModulecalls(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyModulecalls(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoColor(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowColor: false,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-NoColor")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty", "pretty-OutputValues")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyHeaderFromFile(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		file   string
	}{
		{
			name:   "load module header from .adoc",
			golden: "pretty-HeaderFromADOCFile",
			file:   "doc.adoc",
		},
		{
			name:   "load module header from .md",
			golden: "pretty-HeaderFromMDFile",
			file:   "doc.md",
		},
		{
			name:   "load module header from .tf",
			golden: "pretty-HeaderFromTFFile",
			file:   "doc.tf",
		},
		{
			name:   "load module header from .txt",
			golden: "pretty-HeaderFromTXTFile",
			file:   "doc.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := testutil.Settings().WithSections().WithColor().Build()

			expected, err := testutil.GetExpected("pretty", tt.golden)
			assert.Nil(err)

			options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
				HeaderFromFile: tt.file,
			})
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewPretty(settings)
			actual, err := printer.Print(module, settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}

func TestPrettyEmpty(t *testing.T) {
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

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal("", actual)
}
