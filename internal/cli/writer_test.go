/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileWriter(t *testing.T) {
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	tests := map[string]struct {
		file     string
		mode     string
		template string
		begin    string
		end      string
		errMsg   string
	}{
		"ModeInjectNoFile": {
			file:     "file-missing.md",
			mode:     "inject",
			template: OutputTemplate,
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "open testdata/writer/file-missing.md: no such file or directory",
		},
		"EmptyTemplate": {
			file:     "not-applicable.md",
			mode:     "inject",
			template: "",
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "template is missing",
		},
		"EmptyFile": {
			file:     "empty-file.md",
			mode:     "inject",
			template: OutputTemplate,
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "file content is empty",
		},
		"BeginCommentMissing": {
			file:     "begin-comment-missing.md",
			mode:     "inject",
			template: OutputTemplate,
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "begin comment is missing",
		},
		"EndCommentMissing": {
			file:     "end-comment-missing.md",
			mode:     "inject",
			template: OutputTemplate,
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "end comment is missing",
		},
		"EndCommentBeforeBegin": {
			file:     "end-comment-before-begin.md",
			mode:     "inject",
			template: OutputTemplate,
			begin:    outputBeginComment,
			end:      outputEndComment,
			errMsg:   "end comment is before begin comment",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			writer := &fileWriter{
				file: tt.file,
				dir:  filepath.Join("testdata", "writer"),

				mode: tt.mode,

				template: tt.template,
				begin:    tt.begin,
				end:      tt.end,
			}

			_, err := io.WriteString(writer, content)

			assert.NotNil(err)
			assert.Equal(tt.errMsg, err.Error())
		})
	}
}
