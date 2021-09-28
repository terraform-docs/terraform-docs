/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCompatible(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"asciidoc document": {
			expected: true,
		},
		"asciidoc table": {
			expected: true,
		},
		"markdown document": {
			expected: true,
		},
		"markdown table": {
			expected: true,
		},
		"markdown": {
			expected: false,
		},
		"markdown-table": {
			expected: false,
		},
		"md": {
			expected: false,
		},
		"md tbl": {
			expected: false,
		},
		"md-tbl": {
			expected: false,
		},
		"json": {
			expected: false,
		},
		"yaml": {
			expected: false,
		},
		"xml": {
			expected: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			generator := NewGenerator(name, "")
			actual := generator.isCompatible()

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestExecuteTemplate(t *testing.T) {
	header := "this is the header"
	footer := "this is the footer"
	tests := map[string]struct {
		name     string
		content  string
		template string
		expected string
		wantErr  bool
	}{
		"Compatible without template": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "",
			expected: "this is the header\nthis is the footer",
			wantErr:  false,
		},
		"Compatible with template not empty section": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "{{ .Header }}",
			expected: "this is the header",
			wantErr:  false,
		},
		"Compatible with template empty section": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "{{ .Inputs }}",
			expected: "",
			wantErr:  false,
		},
		"Compatible with template and unknown section": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "{{ .Unknown }}",
			expected: "",
			wantErr:  true,
		},
		"Compatible with template include file": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "{{ include \"testdata/sample-file.txt\" }}",
			expected: "Sample file to be included.\n",
			wantErr:  false,
		},
		"Compatible with template include unknown file": {
			name:     "markdown table",
			content:  "this is the header\nthis is the footer",
			template: "{{ include \"file-not-found\" }}",
			expected: "",
			wantErr:  true,
		},
		"Incompatible without template": {
			name:     "yaml",
			content:  "header: \"this is the header\"\nfooter: \"this is the footer\"",
			template: "",
			expected: "header: \"this is the header\"\nfooter: \"this is the footer\"",
			wantErr:  false,
		},
		"Incompatible with template": {
			name:     "yaml",
			content:  "header: \"this is the header\"\nfooter: \"this is the footer\"",
			template: "{{ .Header }}",
			expected: "header: \"this is the header\"\nfooter: \"this is the footer\"",
			wantErr:  false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			generator := NewGenerator(tt.name, "")
			generator.content = tt.content
			generator.header = header
			generator.footer = footer

			actual, err := generator.ExecuteTemplate(tt.template)

			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, actual)
			}
		})
	}
}

func TestGeneratorFunc(t *testing.T) {
	text := "foo"
	tests := map[string]struct {
		fn     func(string) GenerateFunc
		actual func(*Generator) string
	}{
		"WithContent": {
			fn:     WithContent,
			actual: func(r *Generator) string { return r.content },
		},
		"WithHeader": {
			fn:     WithHeader,
			actual: func(r *Generator) string { return r.header },
		},
		"WithFooter": {
			fn:     WithFooter,
			actual: func(r *Generator) string { return r.footer },
		},
		"WithInputs": {
			fn:     WithInputs,
			actual: func(r *Generator) string { return r.inputs },
		},
		"WithModules": {
			fn:     WithModules,
			actual: func(r *Generator) string { return r.modules },
		},
		"WithOutputs": {
			fn:     WithOutputs,
			actual: func(r *Generator) string { return r.outputs },
		},
		"WithProviders": {
			fn:     WithProviders,
			actual: func(r *Generator) string { return r.providers },
		},
		"WithRequirements": {
			fn:     WithRequirements,
			actual: func(r *Generator) string { return r.requirements },
		},
		"WithResources": {
			fn:     WithResources,
			actual: func(r *Generator) string { return r.resources },
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			generator := NewGenerator(name, "", tt.fn(text))

			assert.Equal(text, tt.actual(generator))
		})
	}
}

func TestForEach(t *testing.T) {
	generator := NewGenerator("foo", "")
	generator.ForEach(func(name string) (string, error) {
		return name, nil
	})

	tests := map[string]struct {
		actual string
	}{
		"all":          {actual: generator.content},
		"header":       {actual: generator.header},
		"footer":       {actual: generator.footer},
		"inputs":       {actual: generator.inputs},
		"modules":      {actual: generator.modules},
		"outputs":      {actual: generator.outputs},
		"providers":    {actual: generator.providers},
		"requirements": {actual: generator.requirements},
		"resources":    {actual: generator.resources},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(name, tt.actual)
		})
	}
}
