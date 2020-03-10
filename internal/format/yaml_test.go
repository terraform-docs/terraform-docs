package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestYaml(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("yaml", "yaml")
	assert.Nil(err)

	options := module.NewOptions()
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

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
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

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
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

func TestYamlNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoHeader")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoProviders")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoInputs")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-NoOutputs")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyHeader")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyProviders")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyInputs")
	assert.Nil(err)

	options := module.NewOptions()
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
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyOutputs")
	assert.Nil(err)

	options := module.NewOptions()
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

	options, err := module.NewOptions().With(&module.Options{
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
