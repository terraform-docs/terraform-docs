/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package testutil

import (
	"github.com/imdario/mergo"

	"github.com/terraform-docs/terraform-docs/internal/print"
)

// WithSections appends show all sections to provided Settings.
func WithSections(override ...print.Settings) print.Settings {
	base := print.Settings{
		ShowDataSources:  true,
		ShowFooter:       true,
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,

		ShowDefault: true,
		ShowType:    true,
	}
	return apply(base, override...)
}

// WithHTML appends ShowHTML to provided Settings.
func WithHTML(override ...print.Settings) print.Settings {
	base := print.Settings{
		ShowHTML: true,
	}
	return apply(base, override...)
}

func apply(base print.Settings, override ...print.Settings) print.Settings {
	dest := base
	for i := range override {
		if err := mergo.Merge(&dest, override[i]); err != nil {
			return base
		}
	}
	return dest

}
