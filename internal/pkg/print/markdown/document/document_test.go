package document

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{}

	module, expected, err := testutil.GetExpected("document")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowRequired: true,
	}

	module, expected, err := testutil.GetExpected("document-WithRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName: true,
	}

	module, expected, err := testutil.GetExpected("document-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	module, expected, err := testutil.GetExpected("document-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentEscapeMarkdown(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapeMarkdown: true,
	}

	module, expected, err := testutil.GetExpected("document-EscapeMarkdown")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 0,
	}

	module, expected, err := testutil.GetExpected("document-IndentationBellowAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationAboveAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 10,
	}

	module, expected, err := testutil.GetExpected("document-IndentationAboveAllowed")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationOfFour(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 4,
	}

	module, expected, err := testutil.GetExpected("document-IndentationOfFour")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
