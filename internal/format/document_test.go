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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-WithRequired")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-SortByName")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-SortByRequired")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-NoHeader")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-NoProviders")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-NoInputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-NoOutputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-OnlyHeader")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-OnlyProviders")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-OnlyInputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-OnlyOutputs")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-EscapeCharacters")
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		MarkdownIndent: 0,
	}).Build()

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-IndentationBellowAllowed")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-IndentationAboveAllowed")
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

	module, err := testutil.GetModule(settings)
	assert.Nil(err)

	expected, err := testutil.GetExpected("document", "document-IndentationOfFour")
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
