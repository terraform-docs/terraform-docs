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

func TestString(t *testing.T) {
	values := List{"foo", "42", "false", "true"}
	testPrimitive(t, []testprimitive{
		{
			name:   "value not nil and type string",
			values: values,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.String",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.String",
				hasDefault: true,
			},
		},
	})
}

func TestStringLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected int
	}{
		{
			name:     "string length",
			value:    "foo",
			expected: 3,
		},
		{
			name:     "string length",
			value:    "42",
			expected: 2,
		},
		{
			name:     "string length",
			value:    "false",
			expected: 5,
		},
		{
			name:     "string length",
			value:    "true",
			expected: 4,
		},
		{
			name:     "string length",
			value:    "",
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, String(tt.value).Length())
		})
	}
}

func TestStringUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "string underlying",
			value: "foo",
		},
		{
			name:  "string underlying",
			value: "42",
		},
		{
			name:  "string underlying",
			value: "false",
		},
		{
			name:  "string underlying",
			value: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, String(tt.value).underlying())
		})
	}
}

func TestStringMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "string marshal JSON",
			value:    "foo",
			expected: "\"foo\"",
		},
		{
			name:     "string marshal JSON",
			value:    "lorem \"ipsum\" dolor",
			expected: "\"lorem \\\"ipsum\\\" dolor\"",
		},
		{
			name:     "string marshal JSON",
			value:    "lorem ipsum\ndolor",
			expected: "\"lorem ipsum\\ndolor\"",
		},
		{
			name:     "string marshal a regex",
			value:    "\\.<>[]{}_-",
			expected: "\"\\\\.<>[]{}_-\"",
		},
		{
			name:     "string marshal JSON",
			value:    "",
			expected: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := String(tt.value).MarshalJSON()

			assert.Nil(err)
			assert.Equal(tt.expected, string(actual))
		})
	}
}

func TestStringMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "string marshal XML",
			value:    "foo",
			expected: "<test>foo</test>",
		},
		{
			name:     "string marshal XML",
			value:    "lorem <\"ipsum\"> 'dolor'",
			expected: "<test>lorem &lt;&#34;ipsum&#34;&gt; &#39;dolor&#39;</test>",
		},
		{
			name:     "string marshal XML",
			value:    "",
			expected: "<test xsi:nil=\"true\"></test>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "test"}}

			err := String(tt.value).MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Equal(tt.expected, b.String())
		})
	}
}

func TestStringMarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected interface{}
	}{
		{
			name:     "string marshal YAML",
			value:    "foo",
			expected: "foo",
		},
		{
			name:     "string marshal YAML",
			value:    "lorem ipsum",
			expected: "lorem ipsum",
		},
		{
			name:     "string marshal YAML",
			value:    "",
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := String(tt.value).MarshalYAML()

			assert.Nil(err)
			assert.Equal(tt.expected, actual)
		})
	}
}
