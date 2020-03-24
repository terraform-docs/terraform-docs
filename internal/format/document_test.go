package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("document", "document")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-WithRequired")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-SortByName")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-SortByRequired")
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

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-NoHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-NoInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-NoOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-NoProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-NoRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OnlyHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OnlyInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OnlyOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OnlyProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OnlyRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-EscapeCharacters")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-IndentationBelowAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-IndentationAboveAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("document", "document-IndentationOfFour")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: true,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OutputValues")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentHeaderFromFile(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("document", "document-HeaderFromFile")
	assert.Nil(err)

	options, err := module.NewOptions().WithOverwrite(&module.Options{
		HeaderFromFile: "doc.tf",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOutputValuesNoSensitivity(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: false,
	}).Build()

	expected, err := testutil.GetExpected("document", "document-OutputValuesNoSensitivity")
	assert.Nil(err)

	options := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
