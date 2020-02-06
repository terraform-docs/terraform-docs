package tfconf

import (
	"encoding/json"
)

// Output represents a Terraform output.
type Output struct {
	Name        string      `json:"name" yaml:"name"`
	Description String      `json:"description" yaml:"description"`
	Value       interface{} `json:"value" yaml:"value"`
	Position    Position    `json:"-" yaml:"-"`
}

// GetValue returns JSON representation of the 'Default' value, which is an 'interface'.
// If 'Default' is a primitive type, the primitive value of 'Default' will be returned
// and not the JSON formatted of it.
func (o *Output) GetValue() string {
	marshaled, err := json.MarshalIndent(o.Value, "", "  ")
	if err != nil {
		panic(err)
	}
	if v := string(marshaled); v != "null" {
		return v
	}
	return ""
}

type outputsSortedByName []*Output

func (a outputsSortedByName) Len() int {
	return len(a)
}

func (a outputsSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a outputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type outputsSortedByPosition []*Output

func (a outputsSortedByPosition) Len() int {
	return len(a)
}

func (a outputsSortedByPosition) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a outputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}

// TerraformOutput is used for unmarshalling `terraform outputs --json` into
type TerraformOutput struct {
	Sensitive bool        `json:"sensitive"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
}
