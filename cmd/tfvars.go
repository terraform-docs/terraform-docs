package cmd

import (
	"github.com/spf13/cobra"
)

var tfvarsCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "tfvars [PATH]",
	Short:       "Generate terraform.tfvars of inputs",
	Annotations: formatAnnotations("tfvars"),
}

var tfvarsHCLCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "hcl [PATH]",
	Short:       "Generate HCL format of terraform.tfvars of inputs",
	Annotations: formatAnnotations("tfvars hcl"),
	RunE:        formatRunE,
}

var tfvarsJSONCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "json [PATH]",
	Short:       "Generate JSON format of terraform.tfvars of inputs",
	Annotations: formatAnnotations("tfvars json"),
	RunE:        formatRunE,
}

func init() {
	tfvarsCmd.AddCommand(tfvarsHCLCmd)
	tfvarsCmd.AddCommand(tfvarsJSONCmd)

	rootCmd.AddCommand(tfvarsCmd)
}
