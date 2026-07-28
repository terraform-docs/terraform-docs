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
	"sort"
	"strings"

	"github.com/rquadling/terraform-docs/internal/types"
	"github.com/rquadling/terraform-docs/print"
)

// Input represents a Terraform input.
type Input struct {
	Name        string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Type        types.String `json:"type" toml:"type" xml:"type" yaml:"type"`
	Description types.String `json:"description" toml:"description" xml:"description" yaml:"description"`
	Default     types.Value  `json:"default" toml:"default" xml:"default" yaml:"default"`
	Required    bool         `json:"required" toml:"required" xml:"required" yaml:"required"`
	Validation  types.List   `json:"validation" toml:"validation" xml:"validation" yaml:"validation"`
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

func sortInputsByName(x []*Input) {
	sort.Slice(x, func(i, j int) bool {
		return x[i].Name < x[j].Name
	})
}

func sortInputsByRequired(x []*Input) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].HasDefault() == x[j].HasDefault() {
			return x[i].Name < x[j].Name
		}
		return !x[i].HasDefault() && x[j].HasDefault()
	})
}

func sortInputsByPosition(x []*Input) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Position.Filename == x[j].Position.Filename {
			return x[i].Position.Line < x[j].Position.Line
		}
		return x[i].Position.Filename < x[j].Position.Filename
	})
}

func sortInputsByType(x []*Input) {
	sort.Slice(x, func(i, j int) bool {
		if x[i].Type == x[j].Type {
			return x[i].Name < x[j].Name
		}
		return x[i].Type < x[j].Type
	})
}

type inputs []*Input

func (ii inputs) sort(enabled bool, by string) {
	if !enabled {
		sortInputsByPosition(ii)
	} else {
		switch by {
		case print.SortType:
			sortInputsByType(ii)
		case print.SortRequired:
			sortInputsByRequired(ii)
		case print.SortName:
			sortInputsByName(ii)
		default:
			sortInputsByPosition(ii)
		}
	}
}
