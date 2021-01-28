/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package types

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	nils := List{nil}
	testPrimitive(t, []testprimitive{
		{
			name:   "value nil and type bool",
			values: nils,
			types:  "bool",
			expected: expected{
				typeName:   "bool",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type number",
			values: nils,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type list",
			values: nils,
			types:  "list",
			expected: expected{
				typeName:   "list",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type map",
			values: nils,
			types:  "map",
			expected: expected{
				typeName:   "map",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type string",
			values: nils,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type empty",
			values: nils,
			types:  "",
			expected: expected{
				typeName:   "any",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
	})
}

func TestNilLength(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "nil length",
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, new(Nil).Length())
		})
	}
}

func TestNilMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "nil marshal JSON",
			expected: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := new(Nil).MarshalJSON()

			assert.Nil(err)
			assert.Equal(tt.expected, string(actual))
		})
	}
}

func TestNilMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "nil marshal XML",
			expected: "<test xsi:nil=\"true\"></test>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "test"}}

			err := new(Nil).MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Equal(tt.expected, b.String())
		})
	}
}

func TestNilMarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
	}{
		{
			name:     "nil marshal YAML",
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := new(Nil).MarshalYAML()

			assert.Nil(err)
			assert.Equal(tt.expected, actual)
		})
	}
}
