package cmd

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "version",
	Short: "Print the version number of terraform-docs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(fmt.Sprintf("terraform-docs version %s\n", version.Version()))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
