package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/module"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
	"github.com/terraform-docs/terraform-docs/pkg/print"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().Build()

	expected, err := testutil.GetExpected("markdown", "table")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("markdown", "table-WithRequired")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("markdown", "table-SortByName")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
			Name: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("markdown", "table-SortByRequired")
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

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByType(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		SortByType: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-SortByType")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		SortBy: &module.SortBy{
			Type: true,
		},
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-NoRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       true,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyHeader")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       true,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyInputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      true,
		ShowProviders:    false,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyOutputs")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    true,
		ShowRequirements: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyProviders")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyRequirements(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().With(&print.Settings{
		ShowHeader:       false,
		ShowInputs:       false,
		ShowOutputs:      false,
		ShowProviders:    false,
		ShowRequirements: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OnlyRequirements")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
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

	expected, err := testutil.GetExpected("markdown", "table-EscapeCharacters")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationBelowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 0,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationBelowAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 10,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationAboveAllowed")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		IndentLevel: 4,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-IndentationOfFour")
	assert.Nil(err)

	options := module.NewOptions()
	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOutputValues(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: true,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OutputValues")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableHeaderFromFile(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		file   string
	}{
		{
			name:   "load module header from .adoc",
			golden: "table-HeaderFromADOCFile",
			file:   "doc.adoc",
		},
		{
			name:   "load module header from .md",
			golden: "table-HeaderFromMDFile",
			file:   "doc.md",
		},
		{
			name:   "load module header from .tf",
			golden: "table-HeaderFromTFFile",
			file:   "doc.tf",
		},
		{
			name:   "load module header from .txt",
			golden: "table-HeaderFromTXTFile",
			file:   "doc.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := testutil.Settings().WithSections().Build()

			expected, err := testutil.GetExpected("markdown", tt.golden)
			assert.Nil(err)

			options, err := module.NewOptions().WithOverwrite(&module.Options{
				HeaderFromFile: tt.file,
			})
			assert.Nil(err)

			module, err := testutil.GetModule(options)
			assert.Nil(err)

			printer := NewTable(settings)
			actual, err := printer.Print(module, settings)

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}

func TestTableOutputValuesNoSensitivity(t *testing.T) {
	assert := assert.New(t)
	settings := testutil.Settings().WithSections().With(&print.Settings{
		OutputValues:    true,
		ShowSensitivity: false,
	}).Build()

	expected, err := testutil.GetExpected("markdown", "table-OutputValuesNoSensitivity")
	assert.Nil(err)

	options, err := module.NewOptions().With(&module.Options{
		OutputValues:     true,
		OutputValuesPath: "output_values.json",
	})
	assert.Nil(err)

	module, err := testutil.GetModule(options)
	assert.Nil(err)

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEmpty(t *testing.T) {
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

	printer := NewTable(settings)
	actual, err := printer.Print(module, settings)

	assert.Nil(err)
	assert.Equal("", actual)
}
