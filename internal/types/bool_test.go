package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestBoolUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value bool
	}{
		{
			name:  "bool underlying",
			value: true,
		},
		{
			name:  "bool underlying",
			value: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, Bool(tt.value).underlying())
		})
	}
}
