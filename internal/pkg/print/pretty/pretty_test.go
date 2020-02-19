package pretty

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().Build()

	expected, err := testutil.GetExpected("pretty")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-SortByName")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-SortByRequired")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-NoHeader")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-NoProviders")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-NoInputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("pretty-NoOutputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("pretty-OnlyHeader")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("pretty-OnlyProviders")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("pretty-OnlyInputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithColor().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-OnlyOutputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoColor(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowColor: false,
	}).Build()

	expected, err := testutil.GetExpected("pretty-NoColor")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("pretty-OutputValues")
	assert.Nil(err)

	options := &tfconf.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	}
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
