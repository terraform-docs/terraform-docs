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

			generator := NewGenerator(name)
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

			generator := NewGenerator(tt.name)
			generator.content = tt.content
			generator.Header = header
			generator.Footer = footer

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
			actual: func(g *Generator) string { return g.content },
		},
		"WithHeader": {
			fn:     WithHeader,
			actual: func(g *Generator) string { return g.Header },
		},
		"WithFooter": {
			fn:     WithFooter,
			actual: func(g *Generator) string { return g.Footer },
		},
		"WithInputs": {
			fn:     WithInputs,
			actual: func(g *Generator) string { return g.Inputs },
		},
		"WithModules": {
			fn:     WithModules,
			actual: func(g *Generator) string { return g.Modules },
		},
		"WithOutputs": {
			fn:     WithOutputs,
			actual: func(g *Generator) string { return g.Outputs },
		},
		"WithProviders": {
			fn:     WithProviders,
			actual: func(g *Generator) string { return g.Providers },
		},
		"WithRequirements": {
			fn:     WithRequirements,
			actual: func(g *Generator) string { return g.Requirements },
		},
		"WithResources": {
			fn:     WithResources,
			actual: func(g *Generator) string { return g.Resources },
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			generator := NewGenerator(name, tt.fn(text))

			assert.Equal(text, tt.actual(generator))
		})
	}
}

func TestForEach(t *testing.T) {
	// text := "foo"
	fns := []GenerateFunc{}
	ForEach(func(name string, fn GeneratorCallback) error {
		fns = append(fns, fn(name))
		return nil
	})

	generator := NewGenerator("foo", fns...)

	tests := map[string]struct {
		actual string
	}{
		"all":          {actual: generator.content},
		"header":       {actual: generator.Header},
		"footer":       {actual: generator.Footer},
		"inputs":       {actual: generator.Inputs},
		"modules":      {actual: generator.Modules},
		"outputs":      {actual: generator.Outputs},
		"providers":    {actual: generator.Providers},
		"requirements": {actual: generator.Requirements},
		"resources":    {actual: generator.Resources},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(name, tt.actual)
		})
	}
}
