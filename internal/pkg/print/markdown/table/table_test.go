package table

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowRequired:  true,
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-WithRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:    true,
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:     true,
		SortByRequired: true,
		EscapePipe:     true,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("table-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-NoHeader")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-NoProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-NoInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("table-NoOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("table-OnlyHeader")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("table-OnlyProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("table-OnlyInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapePipe:    true,
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("table-OnlyOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapeCharacters: true,
		EscapePipe:       true,
		ShowHeader:       true,
		ShowProviders:    true,
		ShowInputs:       true,
		ShowOutputs:      true,
	}

	module, expected, err := testutil.GetExpected("table-EscapeCharacters")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 0,
		EscapePipe:     true,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("table-IndentationBellowAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 10,
		EscapePipe:     true,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("table-IndentationAboveAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 4,
		EscapePipe:     true,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("table-IndentationOfFour")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
