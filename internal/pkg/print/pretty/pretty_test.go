package pretty

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:    true,
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:     true,
		SortByRequired: true,
		ShowProviders:  true,
		ShowInputs:     true,
		ShowOutputs:    true,
		ShowColor:      true,
	}

	module, expected, err := testutil.GetExpected("pretty-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   true,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-NoProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   true,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-NoInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   false,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-NoOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyProviders(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    false,
		ShowOutputs:   false,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-OnlyProviders")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyInputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    true,
		ShowOutputs:   false,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-OnlyInputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyOnlyOutputs(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: false,
		ShowInputs:    false,
		ShowOutputs:   true,
		ShowColor:     true,
	}

	module, expected, err := testutil.GetExpected("pretty-OnlyOutputs")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettyNoColor(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		ShowProviders: true,
		ShowInputs:    true,
		ShowOutputs:   true,
		ShowColor:     false,
	}

	module, expected, err := testutil.GetExpected("pretty-NoColor")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
