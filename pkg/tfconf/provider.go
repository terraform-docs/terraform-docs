package tfconf

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/types"
)

// Provider represents a Terraform output.
type Provider struct {
	Name     string       `json:"name" yaml:"name"`
	Alias    types.String `json:"alias" yaml:"alias"`
	Version  types.String `json:"version" yaml:"version"`
	Position Position     `json:"-" yaml:"-"`
}

// FullName returns full name of the provider, with alias if available
func (p *Provider) FullName() string {
	if p.Alias != "" {
		return fmt.Sprintf("%s.%s", p.Name, p.Alias)
	}
	return p.Name
}
