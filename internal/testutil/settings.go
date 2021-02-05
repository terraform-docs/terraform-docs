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

// TestSettings respresents the Settings instance for tests
type TestSettings struct {
	full *print.Settings
}

// Settings returns TestSettings instance with predefined set of print.Settings
func Settings() *TestSettings {
	shared := &print.Settings{
		EscapePipe: true,
	}
	return &TestSettings{
		full: shared,
	}
}

// With appends provided 'override' print.Settings to TestSettings
func (s *TestSettings) With(override *print.Settings) *TestSettings {
	if err := mergo.Merge(override, s.full); err == nil {
		s.full = override
	}
	return s
}

// WithColor appends predefined 'ShowColor: true' to TestSettings
func (s *TestSettings) WithColor() *TestSettings {
	color := &print.Settings{
		ShowColor: true,
	}
	if err := mergo.Merge(color, s.full); err == nil {
		s.full = color
	}
	return s
}

// WithSections appends predefined show all sections ShowHeader, ShowProviders, ShowInputs, ShowModulecalls, ShowOutputs to TestSettings
func (s *TestSettings) WithSections() *TestSettings {
	sections := &print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
		ShowResources:    true,
	}
	if err := mergo.Merge(sections, s.full); err == nil {
		s.full = sections
	}
	return s
}

// Build builds and returns print.Settings based on the provided overrides in TestSettings
func (s *TestSettings) Build() *print.Settings {
	return s.full
}
