/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package markdown

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/markdown/document"
	"github.com/terraform-docs/terraform-docs/cmd/markdown/table"
	"github.com/terraform-docs/terraform-docs/internal/cli"
	"github.com/terraform-docs/terraform-docs/print"
)

// NewCommand returns a new cobra.Command for 'markdown' formatter
func NewCommand(runtime *cli.Runtime, config *print.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "markdown [PATH]",
		Aliases:     []string{"md"},
		Short:       "Generate Markdown of inputs and outputs",
		Annotations: cli.Annotations("markdown"),
		PreRunE:     runtime.PreRunEFunc,
		RunE:        runtime.RunEFunc,
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Anchor, "anchor", true, "create anchor links")
	cmd.PersistentFlags().BoolVar(&config.Settings.Default, "default", true, "show Default column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Escape, "escape", true, "escape special characters")
	cmd.PersistentFlags().BoolVar(&config.Settings.HTML, "html", true, "use HTML tags in genereted output")
	cmd.PersistentFlags().BoolVar(&config.Settings.HideEmpty, "hide-empty", false, "hide empty sections (default false)")
	cmd.PersistentFlags().IntVar(&config.Settings.Indent, "indent", 2, "indention level of Markdown sections [1, 2, 3, 4, 5]")
	cmd.PersistentFlags().BoolVar(&config.Settings.Required, "required", true, "show Required column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Sensitive, "sensitive", true, "show Sensitive column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Type, "type", true, "show Type column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Validation, "validation", true, "show Validation column or section")

	// subcommands
	cmd.AddCommand(document.NewCommand(runtime, config))
	cmd.AddCommand(table.NewCommand(runtime, config))

	return cmd
}
