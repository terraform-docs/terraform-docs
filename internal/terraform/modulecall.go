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

	terraformsdk "github.com/terraform-docs/plugin-sdk/terraform"
)

// ModuleCall represents a submodule called by Terraform module.
type ModuleCall struct {
	Name     string   `json:"name"`
	Source   string   `json:"source"`
	Version  string   `json:"version,omitempty"`
	Position Position `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// FullName returns full name of the modulecall, with version if available
func (mc *ModuleCall) FullName() string {
	if mc.Version != "" {
		return fmt.Sprintf("%s,%s", mc.Source, mc.Version)
	}
	return mc.Source
}

type modulecallsSortedByName []*ModuleCall

func (a modulecallsSortedByName) Len() int           { return len(a) }
func (a modulecallsSortedByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a modulecallsSortedByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type modulecallsSortedBySource []*ModuleCall

func (a modulecallsSortedBySource) Len() int      { return len(a) }
func (a modulecallsSortedBySource) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a modulecallsSortedBySource) Less(i, j int) bool {
	if a[i].Source == a[j].Source {
		return a[i].Name < a[j].Name
	}
	return a[i].Source < a[j].Source
}

type modulecallsSortedByPosition []*ModuleCall

func (a modulecallsSortedByPosition) Len() int      { return len(a) }
func (a modulecallsSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a modulecallsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}

type modulecalls []*ModuleCall

func (mm modulecalls) convert() []*terraformsdk.ModuleCall {
	list := []*terraformsdk.ModuleCall{}
	for _, m := range mm {
		list = append(list, &terraformsdk.ModuleCall{
			Name:    m.Name,
			Source:  m.Source,
			Version: m.Version,
			Position: terraformsdk.Position{
				Filename: m.Position.Filename,
				Line:     m.Position.Line,
			},
		})
	}
	return list
}
