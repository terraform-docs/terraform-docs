package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandAnnotations(t *testing.T) {
	tests := []struct {
		name    string
		command string
	}{
		{
			name:    "command annotations",
			command: "foo",
		},
		{
			name:    "command annotations",
			command: "foo bar",
		},
		{
			name:    "command annotations",
			command: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Annotations(tt.command)
			assert.Equal(tt.command, actual["command"])
			assert.Equal("formatter", actual["kind"])
		})
	}
}
