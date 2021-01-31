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

func TestYaml(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("yaml", "yaml")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-SortByName")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-SortByRequired")
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

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-SortByType")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlNoResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOnlyResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
		ShowResources:    true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OutputValues")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestYamlHeaderFromFile(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		file   string
	}{
		{
			name:   "load module header from .adoc",
			golden: "yaml-HeaderFromADOCFile",
			file:   "doc.adoc",
		},
		{
			name:   "load module header from .md",
			golden: "yaml-HeaderFromMDFile",
			file:   "doc.md",
		},
		{
			name:   "load module header from .tf",
			golden: "yaml-HeaderFromTFFile",
			file:   "doc.tf",
		},
		{
			name:   "load module header from .txt",
			golden: "yaml-HeaderFromTXTFile",
			file:   "doc.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := testutil.Settings().WithSections().Build()

			expected, err := testutil.GetExpected("yaml", tt.golden)
			assert.Nil(err)

			options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
				HeaderFromFile: tt.file,
			})
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewYAML(settings)
			actual, err := printer.Print(module, settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}

func TestYamlEmpty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-Empty")
	assert.Nil(err)

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
