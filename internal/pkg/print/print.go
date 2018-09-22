package print

import (
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

const (
	_ settings.Setting = iota
	// WithRequired prints if inputs are required
	WithRequired
)
