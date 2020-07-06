package tfconf

import (
	"encoding/json"

	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Input represents a Terraform input.
type Input struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Type        types.String `json:"type" toml:"type" xml:"type" yaml:"type"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Default     types.Value  `json:"default" toml:"default" xml:"default" yaml:"default"`
	Required    bool         `json:"required" toml:"required" xml:"required" yaml:"required"`
	Position    Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// GetValue returns JSON representation of the 'Default' value, which is an 'interface'.
// If 'Default' is a primitive type, the primitive value of 'Default' will be returned
// and not the JSON formatted of it.
func (i *Input) GetValue() string {
	marshaled, err := json.MarshalIndent(i.Default, "", "  ")
	if err != nil {
		panic(err)
	}
	value := string(marshaled)
	if value == `null` {
		if i.Required {
			return ""
		}
		return `null` // explicit 'null' value
	}
	return value // everything else
}

// HasDefault indicates if a Terraform variable has a default value set.
func (i *Input) HasDefault() bool {
	return i.Default.HasDefault() || !i.Required
}
