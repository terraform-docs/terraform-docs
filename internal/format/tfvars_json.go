package format

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/iancoleman/orderedmap"

	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/pkg/print"
)

// TfvarsJSON represents Terraform tfvars JSON format.
type TfvarsJSON struct{}

// NewTfvarsJSON returns new instance of TfvarsJSON.
func NewTfvarsJSON(settings *print.Settings) *TfvarsJSON {
	return &TfvarsJSON{}
}

// Print prints a Terraform module as Terraform tfvars JSON document.
func (j *TfvarsJSON) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	copy := orderedmap.New()
	copy.SetEscapeHTML(false)
	for _, i := range module.Inputs {
		copy.Set(i.Name, i.Default)
	}

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}
