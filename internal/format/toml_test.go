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

func TestToml(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("toml", "toml")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-SortByName")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-SortByRequired")
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

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-SortByType")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		SortBy: &terraform.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoHeader(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoInputs(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoModulecalls(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoOutputs(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoProviders(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoRequirements(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-NoRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlNoResources(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    false,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-NoResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyHeader(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyHeader")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyInputs(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyInputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyModulecalls(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyModulecalls")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyOutputs(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyOutputs")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyProviders(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyProviders")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyRequirements(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyRequirements")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOnlyResources(t *testing.T) {
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

	expected, err := testutil.GetExpected("toml", "toml-OnlyResources")
	assert.Nil(err)

	options := terraform.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-OutputValues")
	assert.Nil(err)

	options, err := terraform.NewOptions().With(&terraform.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlHeaderFromFile(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("toml", "toml-HeaderFromFile")
	assert.Nil(err)

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "doc.tf",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTomlEmpty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowModuleCalls:  false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("toml", "toml-Empty")
	assert.Nil(err)

	options, err := terraform.NewOptions().WithOverwrite(&terraform.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTOML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
