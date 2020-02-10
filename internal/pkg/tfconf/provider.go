package tfconf

import (
	"fmt"
)

// Provider represents a Terraform output.
type Provider struct {
	Name     string   `json:"name" yaml:"name"`
	Alias    String   `json:"alias" yaml:"alias"`
	Version  String   `json:"version" yaml:"version"`
	Position Position `json:"-" yaml:"-"`
}

// FullName returns full name of the provider, with alias if available
func (p *Provider) FullName() string {
	if p.Alias != "" {
		return fmt.Sprintf("%s.%s", p.Name, p.Alias)
	}
	return p.Name
}

type providersSortedByName []*Provider

func (a providersSortedByName) Len() int {
	return len(a)
}

func (a providersSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a providersSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}

type providersSortedByPosition []*Provider

func (a providersSortedByPosition) Len() int {
	return len(a)
}

func (a providersSortedByPosition) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a providersSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
