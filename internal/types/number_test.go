package types

import (
	"testing"
)

func TestNumber(t *testing.T) {
	values := List{int(0), int8(42), int16(-1200), int32(1140085647), int64(8922336854775807), float32(13.75), float64(2317483.64)}
	testPrimitive(t, []testprimitive{
		{
			name:   "value not nil and type number",
			values: values,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "types.Number",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: values,
			types:  "",
			expected: expected{
				typeName:   "number",
				valueKind:  "types.Number",
				hasDefault: true,
			},
		},
	})
}
