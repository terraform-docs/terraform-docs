package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/version"
)

// NewCommand returns a new cobra.Command for 'version' command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "Print the version number of terraform-docs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("terraform-docs version %s\n", Full())
		},
	}
	return cmd
}

// Full returns the full version of the binary
func Full() string {
	return version.Full()
}
