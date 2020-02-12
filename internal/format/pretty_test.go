package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().WithColor().Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-SortByName")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-SortByRequired")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-NoHeader")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-NoProviders")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-NoInputs")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-NoOutputs")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyHeader")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyProviders")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyInputs")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-OnlyOutputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("pretty", "pretty-NoColor")
	assert.Nil(err)

	printer := NewPretty(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
