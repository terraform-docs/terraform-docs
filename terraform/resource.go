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
	"strings"

	"github.com/rquadling/terraform-docs/internal/types"
)

// Resource represents a managed or data type that is created by the module
type Resource struct {
	Type           string       `json:"type" toml:"type" xml:"type" yaml:"type"`
	Name           string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	ProviderName   string       `json:"provider" toml:"provider" xml:"provider" yaml:"provider"`
	ProviderSource string       `json:"source" toml:"source" xml:"source" yaml:"source"`
	Mode           string       `json:"mode" toml:"mode" xml:"mode" yaml:"mode"`
	Version        types.String `json:"version" toml:"version" xml:"version" yaml:"version"`
	Description    types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Position       Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// Spec returns the resource spec addresses a specific resource in the config.
// It takes the form: resource_type.resource_name[resource index]
// For more details, see:
// https://www.terraform.io/docs/cli/state/resource-addressing.html#resource-spec
func (r *Resource) Spec() string {
	return r.ProviderName + "_" + r.Type + "." + r.Name
}

// GetMode returns normalized resource type as "resource" or "data source"
func (r *Resource) GetMode() string {
	switch r.Mode {
	case "managed":
		return "resource"
	case "data":
		return "data source"
	default:
		return "invalid"
	}
}

// URL returns a best guess at the URL for resource documentation
func (r *Resource) URL() string {
	kind := ""
	switch r.Mode {
	case "managed":
		kind = "resources"
	case "data":
		kind = "data-sources"
	default:
		return ""
	}

	if strings.Count(r.ProviderSource, "/") > 1 {
		return ""
	}
	return fmt.Sprintf("https://registry.terraform.io/providers/%s/%s/docs/%s/%s", r.ProviderSource, r.Version, kind, r.Type)
}

func sortResourcesByType(x []*Resource) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Mode == x[j].Mode {
			if x[i].Spec() == x[j].Spec() {
				return x[i].Name <= x[j].Name
			}
			return x[i].Spec() < x[j].Spec()
		}
		return x[i].Mode > x[j].Mode
	})
}

type resources []*Resource

func (rr resources) sort(enabled bool, by string) { //nolint:unparam
	// always sort by type
	sortResourcesByType(rr)
}
