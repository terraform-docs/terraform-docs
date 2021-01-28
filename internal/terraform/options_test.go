/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsWith(t *testing.T) {
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

func TestOptionsWithNil(t *testing.T) {
	assert := assert.New(t)
	options := NewOptions()

	_, err := options.With(nil)

	assert.NotNil(err)
}

func TestOptionsWithOverwrite(t *testing.T) {
	assert := assert.New(t)

	options := NewOptions()

	assert.Equal(options.Path, "")
	assert.Equal(options.HeaderFromFile, "main.tf")
	assert.Equal(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "")

	_, err1 := options.With(&Options{
		Path: "/path/to/foo",
	})
	assert.Nil(err1)

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.HeaderFromFile, "main.tf")
	assert.Equal(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "")

	_, err2 := options.WithOverwrite(&Options{
		HeaderFromFile:   "doc.tf",
		OutputValues:     true,
		OutputValuesPath: "/path/to/output/values",
	})
	assert.Nil(err2)

	assert.Equal(options.Path, "/path/to/foo")
	assert.Equal(options.HeaderFromFile, "doc.tf")
	assert.Equal(options.OutputValues, true)
	assert.Equal(options.OutputValuesPath, "/path/to/output/values")

	_, err3 := options.WithOverwrite(&Options{
		Path:         "",
		OutputValues: false,
	})
	assert.Nil(err3)

	assert.NotEqual(options.Path, "")
	assert.Equal(options.HeaderFromFile, "doc.tf")
	assert.NotEqual(options.OutputValues, false)
	assert.Equal(options.OutputValuesPath, "/path/to/output/values")
}

func TestOptionsWithNilOverwrite(t *testing.T) {
	assert := assert.New(t)
	options := NewOptions()

	_, err := options.WithOverwrite(nil)

	assert.NotNil(err)
}
