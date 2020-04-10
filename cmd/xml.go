package cmd

import (
	"github.com/spf13/cobra"
)

var xmlCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "xml [PATH]",
	Short:       "Generate XML of inputs and outputs",
	Annotations: formatAnnotations("xml"),
	RunE:        formatRunE,
}

func init() {
	rootCmd.AddCommand(xmlCmd)
}
