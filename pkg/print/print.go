package print

import (
	"github.com/terraform-docs/terraform-docs/pkg/tfconf"
)

// Format represents a printer format (e.g. json, table, yaml, ...)
type Format interface {
	Print(*tfconf.Module, *Settings) (string, error)
}
