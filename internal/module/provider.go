package module

import (
	"github.com/terraform-docs/terraform-docs/pkg/tfconf"
)

type providersSortedByName []*tfconf.Provider

func (a providersSortedByName) Len() int      { return len(a) }
func (a providersSortedByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a providersSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}

type providersSortedByPosition []*tfconf.Provider

func (a providersSortedByPosition) Len() int      { return len(a) }
func (a providersSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a providersSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
