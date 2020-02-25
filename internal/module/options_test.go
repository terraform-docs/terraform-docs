package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsOverrideWith(t *testing.T) {
	assert := assert.New(t)

	options := NewOptions()

	assert.Equal(options.Path, "")
	assert.Equal(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "")

	options.With(&Options{
		Path: "/path/to/foo",
	})

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "")

	options.With(&Options{
		OutputValues:     true,
		OutputValuesPath: "/path/to/output/values",
	})

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.OutputValues, true)
	assert.Equal(options.OutputValuesPath, "/path/to/output/values")

	options.With(&Options{
		Path:         "",
		OutputValues: false,
	})

	assert.NotEqual(options.Path, "")
	assert.NotEqual(options.OutputValues, false)
}
