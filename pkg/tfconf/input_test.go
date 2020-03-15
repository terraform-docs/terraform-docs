package tfconf

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/stretchr/testify/assert"
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
				Default:     new(types.Null),
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
			assert.Equal(tt.expectRequired, tt.input.Required())
		})
	}
}
