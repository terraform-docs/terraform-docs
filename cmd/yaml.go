package cmd

import (
	"github.com/segmentio/terraform-docs/internal/pkg/print/yaml"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "yaml [PATH]",
	Short: "Generate a YAML of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(module *tfconf.Module) (string, error) {
			return yaml.Print(module, settings)
		})
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
