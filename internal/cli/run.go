/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	pluginsdk "github.com/terraform-docs/plugin-sdk/plugin"
	"github.com/terraform-docs/terraform-docs/internal/format"
	"github.com/terraform-docs/terraform-docs/internal/plugin"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// list of flagset items which are explicitly changed from CLI
var changedfs = make(map[string]bool)

// PreRunEFunc returns actual 'cobra.Command#PreRunE' function for 'formatter'
// commands. This functions reads and normalizes flags and arguments passed
// through CLI execution.
func PreRunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		formatter := cmd.Annotations["command"]

		// root command must have an argument, otherwise we're going to show help
		if formatter == "root" && len(args) == 0 {
			cmd.Help() //nolint:errcheck,gosec
			os.Exit(0)
		}

		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			changedfs[f.Name] = f.Changed
		})

		// read config file if provided and/or available
		if config.File == "" {
			return fmt.Errorf("value of '--config' can't be empty")
		}

		file := filepath.Join(args[0], config.File)
		cfgreader := &cfgreader{
			file:   file,
			config: config,
		}

		if found, err := cfgreader.exist(); !found {
			// config is explicitly provided and file not found, this is an error
			if changedfs["config"] {
				return err
			}
			// config is not provided and file not found, only show an error for the root command
			if formatter == "root" {
				cmd.Help() //nolint:errcheck,gosec
				os.Exit(0)
			}
		} else if err := cfgreader.parse(); err != nil {
			// config file is found, we're now going to parse it
			return err
		}

		// explicitly setting formatter to Config for non-root commands this
		// will effectively override formattter properties from config file
		// if 1) config file exists and 2) formatter is set and 3) explicitly
		// a subcommand was executed in the terminal
		if formatter != "root" {
			config.Formatter = formatter
		}

		config.process()

		if err := config.validate(); err != nil {
			return err
		}

		// set the base module directory
		config.BaseDir = args[0]

		return nil
	}
}

// RunEFunc returns actual 'cobra.Command#RunE' function for 'formatter' commands.
// This functions extract print.Settings and terraform.Options from generated and
// normalized Config and initializes required print.Format instance and executes it.
func RunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		settings, options := config.extract()
		options.Path = config.BaseDir

		module, err := terraform.LoadWithOptions(options)
		if err != nil {
			return err
		}

		printer, err := format.Factory(config.Formatter, settings)
		if err != nil {
			plugins, perr := plugin.Discover()
			if perr != nil {
				return fmt.Errorf("formatter '%s' not found", config.Formatter)
			}

			client, found := plugins.Get(config.Formatter)
			if !found {
				return fmt.Errorf("formatter '%s' not found", config.Formatter)
			}

			content, cerr := client.Execute(pluginsdk.ExecuteArgs{
				Module:   module.Convert(),
				Settings: settings.Convert(),
			})
			if cerr != nil {
				return cerr
			}
			return writeContent(config, content)
		}

		content, err := printer.Print(module, settings)
		if err != nil {
			return err
		}
		return writeContent(config, content)
	}
}

// writeContent to a Writer. This can either be os.Stdout or specific
// file (e.g. README.md) if '--output-file' is provided.
func writeContent(config *Config, content string) error {
	var w io.Writer

	// writing to a file (either inject or replace)
	if config.Output.File != "" {
		w = &fileWriter{
			file: config.Output.File,
			dir:  config.BaseDir,

			mode: config.Output.Mode,

			template: config.Output.Template,
			begin:    config.Output.BeginComment,
			end:      config.Output.EndComment,
		}
	} else {
		// writing to stdout
		w = &stdoutWriter{}
	}

	_, err := io.WriteString(w, content)

	return err
}
