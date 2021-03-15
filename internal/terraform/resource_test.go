/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/types"
)

func TestResourceSpec(t *testing.T) {
	assert := assert.New(t)
	resource := Resource{
		Type:           "private_key",
		Name:           "baz",
		ProviderName:   "tls",
		ProviderSource: "hashicorp/tls",
		Mode:           "managed",
		Version:        types.String("latest"),
	}
	assert.Equal("tls_private_key.baz", resource.Spec())
}

func TestResourceMode(t *testing.T) {
	tests := []struct {
		name        string
		resource    Resource
		expectValue string
	}{
		{
			name: "resource mode as managed",
			resource: Resource{
				Type:           "private_key",
				ProviderName:   "tls",
				ProviderSource: "hashicorp/tls",
				Mode:           "managed",
				Version:        types.String("latest"),
			},
			expectValue: "resource",
		},
		{
			name: "resource mode as data source",
			resource: Resource{
				Type:           "caller_identity",
				ProviderName:   "aws",
				ProviderSource: "hashicorp/aws",
				Mode:           "data",
				Version:        types.String("latest"),
			},
			expectValue: "data source",
		},
		{
			name: "resource mode as invalid",
			resource: Resource{
				Type:           "caller_identity",
				ProviderName:   "aws",
				ProviderSource: "hashicorp/aws",
				Mode:           "",
				Version:        types.String("latest"),
			},
			expectValue: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expectValue, tt.resource.GetMode())
		})
	}
}

func TestResourceURL(t *testing.T) {
	tests := []struct {
		name        string
		resource    Resource
		expectValue string
	}{
		{
			name: "generic URL construction",
			resource: Resource{
				Type:           "private_key",
				ProviderName:   "tls",
				ProviderSource: "hashicorp/tls",
				Mode:           "managed",
				Version:        types.String("latest"),
			},
			expectValue: "https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key",
		},
		{
			name: "unable to construct URL",
			resource: Resource{
				Type:           "custom",
				ProviderName:   "nih",
				ProviderSource: "http://nih.tld/some/path/to/provider/source",
				Mode:           "managed",
				Version:        types.String("latest"),
			},
			expectValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expectValue, tt.resource.URL())
		})
	}
}

func TestResourcesSortedByType(t *testing.T) {
	assert := assert.New(t)
	resources := sampleResources()

	sort.Sort(resourcesSortedByType(resources))

	expected := []string{"a_a.a", "a_f.f", "b_b.b", "b_d.d", "c_c.c", "c_e.c", "c_e.d", "c_e_x.c", "c_e_x.d", "z_z.z", "a_a.a", "z_z.z", "a_a.a", "z_z.z"}
	actual := make([]string, len(resources))

	for k, i := range resources {
		actual[k] = i.Spec()
	}

	assert.Equal(expected, actual)
}

func TestResourcesSortedByTypeAndMode(t *testing.T) {
	assert := assert.New(t)
	resources := sampleResources()

	sort.Sort(resourcesSortedByType(resources))

	expected := []string{"a_a.a (r)", "a_f.f (r)", "b_b.b (r)", "b_d.d (r)", "c_c.c (r)", "c_e.c (r)", "c_e.d (r)", "c_e_x.c (r)", "c_e_x.d (r)", "z_z.z (r)", "a_a.a (d)", "z_z.z (d)", "a_a.a", "z_z.z"}
	actual := make([]string, len(resources))

	for k, i := range resources {
		mode := ""
		switch i.Mode {
		case "managed":
			mode = " (r)"
		case "data":
			mode = " (d)"
		}
		actual[k] = i.Spec() + mode
	}

	assert.Equal(expected, actual)
}

func TestResourceVersion(t *testing.T) {
	tests := []struct {
		name       string
		constraint []string
		expected   string
	}{
		{
			name:       "exact version, without operator",
			constraint: []string{"1.2.3"},
			expected:   "1.2.3",
		},
		{
			name:       "exact version, with operator",
			constraint: []string{"= 1.2.3"},
			expected:   "1.2.3",
		},
		{
			name:       "exact version, with operator, without space",
			constraint: []string{"=1.2.3"},
			expected:   "1.2.3",
		},
		{
			name:       "exclude exact version, with space",
			constraint: []string{"!= 1.2.3"},
			expected:   "latest",
		},
		{
			name:       "exclude exact version, without space",
			constraint: []string{"!=1.2.3"},
			expected:   "latest",
		},
		{
			name:       "comparison version, with space",
			constraint: []string{"> 1.2.3"},
			expected:   "latest",
		},
		{
			name:       "comparison version, without space",
			constraint: []string{">1.2.3"},
			expected:   "latest",
		},
		{
			name:       "range version",
			constraint: []string{"> 1.2.3, < 2.0.0"},
			expected:   "latest",
		},
		{
			name:       "pessimistic version, with space",
			constraint: []string{"~> 1.2.3"},
			expected:   "latest",
		},
		{
			name:       "pessimistic version, without space",
			constraint: []string{"~>1.2.3"},
			expected:   "latest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expected, resourceVersion(tt.constraint))
		})
	}
}

func sampleResources() []*Resource {
	return []*Resource{
		{
			Type:           "e",
			Name:           "d",
			ProviderName:   "c",
			ProviderSource: "hashicorp/e",
			Mode:           "managed",
			Version:        "1.5.0",
		},
		{
			Type:           "e",
			Name:           "c",
			ProviderName:   "c",
			ProviderSource: "hashicorp/e",
			Mode:           "managed",
			Version:        "1.5.0",
		},
		{
			Type:           "e_x",
			Name:           "d",
			ProviderName:   "c",
			ProviderSource: "hashicorp/e",
			Mode:           "managed",
			Version:        "1.5.0",
		},
		{
			Type:           "e_x",
			Name:           "c",
			ProviderName:   "c",
			ProviderSource: "hashicorp/e",
			Mode:           "managed",
			Version:        "1.5.0",
		},
		{
			Type:           "a",
			Name:           "a",
			ProviderName:   "a",
			ProviderSource: "hashicorp/a",
			Mode:           "managed",
			Version:        "1.1.0",
		},
		{
			Type:           "d",
			Name:           "d",
			ProviderName:   "b",
			ProviderSource: "hashicorp/d",
			Mode:           "managed",
			Version:        "1.4.0",
		},
		{
			Type:           "b",
			Name:           "b",
			ProviderName:   "b",
			ProviderSource: "hashicorp/b",
			Mode:           "managed",
			Version:        "1.2.0",
		},
		{
			Type:           "c",
			Name:           "c",
			ProviderName:   "c",
			ProviderSource: "hashicorp/c",
			Mode:           "managed",
			Version:        "1.3.0",
		},
		{
			Type:           "f",
			Name:           "f",
			ProviderName:   "a",
			ProviderSource: "hashicorp/f",
			Mode:           "managed",
			Version:        "1.6.0",
		},
		{
			ProviderName:   "z",
			Type:           "z",
			Name:           "z",
			ProviderSource: "hashicorp/a",
			Mode:           "managed",
			Version:        "1.5.0",
		},
		{
			ProviderName:   "z",
			Type:           "z",
			Name:           "z",
			ProviderSource: "hashicorp/a",
			Mode:           "data",
			Version:        "1.5.0",
		},
		{
			ProviderName:   "a",
			Type:           "a",
			Name:           "a",
			ProviderSource: "hashicorp/a",
			Mode:           "data",
			Version:        "1.5.0",
		},
		{
			ProviderName:   "z",
			Type:           "z",
			Name:           "z",
			ProviderSource: "hashicorp/a",
			Mode:           "",
			Version:        "1.5.0",
		},
		{
			ProviderName:   "a",
			Type:           "a",
			Name:           "a",
			ProviderSource: "hashicorp/a",
			Mode:           "",
			Version:        "1.5.0",
		},
	}
}
