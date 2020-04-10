package cmd

import (
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "json [PATH]",
	Short:       "Generate JSON of inputs and outputs",
	Annotations: formatAnnotations("json"),
	RunE:        formatRunE,
}

func init() {
	jsonCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")

	rootCmd.AddCommand(jsonCmd)
}
