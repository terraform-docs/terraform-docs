package terraform

import (
	"fmt"
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Resource represents a managed or data type that is created by the module
type Resource struct {
	Type           string       `json:"type" toml:"type" xml:"type" yaml:"type"`
	ProviderName   string       `json:"providerName" toml:"providerName" xml:"providerName" yaml:"providerName"`
	ProviderSource string       `json:"provicerSource" toml:"providerSource" xml:"providerSource" yaml:"providerSource"`
	Mode           string       `json:"mode" toml:"mode" xml:"mode" yaml:"mode"`
	Version        types.String `json:"version" toml:"version" xml:"version" yaml:"version"`
}

type resourcesSortedByType []*Resource

// FullType returns full name of the type of the resource, including the provider name
func (r *Resource) FullType() string {
	return r.ProviderName + "_" + r.Type
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

func (a resourcesSortedByType) Len() int      { return len(a) }
func (a resourcesSortedByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a resourcesSortedByType) Less(i, j int) bool {
	return a[i].FullType() < a[j].FullType() || (a[i].FullType() == a[j].FullType())
}
