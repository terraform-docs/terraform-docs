package types

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	values := []List{
		List{"foo", "bar", "baz"},
		List{"1", "2", "3"},
		List{true, false, true},
		List{10, float64(1000), int8(42)},
	}
	testList(t, []testlist{
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
	})
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
