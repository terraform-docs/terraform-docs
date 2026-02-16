/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	jsonsdk "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rquadling/terraform-docs/internal/testutil"
	"github.com/rquadling/terraform-docs/print"
)

func TestCommonSort(t *testing.T) {
	tests := map[string]struct {
		config print.Config
	}{
		"NoSort": {
			config: *print.NewConfig(),
		},
		"SortByName": {
			config: testutil.With(func(c *print.Config) {
				c.Sort.Enabled = true
				c.Sort.By = print.SortName
			}),
		},
		"SortByRequired": {
			config: testutil.With(func(c *print.Config) {
				c.Sort.Enabled = true
				c.Sort.By = print.SortRequired
			}),
		},
		"SortByType": {
			config: testutil.With(func(c *print.Config) {
				c.Sort.Enabled = true
				c.Sort.By = print.SortType
			}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			module, err := testutil.GetModule(&tt.config)
			assert.Nil(err)

			type Expected struct {
				Inputs       []string `json:"inputs"`
				Modules      []string `json:"modules"`
				Outputs      []string `json:"outputs"`
				Providers    []string `json:"providers"`
				Requirements []string `json:"requirements"`
				Resources    []string `json:"resources"`
			}

			golden, err := testutil.GetExpected("common", "sort-"+name)
			assert.Nil(err)

			var expected Expected

			err = jsonsdk.Unmarshal([]byte(golden), &expected)
			assert.Nil(err)

			for ii, i := range module.Inputs {
				assert.Equal(expected.Inputs[ii], i.Name)
			}
			for ii, m := range module.ModuleCalls {
				assert.Equal(expected.Modules[ii], m.Name+"-"+m.Source)
			}
			for ii, o := range module.Outputs {
				assert.Equal(expected.Outputs[ii], o.Name)
			}
			for ii, p := range module.Providers {
				assert.Equal(expected.Providers[ii], p.FullName())
			}
			for ii, r := range module.Requirements {
				assert.Equal(expected.Requirements[ii], r.Name)
			}
			for ii, r := range module.Resources {
				assert.Equal(expected.Resources[ii], r.Spec()+"__"+r.Mode)
			}
		})
	}
}

func TestCommonHeaderFrom(t *testing.T) {
	tests := map[string]struct {
		config print.Config
	}{
		"HeaderFromADOCFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Header = true
				c.HeaderFrom = "doc.adoc"
			}),
		},
		"HeaderFromMDFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Header = true
				c.HeaderFrom = "doc.md"
			}),
		},
		"HeaderFromTFFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Header = true
				c.HeaderFrom = "doc.tf"
			}),
		},
		"HeaderFromTXTFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Header = true
				c.HeaderFrom = "doc.txt"
			}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("common", "header-"+name)
			assert.Nil(err)

			module, err := testutil.GetModule(&tt.config)
			assert.Nil(err)

			assert.Equal(expected, module.Header)
		})
	}
}

func TestCommonFooterFrom(t *testing.T) {
	tests := map[string]struct {
		config print.Config
	}{
		"FooterFromADOCFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Footer = true
				c.FooterFrom = "doc.adoc"
			}),
		},
		"FooterFromMDFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Footer = true
				c.FooterFrom = "doc.md"
			}),
		},
		"FooterFromTFFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Footer = true
				c.FooterFrom = "doc.tf"
			}),
		},
		"FooterFromTXTFile": {
			config: testutil.With(func(c *print.Config) {
				c.Sections.Footer = true
				c.FooterFrom = "doc.txt"
			}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("common", "footer-"+name)
			assert.Nil(err)

			module, err := testutil.GetModule(&tt.config)
			assert.Nil(err)

			assert.Equal(expected, module.Footer)
		})
	}
}
