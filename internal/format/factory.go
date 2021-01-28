/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"fmt"

	"github.com/terraform-docs/terraform-docs/pkg/print"
)

// Factory initializes and returns the conceret implementation of
// print.Format based on the provided 'name', for example for name
// of 'json' it will return '*format.JSON' through 'format.NewJSON'
// function.
func Factory(name string, settings *print.Settings) (print.Format, error) {
	switch name {
	case "asciidoc", "adoc":
		return NewAsciidocTable(settings), nil
	case "asciidoc document", "asciidoc doc", "adoc document", "adoc doc":
		return NewAsciidocDocument(settings), nil
	case "asciidoc table", "asciidoc tbl", "adoc table", "adoc tbl":
		return NewAsciidocTable(settings), nil
	case "json":
		return NewJSON(settings), nil
	case "markdown", "md":
		return NewTable(settings), nil
	case "markdown document", "markdown doc", "md document", "md doc":
		return NewDocument(settings), nil
	case "markdown table", "markdown tbl", "md table", "md tbl":
		return NewTable(settings), nil
	case "pretty":
		return NewPretty(settings), nil
	case "tfvars hcl":
		return NewTfvarsHCL(settings), nil
	case "tfvars json":
		return NewTfvarsJSON(settings), nil
	case "toml":
		return NewTOML(settings), nil
	case "xml":
		return NewXML(settings), nil
	case "yaml":
		return NewYAML(settings), nil
	}
	return nil, fmt.Errorf("formatter '%s' not found", name)
}
