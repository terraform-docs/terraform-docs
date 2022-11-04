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

	"github.com/terraform-docs/terraform-docs/internal/types"
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

// GetResourceType returns the type of a specific resource in the config.
// Concatenating it with the provider name
// It takes the form: provider-name_resource-type
// e.g. aws_iam_role
func (r *Resource) GetResourceType() string {
	return r.ProviderName + "_" + r.Type
}

// GetResourceName returns the name of a specific resource in the config.
func (r *Resource) GetResourceName() string {
	return r.Name
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
			if x[i].GetResourceType() == x[j].GetResourceType() {
				return x[i].Name <= x[j].Name
			}
			return x[i].GetResourceType() < x[j].GetResourceType()
		}
		return x[i].Mode > x[j].Mode
	})
}

type resources []*Resource

func (rr resources) sort(enabled bool, by string) { //nolint:unparam
	// always sort by type
	sortResourcesByType(rr)
}
