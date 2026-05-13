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

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/stretchr/testify/assert"
)

func parseHCLForTest(t *testing.T, src string) map[string]*hcl.File {
	t.Helper()
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL([]byte(src), "test.tf")
	if diags.HasErrors() {
		t.Fatalf("parse error: %s", diags)
	}
	return map[string]*hcl.File{"test.tf": file}
}

func TestExtractProviderBlocks(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		expected []rawProviderBlock
	}{
		{
			name:     "no provider blocks",
			src:      `resource "aws_s3_bucket" "b" {}`,
			expected: nil,
		},
		{
			name: "single provider block without alias",
			src: `provider "aws" {
  region = "us-east-1"
}`,
			expected: []rawProviderBlock{
				{LocalName: "aws", Alias: "", Filename: "test.tf", Line: 1},
			},
		},
		{
			name: "provider block with alias",
			src: `provider "aws" {
  alias  = "ident"
  region = "us-east-1"
}`,
			expected: []rawProviderBlock{
				{LocalName: "aws", Alias: "ident", Filename: "test.tf", Line: 1},
			},
		},
		{
			name: "two provider blocks for same name with and without alias",
			src: `provider "aws" {}
provider "aws" {
  alias = "west"
}`,
			expected: []rawProviderBlock{
				{LocalName: "aws", Alias: "", Filename: "test.tf", Line: 1},
				{LocalName: "aws", Alias: "west", Filename: "test.tf", Line: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			files := parseHCLForTest(t, tt.src)
			actual := extractProviderBlocks(files)
			assert.ElementsMatch(tt.expected, actual)
		})
	}
}

func TestExtractResourcesProviderForms(t *testing.T) {
	tests := []struct {
		name             string
		src              string
		expectedName     string
		expectedAlias    string
	}{
		{
			name:          "no provider attribute defaults to type prefix",
			src:           `resource "aws_s3_bucket" "b" {}`,
			expectedName:  "aws",
			expectedAlias: "",
		},
		{
			name: "traversal form provider = aws",
			src: `resource "aws_s3_bucket" "b" {
  provider = aws
}`,
			expectedName:  "aws",
			expectedAlias: "",
		},
		{
			name: "traversal form with alias provider = aws.ident",
			src: `resource "aws_s3_bucket" "b" {
  provider = aws.ident
}`,
			expectedName:  "aws",
			expectedAlias: "ident",
		},
		{
			name: "legacy string-literal form provider = \"aws\"",
			src: `data "aws_caller_identity" "c" {
  provider = "aws"
}`,
			expectedName:  "aws",
			expectedAlias: "",
		},
		{
			name: "legacy string-literal form with alias provider = \"aws.ident\"",
			src: `data "aws_caller_identity" "c" {
  provider = "aws.ident"
}`,
			expectedName:  "aws",
			expectedAlias: "ident",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			files := parseHCLForTest(t, tt.src)
			resources := extractResources(files)
			assert.Len(resources, 1)
			assert.Equal(tt.expectedName, resources[0].ProviderName)
			assert.Equal(tt.expectedAlias, resources[0].ProviderAlias)
		})
	}
}
