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

	_, err1 := options.With(&Options{
		Path: "/path/to/foo",
	})
	assert.Nil(err1)

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "")

	_, err2 := options.With(&Options{
		OutputValues:     true,
		OutputValuesPath: "/path/to/output/values",
	})
	assert.Nil(err2)

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.OutputValues, true)
	assert.Equal(options.OutputValuesPath, "/path/to/output/values")

	_, err3 := options.With(&Options{
		Path:         "",
		OutputValues: false,
	})
	assert.Nil(err3)

	assert.NotEqual(options.Path, "")
	assert.NotEqual(options.OutputValues, false)
}

func TestOptionsOverrideWithNil(t *testing.T) {
	assert := assert.New(t)
	options := NewOptions()

	_, err := options.With(nil)

	assert.NotNil(err)
}
