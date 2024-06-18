/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/asciidoc"
	"github.com/terraform-docs/terraform-docs/cmd/completion"
	"github.com/terraform-docs/terraform-docs/cmd/json"
	"github.com/terraform-docs/terraform-docs/cmd/markdown"
	"github.com/terraform-docs/terraform-docs/cmd/pretty"
	"github.com/terraform-docs/terraform-docs/cmd/tfvars"
	"github.com/terraform-docs/terraform-docs/cmd/toml"
	versioncmd "github.com/terraform-docs/terraform-docs/cmd/version"
	"github.com/terraform-docs/terraform-docs/cmd/xml"
	"github.com/terraform-docs/terraform-docs/cmd/yaml"
	"github.com/terraform-docs/terraform-docs/internal/cli"
	"github.com/terraform-docs/terraform-docs/internal/version"
	"github.com/terraform-docs/terraform-docs/print"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := NewCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}
	return nil
}

// NewCommand returns a new cobra.Command for 'root' command
func NewCommand() *cobra.Command {
	config := print.DefaultConfig()
	runtime := cli.NewRuntime(config)
	cmd := &cobra.Command{
		Args:          cobra.MaximumNArgs(1),
		Use:           "terraform-docs [PATH]",
		Short:         "A utility to generate documentation from Terraform modules in various output formats",
		Long:          "A utility to generate documentation from Terraform modules in various output formats",
		Version:       version.Full(),
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations:   cli.Annotations("root"),
		PreRunE:       runtime.PreRunEFunc,
		RunE:          runtime.RunEFunc,
	}

	// flags
	cmd.PersistentFlags().StringVarP(&config.File, "config", "c", ".terraform-docs.yml", "config file name")
	cmd.PersistentFlags().BoolVar(&config.Recursive.Enabled, "recursive", false, "update submodules recursively (default false)")
	cmd.PersistentFlags().StringVar(&config.Recursive.Path, "recursive-path", "modules", "submodules path to recursively update")
	cmd.PersistentFlags().BoolVar(&config.Recursive.IncludeMain, "recursive-include-main", true, "include the main module")

	cmd.PersistentFlags().StringSliceVar(&config.Sections.Show, "show", []string{}, "show section ["+print.AllSections+"]")
	cmd.PersistentFlags().StringSliceVar(&config.Sections.Hide, "hide", []string{}, "hide section ["+print.AllSections+"]")

	cmd.PersistentFlags().StringVar(&config.Output.File, "output-file", "", "file path to insert output into (default \"\")")
	cmd.PersistentFlags().StringVar(&config.Output.Mode, "output-mode", "inject", "output to file method ["+print.OutputModes+"]")
	cmd.PersistentFlags().StringVar(&config.Output.Template, "output-template", print.OutputTemplate, "output template")
	cmd.PersistentFlags().BoolVar(&config.Output.Check, "output-check", false, "check if content of output file is up to date (default false)")

	cmd.PersistentFlags().BoolVar(&config.Sort.Enabled, "sort", true, "sort items")
	cmd.PersistentFlags().StringVar(&config.Sort.By, "sort-by", "name", "sort items by criteria ["+print.SortTypes+"]")

	cmd.PersistentFlags().StringVar(&config.HeaderFrom, "header-from", "main.tf", "relative path of a file to read header from")
	cmd.PersistentFlags().StringVar(&config.FooterFrom, "footer-from", "", "relative path of a file to read footer from (default \"\")")

	cmd.PersistentFlags().BoolVar(&config.Settings.LockFile, "lockfile", true, "read .terraform.lock.hcl if exist")

	cmd.PersistentFlags().BoolVar(&config.OutputValues.Enabled, "output-values", false, "inject output values into outputs (default false)")
	cmd.PersistentFlags().StringVar(&config.OutputValues.From, "output-values-from", "", "inject output values from file into outputs (default \"\")")

	cmd.PersistentFlags().BoolVar(&config.Settings.ReadComments, "read-comments", true, "use comments as description when description is empty")

	// formatter subcommands
	cmd.AddCommand(asciidoc.NewCommand(runtime, config))
	cmd.AddCommand(json.NewCommand(runtime, config))
	cmd.AddCommand(markdown.NewCommand(runtime, config))
	cmd.AddCommand(pretty.NewCommand(runtime, config))
	cmd.AddCommand(tfvars.NewCommand(runtime, config))
	cmd.AddCommand(toml.NewCommand(runtime, config))
	cmd.AddCommand(xml.NewCommand(runtime, config))
	cmd.AddCommand(yaml.NewCommand(runtime, config))

	// other subcommands
	cmd.AddCommand(completion.NewCommand())
	cmd.AddCommand(versioncmd.NewCommand())

	return cmd
}
