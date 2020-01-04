package pretty

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{}

	module, expected, err := testutil.GetExpexted("pretty")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByName(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName: true,
	}

	module, expected, err := testutil.GetExpexted("pretty-SortByName")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestPrettySortByRequired(t *testing.T) {
	assert := assert.New(t)
	settings := &print.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	module, expected, err := testutil.GetExpexted("pretty-SortByRequired")
	assert.Nil(err)

	actual, err := Print(module, settings)

	assert.Nil(err)
	assert.Equal(expected, actual)
}
