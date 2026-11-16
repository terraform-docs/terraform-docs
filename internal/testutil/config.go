/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package testutil

import (
	"dario.cat/mergo"

	"github.com/rquadling/terraform-docs/print"
)

func baseConfig() print.Config {
	base := print.NewConfig()
	base.Settings.ReadComments = true

	return *base
}

func baseSections() print.Config {
	base := baseConfig()

	base.Sections.DataSources = true
	base.Sections.Header = true
	base.Sections.Inputs = true
	base.Sections.ModuleCalls = true
	base.Sections.Outputs = true
	base.Sections.Providers = true
	base.Sections.Requirements = true
	base.Sections.Resources = true

	base.Settings.Default = true
	base.Settings.Type = true

	return base
}

// With appends items to provided print.Config.
func With(fn func(*print.Config)) print.Config {
	base := baseConfig()
	fn(&base)

	return base
}

// WithSections shows all sections (including footer) to provided print.Config.
func WithSections(override ...print.Config) print.Config {
	base := baseSections()

	base.Sections.Footer = true
	base.FooterFrom = "footer.md"

	return apply(base, override...)
}

// WithDefaultSections shows default sections (everything except footer) to provided print.Config.
func WithDefaultSections(override ...print.Config) print.Config {
	base := baseSections()

	return apply(base, override...)
}

// WithHTML sets HTML to provided print.Config.
func WithHTML(override ...print.Config) print.Config {
	base := baseConfig()
	base.Settings.HTML = true

	return apply(base, override...)
}

// WithHideEmpty sets HideEmpty to provided print.Config.
func WithHideEmpty(override ...print.Config) print.Config {
	base := baseConfig()
	base.Settings.HideEmpty = true

	return apply(base, override...)
}

func apply(base print.Config, override ...print.Config) print.Config { //nolint:gocritic
	dest := base
	for i := range override {
		if err := mergo.Merge(&dest, override[i]); err != nil {
			return base
		}
	}
	return dest
}
