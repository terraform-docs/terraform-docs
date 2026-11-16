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

	"github.com/rquadling/terraform-docs/internal/types"
)

func TestProviderName(t *testing.T) {
	tests := map[string]struct {
		provider Provider
		expected string
	}{
		"WithoutAlias": {
			provider: Provider{
				Name:     "provider",
				Alias:    types.String(""),
				Version:  types.String(">= 1.2.3"),
				Position: Position{Filename: "foo.tf", Line: 13},
			},
			expected: "provider",
		},
		"WithAlias": {
			provider: Provider{
				Name:     "provider",
				Alias:    types.String("alias"),
				Version:  types.String(">= 1.2.3"),
				Position: Position{Filename: "foo.tf", Line: 13},
			},
			expected: "provider.alias",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.expected, tt.provider.FullName())
		})
	}
}

func TestProvidersSort(t *testing.T) {
	providers := sampleProviders()
	tests := map[string]struct {
		sortType func([]*Provider)
		expected []string
	}{
		"ByName": {
			sortType: sortProvidersByName,
			expected: []string{"a", "b", "c", "d", "d.a", "e", "e.a"},
		},
		"ByPosition": {
			sortType: sortProvidersByPosition,
			expected: []string{"e.a", "b", "d", "d.a", "a", "e", "c"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			tt.sortType(providers)

			actual := make([]string, len(providers))

			for k, p := range providers {
				actual[k] = p.FullName()
			}

			assert.Equal(tt.expected, actual)
		})
	}
}

func sampleProviders() []*Provider {
	return []*Provider{
		{
			Name:     "d",
			Alias:    types.String(""),
			Version:  types.String("1.3.2"),
			Position: Position{Filename: "foo/main.tf", Line: 21},
		},
		{
			Name:     "d",
			Alias:    types.String("a"),
			Version:  types.String("> 1.x"),
			Position: Position{Filename: "foo/main.tf", Line: 25},
		},
		{
			Name:     "b",
			Alias:    types.String(""),
			Version:  types.String("= 2.1.0"),
			Position: Position{Filename: "foo/main.tf", Line: 13},
		},
		{
			Name:     "a",
			Alias:    types.String(""),
			Version:  types.String(""),
			Position: Position{Filename: "foo/main.tf", Line: 39},
		},
		{
			Name:     "c",
			Alias:    types.String(""),
			Version:  types.String("~> 0.5.0"),
			Position: Position{Filename: "foo/main.tf", Line: 53},
		},
		{
			Name:     "e",
			Alias:    types.String(""),
			Version:  types.String(""),
			Position: Position{Filename: "foo/main.tf", Line: 47},
		},
		{
			Name:     "e",
			Alias:    types.String("a"),
			Version:  types.String("> 1.0"),
			Position: Position{Filename: "foo/main.tf", Line: 5},
		},
	}
}
