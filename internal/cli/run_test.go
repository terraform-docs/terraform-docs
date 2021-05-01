/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionConstraint(t *testing.T) {
	type tuple struct {
		constraint string
		version    string
	}
	tests := map[string]struct {
		versions []tuple
		wantErr  bool
	}{
		"NoRange": {
			versions: []tuple{
				{"", "1.2.3"},
			},
			wantErr: false,
		},
		"ValidConstraint": {
			versions: []tuple{
				{">= 1.0, < 1.2", "1.1.5"},
				{"= 1.0", "1.0.0"},
				{"1.0", "1.0.0"},
				{">= 1.0", "1.2.3"},
				{"~> 1.0", "1.1"},
				{"~> 1.0", "1.2.3"},
				{"~> 1.0.0", "1.0.7"},
				{"~> 1.0.7", "1.0.7"},
				{"~> 1.0.7", "1.0.8"},
				{"~> 2.1.0-a", "2.1.0-beta"},
				{">= 2.1.0-a", "2.1.0-beta"},
				{">= 2.1.0-a", "2.1.1"},
				{">= 2.1.0-a", "2.1.0"},
				{"<= 2.1.0-a", "2.0.0"},
			},
			wantErr: false,
		},
		"MalformedCurrent": {
			versions: []tuple{
				{"> 1.0", "1.2.x"},
			},
			wantErr: true,
		},
		"InvalidConstraint": {
			versions: []tuple{
				{"< 1.0, < 1.2", "1.1.5"},
				{"> 1.1, <= 1.2", "1.2.3"},
				{"> 1.2, <= 1.1", "1.2.3"},
				{"= 1.0", "1.1.5"},
				{"~> 1.0", "2.0"},
				{"~> 1.0.0", "1.2.3"},
				{"~> 1.0.0", "1.1.0"},
				{"~> 1.0.7", "1.0.4"},
				{"~> 2.0", "2.1.0-beta"},
				{"~> 2.1.0-a", "2.2.0"},
				{"~> 2.1.0-a", "2.1.0"},
				{"~> 2.1.0-a", "2.2.0-alpha"},
				{"> 2.0", "2.1.0-beta"},
				{">= 2.1.0-a", "2.1.1-beta"},
				{">= 2.0.0", "2.1.0-beta"},
				{">= 2.1.0-a", "2.1.1-beta"},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			for _, v := range tt.versions {
				err := checkConstraint(v.constraint, v.version)

				if tt.wantErr {
					assert.NotNil(err)
				} else {
					assert.Nil(err)
				}
			}
		})
	}
}
