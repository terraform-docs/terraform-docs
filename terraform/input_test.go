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

func TestInputValue(t *testing.T) {
	inputName := "input"
	inputType := types.String("type")
	inputDescr := types.String("description")
	inputPos := Position{Filename: "foo.tf", Line: 13}

	tests := []struct {
		name           string
		input          Input
		expectValue    string
		expectDefault  bool
		expectRequired bool
	}{
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(nil),
				Required:    true,
				Position:    inputPos,
			},
			expectValue:    "",
			expectDefault:  false,
			expectRequired: true,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(nil),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "null",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(true),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "true",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(false),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "false",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(""),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "\"\"",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf("foo"),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "\"foo\"",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(42),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "42",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(13.75),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "13.75",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(types.List{"a", "b", "c"}.Underlying()),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "[\n  \"a\",\n  \"b\",\n  \"c\"\n]",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(types.List{}.Underlying()),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "[]",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(types.Map{"a": 1, "b": 2, "c": 3}.Underlying()),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "{\n  \"a\": 1,\n  \"b\": 2,\n  \"c\": 3\n}",
			expectDefault:  true,
			expectRequired: false,
		},
		{
			name: "input Value and HasDefault",
			input: Input{
				Name:        inputName,
				Type:        inputType,
				Description: inputDescr,
				Default:     types.ValueOf(types.Map{}.Underlying()),
				Required:    false,
				Position:    inputPos,
			},
			expectValue:    "{}",
			expectDefault:  true,
			expectRequired: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expectValue, tt.input.GetValue())
			assert.Equal(tt.expectDefault, tt.input.HasDefault())
		})
	}
}

func TestInputsSorted(t *testing.T) {
	inputs := sampleInputs()
	tests := map[string]struct {
		sortType func([]*Input)
		expected []string
	}{
		"ByName": {
			sortType: sortInputsByName,
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		"ByRequired": {
			sortType: sortInputsByRequired,
			expected: []string{"b", "d", "a", "c", "e", "f"},
		},
		"ByPosition": {
			sortType: sortInputsByPosition,
			expected: []string{"a", "d", "e", "b", "c", "f"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			tt.sortType(inputs)

			actual := make([]string, len(inputs))

			for k, i := range inputs {
				actual[k] = i.Name
			}

			assert.Equal(tt.expected, actual)
		})
	}
}

func sampleInputs() []*Input {
	return []*Input{
		{
			Name:        "e",
			Type:        types.String(""),
			Description: types.String("description of e"),
			Default:     types.ValueOf(true),
			Required:    false,
			Position:    Position{Filename: "foo/variables.tf", Line: 35},
		},
		{
			Name:        "a",
			Type:        types.String("string"),
			Description: types.String(""),
			Default:     types.ValueOf("a"),
			Required:    false,
			Position:    Position{Filename: "foo/variables.tf", Line: 10},
		},
		{
			Name:        "d",
			Type:        types.String("string"),
			Description: types.String("description for d"),
			Default:     types.ValueOf(nil),
			Required:    true,
			Position:    Position{Filename: "foo/variables.tf", Line: 23},
		},
		{
			Name:        "b",
			Type:        types.String("number"),
			Description: types.String("description of b"),
			Default:     types.ValueOf(nil),
			Required:    true,
			Position:    Position{Filename: "foo/variables.tf", Line: 42},
		},
		{
			Name:        "c",
			Type:        types.String("list"),
			Description: types.String("description of c"),
			Default:     types.ValueOf("c"),
			Required:    false,
			Position:    Position{Filename: "foo/variables.tf", Line: 51},
		},
		{
			Name:        "f",
			Type:        types.String("string"),
			Description: types.String("description of f"),
			Default:     types.ValueOf(nil),
			Required:    false,
			Position:    Position{Filename: "foo/variables.tf", Line: 59},
		},
	}
}
