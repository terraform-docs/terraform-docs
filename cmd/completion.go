package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion SHELL",
	Short: "Generate autocomplete for terraform-docs",
	Long:  "Generate autocomplete for terraform-docs",
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate autocomplete for bash",
	Long:  "Generate autocomplete for bash",
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenBashCompletion(os.Stdout)
	},
}
var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate autocomplete for zsh",
	Long:  "Generate autocomplete for zsh",
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)

	rootCmd.AddCommand(completionCmd)
}
