package tfconf

import (
	"encoding/json"
)

// Input represents a Terraform input.
type Input struct {
	Name        string      `json:"name"`
	Type        String      `json:"type"`
	Description String      `json:"description"`
	Default     interface{} `json:"default"`
	Position    Position    `json:"-"`
}

// ValueOf returns JSON representation of the 'Default' value, which is an 'interface'.
func ValueOf(v interface{}) string {
	marshaled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(marshaled)
}

// HasDefault indicates if a Terraform variable has a default value set.
func (i *Input) HasDefault() bool {
	return i.Default != nil
}

type inputsSortedByName []*Input

func (a inputsSortedByName) Len() int {
	return len(a)
}

func (a inputsSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a inputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type inputsSortedByRequired []*Input

func (a inputsSortedByRequired) Len() int {
	return len(a)
}

func (a inputsSortedByRequired) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a inputsSortedByRequired) Less(i, j int) bool {
	switch {
	// i required, j not: i gets priority
	case !a[i].HasDefault() && a[j].HasDefault():
		return true
	// j required, i not: i does not get priority
	case a[i].HasDefault() && !a[j].HasDefault():
		return false
	// Otherwise, sort by name
	default:
		return a[i].Name < a[j].Name
	}
}

type inputsSortedByPosition []*Input

func (a inputsSortedByPosition) Len() int {
	return len(a)
}

func (a inputsSortedByPosition) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a inputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
