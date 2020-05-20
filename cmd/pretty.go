package cmd

import (
	"github.com/spf13/cobra"
)

var prettyCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "pretty [PATH]",
	Short:       "Generate colorized pretty of inputs and outputs",
	Annotations: formatAnnotations("pretty"),
	RunE:        formatRunE,
}

func init() {
	prettyCmd.PersistentFlags().BoolVar(new(bool), "no-color", false, "do not colorize printed result")
	prettyCmd.PersistentFlags().MarkDeprecated("no-color", "use '--color=false' instead") //nolint:errcheck

	prettyCmd.PersistentFlags().BoolVar(&settings.ShowColor, "color", true, "colorize printed result")

	rootCmd.AddCommand(prettyCmd)
}
