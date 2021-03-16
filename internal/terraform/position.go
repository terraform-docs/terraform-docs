/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

// Position represents position of Terraform item (input, output, provider, etc) in a file.
type Position struct {
	Filename string `json:"-" toml:"-" xml:"-" yaml:"-"`
	Line     int    `json:"-" toml:"-" xml:"-" yaml:"-"`
}
