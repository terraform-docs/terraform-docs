package doc

import "sort"

// Output represents a Terraform output.
type Module struct {
	Name        string
	Description string
	Source      string
}

// HasDescription indicates if a Terraform input has a description.
func (i *Module) HasDescription() bool {
	return i.Description != ""
}

type modulesSortedByName []Module

func (a modulesSortedByName) Len() int {
	return len(a)
}

func (a modulesSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a modulesSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

// SortModulesByName sorts a list of inputs by name.
func SortModulesByName(modules []Module) {
	sort.Sort(modulesSortedByName(modules))
}
