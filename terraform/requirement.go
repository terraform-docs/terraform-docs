/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"github.com/rquadling/terraform-docs/internal/types"
)

// Requirement represents a requirement for Terraform module.
type Requirement struct {
	Name    string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Version types.String `json:"version" toml:"version" xml:"version" yaml:"version"`
}
