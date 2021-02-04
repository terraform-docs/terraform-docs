/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestNumberLength(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected int
	}{
		{
			name:     "number length",
			value:    float64(int(0)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(int8(42)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(int16(-1200)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(int32(1140085647)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(int64(8922336854775807)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(float32(13.75)),
			expected: 0,
		},
		{
			name:     "number length",
			value:    float64(2317483.64),
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, Number(tt.value).Length())
		})
	}
}

func TestNumberUnderlying(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{
			name:  "number underlying",
			value: float64(int(0)),
		},
		{
			name:  "number underlying",
			value: float64(int8(42)),
		},
		{
			name:  "number underlying",
			value: float64(int16(-1200)),
		},
		{
			name:  "number underlying",
			value: float64(int32(1140085647)),
		},
		{
			name:  "number underlying",
			value: float64(int64(8922336854775807)),
		},
		{
			name:  "number underlying",
			value: float64(float32(13.75)),
		},
		{
			name:  "number underlying",
			value: float64(2317483.64),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.value, Number(tt.value).underlying())
		})
	}
}
