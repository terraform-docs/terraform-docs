package doc

import "sort"

// Output represents a Terraform output.
type Output struct {
	Name        string
	Description string
}

// HasDescription indicates if a Terraform output has a description.
func (o *Output) HasDescription() bool {
	return o.Description != ""
}

type outputsSortedByName []Output

func (a outputsSortedByName) Len() int {
	return len(a)
}

func (a outputsSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a outputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

// SortOutputsByName sorts a list of outputs by name.
func SortOutputsByName(outputs []Output) {
	sort.Sort(outputsSortedByName(outputs))
}
