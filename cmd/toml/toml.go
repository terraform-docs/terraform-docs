/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package toml

import (
	"github.com/spf13/cobra"

	"github.com/rquadling/terraform-docs/internal/cli"
	"github.com/rquadling/terraform-docs/print"
)

// NewCommand returns a new cobra.Command for 'toml' formatter
func NewCommand(runtime *cli.Runtime, config *print.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "toml [PATH]",
		Short:       "Generate TOML of inputs and outputs",
		Annotations: cli.Annotations("toml"),
		PreRunE:     runtime.PreRunEFunc,
		RunE:        runtime.RunEFunc,
	}
	return cmd
}
