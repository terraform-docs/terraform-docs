package json

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("json")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonSortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("json-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonSortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:     true,
		SortByRequired: true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
	}

	module, expected, err := testutil.GetExpected("json-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("json-NoProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("json-NoInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("json-NoOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("json-OnlyProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
	}

	module, expected, err := testutil.GetExpected("json-OnlyInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
	}

	module, expected, err := testutil.GetExpected("json-OnlyOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestJsonEscapeCharacters(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		EscapeCharacters: true,
		ShowProviders:    true,
		ShowInputs:       true,
		ShowOutputs:      true,
	}

	module, expected, err := testutil.GetExpected("json-EscapeCharacters")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
