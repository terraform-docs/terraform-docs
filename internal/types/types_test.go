/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expected struct {
	typeName   string
	valueKind  string
	hasDefault bool
}

type testprimitive struct {
	name     string
	values   List
	types    string
	expected expected
}

type testlist struct {
	name     string
	values   []List
	types    string
	expected expected
}

type testmap struct {
	name     string
	values   []Map
	types    string
	expected expected
}

func testPrimitive(t *testing.T, tests []testprimitive) {
	for i := range tests {
		tt := tests[i]
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv)
				actualType := TypeOf(tt.types, tv)

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}
