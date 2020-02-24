package types

import (
	"testing"
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
