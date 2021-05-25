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

	"github.com/terraform-docs/terraform-docs/internal/print"
)

// createAnchorMarkdown
func createAnchorMarkdown(t string, s string, settings *print.Settings) string {
	sanitizedName := sanitizeName(s, settings)

	if settings.ShowAnchor {
		anchorName := fmt.Sprintf("%s_%s", t, s)
		sanitizedAnchorName := sanitizeName(anchorName, settings)
		// the <a> link is purposely not sanitized as this breaks markdown formatting
		return fmt.Sprintf("<a name=\"%s\"></a> [%s](#%s)", anchorName, sanitizedName, sanitizedAnchorName)
	}

	return sanitizedName
}

// createAnchorAsciidoc
func createAnchorAsciidoc(t string, s string, settings *print.Settings) string {
	sanitizedName := sanitizeName(s, settings)

	if settings.ShowAnchor {
		anchorName := fmt.Sprintf("%s_%s", t, s)
		sanitizedAnchorName := sanitizeName(anchorName, settings)
		return fmt.Sprintf("[[%s]] <<%s,%s>>", sanitizedAnchorName, sanitizedAnchorName, sanitizedName)
	}

	return sanitizedName
}
