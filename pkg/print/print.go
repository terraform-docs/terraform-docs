package print

import (
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// Format represents a printer format (e.g. json, table, yaml, ...)
type Format interface {
	Print(*terraform.Module, *Settings) (string, error)
}
