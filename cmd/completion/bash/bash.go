/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package bash

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for 'completion bash' command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "bash",
		Short: "Generate shell completion for bash",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenBashCompletion(os.Stdout)
		},
	}
	return cmd
}
