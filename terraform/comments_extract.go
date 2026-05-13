/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/reader"
)

const ignoreMarker = "terraform-docs-ignore"

// loadComments reads contiguous leading comment lines (`#` or `//`) immediately
// preceding the given source line and returns them joined with single spaces.
// It absorbs read errors and returns an empty string on failure.
func loadComments(filename string, lineNum int) string {
	lines := reader.Lines{
		FileName: filename,
		LineNum:  lineNum,
		Condition: func(line string) bool {
			return strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//")
		},
		Parser: func(line string) (string, bool) {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "#")
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			return line, true
		},
	}
	comment, err := lines.Extract()
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(comment, " ")
}

// isIgnored returns the leading comments for the given source position and
// whether those comments contain the `terraform-docs-ignore` marker.
func isIgnored(filename string, line int) (comments string, ignored bool) {
	comments = loadComments(filename, line)
	return comments, strings.Contains(comments, ignoreMarker)
}
