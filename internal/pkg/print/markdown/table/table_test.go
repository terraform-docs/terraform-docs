package table

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{}

	module, expected, err := testutil.GetExpexted("table")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowRequired: true,
	}

	module, expected, err := testutil.GetExpexted("table-WithRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName: true,
	}

	module, expected, err := testutil.GetExpexted("table-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	module, expected, err := testutil.GetExpexted("table-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableEscapeMarkdown(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapeMarkdown: true,
	}

	module, expected, err := testutil.GetExpexted("table-EscapeMarkdown")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 0,
	}

	module, expected, err := testutil.GetExpexted("table-IndentationBellowAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 10,
	}

	module, expected, err := testutil.GetExpexted("table-IndentationAboveAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestTableIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 4,
	}

	module, expected, err := testutil.GetExpexted("table-IndentationOfFour")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
