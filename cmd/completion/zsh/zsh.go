package zsh

import (
	"os"

	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for 'completion zsh' command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "zsh",
		Short: "Generate shel completion for zsh",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenZshCompletion(os.Stdout)
		},
	}
	return cmd
}
