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

	"github.com/terraform-docs/terraform-docs/internal/types"
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

type providersSortedByName []*Provider

func (a providersSortedByName) Len() int      { return len(a) }
func (a providersSortedByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a providersSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}

type providersSortedByPosition []*Provider

func (a providersSortedByPosition) Len() int      { return len(a) }
func (a providersSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a providersSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
