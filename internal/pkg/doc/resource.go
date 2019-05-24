package doc

import (
	"sort"
	"strings"
)

type ResourceType string

// Output represents a Terraform output.
type Resource struct {
	Name        string
	Description string
	Type        ResourceType
}

// HasDescription indicates if a Terraform resource has a description.
func (i *Resource) HasDescription() bool {
	return i.Description != ""
}

type resourcesSortedByName []Resource

func (a resourcesSortedByName) Len() int {
	return len(a)
}

func (a resourcesSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a resourcesSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

// SortResourcesByName sorts a list of resources by name.
func SortResourcesByName(resources []Resource) {
	sort.Sort(resourcesSortedByName(resources))
}

type resourcesSortedByType []Resource

func (a resourcesSortedByType) Len() int {
	return len(a)
}

func (a resourcesSortedByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a resourcesSortedByType) Less(i, j int) bool {
	return a[i].Type < a[j].Type
}

// SortResourcesByType sorts a list of resources by type.
func SortResourcesByType(resources []Resource) {
	sort.Sort(resourcesSortedByType(resources))
}

func (rt ResourceType) Provider() string {
	return strings.SplitN(string(rt), "_", 2)[0]
}

func (rt ResourceType) Name() string {
	return strings.SplitN(string(rt), "_", 2)[1]
}
