/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package xml

import (
	"github.com/spf13/cobra"

	"github.com/rquadling/terraform-docs/internal/cli"
	"github.com/rquadling/terraform-docs/print"
)

// NewCommand returns a new cobra.Command for 'xml' formatter
func NewCommand(runtime *cli.Runtime, config *print.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "xml [PATH]",
		Short:       "Generate XML of inputs and outputs",
		Annotations: cli.Annotations("xml"),
		PreRunE:     runtime.PreRunEFunc,
		RunE:        runtime.RunEFunc,
	}
	return cmd
}
