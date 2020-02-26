package tfconf

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestOutputValue(t *testing.T) {
	outputName := "output"
	outputDescr := types.String("description")
	outputPos := Position{Filename: "foo.tf", Line: 13}

	tests := []struct {
		name          string
		output        Output
		expectValue   string
		expectDefault bool
	}{
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(nil),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "",
			expectDefault: false,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(true),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "true",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(false),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "false",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(""),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "\"\"",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf("foo"),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "\"foo\"",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(42),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "42",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(13.75),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "13.75",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(types.List{"a", "b", "c"}.Underlying()),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "[\n  \"a\",\n  \"b\",\n  \"c\"\n]",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(types.List{}.Underlying()),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "[]",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(types.Map{"a": 1, "b": 2, "c": 3}.Underlying()),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "{\n  \"a\": 1,\n  \"b\": 2,\n  \"c\": 3\n}",
			expectDefault: true,
		},
		{
			name: "output Value and HasDefault",
			output: Output{
				Name:        outputName,
				Description: outputDescr,
				Value:       types.ValueOf(types.Map{}.Underlying()),
				Sensitive:   false,
				Position:    outputPos,
				ShowValue:   false,
			},
			expectValue:   "{}",
			expectDefault: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expectValue, tt.output.GetValue())
			assert.Equal(tt.expectDefault, tt.output.HasDefault())
		})
	}
}
