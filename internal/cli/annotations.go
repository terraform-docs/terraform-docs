/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

// Annotations returns set of annotations for cobra.Commands,
// specifically the command 'name' and command 'kind'
func Annotations(cmd string) map[string]string {
	return map[string]string{
		"command": cmd,
		"kind":    "formatter",
	}
}
