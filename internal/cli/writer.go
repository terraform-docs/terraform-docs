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
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	errFileEmpty             = "file content is empty"
	errTemplateEmpty         = "template is missing"
	errBeginCommentMissing   = "begin comment is missing"
	errEndCommentMissing     = "end comment is missing"
	errEndCommentBeforeBegin = "end comment is before begin comment"
)

// stdoutWriter writes content to os.Stdout.
type stdoutWriter struct{}

func (sw *stdoutWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write([]byte(string(p) + "\n"))
}

// fileWriter writes content to file.
//
// First of all it will process 'content' into provided 'template'.
//
// If 'mode' is 'replace' it replaces the whole content of 'dir/file'
// with output of executed template. Note that this will create 'dir/file'
// if it doesn't exist.
//
// If 'mode' is 'inject' it will attempt to inject the output of executed
// template into 'dir/file' between the 'begin' and 'end' comment. Note that
// this will fail if 'dir/file' doesn't exist, or doesn't contain 'begin' or
// 'end' comment.
type fileWriter struct {
	file string
	dir  string

	mode string

	template string
	begin    string
	end      string
}

func (fw *fileWriter) Write(p []byte) (int, error) {
	var buf bytes.Buffer

	if fw.template == "" {
		return 0, errors.New(errTemplateEmpty)
	}

	tmpl := template.Must(template.New("content").Parse(fw.template))
	if err := tmpl.ExecuteTemplate(&buf, "content", struct {
		Content string
	}{
		Content: string(p),
	}); err != nil {
		return 0, err
	}

	content := buf.String()
	filename := filepath.Join(fw.dir, fw.file)

	if fw.mode == outputModeInject {
		f, err := os.ReadFile(filename)
		if err != nil {
			return 0, err
		}

		fc := string(f)
		if fc == "" {
			return 0, errors.New(errFileEmpty)
		}

		before := strings.Index(fc, fw.begin)
		if before < 0 {
			return 0, errors.New(errBeginCommentMissing)
		}
		content = fc[:before] + content

		after := strings.Index(fc, fw.end)
		if after < 0 {
			return 0, errors.New(errEndCommentMissing)
		}
		if after < before {
			return 0, errors.New(errEndCommentBeforeBegin)
		}
		content += fc[after+len(fw.end):]
	}

	n := len(content)
	err := os.WriteFile(filename, []byte(content), 0644)

	return n, err
}
