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

func TestMap(t *testing.T) {
	values := []Map{
		{
			"a": 1,
			"b": 2,
			"c": 3,
		},
		{
			"name": "hello",
			"foo": map[string]string{
				"foo": "foo",
				"bar": "foo",
			},
			"bar": map[string]string{
				"foo": "bar",
				"bar": "bar",
			},
			"buzz": []string{"fizz", "buzz"},
		},
		{},
	}
	tests := []testmap{
		{
			name:   "value not nil and type map",
			values: values,
			types:  "map",
			expected: expected{
				typeName:   "map",
				valueKind:  "types.Map",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "map",
				valueKind:  "types.Map",
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

func TestMapLength(t *testing.T) {
	tests := []struct {
		name     string
		value    map[string]interface{}
		expected int
	}{
		{
			name: "map length",
			value: Map{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: 3,
		},
		{
			name: "map length",
			value: Map{
				"name": "hello",
				"foo": map[string]string{
					"foo": "foo",
					"bar": "foo",
				},
				"bar": map[string]string{
					"foo": "bar",
					"bar": "bar",
				},
				"buzz": []string{"fizz", "buzz"},
			},
			expected: 4,
		},
		{
			name:     "map length",
			value:    Map{},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, Map(tt.value).Length())
		})
	}
}

func TestMapUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value map[string]interface{}
	}{
		{
			name: "map underlying",
			value: Map{
				"a": 1,
				"b": 2,
				"c": 3,
			},
		},
		{
			name: "map underlying",
			value: Map{
				"name": "hello",
				"foo": map[string]string{
					"foo": "foo",
					"bar": "foo",
				},
				"bar": map[string]string{
					"foo": "bar",
					"bar": "bar",
				},
				"buzz": []string{"fizz", "buzz"},
			},
		},
		{
			name:  "map underlying",
			value: Map{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, Map(tt.value).Underlying())
		})
	}
}

func TestMapMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		value    map[string]interface{}
		expected string
	}{
		{
			name: "map marshal XML",
			value: Map{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: "<test><a>1</a><b>2</b><c>3</c></test>",
		},
		{
			name: "map marshal XML",
			value: Map{
				"name": "hello",
				"foo": Map{
					"foo": "foo",
					"bar": "foo",
				}.Underlying(),
				"bar": Map{
					"foo": "bar",
					"bar": "bar",
				}.Underlying(),
				"buzz": List{"fizz", "buzz"}.Underlying(),
			},
			expected: "<test><bar><bar>bar</bar><foo>bar</foo></bar><buzz><item>fizz</item><item>buzz</item></buzz><foo><bar>foo</bar><foo>foo</foo></foo><name>hello</name></test>",
		},
		{
			name:     "map marshal XML",
			value:    Map{},
			expected: "<test></test>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "test"}}

			err := Map(tt.value).MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Equal(tt.expected, b.String())
		})
	}
}
