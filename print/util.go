/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

func contains(list []string, name string) bool {
	for _, v := range list {
		if v == name {
			return true
		}
	}
	return false
}

// nolint
func index(list []string, name string) int {
	for i, v := range list {
		if v == name {
			return i
		}
	}
	return -1
}

// nolint
func remove(list []string, name string) []string {
	index := index(list, name)
	if index < 0 {
		return list
	}
	list[index] = list[len(list)-1]
	list[len(list)-1] = ""
	return list[:len(list)-1]
}
