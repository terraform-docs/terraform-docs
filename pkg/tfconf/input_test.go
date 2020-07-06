package tfconf

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/types"
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
