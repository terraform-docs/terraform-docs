package types

import (
	"testing"
)

func TestNil(t *testing.T) {
	nils := List{nil}
	testPrimitive(t, []testprimitive{
		{
			name:   "value nil and type bool",
			values: nils,
			types:  "bool",
			expected: expected{
				typeName:   "bool",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type number",
			values: nils,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type list",
			values: nils,
			types:  "list",
			expected: expected{
				typeName:   "list",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type map",
			values: nil,
			types:  "map",
			expected: expected{
				typeName:   "map",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type string",
			values: nils,
			types:  "string",
			expected: expected{
				typeName:   "string",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type empty",
			values: nils,
			types:  "",
			expected: expected{
				typeName:   "any",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
	})
}
