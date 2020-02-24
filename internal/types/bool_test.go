package types

import (
	"testing"
)

func TestBool(t *testing.T) {
	values := List{true, false}
	testPrimitive(t, []testprimitive{
		{
			name:   "value not nil and type bool",
			values: values,
			types:  "bool",
			expected: expected{
				typeName:   "bool",
				valueKind:  "types.Bool",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "bool",
				valueKind:  "types.Bool",
				hasDefault: true,
			},
		},
	})
}
