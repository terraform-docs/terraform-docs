package tfconf

import (
	"encoding/json"

	"github.com/segmentio/terraform-docs/internal/types"
)

// Input represents a Terraform input.
type Input struct {
	Name        string        `json:"name" yaml:"name"`
	Type        types.String  `json:"type" yaml:"type"`
	Description types.String  `json:"description" yaml:"description"`
	Default     types.Default `json:"default" yaml:"default"`
	Position    Position      `json:"-" yaml:"-"`
}

// Value returns JSON representation of the 'Default' value, which is an 'interface'.
// If 'Default' is a primitive type, the primitive value of 'Default' will be returned
// and not the JSON formatted of it.
func (i *Input) Value() string {
	marshaled, err := json.MarshalIndent(i.Default, "", "  ")
	if err != nil {
		panic(err)
	}
	if value := string(marshaled); value != "null" {
		return value
	}
	return ""
}

// HasDefault indicates if a Terraform variable has a default value set.
func (i *Input) HasDefault() bool {
	return i.Default.HasDefault()
}
