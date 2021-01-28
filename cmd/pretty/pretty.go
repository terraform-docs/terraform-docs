/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package pretty

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for pretty formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "pretty [PATH]",
		Short:       "Generate colorized pretty of inputs and outputs",
		Annotations: cli.Annotations("pretty"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Color, "color", true, "colorize printed result")

	// deprecation
	cmd.PersistentFlags().BoolVar(&config.Settings.Deprecated.NoColor, "no-color", false, "do not colorize printed result")
	cmd.PersistentFlags().MarkDeprecated("no-color", "use '--color=false' instead") //nolint:errcheck

	return cmd
}
