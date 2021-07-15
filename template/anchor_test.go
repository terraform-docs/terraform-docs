/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/print"
)

func TestAnchorMarkdown(t *testing.T) {
	tests := []struct {
		typeSection string
		name        string
		anchor      bool
		escape      bool
		expected    string
	}{
		{
			typeSection: "module",
			name:        "banana_anchor_escape",
			anchor:      true,
			escape:      true,
			expected:    "<a name=\"module_banana_anchor_escape\"></a> [banana\\_anchor\\_escape](#module\\_banana\\_anchor\\_escape)",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_noescape",
			anchor:      true,
			escape:      false,
			expected:    "<a name=\"module_banana_anchor_noescape\"></a> [banana_anchor_noescape](#module_banana_anchor_noescape)",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_escape",
			anchor:      false,
			escape:      true,
			expected:    "banana\\_anchor\\_escape",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_noescape",
			anchor:      false,
			escape:      false,
			expected:    "banana_anchor_noescape",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := &print.Settings{
				ShowAnchor:       tt.anchor,
				EscapeCharacters: tt.escape,
			}
			actual := createAnchorMarkdown(tt.typeSection, tt.name, settings)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestAnchorAsciidoc(t *testing.T) {
	tests := []struct {
		typeSection string
		name        string
		anchor      bool
		escape      bool
		expected    string
	}{
		{
			typeSection: "module",
			name:        "banana_anchor_escape",
			anchor:      true,
			escape:      true,
			expected:    "[[module\\_banana\\_anchor\\_escape]] <<module\\_banana\\_anchor\\_escape,banana\\_anchor\\_escape>>",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_noescape",
			anchor:      true,
			escape:      false,
			expected:    "[[module_banana_anchor_noescape]] <<module_banana_anchor_noescape,banana_anchor_noescape>>",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_escape",
			anchor:      false,
			escape:      true,
			expected:    "banana\\_anchor\\_escape",
		},
		{
			typeSection: "module",
			name:        "banana_anchor_noescape",
			anchor:      false,
			escape:      false,
			expected:    "banana_anchor_noescape",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := &print.Settings{
				ShowAnchor:       tt.anchor,
				EscapeCharacters: tt.escape,
			}
			actual := createAnchorAsciidoc(tt.typeSection, tt.name, settings)

			assert.Equal(tt.expected, actual)
		})
	}
}
