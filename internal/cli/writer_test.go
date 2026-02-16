/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"bytes"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rquadling/terraform-docs/internal/testutil"
	"github.com/rquadling/terraform-docs/print"
)

func TestFileWriterFullPath(t *testing.T) {
	tests := map[string]struct {
		file     string
		dir      string
		expected string
	}{
		"Relative": {
			file:     "file.md",
			dir:      "/path/to/module",
			expected: "/path/to/module/file.md",
		},
		"Absolute": {
			file:     "/path/to/module/file.md",
			dir:      ".",
			expected: "/path/to/module/file.md",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			writer := &fileWriter{
				file: tt.file,
				dir:  tt.dir,
			}

			actual := writer.fullFilePath()
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestFileWriter(t *testing.T) {
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	tests := map[string]struct {
		file     string
		mode     string
		check    bool
		template string
		begin    string
		end      string
		writer   io.Writer

		expected string
		wantErr  bool
		errMsg   string
	}{
		// Successful writes
		"ModeInject": {
			file:     "mode-inject.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-inject",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeInjectEmptyFile": {
			file:     "empty-file.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-inject-empty-file",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeInjectNoCommentAppend": {
			file:     "mode-inject-no-comment.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-inject-no-comment",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeInjectFileMissing": {
			file:     "file-missing.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-inject-file-missing",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeReplaceWithComment": {
			file:     "mode-replace.md",
			mode:     "replace",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-replace-with-comment",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeReplaceWithCommentEmptyFile": {
			file:     "mode-replace.md",
			mode:     "replace",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-replace-with-comment",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeReplaceWithCommentFileMissing": {
			file:     "file-missing.md",
			mode:     "replace",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   &bytes.Buffer{},

			expected: "mode-replace-with-comment",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeReplaceWithoutComment": {
			file:     "mode-replace.md",
			mode:     "replace",
			check:    false,
			template: print.OutputContent,
			begin:    "",
			end:      "",
			writer:   &bytes.Buffer{},

			expected: "mode-replace-without-comment",
			wantErr:  false,
			errMsg:   "",
		},
		"ModeReplaceWithoutTemplate": {
			file:     "mode-replace.md",
			mode:     "replace",
			check:    false,
			template: "",
			begin:    "",
			end:      "",
			writer:   &bytes.Buffer{},

			expected: "mode-replace-without-template",
			wantErr:  false,
			errMsg:   "",
		},

		// Error writes
		"EmptyTemplate": {
			file:     "not-applicable.md",
			mode:     "inject",
			check:    false,
			template: "",
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   nil,

			expected: "",
			wantErr:  true,
			errMsg:   "template is missing",
		},
		"BeginCommentMissing": {
			file:     "begin-comment-missing.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   nil,

			expected: "",
			wantErr:  true,
			errMsg:   "begin comment is missing",
		},
		"EndCommentMissing": {
			file:     "end-comment-missing.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   nil,

			expected: "",
			wantErr:  true,
			errMsg:   "end comment is missing",
		},
		"EndCommentBeforeBegin": {
			file:     "end-comment-before-begin.md",
			mode:     "inject",
			check:    false,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   nil,

			expected: "",
			wantErr:  true,
			errMsg:   "end comment is before begin comment",
		},
		"ModeReplaceOutOfDate": {
			file:     "mode-replace.md",
			mode:     "replace",
			check:    true,
			template: print.OutputTemplate,
			begin:    print.OutputBeginComment,
			end:      print.OutputEndComment,
			writer:   nil,

			expected: "",
			wantErr:  true,
			errMsg:   "testdata/writer/mode-replace.md is out of date",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			writer := &fileWriter{
				file: tt.file,
				dir:  filepath.Join("testdata", "writer"),

				mode: tt.mode,

				check: tt.check,

				template: tt.template,
				begin:    tt.begin,
				end:      tt.end,

				writer: tt.writer,
			}

			_, err := io.WriteString(writer, content)

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)

				w, ok := tt.writer.(*bytes.Buffer)
				assert.True(ok, "tt.writer is not a valid bytes.Buffer")

				expected, err := testutil.GetExpected("writer", tt.expected)
				assert.Nil(err)

				assert.Equal(expected, w.String())
			}
		})
	}
}
