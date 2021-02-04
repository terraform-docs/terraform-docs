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
	"fmt"
	"strings"

	terraformsdk "github.com/terraform-docs/plugin-sdk/terraform"
	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Input represents a Terraform input.
type Input struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Type        types.String `json:"type" toml:"type" xml:"type" yaml:"type"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Default     types.Value  `json:"default" toml:"default" xml:"default" yaml:"default"`
	Required    bool         `json:"required" toml:"required" xml:"required" yaml:"required"`
	Position    Position     `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// GetValue returns JSON representation of the 'Default' value, which is an 'interface'.
// If 'Default' is a primitive type, the primitive value of 'Default' will be returned
// and not the JSON formatted of it.
func (i *Input) GetValue() string {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i.Default)
	if err != nil {
		panic(err)
	}
	value := strings.TrimSpace(buf.String())
	if value == `null` {
		if i.Required {
			return ""
		}
		return `null` // explicit 'null' value
	}
	return value // everything else
}

// HasDefault indicates if a Terraform variable has a default value set.
func (i *Input) HasDefault() bool {
	return i.Default.HasDefault() || !i.Required
}

type inputs []*Input

func (ii inputs) convert() []*terraformsdk.Input {
	list := []*terraformsdk.Input{}
	for _, i := range ii {
		list = append(list, &terraformsdk.Input{
			Name:        i.Name,
			Type:        fmt.Sprintf("%v", i.Type.Raw()),
			Description: fmt.Sprintf("%v", i.Description.Raw()),
			Default:     i.Default.Raw(),
			Required:    i.Required,
			Position: terraformsdk.Position{
				Filename: i.Position.Filename,
				Line:     i.Position.Line,
			},
		})
	}
	return list
}

type inputsSortedByName []*Input

func (a inputsSortedByName) Len() int           { return len(a) }
func (a inputsSortedByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type inputsSortedByRequired []*Input

func (a inputsSortedByRequired) Len() int      { return len(a) }
func (a inputsSortedByRequired) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByRequired) Less(i, j int) bool {
	if a[i].HasDefault() == a[j].HasDefault() {
		return a[i].Name < a[j].Name
	}
	return !a[i].HasDefault() && a[j].HasDefault()
}

type inputsSortedByPosition []*Input

func (a inputsSortedByPosition) Len() int      { return len(a) }
func (a inputsSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}

type inputsSortedByType []*Input

func (a inputsSortedByType) Len() int      { return len(a) }
func (a inputsSortedByType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByType) Less(i, j int) bool {
	if a[i].Type == a[j].Type {
		return a[i].Name < a[j].Name
	}
	return a[i].Type < a[j].Type
}
