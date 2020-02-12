package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowRequired: true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-WithRequired")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-SortByName")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-SortByRequired")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-NoHeader")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-NoProviders")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-NoInputs")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-NoOutputs")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-OnlyHeader")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-OnlyProviders")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-OnlyInputs")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-OnlyOutputs")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		EscapeCharacters: true,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-EscapeCharacters")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 0,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-IndentationBellowAllowed")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 10,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-IndentationAboveAllowed")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 4,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("table", "table-IndentationOfFour")
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
