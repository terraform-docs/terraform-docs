/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package plugin

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// TestRepairModule sends a module across gob the way the plugin RPC does, and
// checks that the empty collection defaults still render as `[]` and `{}`.
func TestRepairModule(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{name: "empty list", value: []interface{}{}, expected: `[]`},
		{name: "empty map", value: map[string]interface{}{}, expected: `{}`},
		{name: "nested empty list", value: map[string]interface{}{"team": []interface{}{}}, expected: `{"team":[]}`},
		{name: "list of empty maps", value: []interface{}{map[string]interface{}{}}, expected: `[{}]`},
		{name: "populated list", value: []interface{}{"a"}, expected: `["a"]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			module := &terraform.Module{
				Inputs: []*terraform.Input{{Name: "example", Default: types.ValueOf(tt.value)}},
			}

			var buffer bytes.Buffer
			assert.Nil(gob.NewEncoder(&buffer).Encode(module))

			decoded := new(terraform.Module)
			assert.Nil(gob.NewDecoder(&buffer).Decode(decoded))

			repairModule(decoded)

			encoded, err := json.Marshal(decoded.Inputs[0].Default)
			assert.Nil(err)
			assert.Equal(tt.expected, string(encoded))
		})
	}
}

// TestRepairModuleOutputs checks that output values round-trip through gob properly.
func TestRepairModuleOutputs(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{name: "empty list", value: []interface{}{}, expected: `[]`},
		{name: "empty map", value: map[string]interface{}{}, expected: `{}`},
		{name: "nested empty list", value: map[string]interface{}{"team": []interface{}{}}, expected: `{"team":[]}`},
		{name: "list of empty maps", value: []interface{}{map[string]interface{}{}}, expected: `[{}]`},
		{name: "populated list", value: []interface{}{"a"}, expected: `["a"]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			module := &terraform.Module{
				Outputs: []*terraform.Output{{Name: "example", Value: types.ValueOf(tt.value)}},
			}

			var buffer bytes.Buffer
			assert.Nil(gob.NewEncoder(&buffer).Encode(module))

			decoded := new(terraform.Module)
			assert.Nil(gob.NewDecoder(&buffer).Decode(decoded))

			repairModule(decoded)

			encoded, err := json.Marshal(decoded.Outputs[0].Value)
			assert.Nil(err)
			assert.Equal(tt.expected, string(encoded))
		})
	}
}

// TestRepairModuleValueReceivers checks that non-pointer types.List and types.Map are also repaired.
func TestRepairModuleValueReceivers(t *testing.T) {
	assert := assert.New(t)

	module := &terraform.Module{
		Inputs: []*terraform.Input{
			{Name: "val_list", Default: types.List(nil)},
			{Name: "val_map", Default: types.Map(nil)},
		},
	}

	repairModule(module)

	encodedList, err := json.Marshal(module.Inputs[0].Default)
	assert.Nil(err)
	assert.Equal(`[]`, string(encodedList))

	encodedMap, err := json.Marshal(module.Inputs[1].Default)
	assert.Nil(err)
	assert.Equal(`{}`, string(encodedMap))
}

// TestRepairModuleTolerance checks the guards, since a plugin must not panic
// on a module that arrives without inputs or outputs.
func TestRepairModuleTolerance(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() { repairModule(nil) })
	assert.NotPanics(func() { repairModule(&terraform.Module{}) })
	assert.NotPanics(func() {
		repairModule(&terraform.Module{
			Inputs:         []*terraform.Input{nil, {Name: "no-default"}},
			RequiredInputs: []*terraform.Input{nil, {Name: "no-default-req"}},
			OptionalInputs: []*terraform.Input{nil, {Name: "no-default-opt"}},
			Outputs:        []*terraform.Output{nil, {Name: "no-value"}},
		})
	})
}
