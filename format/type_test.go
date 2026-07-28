// /*
// Copyright 2021 The terraform-docs Authors.

// Licensed under the MIT license (the "License"); you may not
// use this file except in compliance with the License.

// You may obtain a copy of the License at the LICENSE file in
// the root directory of this source tree.
// */

package format

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rquadling/terraform-docs/print"
)

func TestFormatType(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
		wantErr  bool
	}{
		{
			name:     "format type from name",
			format:   "asciidoc",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "adoc",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "asciidoc document",
			expected: "*format.asciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "asciidoc doc",
			expected: "*format.asciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "adoc document",
			expected: "*format.asciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "adoc doc",
			expected: "*format.asciidocDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "asciidoc table",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "asciidoc tbl",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "adoc table",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "adoc tbl",
			expected: "*format.asciidocTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "json",
			expected: "*format.json",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "markdown",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "md",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "markdown document",
			expected: "*format.markdownDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "markdown doc",
			expected: "*format.markdownDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "md document",
			expected: "*format.markdownDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "md doc",
			expected: "*format.markdownDocument",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "markdown table",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "markdown tbl",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "md table",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "md tbl",
			expected: "*format.markdownTable",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "pretty",
			expected: "*format.pretty",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "tfvars hcl",
			expected: "*format.tfvarsHCL",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "tfvars json",
			expected: "*format.tfvarsJSON",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "toml",
			expected: "*format.toml",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "xml",
			expected: "*format.xml",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "yaml",
			expected: "*format.yaml",
			wantErr:  false,
		},
		{
			name:     "format type from name",
			format:   "unknown",
			expected: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			config := print.DefaultConfig()
			config.Formatter = tt.format
			actual, err := New(config)
			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, reflect.TypeOf(actual).String())
			}
		})
	}
}
