package tfconf

import (
	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Requirement represents a requirement for Terraform module.
type Requirement struct {
	Name    string       `json:"name" xml:"name" yaml:"name"`
	Version types.String `json:"version" xml:"version" yaml:"version"`
}
