/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulecallName(t *testing.T) {
	tests := map[string]struct {
		module   ModuleCall
		expected string
	}{
		"WithoutVersion": {
			module: ModuleCall{
				Name:   "provider",
				Source: "bar",
			},
			expected: "bar",
		},
		"WithVersion": {
			module: ModuleCall{
				Name:    "provider",
				Source:  "bar",
				Version: "1.2.3",
			},
			expected: "bar,1.2.3",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, tt.module.FullName())
		})
	}
}

func TestModulecallSort(t *testing.T) {
	modules := sampleModulecalls()
	tests := map[string]struct {
		sortType func([]*ModuleCall)
		expected []string
	}{
		"ByName": {
			sortType: sortModulecallsByName,
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		"BySource": {
			sortType: sortModulecallsBySource,
			expected: []string{"f", "d", "c", "e", "a", "b"},
		},
		"ByPosition": {
			sortType: sortModulecallsByPosition,
			expected: []string{"b", "c", "a", "e", "d", "f"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			tt.sortType(modules)

			actual := make([]string, len(modules))

			for k, i := range modules {
				actual[k] = i.Name
			}

			assert.Equal(tt.expected, actual)
		})
	}
}

func sampleModulecalls() []*ModuleCall {
	return []*ModuleCall{
		{
			Name:     "a",
			Source:   "z",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 35},
		},
		{
			Name:     "b",
			Source:   "z",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 10},
		},
		{
			Name:     "c",
			Source:   "m",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 23},
		},
		{
			Name:     "e",
			Source:   "x",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 42},
		},
		{
			Name:     "d",
			Source:   "l",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 51},
		},
		{
			Name:     "f",
			Source:   "a",
			Version:  "1.2.3",
			Position: Position{Filename: "foo/main.tf", Line: 59},
		},
	}
}
