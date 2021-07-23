/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/
package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
				Mode:     outputModeReplace,
				Template: "",
			},
			wantErr: false,
			errMsg:  "",
		},
		"TemplateLiteralLineBreak": {
			output: output{
				File:     "README.md",
				Mode:     outputModeInject,
				Template: fmt.Sprintf("%s\\n%s\\n%s", outputBeginComment, outputContent, outputEndComment),
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
				Mode:     outputModeInject,
				Template: "",
			},
			wantErr: true,
			errMsg:  "value of '--output-template' can't be empty",
		},
		"TemplateNotContent": {
			output: output{
				File:     "README.md",
				Mode:     outputModeInject,
				Template: fmt.Sprintf("%s\n%s", outputBeginComment, outputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)",
		},
		"TemplateNotThreeLines": {
			output: output{
				File:     "README.md",
				Mode:     outputModeInject,
				Template: fmt.Sprintf("%s%s%s", outputBeginComment, outputContent, outputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' should contain at least 3 lines (begin comment, {{ .Content }}, and end comment)",
		},
		"TemplateBeginCommentMissing": {
			output: output{
				File:     "README.md",
				Mode:     outputModeInject,
				Template: fmt.Sprintf("no-begin-comment\n%s\n%s", outputContent, outputEndComment),
			},
			wantErr: true,
			errMsg:  "value of '--output-template' is missing begin comment",
		},
		"TemplateEndCommentMissing": {
			output: output{
				File:     "README.md",
				Mode:     outputModeInject,
				Template: fmt.Sprintf("%s\n%s\nno-end-comment", outputBeginComment, outputContent),
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
