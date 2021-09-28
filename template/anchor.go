/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"fmt"
)

// CreateAnchorMarkdown creates HTML anchor for Markdown format.
func CreateAnchorMarkdown(prefix string, value string, anchor bool, escape bool) string {
	sanitizedName := SanitizeName(value, escape)

	if anchor {
		anchorName := fmt.Sprintf("%s_%s", prefix, value)
		sanitizedAnchorName := SanitizeName(anchorName, escape)
		// the <a> link is purposely not sanitized as this breaks markdown formatting
		return fmt.Sprintf("<a name=\"%s\"></a> [%s](#%s)", anchorName, sanitizedName, sanitizedAnchorName)
	}

	return sanitizedName
}

// CreateAnchorAsciidoc creates HTML anchor for AsciiDoc format.
func CreateAnchorAsciidoc(prefix string, value string, anchor bool, escape bool) string {
	sanitizedName := SanitizeName(value, escape)

	if anchor {
		anchorName := fmt.Sprintf("%s_%s", prefix, value)
		sanitizedAnchorName := SanitizeName(anchorName, escape)
		return fmt.Sprintf("[[%s]] <<%s,%s>>", sanitizedAnchorName, sanitizedAnchorName, sanitizedName)
	}

	return sanitizedName
}
