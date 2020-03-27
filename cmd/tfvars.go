package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var tfvarsCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "tfvars [PATH]",
	Short: "Generate terraform.tfvars of inputs",
}

var tfvarsHCLCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "hcl [PATH]",
	Short: "Generate HCL format of terraform.tfvars of inputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewTfvarsHCL(settings))
	},
}

var tfvarsJSONCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "json [PATH]",
	Short: "Generate JSON format of terraform.tfvars of inputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewTfvarsJSON(settings))
	},
}

func init() {
	tfvarsCmd.AddCommand(tfvarsHCLCmd)
	tfvarsCmd.AddCommand(tfvarsJSONCmd)

	rootCmd.AddCommand(tfvarsCmd)
}
