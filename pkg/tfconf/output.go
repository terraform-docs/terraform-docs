package tfconf

import (
	"github.com/segmentio/terraform-docs/internal/types"
)

// Output represents a Terraform output.
type Output struct {
	Name        string         `json:"name" yaml:"name"`
	Description types.TFString `json:"description" yaml:"description"`
	Value       interface{}    `json:"value,omitempty" yaml:"value,omitempty"`
	Position    Position       `json:"-" yaml:"-"`
}
