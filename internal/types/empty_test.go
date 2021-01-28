/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	values := List{""}
	testPrimitive(t, []testprimitive{
		{
			name:   "value empty and type string",
			values: values,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
		{
			name:   "value empty and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
	})
}

func TestEmptyLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected int
	}{
		{
			name:     "empty length",
			value:    "",
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, Empty(tt.value).Length())
		})
	}
}

func TestEmptyUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "empty underlying",
			value: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, Empty(tt.value).underlying())
		})
	}
}

func TestEmptyMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "string marshal JSON",
			value:    "foo",
			expected: "\"\"",
		},
		{
			name:     "string marshal JSON",
			value:    "",
			expected: "\"\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := Empty(tt.value).MarshalJSON()

			assert.Nil(err)
			assert.Equal(tt.expected, string(actual))
		})
	}
}
