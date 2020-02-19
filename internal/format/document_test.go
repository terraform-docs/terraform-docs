package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("document")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("document-WithRequired")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("document-SortByName")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("document-SortByRequired")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("document-NoHeader")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("document-NoProviders")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("document-NoInputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("document-NoOutputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("document-OnlyHeader")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("document-OnlyProviders")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}).Build()

	expected, err := testutil.GetExpected("document-OnlyInputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}).Build()

	expected, err := testutil.GetExpected("document-OnlyOutputs")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		EscapeCharacters: true,
	}).Build()

	expected, err := testutil.GetExpected("document-EscapeCharacters")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationBelowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 0,
	}).Build()

	expected, err := testutil.GetExpected("document-IndentationBelowAllowed")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 10,
	}).Build()

	expected, err := testutil.GetExpected("document-IndentationAboveAllowed")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 4,
	}).Build()

	expected, err := testutil.GetExpected("document-IndentationOfFour")
	assert.Nil(err)

	module, err := testutil.GetModule(new(tfconf.Options))
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues: true,
	}).Build()

	expected, err := testutil.GetExpected("document-OutputValues")
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
