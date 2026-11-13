/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package plugin

import (
	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// repairModule restores what gob loses while the module travels to a plugin.
//
// gob decodes a zero length slice or map back as a nil one, so a variable
// declared `default = []`, or `default = { team = [] }`, reaches the plugin
// holding a nil collection. That marshals to `null` rather than `[]`, and the
// plugin then documents the wrong default value.
func repairModule(module *terraform.Module) {
	if module == nil {
		return
	}

	for _, input := range module.Inputs {
		if input != nil {
			input.Default = repairDefault(input.Default)
		}
	}

	for _, input := range module.RequiredInputs {
		if input != nil {
			input.Default = repairDefault(input.Default)
		}
	}

	for _, input := range module.OptionalInputs {
		if input != nil {
			input.Default = repairDefault(input.Default)
		}
	}

	for _, output := range module.Outputs {
		if output != nil {
			output.Value = repairDefault(output.Value)
		}
	}
}

// repairDefault repairs a decoded value.
func repairDefault(value types.Value) types.Value {
	switch v := value.(type) {
	case *types.List:
		if v != nil {
			*v = types.List(repairSlice(*v))
		}
		return v
	case types.List:
		return types.List(repairSlice(v))
	case *types.Map:
		if v != nil {
			*v = types.Map(repairMap(*v))
		}
		return v
	case types.Map:
		return types.Map(repairMap(v))
	default:
		return value
	}
}

// repairAny normalizes the bare collections nested inside a decoded value.
func repairAny(value interface{}) interface{} {
	switch v := value.(type) {
	case []interface{}:
		return repairSlice(v)
	case map[string]interface{}:
		return repairMap(v)
	}
	return value
}

// repairSlice returns a non-nil slice, with every element repaired.
func repairSlice(list []interface{}) []interface{} {
	repaired := make([]interface{}, len(list))
	for i, item := range list {
		repaired[i] = repairAny(item)
	}
	return repaired
}

// repairMap returns a non-nil map, with every value repaired.
func repairMap(m map[string]interface{}) map[string]interface{} {
	repaired := make(map[string]interface{}, len(m))
	for key, item := range m {
		repaired[key] = repairAny(item)
	}
	return repaired
}
