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
	if len(override) != 1 {
		return base
	}
	dest := override[0]
	if err := mergo.Merge(&dest, base); err != nil {
		return base
	}
	return dest
}
