package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestYaml(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-SortByName")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-SortByRequired")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-NoHeader")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-NoProviders")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-NoInputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-NoOutputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyHeader")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyProviders")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyInputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("yaml", "yaml-OnlyOutputs")
	assert.Nil(err)

	printer := NewYAML(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
