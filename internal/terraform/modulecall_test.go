/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModulecallNameWithoutVersion(t *testing.T) {
	assert := assert.New(t)
	modulecall := ModuleCall{
		Name:   "provider",
		Source: "bar",
	}
	assert.Equal("bar", modulecall.FullName())
}

func TestModulecallNameWithVersion(t *testing.T) {
	assert := assert.New(t)
	modulecall := ModuleCall{
		Name:    "provider",
		Source:  "bar",
		Version: "1.2.3",
	}
	assert.Equal("bar,1.2.3", modulecall.FullName())
}
