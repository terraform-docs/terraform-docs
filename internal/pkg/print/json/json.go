package json

import (
	"bytes"
	"encoding/json"
	"strings"

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

	copy := &tfconf.Module{
		Header:    "",
		Providers: make([]*tfconf.Provider, 0),
		Inputs:    make([]*tfconf.Input, 0),
		Outputs:   make([]*tfconf.Output, 0),
	}

	if settings.ShowHeader {
		copy.Header = module.Header
	}
	if settings.ShowProviders {
		copy.Providers = module.Providers
	}
	if settings.ShowInputs {
		copy.Inputs = module.Inputs
	}
	if settings.ShowOutputs {
		copy.Outputs = module.Outputs
	}

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(prefix, indent)
	encoder.SetEscapeHTML(settings.EscapeCharacters)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}
