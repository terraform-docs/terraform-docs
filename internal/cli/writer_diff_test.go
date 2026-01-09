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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/print"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestFileWriter_Diff(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test-diff.txt")
	initialContent := []byte("line 1\nline 2\nline 3\n")
	newContent := []byte("line 1\nline 2 changed\nline 3\n")

	// Create initial file
	err := os.WriteFile(filePath, initialContent, 0644)
	assert.NoError(t, err)

	t.Run("CheckMode_PrintsDiff", func(t *testing.T) {
		fw := &fileWriter{
			file:  filePath,
			dir:   "",
			check: true,
			mode:  print.OutputModeReplace,
		}

		output := captureStdout(func() {
			_, err := fw.Write(newContent)
			assert.Error(t, err)
			assert.Equal(t, fmt.Sprintf("%s is out of date", filePath), err.Error())
		})

		assert.Contains(t, output, "--- "+filePath)
		assert.Contains(t, output, "+++ "+filePath)
		assert.Contains(t, output, "-line 2")
		assert.Contains(t, output, "+line 2 changed")
	})

	t.Run("WriteMode_PrintsDiff_OnOverwrite", func(t *testing.T) {
		// Reset file content
		err := os.WriteFile(filePath, initialContent, 0644)
		assert.NoError(t, err)

		fw := &fileWriter{
			file:  filePath,
			dir:   "",
			check: false,
			mode:  print.OutputModeReplace,
		}

		output := captureStdout(func() {
			_, err := fw.Write(newContent)
			assert.NoError(t, err)
		})

		assert.Contains(t, output, "--- "+filePath)
		assert.Contains(t, output, "+++ "+filePath)
		assert.Contains(t, output, "-line 2")
		assert.Contains(t, output, "+line 2 changed")
		assert.Contains(t, output, filePath+" updated successfully")
	})

	t.Run("WriteMode_NoDiff_OnIdentical", func(t *testing.T) {
		// Reset file content
		err := os.WriteFile(filePath, initialContent, 0644)
		assert.NoError(t, err)

		fw := &fileWriter{
			file:  filePath,
			dir:   "",
			check: false,
			mode:  print.OutputModeReplace,
		}

		output := captureStdout(func() {
			_, err := fw.Write(initialContent)
			assert.NoError(t, err)
		})

		assert.NotContains(t, output, "--- ")
		assert.NotContains(t, output, "+++ ")
		assert.Contains(t, output, filePath+" updated successfully")
	})
}
