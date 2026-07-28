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
	"github.com/rquadling/terraform-docs/print"
)

// ModuleCall represents a submodule called by Terraform module.
type ModuleCall struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Source      string       `json:"source" toml:"source" xml:"source" yaml:"source"`
	Version     string       `json:"version" toml:"version" xml:"version" yaml:"version"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Position    Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// FullName returns full name of the modulecall, with version if available
func (mc *ModuleCall) FullName() string {
	if mc.Version != "" {
		return fmt.Sprintf("%s,%s", mc.Source, mc.Version)
	}
	return mc.Source
}

func sortModulecallsByName(x []*ModuleCall) {
	sort.Slice(x, func(i, j int) bool {
		return x[i].Name < x[j].Name
	})
}

func sortModulecallsBySource(x []*ModuleCall) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Source == x[j].Source {
			return x[i].Name < x[j].Name
		}
		return x[i].Source < x[j].Source
	})
}

func sortModulecallsByPosition(x []*ModuleCall) {
	sort.Slice(x, func(i, j int) bool {
		return x[i].Position.Filename < x[j].Position.Filename || x[i].Position.Line < x[j].Position.Line
	})
}

type modulecalls []*ModuleCall

func (mm modulecalls) sort(enabled bool, by string) {
	if !enabled {
		sortModulecallsByPosition(mm)
	} else {
		switch by {
		case print.SortName, print.SortRequired:
			sortModulecallsByName(mm)
		case print.SortType:
			sortModulecallsBySource(mm)
		default:
			sortModulecallsByPosition(mm)
		}
	}
}
