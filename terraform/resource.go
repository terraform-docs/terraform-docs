/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"bytes"
	"sort"
	"strings"
	gotemplate "text/template"

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
	RegistryURL    string       `json:"-" toml:"-" xml:"-" yaml:"-"`
}

const defaultRegistryURLTemplate = "https://registry.terraform.io/providers/{{.Namespace}}/{{.Provider}}/{{.Version}}/docs/{{.Kind}}/{{.Type}}"

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

// URL returns the URL for resource documentation using the configured registry URL template.
// If no custom template is set, it defaults to the Terraform public registry.
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

	parts := strings.SplitN(r.ProviderSource, "/", 2)
	namespace := parts[0]
	provider := parts[0]
	if len(parts) == 2 {
		provider = parts[1]
	}

	version := string(r.Version)
	versionWithV := version
	if version != "latest" && version != "" {
		versionWithV = "v" + version
	}

	urlTemplate := r.RegistryURL
	if urlTemplate == "" {
		urlTemplate = defaultRegistryURLTemplate
	}

	tpl, err := gotemplate.New("registry-url").Parse(urlTemplate)
	if err != nil {
		return ""
	}

	data := struct {
		Namespace    string
		Provider     string
		Version      string
		VersionWithV string
		Kind         string
		Type         string
	}{
		Namespace:    namespace,
		Provider:     provider,
		Version:      version,
		VersionWithV: versionWithV,
		Kind:         kind,
		Type:         r.Type,
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return ""
	}
	return buf.String()
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
