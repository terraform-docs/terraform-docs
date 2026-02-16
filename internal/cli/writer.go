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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rquadling/terraform-docs/print"
)

// stdoutWriter writes content to os.Stdout.
type stdoutWriter struct{}

// Write content to Stdout
func (sw *stdoutWriter) Write(p []byte) (int, error) {
	return os.Stdout.WriteString(string(p) + "\n")
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

	check bool

	template string
	begin    string
	end      string

	writer io.Writer
}

// Write content to target file
func (fw *fileWriter) Write(p []byte) (int, error) {
	filename := fw.fullFilePath()

	if fw.template == "" {
		// template is optional for mode replace
		if fw.mode == print.OutputModeReplace {
			return fw.write(filename, p)
		}
		return 0, errors.New("template is missing")
	}

	// apply template to generated output
	buf, err := fw.apply(p)
	if err != nil {
		return 0, err
	}

	// Replace the content of 'filename' with generated output,
	// no further processing is required for mode 'replace'.
	if fw.mode == print.OutputModeReplace {
		return fw.write(filename, buf.Bytes())
	}

	content, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		// In mode 'inject', if target file not found:
		// create it and save the generated output into it.
		return fw.write(filename, buf.Bytes())
	}

	if len(content) == 0 {
		// In mode 'inject', if target file is found BUT it's empty:
		// save the generated output into it.
		return fw.write(filename, buf.Bytes())
	}

	return fw.inject(filename, string(content), buf.String())
}

// fullFilePath of the target file. If file is absolute path it will be
// used as is, otherwise dir (i.e. module root folder) will be joined to
// it.
func (fw *fileWriter) fullFilePath() string {
	if filepath.IsAbs(fw.file) {
		return fw.file
	}
	return filepath.Join(fw.dir, fw.file)
}

// apply template to generated output
func (fw *fileWriter) apply(p []byte) (bytes.Buffer, error) {
	type content struct {
		Content string
	}

	var buf bytes.Buffer

	tmpl := template.Must(template.New("content").Parse(fw.template))
	err := tmpl.ExecuteTemplate(&buf, "content", content{string(p)})

	return buf, err
}

// inject generated output into file.
func (fw *fileWriter) inject(filename string, content string, generated string) (int, error) {
	before := strings.Index(content, fw.begin)
	after := strings.Index(content, fw.end)

	// current file content doesn't have surrounding
	// so we're going to append the generated output
	// to current file.
	if before < 0 && after < 0 {
		return fw.write(filename, []byte(content+"\n"+generated))
	}

	// begin comment is missing
	if before < 0 {
		return 0, errors.New("begin comment is missing")
	}

	generated = content[:before] + generated

	// end comment is missing
	if after < 0 {
		return 0, errors.New("end comment is missing")
	}

	// end comment is before begin comment
	if after < before {
		return 0, errors.New("end comment is before begin comment")
	}

	generated += content[after+len(fw.end):]

	return fw.write(filename, []byte(generated))
}

// write the content to io.Writer. If no io.Writer is available,
// it will be written to 'filename'.
func (fw *fileWriter) write(filename string, p []byte) (int, error) {
	// if run in check mode return exit 1
	if fw.check {
		f, err := os.ReadFile(filepath.Clean(filename))
		if err != nil {
			return 0, err
		}

		// check for changes and print changed file
		if !bytes.Equal(f, p) {
			return 0, fmt.Errorf("%s is out of date", filename)
		}

		fmt.Printf("%s is up to date\n", filename)
		return 0, nil
	}

	if fw.writer != nil {
		return fw.writer.Write(p)
	}

	fmt.Printf("%s updated successfully\n", filename)
	return len(p), os.WriteFile(filename, p, 0644)
}
