package format

import (
	"fmt"

	"github.com/segmentio/terraform-docs/pkg/print"
)

// Factory initializes and returns the conceret implementation of
// print.Format based on the provided 'name', for example for name
// of 'json' it will return '*format.JSON' through 'format.NewJSON'
// function.
func Factory(name string, settings *print.Settings) (print.Format, error) {
	switch name {
	case "json":
		return NewJSON(settings), nil
	case "markdown":
		return NewTable(settings), nil
	case "markdown document":
		return NewDocument(settings), nil
	case "markdown table":
		return NewTable(settings), nil
	case "pretty":
		return NewPretty(settings), nil
	case "tfvars hcl":
		return NewTfvarsHCL(settings), nil
	case "tfvars json":
		return NewTfvarsJSON(settings), nil
	case "xml":
		return NewXML(settings), nil
	case "yaml":
		return NewYAML(settings), nil
	}
	return nil, fmt.Errorf("formatter '%s' not found", name)
}
