package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	notNilValues := make([]interface{}, 0)
	notNilValues = append(notNilValues, int(0), int8(42), int16(-1200), int32(1140085647), int64(8922336854775807), float32(13.75), float64(2317483.64))

	nilValues := make([]interface{}, 0)
	nilValues = append(nilValues, nil)

	type expected struct {
		typeName   string
		valueKind  string
		hasDefault bool
	}
	tests := []struct {
		name     string
		values   []interface{}
		types    string
		expected expected
	}{
		{
			name:   "value not nil and type number",
			values: notNilValues,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "types.Number",
				hasDefault: true,
			},
		},
		{
			name:   "value not nil and type empty",
			values: notNilValues,
			types:  "",
			expected: expected{
				typeName:   "number",
				valueKind:  "types.Number",
				hasDefault: true,
			},
		},
		{
			name:   "value nil and type number",
			values: nilValues,
			types:  "number",
			expected: expected{
				typeName:   "number",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
		{
			name:   "value nil and type empty",
			values: nilValues,
			types:  "",
			expected: expected{
				typeName:   "any",
				valueKind:  "*types.Nil",
				hasDefault: false,
			},
		},
	}
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv)
				actualType := TypeOf(tt.types, tv)

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}
