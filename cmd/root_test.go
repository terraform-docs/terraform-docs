/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cmd

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-docs/terraform-docs/internal/testutil"
)

func executeCommand(root *cobra.Command, args []string) (output string, err error) {
	_, output, err = executeCommandC(root, args)
	return output, err
}

func executeCommandC(root *cobra.Command, args []string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func TestMarkdownTableCommandExecute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name   string
		format string
		params []string
	}{
		{
			name:   "table",
			format: "markdown",
			params: []string{"md", "table", "--show-all", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-HideAll",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyHeader",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "header", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyInputs",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "inputs", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyModules",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "modules", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyOutputs",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "outputs", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyProviders",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "providers", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyRequirements",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "requirements", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-OnlyResources",
			format: "markdown",
			params: []string{"md", "table", "--hide-all", "--show", "resources", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoHeader",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "header", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoInputs",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "inputs", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoModules",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "modules", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoOutputs",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "outputs", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoProviders",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "providers", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoRequirements",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "requirements", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoResources",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--hide", "resources", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-NoSort",
			format: "markdown",
			params: []string{"md", "table", "--show-all", "--sort=false", filepath.Join("testdata", "example")},
		},
		{
			name:   "table-Minimal",
			format: "markdown",
			params: []string{"md", "table", "--show-all", filepath.Join("testdata", "example-minimal")},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.format, tt.name), func(t *testing.T) {
			expected, err := testutil.GetExpected(tt.format, tt.name)
			assert.Nil(err)

			cmd := NewCommand()
			actual, err := executeCommand(cmd, tt.params)
			assert.Nil(err)

			assert.Equal(expected, actual)
		})
	}
}
