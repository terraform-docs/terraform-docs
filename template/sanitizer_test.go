/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		escape   bool
		expected string
	}{
		{
			name:     "sanitize name with escape character",
			input:    "abcdefgh",
			escape:   true,
			expected: "abcdefgh",
		},
		{
			name:     "sanitize name with escape character",
			input:    "abcd_efgh",
			escape:   true,
			expected: "abcd\\_efgh",
		},
		{
			name:     "sanitize name with escape character",
			input:    "_abcdefgh",
			escape:   true,
			expected: "\\_abcdefgh",
		},
		{
			name:     "sanitize name with escape character",
			input:    "abcd__efgh",
			escape:   true,
			expected: "abcd\\_\\_efgh",
		},
		{
			name:     "sanitize name with escape character",
			input:    "_",
			escape:   true,
			expected: "\\_",
		},
		{
			name:     "sanitize name with escape character",
			input:    "",
			escape:   true,
			expected: "",
		},
		{
			name:     "sanitize name without escape character",
			input:    "abcdefgh",
			escape:   false,
			expected: "abcdefgh",
		},
		{
			name:     "sanitize name without escape character",
			input:    "abcd_efgh",
			escape:   false,
			expected: "abcd_efgh",
		},
		{
			name:     "sanitize name without escape character",
			input:    "abcd__efgh",
			escape:   false,
			expected: "abcd__efgh",
		},
		{
			name:     "sanitize name without escape character",
			input:    "_",
			escape:   false,
			expected: "_",
		},
		{
			name:     "sanitize name without escape character",
			input:    "",
			escape:   false,
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := SanitizeName(tt.input, tt.escape)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestSanitizeSection(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		escape   bool
	}{
		{
			name:     "sanitize section item empty",
			filename: "empty",
			escape:   true,
		},
		{
			name:     "sanitize section item complex",
			filename: "complex",
			escape:   true,
		},
		{
			name:     "sanitize section item codeblock",
			filename: "codeblock",
			escape:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			bytes, err := os.ReadFile(filepath.Join("testdata", "section", tt.filename+".golden"))
			assert.Nil(err)

			actual := SanitizeSection(string(bytes), tt.escape, false)

			expected, err := os.ReadFile(filepath.Join("testdata", "section", tt.filename+".expected"))
			assert.Nil(err)

			assert.Equal(string(expected), actual)
		})
	}
}

func TestSanitizeDocument(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		escape   bool
	}{
		{
			name:     "sanitize document item empty",
			filename: "empty",
			escape:   true,
		},
		{
			name:     "sanitize document item complex",
			filename: "complex",
			escape:   true,
		},
		{
			name:     "sanitize document item codeblock",
			filename: "codeblock",
			escape:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			bytes, err := os.ReadFile(filepath.Join("testdata", "document", tt.filename+".golden"))
			assert.Nil(err)

			actual := SanitizeDocument(string(bytes), tt.escape, false)

			expected, err := os.ReadFile(filepath.Join("testdata", "document", tt.filename+".expected"))
			assert.Nil(err)

			assert.Equal(string(expected), actual)
		})
	}
}

func TestSanitizeMarkdownTable(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected string
		html     bool
		escape   bool
	}{
		{
			name:     "sanitize table item empty",
			filename: "empty",
			expected: "empty",
			html:     true,
			escape:   true,
		},
		{
			name:     "sanitize table item complex with html",
			filename: "complex",
			expected: "complex-html",
			html:     true,
			escape:   true,
		},
		{
			name:     "sanitize table item complex without html",
			filename: "complex",
			expected: "complex-nohtml",
			html:     false,
			escape:   true,
		},
		{
			name:     "sanitize table item codeblock with html",
			filename: "codeblock",
			expected: "codeblock-html",
			html:     true,
			escape:   true,
		},
		{
			name:     "sanitize table item codeblock without html",
			filename: "codeblock",
			expected: "codeblock-nohtml",
			html:     false,
			escape:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			bytes, err := os.ReadFile(filepath.Join("testdata", "table", tt.filename+".golden"))
			assert.Nil(err)

			actual := SanitizeMarkdownTable(string(bytes), tt.escape, tt.html)

			expected, err := os.ReadFile(filepath.Join("testdata", "table", tt.expected+".markdown.expected"))
			assert.Nil(err)

			assert.Equal(string(expected), actual)
		})
	}
}

func TestSanitizeAsciidocTable(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		escape   bool
	}{
		{
			name:     "sanitize table item empty",
			filename: "empty",
			escape:   false,
		},
		{
			name:     "sanitize table item complex",
			filename: "complex",
			escape:   false,
		},
		{
			name:     "sanitize table item codeblock",
			filename: "codeblock",
			escape:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			bytes, err := os.ReadFile(filepath.Join("testdata", "table", tt.filename+".golden"))
			assert.Nil(err)

			actual := SanitizeAsciidocTable(string(bytes), tt.escape, false)

			expected, err := os.ReadFile(filepath.Join("testdata", "table", tt.filename+".asciidoc.expected"))
			assert.Nil(err)

			assert.Equal(string(expected), actual)
		})
	}
}

func TestConvertMultiLineText(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		isTable  bool
		showHTML bool
		expected string
	}{
		{
			name:     "convert multi-line newline-single",
			filename: "newline-single",
			isTable:  false,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit,\n\nsed do eiusmod tempor incididunt\n\nut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line newline-single",
			filename: "newline-single",
			isTable:  true,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,<br/><br/>consectetur adipiscing elit,<br/><br/>sed do eiusmod tempor incididunt<br/><br/>ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line newline-single",
			filename: "newline-single",
			isTable:  true,
			showHTML: false,
			expected: "Lorem ipsum dolor sit amet,  consectetur adipiscing elit,  sed do eiusmod tempor incididunt  ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line newline-double",
			filename: "newline-double",
			isTable:  false,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,\n\n\nconsectetur adipiscing elit,\n\n\nsed do eiusmod tempor incididunt\n\n\nut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line newline-double",
			filename: "newline-double",
			isTable:  true,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,<br/><br/><br/>consectetur adipiscing elit,<br/><br/><br/>sed do eiusmod tempor incididunt<br/><br/><br/>ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line newline-double",
			filename: "newline-double",
			isTable:  true,
			showHTML: false,
			expected: "Lorem ipsum dolor sit amet,   consectetur adipiscing elit,   sed do eiusmod tempor incididunt   ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line paragraph",
			filename: "paragraph",
			isTable:  false,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,  \nconsectetur adipiscing elit,  \nsed do eiusmod tempor incididunt  \nut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line paragraph",
			filename: "paragraph",
			isTable:  true,
			showHTML: true,
			expected: "Lorem ipsum dolor sit amet,<br/>consectetur adipiscing elit,<br/>sed do eiusmod tempor incididunt<br/>ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line paragraph",
			filename: "paragraph",
			isTable:  true,
			showHTML: false,
			expected: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line list",
			filename: "list",
			isTable:  false,
			showHTML: true,
			expected: "- Lorem ipsum dolor sit amet,\n  * Lorem ipsum dolor sit amet,\n  * consectetur adipiscing elit,\n- consectetur adipiscing elit,\n- sed do eiusmod tempor incididunt\n- ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line list",
			filename: "list",
			isTable:  true,
			showHTML: true,
			expected: "- Lorem ipsum dolor sit amet,<br/>  * Lorem ipsum dolor sit amet,<br/>  * consectetur adipiscing elit,<br/>- consectetur adipiscing elit,<br/>- sed do eiusmod tempor incididunt<br/>- ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line list",
			filename: "list",
			isTable:  true,
			showHTML: false,
			expected: "- Lorem ipsum dolor sit amet,   * Lorem ipsum dolor sit amet,   * consectetur adipiscing elit, - consectetur adipiscing elit, - sed do eiusmod tempor incididunt - ut labore et dolore magna aliqua.",
		},
		{
			name:     "convert multi-line indentations",
			filename: "indentations",
			isTable:  false,
			showHTML: true,
			expected: "This is is a multline test which works\n\nKey  \n  Foo1: blah  \n  Foo2: blah\n\nKey2  \nFoo1: bar1  \nFoo2: bar2",
		},
		{
			name:     "convert multi-line indentations",
			filename: "indentations",
			isTable:  true,
			showHTML: true,
			expected: "This is is a multline test which works<br/><br/>Key<br/>  Foo1: blah<br/>  Foo2: blah<br/><br/>Key2<br/>Foo1: bar1<br/>Foo2: bar2",
		},
		{
			name:     "convert multi-line indentations",
			filename: "indentations",
			isTable:  true,
			showHTML: false,
			expected: "This is is a multline test which works  Key   Foo1: blah   Foo2: blah  Key2 Foo1: bar1 Foo2: bar2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			path := filepath.Join("testdata", "multiline", tt.filename+".golden")
			bytes, err := os.ReadFile(path)
			assert.Nil(err)

			actual := ConvertMultiLineText(string(bytes), tt.isTable, false, tt.showHTML)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestEscapeCharacters(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		escapePipe  bool
		escapeChars bool
		expected    string
	}{
		{
			name:        "escape pipe",
			input:       "lorem | ipsum |dolor|consectetur| `adipi | scing |elit|sit|`",
			escapePipe:  true,
			escapeChars: false,
			expected:    "lorem \\| ipsum \\|dolor\\|consectetur\\| `adipi | scing |elit|sit|`",
		},
		{
			name:        "escape pipe",
			input:       "lorem || ipsum ||dolor||consectetur|| `adipi || scing ||elit||sit||`",
			escapePipe:  true,
			escapeChars: false,
			expected:    "lorem \\|\\| ipsum \\|\\|dolor\\|\\|consectetur\\|\\| `adipi || scing ||elit||sit||`",
		},
		{
			name:        "escape pipe",
			input:       "lorem ||| ipsum |||dolor|||consectetur||| `adipi ||| scing |||elit|||sit|||`",
			escapePipe:  true,
			escapeChars: false,
			expected:    "lorem \\|\\|\\| ipsum \\|\\|\\|dolor\\|\\|\\|consectetur\\|\\|\\| `adipi ||| scing |||elit|||sit|||`",
		},
		{
			name:        "do not escape pipe",
			input:       "lorem | ipsum |dolor|consectetur| `adipi | scing |elit|sit|`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem | ipsum |dolor|consectetur| `adipi | scing |elit|sit|`",
		},
		{
			name:        "do not escape pipe",
			input:       "lorem || ipsum ||dolor||consectetur|| `adipi || scing ||elit||sit||`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem || ipsum ||dolor||consectetur|| `adipi || scing ||elit||sit||`",
		},
		{
			name:        "do not escape pipe",
			input:       "lorem ||| ipsum |||dolor|||consectetur||| `adipi ||| scing |||elit|||sit|||`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem ||| ipsum |||dolor|||consectetur||| `adipi ||| scing |||elit|||sit|||`",
		},
		{
			name:        "escape underscore",
			input:       "lorem _ ipsum _dolor_consectetur_ incid_idunt `adipi _ scing _elit_sit_`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "lorem \\_ ipsum _dolor\\_consectetur_ incid\\_idunt `adipi _ scing _elit_sit_`",
		},
		{
			name:        "escape underscore",
			input:       "lorem __ ipsum __dolor__consectetur__ incid__idunt `adipi __ scing __elit__sit__`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "lorem \\_\\_ ipsum __dolor\\_\\_consectetur__ incid\\_\\_idunt `adipi __ scing __elit__sit__`",
		},
		{
			name:        "escape underscore",
			input:       "lorem ___ ipsum ___dolor___consectetur___ incid___idunt `adipi ___ scing ___elit___sit___`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "lorem \\_\\_\\_ ipsum ___dolor\\_\\_\\_consectetur___ incid\\_\\_\\_idunt `adipi ___ scing ___elit___sit___`",
		},
		{
			name:        "do not escape underscore",
			input:       "lorem _ ipsum _dolor_consectetur_ incid_idunt `adipi _ scing _elit_sit_`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem _ ipsum _dolor_consectetur_ incid_idunt `adipi _ scing _elit_sit_`",
		},
		{
			name:        "do not escape underscore",
			input:       "lorem __ ipsum __dolor__consectetur__ incid__idunt `adipi __ scing __elit__sit__`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem __ ipsum __dolor__consectetur__ incid__idunt `adipi __ scing __elit__sit__`",
		},
		{
			name:        "do not escape underscore",
			input:       "lorem ___ ipsum ___dolor___consectetur___ incid___idunt `adipi ___ scing ___elit___sit___`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "lorem ___ ipsum ___dolor___consectetur___ incid___idunt `adipi ___ scing ___elit___sit___`",
		},
		{
			name:        "escape asterisk",
			input:       "* lorem * ipsum *dolor*consectetur* `adipi * scing *elit*sit*`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "* lorem * ipsum *dolor*consectetur* `adipi * scing *elit*sit*`",
		},
		{
			name:        "escape asterisk",
			input:       "** lorem ** ipsum **dolor**consectetur** `adipi ** scing **elit**sit**`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "** lorem ** ipsum **dolor**consectetur** `adipi ** scing **elit**sit**`",
		},
		{
			name:        "escape asterisk",
			input:       "*** lorem *** ipsum ***dolor***consectetur*** `adipi *** scing ***elit***sit***`",
			escapePipe:  false,
			escapeChars: true,
			expected:    "*** lorem *** ipsum ***dolor***consectetur*** `adipi *** scing ***elit***sit***`",
		},
		{
			name:        "escape asterisk",
			input:       "**lorem ipsum dolor consectetur adipi scing elit sit**",
			escapePipe:  false,
			escapeChars: true,
			expected:    "**lorem ipsum dolor consectetur adipi scing elit sit**",
		},
		{
			name:        "do not escape asterisk",
			input:       "* lorem * ipsum *dolor*consectetur* `adipi * scing *elit*sit*`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "* lorem * ipsum *dolor*consectetur* `adipi * scing *elit*sit*`",
		},
		{
			name:        "do not escape asterisk",
			input:       "** lorem ** ipsum **dolor**consectetur** `adipi ** scing **elit**sit**`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "** lorem ** ipsum **dolor**consectetur** `adipi ** scing **elit**sit**`",
		},
		{
			name:        "do not escape asterisk",
			input:       "*** lorem *** ipsum ***dolor***consectetur*** `adipi *** scing ***elit***sit***`",
			escapePipe:  false,
			escapeChars: false,
			expected:    "*** lorem *** ipsum ***dolor***consectetur*** `adipi *** scing ***elit***sit***`",
		},
		{
			name:        "do not escape asterisk",
			input:       "**lorem ipsum dolor consectetur adipi scing elit sit**",
			escapePipe:  false,
			escapeChars: false,
			expected:    "**lorem ipsum dolor consectetur adipi scing elit sit**",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := EscapeCharacters(tt.input, tt.escapeChars, tt.escapePipe)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestNormalizeURLs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		escape   bool
		expected string
	}{
		{
			name:     "normalize url with escape character",
			input:    "https://www.foo\\_bar.com/",
			escape:   true,
			expected: "https://www.foo_bar.com/",
		},
		{
			name:     "normalize url with escape character",
			input:    "lorem ipsum https://www.foo\\_bar.com/ dolor sit amet",
			escape:   true,
			expected: "lorem ipsum https://www.foo_bar.com/ dolor sit amet",
		},
		{
			name:     "normalize url with escape character",
			input:    "lorem\\_ipsum https://www.foo\\_bar.com/ dolor sit amet",
			escape:   true,
			expected: "lorem\\_ipsum https://www.foo_bar.com/ dolor sit amet",
		},
		{
			name:     "normalize url without escape character",
			input:    "https://www.foo_bar.com/",
			escape:   false,
			expected: "https://www.foo_bar.com/",
		},
		{
			name:     "normalize url without escape character",
			input:    "https://www.foo\\_bar.com/",
			escape:   false,
			expected: "https://www.foo\\_bar.com/",
		},
		{
			name:     "normalize url with escape character",
			input:    "lorem\\_ipsum https://www.foo\\_bar.com/ dolor sit amet",
			escape:   false,
			expected: "lorem\\_ipsum https://www.foo\\_bar.com/ dolor sit amet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := NormalizeURLs(tt.input, tt.escape)

			assert.Equal(tt.expected, actual)
		})
	}
}
