package tfconf

import (
	"bytes"
	"encoding/json"

	"github.com/segmentio/terraform-docs/internal/types"
)

// Output represents a Terraform output.
type Output struct {
	Name        string       `json:"name" yaml:"name"`
	Description types.String `json:"description" yaml:"description"`
	Value       types.Value  `json:"value,omitempty" yaml:"value,omitempty"`
	Sensitive   bool         `json:"sensitive,omitempty" yaml:"sensitive,omitempty"`
	Position    Position     `json:"-" yaml:"-"`
	ShowValue   bool         `json:"-" yaml:"-"`
}

type withvalue struct {
	Name        string       `json:"name" yaml:"name"`
	Description types.String `json:"description" yaml:"description"`
	Value       types.Value  `json:"value" yaml:"value"`
	Sensitive   bool         `json:"sensitive" yaml:"sensitive"`
	Position    Position     `json:"-" yaml:"-"`
	ShowValue   bool         `json:"-" yaml:"-"`
}

// GetValue returns JSON representation of the 'Value', which is an 'interface'.
// If 'Value' is a primitive type, the primitive value of 'Value' will be returned
// and not the JSON formatted of it.
func (o *Output) GetValue() string {
	marshaled, err := json.MarshalIndent(o.Value, "", "  ")
	if err != nil {
		panic(err)
	}
	if value := string(marshaled); value != "null" {
		return value
	}
	return ""
}

// HasDefault indicates if a Terraform output has a default value set.
func (o *Output) HasDefault() bool {
	return o.Value.HasDefault()
}

// MarshalJSON custom yaml marshal function to take
// '--output-values' flag into consideration. It means
// if the flag is not set Value and Sensitive fields
// are set to 'omitempty', otherwise if output values
// are being shown 'omitempty' gets explicitly removed
// to show even empty and false values.
func (o *Output) MarshalJSON() ([]byte, error) {
	fn := func(oo interface{}) ([]byte, error) {
		buf := new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(oo); err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	}
	if o.ShowValue {
		return fn(withvalue(*o))
	}
	return fn(*o)

}

// MarshalYAML custom yaml marshal function to take
// '--output-values' flag into consideration. It means
// if the flag is not set Value and Sensitive fields
// are set to 'omitempty', otherwise if output values
// are being shown 'omitempty' gets explicitly removed
// to show even empty and false values.
func (o *Output) MarshalYAML() (interface{}, error) {
	if o.ShowValue {
		return withvalue(*o), nil
	}
	return o, nil
}
