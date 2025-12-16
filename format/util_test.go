/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		{
			name:     "preserve double spaces",
			markdown: "Lorem ipsum dolor sit amet,  \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,  \nconsectetur adipiscing elit",
		},
		{
			name:     "remove trailing space",
			markdown: "Lorem ipsum dolor sit amet, \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit",
		},
		{
			name:     "remove blank line with only doubl spaces",
			markdown: "Lorem ipsum dolor sit amet,\n  \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "sanitize link",
			markdown: "https://www.domain.com/",
			expected: "<https://www.domain.com/>",
		},
		{
			name:     "sanitize link in paragraph",
			markdown: "This is a domain https://www.domain.com/ inline in a paragraph of text",
			expected: "This is a domain <https://www.domain.com/> inline in a paragraph of text",
		},
		{
			name:     "sanitize link with valid format",
			markdown: "This is a valid link <https://www.domain.com/> already",
			expected: "This is a valid link <https://www.domain.com/> already",
		},
		{
			name:     "sanitize link with valid markdown format",
			markdown: "This is a valid [link](https://www.domain.com/) already",
			expected: "This is a valid [link](https://www.domain.com/) already",
		},
		{
			name:     "sanitize link with multiple occurrences",
			markdown: "Link 1: [link](https://www.domain.com/). Link 2: https://www.domain.com/",
			expected: "Link 1: [link](https://www.domain.com/). Link 2: <https://www.domain.com/>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := sanitize(tt.markdown)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestSanitizeBareLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sanitize link",
			input:    "https://www.domain.com/",
			expected: "<https://www.domain.com/>",
		},
		{
			name:     "sanitize link in paragraph",
			input:    "This is a domain https://www.domain.com/ inline in a paragraph of text",
			expected: "This is a domain <https://www.domain.com/> inline in a paragraph of text",
		},
		{
			name:     "sanitize link with valid format",
			input:    "This is a valid link <https://www.domain.com/> already",
			expected: "This is a valid link <https://www.domain.com/> already",
		},
		{
			name:     "sanitize link with valid markdown format",
			input:    "This is a valid [link](https://www.domain.com/) already",
			expected: "This is a valid [link](https://www.domain.com/) already",
		},
		{
			name:     "sanitize link with multiple occurrences",
			input:    "Link 1: [link](https://www.domain.com/). Link 2: https://www.domain.com/",
			expected: "Link 1: [link](https://www.domain.com/). Link 2: <https://www.domain.com/>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := SanitizeBareLinks(tt.input)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestFenceCodeBlock(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		language  string
		expected  string
		extraline bool
	}{
		{
			name:      "single line",
			code:      "foo",
			language:  "json",
			expected:  "`foo`",
			extraline: false,
		},
		{
			name:      "single line",
			code:      "\"bar\"",
			language:  "hcl",
			expected:  "`\"bar\"`",
			extraline: false,
		},
		{
			name:      "single line",
			code:      "fuzz_buzz",
			language:  "",
			expected:  "`fuzz_buzz`",
			extraline: false,
		},
		{
			name:      "multi lines",
			code:      "[\n  \"foo\",\n  \"bar\",\n  \"baz\"\n]",
			language:  "json",
			expected:  "\n\n```json\n[\n  \"foo\",\n  \"bar\",\n  \"baz\"\n]\n```\n",
			extraline: true,
		},
		{
			name:      "multi lines",
			code:      "variable \"foo\" {\n  default = true\n}",
			language:  "hcl",
			expected:  "\n\n```hcl\nvariable \"foo\" {\n  default = true\n}\n```\n",
			extraline: true,
		},
		{
			name:      "multi lines",
			code:      "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2",
			language:  "",
			expected:  "\n\n```\nUsage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n```\n",
			extraline: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, extraline := PrintFencedCodeBlock(tt.code, tt.language)

			assert.Equal(tt.expected, actual)
			assert.Equal(tt.extraline, extraline)
		})
	}
}
