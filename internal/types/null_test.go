package types

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	nulls := List{new(Null)}
	tests := []testprimitive{
		{
			name:   "value null and type bool",
			values: nulls,
			types:  "bool",
			expected: expected{
				typeName:   "bool",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
		{
			name:   "value null and type number",
			values: nulls,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
		{
			name:   "value null and type list",
			values: nulls,
			types:  "list",
			expected: expected{
				typeName:   "list",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
		{
			name:   "value null and type map",
			values: nulls,
			types:  "map",
			expected: expected{
				typeName:   "map",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
		{
			name:   "value null and type string",
			values: nulls,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
		{
			name:   "value null and type empty",
			values: nulls,
			types:  "",
			expected: expected{
				typeName:   "any",
				valueKind:  "*types.Null",
				hasDefault: true,
			},
		},
	}
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := new(Null)
				actualType := TypeOf(tt.types, tv)

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}

func TestNUllMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "null marshal JSON",
			expected: "\"null\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := new(Null).MarshalJSON()

			assert.Nil(err)
			assert.Equal(tt.expected, string(actual))
		})
	}
}

func TestNullMarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "null marshal XML",
			expected: "<test>null</test>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "test"}}

			err := new(Null).MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Equal(tt.expected, b.String())
		})
	}
}

func TestNullMarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
	}{
		{
			name:     "null marshal YAML",
			expected: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := new(Null).MarshalYAML()

			assert.Nil(err)
			assert.Equal(tt.expected, actual)
		})
	}
}
