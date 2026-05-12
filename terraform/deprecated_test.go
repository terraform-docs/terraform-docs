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
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/print"
)

func TestInputIsDeprecated(t *testing.T) {
	assert := assert.New(t)

	assert.False((&Input{}).IsDeprecated(), "empty Deprecated should report false")
	assert.False((&Input{Deprecated: types.String("")}).IsDeprecated(), "empty string Deprecated should report false")
	assert.True((&Input{Deprecated: types.String("use var.new")}).IsDeprecated(), "non-empty Deprecated should report true")
}

func TestOutputIsDeprecated(t *testing.T) {
	assert := assert.New(t)

	assert.False((&Output{}).IsDeprecated(), "empty Deprecated should report false")
	assert.False((&Output{Deprecated: types.String("")}).IsDeprecated(), "empty string Deprecated should report false")
	assert.True((&Output{Deprecated: types.String("use output.new")}).IsDeprecated(), "non-empty Deprecated should report true")
}

func TestInputMarshalDeprecated(t *testing.T) {
	tests := []struct {
		name           string
		input          Input
		jsonContains   string
		jsonOmits      string
		yamlContains   string
		xmlContains    string
		xmlOmits       string
	}{
		{
			name:         "deprecated set surfaces in every encoding",
			input:        Input{Name: "old", Type: types.String("string"), Default: types.ValueOf(nil), Deprecated: types.String("use var.new")},
			jsonContains: `"deprecated":"use var.new"`,
			yamlContains: "deprecated: use var.new",
			xmlContains:  "<deprecated>use var.new</deprecated>",
		},
		{
			name:      "deprecated empty is omitted from every encoding",
			input:     Input{Name: "current", Type: types.String("string"), Default: types.ValueOf(nil)},
			jsonOmits: `deprecated`,
			xmlOmits:  "<deprecated>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			// Input has no custom Marshal* methods — struct tags drive the output.
			jsonBytes, err := json.Marshal(tt.input)
			assert.Nil(err)
			if tt.jsonContains != "" {
				assert.Contains(string(jsonBytes), tt.jsonContains)
			}
			if tt.jsonOmits != "" {
				assert.NotContains(string(jsonBytes), tt.jsonOmits)
			}

			if tt.yamlContains != "" {
				yamlBytes, err := yaml.Marshal(tt.input)
				assert.Nil(err)
				assert.Contains(string(yamlBytes), tt.yamlContains)
			}

			var xmlBuf bytes.Buffer
			err = xml.NewEncoder(&xmlBuf).Encode(tt.input)
			assert.Nil(err)
			if tt.xmlContains != "" {
				assert.Contains(xmlBuf.String(), tt.xmlContains)
			}
			if tt.xmlOmits != "" {
				assert.NotContains(xmlBuf.String(), tt.xmlOmits)
			}
		})
	}
}

func TestOutputMarshalDeprecated(t *testing.T) {
	tests := []struct {
		name         string
		output       Output
		jsonContains string
		jsonOmits    string
		yamlContains string
		xmlContains  string
		xmlOmits     string
	}{
		{
			name:         "deprecated + ShowValue=false surfaces in every encoding",
			output:       Output{Name: "legacy", Description: types.String("d"), Deprecated: types.String("use output.new")},
			jsonContains: `"deprecated":"use output.new"`,
			yamlContains: "deprecated: use output.new",
			xmlContains:  "<deprecated>use output.new</deprecated>",
		},
		{
			name:         "deprecated + ShowValue=true still surfaces (regression for Output.MarshalXML)",
			output:       Output{Name: "legacy", Description: types.String("d"), Value: types.ValueOf("v"), ShowValue: true, Deprecated: types.String("use output.new")},
			jsonContains: `"deprecated":"use output.new"`,
			yamlContains: "deprecated: use output.new",
			xmlContains:  "<deprecated>use output.new</deprecated>",
		},
		{
			name:      "deprecated empty + ShowValue=false is omitted",
			output:    Output{Name: "current", Description: types.String("d")},
			jsonOmits: `deprecated`,
			xmlOmits:  "<deprecated>",
		},
		{
			name:      "deprecated empty + ShowValue=true is omitted",
			output:    Output{Name: "current", Description: types.String("d"), Value: types.ValueOf("v"), ShowValue: true},
			jsonOmits: `deprecated`,
			xmlOmits:  "<deprecated>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			out := tt.output

			jsonBytes, err := out.MarshalJSON()
			assert.Nil(err)
			if tt.jsonContains != "" {
				assert.Contains(string(jsonBytes), tt.jsonContains)
			}
			if tt.jsonOmits != "" {
				assert.NotContains(string(jsonBytes), tt.jsonOmits)
			}

			// Reset ShowValue-mutated state for the YAML/XML passes.
			out = tt.output
			if tt.yamlContains != "" {
				yamlValue, err := out.MarshalYAML()
				assert.Nil(err)
				yamlBytes, err := yaml.Marshal(yamlValue)
				assert.Nil(err)
				assert.Contains(string(yamlBytes), tt.yamlContains)
			}

			out = tt.output
			var xmlBuf bytes.Buffer
			err = xml.NewEncoder(&xmlBuf).EncodeElement(&out, xml.StartElement{Name: xml.Name{Local: "output"}})
			assert.Nil(err)
			if tt.xmlContains != "" {
				assert.Contains(xmlBuf.String(), tt.xmlContains)
			}
			if tt.xmlOmits != "" {
				assert.NotContains(xmlBuf.String(), tt.xmlOmits)
			}
		})
	}
}

func TestLoadInputsDeprecated(t *testing.T) {
	assert := assert.New(t)

	meta, files, err := loadModule("testdata/deprecated")
	assert.Nil(err)

	positions := extractBlockPositions(files, "variable")
	all, _, _ := loadInputs(meta, positions, print.NewConfig())

	byName := map[string]*Input{}
	for _, in := range all {
		byName[in.Name] = in
	}

	assert.True(byName["old_name"].IsDeprecated(), "old_name should be deprecated")
	assert.Equal("Will be removed in v2.0; use var.new_name.", string(byName["old_name"].Deprecated))
	assert.False(byName["new_name"].IsDeprecated(), "new_name should not be deprecated")
}

func TestLoadOutputsDeprecated(t *testing.T) {
	assert := assert.New(t)

	meta, files, err := loadModule("testdata/deprecated")
	assert.Nil(err)

	positions := extractBlockPositions(files, "output")
	outs, err := loadOutputs(meta, positions, print.NewConfig())
	assert.Nil(err)

	byName := map[string]*Output{}
	for _, o := range outs {
		byName[o.Name] = o
	}

	assert.True(byName["legacy_id"].IsDeprecated(), "legacy_id should be deprecated")
	assert.Equal("Use output.id; will be removed in v2.0.", string(byName["legacy_id"].Deprecated))
	assert.False(byName["id"].IsDeprecated(), "id should not be deprecated")
}
