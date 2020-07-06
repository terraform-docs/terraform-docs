package tfconf

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
