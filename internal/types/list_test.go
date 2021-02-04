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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	values := []List{
		{"foo", "bar", "baz"},
		{"1", "2", "3"},
		{true, false, true},
		{10, float64(1000), int8(42)},
	}
	tests := []testlist{
		{
			name:   "value not nil and type list",
			values: values,
			types:  "list",
			expected: expected{
				typeName:   "list",
				valueKind:  "types.List",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "list",
				valueKind:  "types.List",
				hasDefault: true,
			},
		},
	}
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv.Underlying())
				actualType := TypeOf(tt.types, tv.Underlying())

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}

func TestListLength(t *testing.T) {
	tests := []struct {
		name     string
		value    []interface{}
		expected int
	}{
		{
			name:     "list length",
			value:    List{"foo", "bar", "baz"},
			expected: 3,
		},
		{
			name:     "list length",
			value:    List{"1", "2", "3"},
			expected: 3,
		},
		{
			name:     "list length",
			value:    List{true, false, true},
			expected: 3,
		},
		{
			name:     "list length",
			value:    List{10, float64(1000), int8(42)},
			expected: 3,
		},
		{
			name:     "list length",
			value:    List{},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, List(tt.value).Length())
		})
	}
}

func TestListUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value []interface{}
	}{
		{
			name:  "list underlying",
			value: List{"foo", "bar", "baz"},
		},
		{
			name:  "list underlying",
			value: List{"1", "2", "3"},
		},
		{
			name:  "list underlying",
			value: List{true, false, true},
		},
		{
			name:  "list underlying",
			value: List{10, float64(1000), int8(42)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, List(tt.value).Underlying())
		})
	}
}

func TestListMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		value    []interface{}
		expected string
	}{
		{
			name:     "list marshal XML",
			value:    List{"foo", "bar", "baz"},
			expected: "<test><item>foo</item><item>bar</item><item>baz</item></test>",
		},
		{
			name:     "list marshal XML",
			value:    List{"1", "2", "3"},
			expected: "<test><item>1</item><item>2</item><item>3</item></test>",
		},
		{
			name:     "list marshal XML",
			value:    List{true, false, true},
			expected: "<test><item>true</item><item>false</item><item>true</item></test>",
		},
		{
			name:     "list marshal XML",
			value:    List{10, float64(1000), int8(42)},
			expected: "<test><item>10</item><item>1000</item><item>42</item></test>",
		},
		{
			name:     "list marshal XML",
			value:    List{},
			expected: "<test></test>",
		},
		{
			name:     "list marshal XML",
			value:    List{nil, "something"},
			expected: "<test><item></item><item>something</item></test>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "test"}}

			err := List(tt.value).MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Equal(tt.expected, b.String())
		})
	}
}
