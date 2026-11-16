/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/rquadling/terraform-docs/internal/plugin"
	"github.com/rquadling/terraform-docs/internal/version"
)

// NewCommand returns a new cobra.Command for 'version' command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "Print the version number of terraform-docs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("terraform-docs version %s\n", version.Full())
			plugins, err := plugin.Discover()
			if err != nil {
				return
			}
			for _, f := range plugins.All() {
				name, err := f.Name()
				if err != nil {
					name = "unknown"
				}
				version, err := f.Version()
				if err != nil {
					version = "unknown"
				}
				fmt.Printf("- plugin %s %s\n", name, version)
			}
		},
	}
	return cmd
}
