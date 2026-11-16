/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

// Mappings of CLI flags to Viper config
var flagMappings = map[string]string{
	"header-from": "header-from",
	"footer-from": "footer-from",

	"hide-empty": "hide-empty",

	"show": "sections.show",
	"hide": "sections.hide",

	"output-file":     "output.file",
	"output-mode":     "output.mode",
	"output-template": "output.template",

	"output-values":      "output-values.enabled",
	"output-values-from": "output-values.from",

	"sort":             "sort.enabled",
	"sort-by":          "sort.by",
	"sort-by-required": "required",
	"sort-by-type":     "type",

	"anchor":        "settings.anchor",
	"color":         "settings.color",
	"default":       "settings.default",
	"description":   "settings.description",
	"escape":        "settings.escape",
	"indent":        "settings.indent",
	"read-comments": "settings.read-comments",
	"required":      "settings.required",
	"sensitive":     "settings.sensitive",
	"type":          "settings.type",
	"validation":    "settings.validation",
}
