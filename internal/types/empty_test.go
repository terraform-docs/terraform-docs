package types

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	values := List{""}
	testPrimitive(t, []testprimitive{
		{
			name:   "value empty and type string",
			values: values,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
		{
			name:   "value empty and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "string",
				valueKind:  "types.Empty",
				hasDefault: true,
			},
		},
	})
}
