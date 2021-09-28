/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	list := []string{"foo", "bar", "buzz"}
	tests := []struct {
		name     string
		item     string
		expected bool
	}{
		{
			name:     "item exists in slice",
			item:     "foo",
			expected: true,
		},
		{
			name:     "item exists in slice",
			item:     "bar",
			expected: true,
		},
		{
			name:     "item not exist in slice",
			item:     "fizz",
			expected: false,
		},
		{
			name:     "empty item",
			item:     "",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := contains(list, tt.item)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestSliceIndex(t *testing.T) {
	list := []string{"foo", "bar", "buzz"}
	tests := []struct {
		name     string
		item     string
		expected int
	}{
		{
			name:     "index of item exists in slice",
			item:     "foo",
			expected: 0,
		},
		{
			name:     "index of item exists in slice",
			item:     "bar",
			expected: 1,
		},
		{
			name:     "index of item not exist in slice",
			item:     "fizz",
			expected: -1,
		},
		{
			name:     "index of empty item",
			item:     "",
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := index(list, tt.item)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestSliceRemove(t *testing.T) {
	list := []string{"foo", "bar", "buzz"}
	tests := []struct {
		name     string
		item     string
		expected []string
	}{
		{
			name:     "remove item exists in slice",
			item:     "foo",
			expected: []string{"buzz", "bar"},
		},
		{
			name:     "remove item exists in slice",
			item:     "bar",
			expected: []string{"foo", "buzz"},
		},
		{
			name:     "remove item not exist in slice",
			item:     "fizz",
			expected: []string{"foo", "bar", "buzz"},
		},
		{
			name:     "remove empty item",
			item:     "",
			expected: []string{"foo", "bar", "buzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			cpy := make([]string, len(list))
			copy(cpy, list)
			actual := remove(cpy, tt.item)
			assert.Equal(len(tt.expected), len(actual))
			assert.Equal(tt.expected, actual)
		})
	}
}
