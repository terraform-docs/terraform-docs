package json

import (
	"encoding/json"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

const (
	indent string = "  "
	prefix string = ""
)

// Print prints a document as json.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	module.Sort(settings)

	buffer, err := json.MarshalIndent(module, prefix, indent)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
