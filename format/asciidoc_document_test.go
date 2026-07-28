/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rquadling/terraform-docs/internal/testutil"
	"github.com/rquadling/terraform-docs/print"
)

func TestAsciidocDocument(t *testing.T) {
	tests := map[string]struct {
		config print.Config
	}{
		// Base
		"Base": {
			config: testutil.WithSections(),
		},
		"Empty": {
			config: testutil.WithDefaultSections(
				testutil.With(func(c *print.Config) {
					c.ModuleRoot = "empty"
				}),
			),
		},
		"HideEmpty": {
			config: testutil.WithDefaultSections(
				testutil.WithHideEmpty(),
				testutil.With(func(c *print.Config) {
					c.ModuleRoot = "empty"
				}),
			),
		},
		"HideAll": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Header = false // Since we don't show the header, the file won't be loaded at all
				c.HeaderFrom = "bad.tf"
			}),
		},

		// Settings
		"WithRequired": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Settings.Required = true
				}),
			),
		},
		"WithAnchor": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Settings.Anchor = true
				}),
			),
		},
		"WithoutDefault": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Inputs = true
				c.Settings.Default = false
				c.Settings.Type = true
			}),
		},
		"WithoutType": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Inputs = true
				c.Settings.Default = true
				c.Settings.Type = false
			}),
		},
		"IndentationOfFour": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Settings.Indent = 4
				}),
			),
		},
		"OutputValues": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Outputs = true
				c.OutputValues.Enabled = true
				c.OutputValues.From = "output_values.json"
				c.Settings.Sensitive = true
			}),
		},
		"OutputValuesNoSensitivity": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Outputs = true
				c.OutputValues.Enabled = true
				c.OutputValues.From = "output_values.json"
				c.Settings.Sensitive = false
			}),
		},

		// Only section
		"OnlyDataSources": {
			config: testutil.With(func(c *print.Config) { c.Sections.DataSources = true }),
		},
		"OnlyHeader": {
			config: testutil.With(func(c *print.Config) { c.Sections.Header = true }),
		},
		"OnlyFooter": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Footer = true
				c.FooterFrom = "footer.md"
			}),
		},
		"OnlyInputs": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Inputs = true
				c.Settings.Default = true
				c.Settings.Type = true
			}),
		},
		"OnlyOutputs": {
			config: testutil.With(func(c *print.Config) { c.Sections.Outputs = true }),
		},
		"OnlyModulecalls": {
			config: testutil.With(func(c *print.Config) { c.Sections.ModuleCalls = true }),
		},
		"OnlyProviders": {
			config: testutil.With(func(c *print.Config) { c.Sections.Providers = true }),
		},
		"OnlyRequirements": {
			config: testutil.With(func(c *print.Config) { c.Sections.Requirements = true }),
		},
		"OnlyResources": {
			config: testutil.With(func(c *print.Config) { c.Sections.Resources = true }),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("asciidoc", "document-"+name)
			assert.Nil(err)

			module, err := testutil.GetModule(&tt.config)
			assert.Nil(err)

			formatter := NewAsciidocDocument(&tt.config)

			err = formatter.Generate(module)
			assert.Nil(err)

			assert.Equal(expected, formatter.Content())
		})
	}
}
