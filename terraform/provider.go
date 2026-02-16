/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"fmt"
	"sort"

	"github.com/rquadling/terraform-docs/internal/types"
)

// Provider represents a Terraform output.
type Provider struct {
	Name     string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Alias    types.String `json:"alias" toml:"alias" xml:"alias" yaml:"alias"`
	Version  types.String `json:"version" toml:"version" xml:"version" yaml:"version"`
	Position Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// FullName returns full name of the provider, with alias if available
func (p *Provider) FullName() string {
	if p.Alias != "" {
		return fmt.Sprintf("%s.%s", p.Name, p.Alias)
	}
	return p.Name
}

func sortProvidersByName(x []*Provider) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Name == x[j].Name {
			return x[i].Name == x[j].Name && x[i].Alias < x[j].Alias
		}
		return x[i].Name < x[j].Name
	})
}

func sortProvidersByPosition(x []*Provider) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Position.Filename == x[j].Position.Filename {
			return x[i].Position.Line < x[j].Position.Line
		}
		return x[i].Position.Filename < x[j].Position.Filename
	})
}

type providers []*Provider

func (pp providers) sort(enabled bool, by string) { //nolint:unparam
	if !enabled {
		sortProvidersByPosition(pp)
	} else {
		// always sort by name if sorting is enabled
		sortProvidersByName(pp)
	}
}
