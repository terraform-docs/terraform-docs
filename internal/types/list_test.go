package types

import (
	"testing"
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
