package format

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/testutil"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/stretchr/testify/assert"
)

func TestAsciidocDocument(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("asciidoc", "document")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		ShowRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-WithRequired")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-SortByName")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByName:     true,
		SortByRequired: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-SortByRequired")
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

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-SortByType")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-NoHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-NoInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-NoOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-NoProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-NoRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OnlyHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OnlyInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OnlyOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OnlyProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OnlyRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentIndentationBelowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 0,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-IndentationBelowAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 10,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-IndentationAboveAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 4,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-IndentationOfFour")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: true,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OutputValues")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentHeaderFromTFFile(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("asciidoc", "document-HeaderFromTFFile")
	assert.Nil(err)

	options, err := module.NewOptions().WithOverwrite(&module.Options{
		HeaderFromFile: "doc.tf",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentHeaderFromMDFile(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("asciidoc", "document-HeaderFromMDFile")
	assert.Nil(err)

	options, err := module.NewOptions().WithOverwrite(&module.Options{
		HeaderFromFile: "doc.md",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentOutputValuesNoSensitivity(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: false,
	}).Build()

	expected, err := testutil.GetExpected("asciidoc", "document-OutputValuesNoSensitivity")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestAsciidocDocumentEmpty(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}).Build()

	options, err := module.NewOptions().WithOverwrite(&module.Options{
		HeaderFromFile: "bad.tf",
	})
	options.ShowHeader = false // Since we don't show the header, the file won't be loaded at all
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewAsciidocDocument(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal("", actual)
}
