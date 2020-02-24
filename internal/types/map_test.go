package types

import (
	"testing"
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
	testMap(t, []testmap{
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
	})
}
