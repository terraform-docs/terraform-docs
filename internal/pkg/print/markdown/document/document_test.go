package document

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("document")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentWithRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowRequired:  true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
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
		SortByName:    true,
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
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
		SortByName:     true,
		SortByRequired: true,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("document-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoHeader(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("document-NoHeader")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("document-NoProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("document-NoInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("document-NoOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyHeader(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    true,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("document-OnlyHeader")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    false,
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("document-OnlyProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("document-OnlyInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowHeader:    false,
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("document-OnlyOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapeCharacters: true,
		ShowHeader:       true,
		ShowProviders:    true,
		ShowInputs:       true,
		ShowOutputs:      true,
	}

	module, expected, err := testutil.GetExpected("document-EscapeCharacters")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestDocumentIndentationBellowAllowed(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		MarkdownIndent: 0,
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
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
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
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
		ShowHeader:     true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("document-IndentationOfFour")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
