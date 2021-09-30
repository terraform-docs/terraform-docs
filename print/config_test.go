/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSections(t *testing.T) {
	tests := map[string]struct {
		sections sections
		wantErr  bool
		errMsg   string
	}{
		"OnlyShows": {
			sections: sections{
				Show: []string{sectionHeader, sectionInputs},
				Hide: []string{},
			},
			wantErr: false,
			errMsg:  "",
		},
		"OnlyHide": {
			sections: sections{
				Show: []string{},
				Hide: []string{sectionHeader, sectionInputs},
			},
			wantErr: false,
			errMsg:  "",
		},
		"BothShowAndHide": {
			sections: sections{
				Show: []string{sectionHeader},
				Hide: []string{sectionInputs},
			},
			wantErr: true,
			errMsg:  "'--show' and '--hide' can't be used together",
		},
		"UnknownShow": {
			sections: sections{
				Show: []string{"foo"},
				Hide: []string{},
			},
			wantErr: true,
			errMsg:  "'foo' is not a valid section",
		},
		"UnknownHide": {
			sections: sections{
				Show: []string{},
				Hide: []string{"foo"},
			},
			wantErr: true,
			errMsg:  "'foo' is not a valid section",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.sections.validate()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}
		})
	}
}

func TestConfigVisibility(t *testing.T) {
	tests := []struct {
		sections sections
		name     string
		expected bool
	}{
		{
			sections: sections{},
			name:     "header",
			expected: true,
		},
		{
			sections: sections{
				Show: []string{"header"},
				Hide: []string{},
			},
			name:     "header",
			expected: true,
		},
		{
			sections: sections{
				Show: []string{"all"},
				Hide: []string{},
			},
			name:     "header",
			expected: true,
		},
		{
			sections: sections{
				Show: []string{},
				Hide: []string{"inputs"},
			},
			name:     "header",
			expected: true,
		},

		{
			sections: sections{
				Show: []string{},
				Hide: []string{"header"},
			},
			name:     "header",
			expected: false,
		},
		{
			sections: sections{
				Show: []string{},
				Hide: []string{"all"},
			},
			name:     "header",
			expected: false,
		},
		{
			sections: sections{
				Show: []string{"inputs"},
				Hide: []string{},
			},
			name:     "header",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run("section visibility", func(t *testing.T) {
			assert := assert.New(t)

			visible := tt.sections.visibility(tt.name)
			assert.Equal(tt.expected, visible)
		})
	}
}

func TestConfigOutput(t *testing.T) {
	tests := map[string]struct {
		output  output
		wantErr bool
		errMsg  string
	}{
		"FileEmpty": {
			output: output{
				File:     "",
				Mode:     "",
				Template: "",
			},
			wantErr: false,
			errMsg:  "",
		},
		"TemplateEmptyModeReplace": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeReplace,
				Template: "",
			},
			wantErr: false,
			errMsg:  "",
		},
		"TemplateLiteralLineBreak": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: fmt.Sprintf("%s\\n%s\\n%s", OutputBeginComment, OutputContent, OutputEndComment),
			},
			wantErr: false,
			errMsg:  "",
		},
		"NoExtraValidationModeReplace": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeReplace,
				Template: fmt.Sprintf("%s\\n%s\\n%s", OutputBeginComment, OutputContent, OutputEndComment),
			},
			wantErr: false,
			errMsg:  "",
		},

		"ModeEmpty": {
			output: output{
				File:     "README.md",
				Mode:     "",
				Template: "",
			},
			wantErr: true,
			errMsg:  "value of '--output-mode' can't be empty",
		},
		"TemplateEmptyModeInject": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: "",
			},
			wantErr: true,
			errMsg:  "value of '--output-template' can't be empty",
		},
		"TemplateNotContent": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: fmt.Sprintf("%s\n%s", OutputBeginComment, OutputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)",
		},
		"TemplateNotThreeLines": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: fmt.Sprintf("%s%s%s", OutputBeginComment, OutputContent, OutputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' should contain at least 3 lines (begin comment, {{ .Content }}, and end comment)",
		},
		"TemplateBeginCommentMissing": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: fmt.Sprintf("no-begin-comment\n%s\n%s", OutputContent, OutputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' is missing begin comment",
		},
		"TemplateEndCommentMissing": {
			output: output{
				File:     "README.md",
				Mode:     OutputModeInject,
				Template: fmt.Sprintf("%s\n%s\nno-end-comment", OutputBeginComment, OutputContent),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' is missing end comment",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.output.validate()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}
		})
	}
}

func TestIsInlineComment(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		{
			name:     "markdown comment variant",
			line:     "<!-- this is a comment -->",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "<!-- this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "<!-- this is not a --> comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this <!-- is not a comment -->",
			expected: false,
		},

		{
			name:     "markdown comment variant",
			line:     "[]: # (this is a comment)",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # (this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # (this is not a) comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this []: # (is not a comment)",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]:#(this is not a comment)",
			expected: false,
		},

		{
			name:     "markdown comment variant",
			line:     "[]: # \"this is a comment\"",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # \"this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # \"this is not a\" comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this []: # \"is not a comment\"",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]:#\"this is not a comment\"",
			expected: false,
		},

		{
			name:     "markdown comment variant",
			line:     "[]: # 'this is a comment'",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # 'this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]: # 'this is not a' comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this []: # 'is not a comment'",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[]:#'this is not a comment'",
			expected: false,
		},

		{
			name:     "markdown comment variant",
			line:     "[//]: # (this is a comment)",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "[//]: # (this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[//]: # (this is not a) comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this [//]: # (is not a comment)",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[//]:#(this is not a comment)",
			expected: false,
		},

		{
			name:     "markdown comment variant",
			line:     "[comment]: # (this is a comment)",
			expected: true,
		},
		{
			name:     "markdown comment variant",
			line:     "[comment]: # (this is not a comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[comment]: # (this is not a) comment",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "this [comment]: # (is not a comment)",
			expected: false,
		},
		{
			name:     "markdown comment variant",
			line:     "[comment]:#(this is not a comment)",
			expected: false,
		},

		{
			name:     "asciidoc comment variant",
			line:     "// this is a comment",
			expected: true,
		},
		{
			name:     "asciidoc comment variant",
			line:     "this // is not a comment",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := isInlineComment(tt.line)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestConfigSort(t *testing.T) {
	tests := map[string]struct {
		sort    sort
		wantErr bool
		errMsg  string
	}{
		"name": {
			sort: sort{
				By: SortName,
			},
			wantErr: false,
			errMsg:  "",
		},
		"required": {
			sort: sort{
				By: SortRequired,
			},
			wantErr: false,
			errMsg:  "",
		},
		"type": {
			sort: sort{
				By: SortType,
			},
			wantErr: false,
			errMsg:  "",
		},

		"foo": {
			sort: sort{
				By: "foo",
			},
			wantErr: true,
			errMsg:  "'foo' is not a valid sort type",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.sort.validate()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}
		})
	}
}

func TestConfigOutputvalues(t *testing.T) {
	tests := map[string]struct {
		outputvalues outputvalues
		wantErr      bool
		errMsg       string
	}{
		"OK": {
			outputvalues: outputvalues{
				Enabled: true,
				From:    "file.json",
			},
			wantErr: false,
			errMsg:  "",
		},
		"Disabled": {
			outputvalues: outputvalues{
				Enabled: false,
			},
			wantErr: false,
			errMsg:  "",
		},
		"FromEmpty": {
			outputvalues: outputvalues{
				Enabled: true,
				From:    "",
			},
			wantErr: true,
			errMsg:  "value of '--output-values-from' is missing",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.outputvalues.validate()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}
		})
	}
}

func TestConfigValidate(t *testing.T) {
	tests := map[string]struct {
		config  func(c *Config)
		wantErr bool
		errMsg  string
	}{
		"OK": {
			config:  func(c *Config) {},
			wantErr: false,
			errMsg:  "",
		},
		"FormatterEmpty": {
			config: func(c *Config) {
				c.Formatter = ""
			},
			wantErr: true,
			errMsg:  "value of 'formatter' can't be empty",
		},
		"RecursivePathEmpty": {
			config: func(c *Config) {
				c.Recursive.Enabled = true
				c.Recursive.Path = ""
			},
			wantErr: true,
			errMsg:  "value of '--recursive-path' can't be empty",
		},
		"HeaderFromEmpty": {
			config: func(c *Config) {
				c.HeaderFrom = ""
			},
			wantErr: true,
			errMsg:  "value of '--header-from' can't be empty",
		},
		"FooterFrom": {
			config: func(c *Config) {
				c.FooterFrom = ""
				c.Sections.Footer = true
			},
			wantErr: true,
			errMsg:  "value of '--footer-from' can't be empty",
		},
		"SameHeaderFooterFrom": {
			config: func(c *Config) {
				c.Formatter = "foo"
				c.HeaderFrom = "README.md"
				c.FooterFrom = "README.md"
			},
			wantErr: true,
			errMsg:  "value of '--footer-from' can't equal value of '--header-from",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			config := DefaultConfig()
			config.Formatter = "foo"
			tt.config(config)
			err := config.Validate()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}
		})
	}
}
