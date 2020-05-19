package format

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/segmentio/terraform-docs/pkg/print"
)

func TestFormatFactory(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
		wantErr  bool
	}{
		{
			name:     "format factory from name",
			format:   "asciidoc",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc document",
			expected: "*format.AsciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "asciidoc table",
			expected: "*format.AsciidocTable",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "json",
			expected: "*format.JSON",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown",
			expected: "*format.Table",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown document",
			expected: "*format.Document",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "markdown table",
			expected: "*format.Table",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "pretty",
			expected: "*format.Pretty",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "tfvars hcl",
			expected: "*format.TfvarsHCL",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "tfvars json",
			expected: "*format.TfvarsJSON",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "toml",
			expected: "*format.TOML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "xml",
			expected: "*format.XML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "yaml",
			expected: "*format.YAML",
			wantErr:  false,
		},
		{
			name:     "format factory from name",
			format:   "unknown",
			expected: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			settings := &print.Settings{}
			actual, err := Factory(tt.format, settings)
			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, reflect.TypeOf(actual).String())
			}
		})
	}
}
