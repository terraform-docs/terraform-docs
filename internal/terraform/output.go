/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"bytes"
	"encoding/json"
	"encoding/xml"

	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Output represents a Terraform output.
type Output struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Value       types.Value  `json:"value,omitempty" toml:"value,omitempty" xml:"value,omitempty" yaml:"value,omitempty"`
	Sensitive   bool         `json:"sensitive,omitempty" toml:"sensitive,omitempty" xml:"sensitive,omitempty" yaml:"sensitive,omitempty"`
	Position    Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
	ShowValue   bool         `json:"-" toml:"-" xml:"-" yaml:"-"`
}

type withvalue struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Value       types.Value  `json:"value" toml:"value" xml:"value" yaml:"value"`
	Sensitive   bool         `json:"sensitive" toml:"sensitive" xml:"sensitive" yaml:"sensitive"`
	Position    Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
	ShowValue   bool         `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// GetValue returns JSON representation of the 'Value', which is an 'interface'.
// If 'Value' is a primitive type, the primitive value of 'Value' will be returned
// and not the JSON formatted of it.
func (o *Output) GetValue() string {
	if !o.ShowValue || o.Value == nil {
		return ""
	}
	marshaled, err := json.MarshalIndent(o.Value, "", "  ")
	if err != nil {
		panic(err)
	}
	value := string(marshaled)
	if value == `null` {
		return "" // types.Nil
	}
	return value // everything else
}

// HasDefault indicates if a Terraform output has a default value set.
func (o *Output) HasDefault() bool {
	if !o.ShowValue || o.Value == nil {
		return false
	}
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
	o.Value = nil       // explicitly make empty
	o.Sensitive = false // explicitly make empty
	return fn(*o)
}

// MarshalXML custom xml marshal function to take
// '--output-values' flag into consideration. It means
// if the flag is not set Value and Sensitive fields
// are set to 'omitempty', otherwise if output values
// are being shown 'omitempty' gets explicitly removed
// to show even empty and false values.
func (o *Output) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	fn := func(v interface{}, name string) error {
		return e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: name}})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	fn(o.Name, "name")               //nolint: errcheck
	fn(o.Description, "description") //nolint: errcheck
	if o.ShowValue {
		fn(o.Value, "value")         //nolint: errcheck
		fn(o.Sensitive, "sensitive") //nolint: errcheck
	}
	return e.EncodeToken(start.End())
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
	o.Value = nil       // explicitly make empty
	o.Sensitive = false // explicitly make empty
	return *o, nil
}

// output is used for unmarshalling `terraform outputs --json` into
type output struct {
	Sensitive bool        `json:"sensitive"`
	Type      interface{} `json:"type"`
	Value     interface{} `json:"value"`
}

type outputsSortedByName []*Output

func (a outputsSortedByName) Len() int      { return len(a) }
func (a outputsSortedByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a outputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type outputsSortedByPosition []*Output

func (a outputsSortedByPosition) Len() int      { return len(a) }
func (a outputsSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a outputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
