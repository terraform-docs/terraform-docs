/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"errors"
	"fmt"
	"io"
	"os"

	goversion "github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	pluginsdk "github.com/terraform-docs/plugin-sdk/plugin"
	"github.com/terraform-docs/terraform-docs/internal/format"
	"github.com/terraform-docs/terraform-docs/internal/plugin"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
	"github.com/terraform-docs/terraform-docs/internal/version"
)

// PreRunEFunc returns actual 'cobra.Command#PreRunE' function for 'formatter'
// commands. This functions reads and normalizes flags and arguments passed
// through CLI execution.
func PreRunEFunc(config *Config) func(*cobra.Command, []string) error { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	return func(cmd *cobra.Command, args []string) error {
		config.isFlagChanged = cmd.Flags().Changed

		formatter := cmd.Annotations["command"]

		// root command must have an argument, otherwise we're going to show help
		if formatter == "root" && len(args) == 0 {
			cmd.Help() //nolint:errcheck,gosec
			os.Exit(0)
		}

		// this can only happen in one way: terraform-docs -c "" /path/to/module
		if config.File == "" {
			return errors.New("value of '--config' can't be empty")
		}

		v := viper.New()

		if config.isFlagChanged("config") {
			v.SetConfigFile(config.File)
		} else {
			v.SetConfigName(".terraform-docs")
			v.SetConfigType("yml")
		}

		v.AddConfigPath(args[0])           // first look at module root
		v.AddConfigPath(".")               // then current directory
		v.AddConfigPath("$HOME/.tfdocs.d") // and finally $HOME/.tfdocs.d/

		if err := v.ReadInConfig(); err != nil {
			var perr *os.PathError
			if errors.As(err, &perr) {
				return fmt.Errorf("config file %s not found", config.File)
			}

			var cerr viper.ConfigFileNotFoundError
			if !errors.As(err, &cerr) {
				return err
			}

			// config is not provided, only show error for root command
			if formatter == "root" {
				cmd.Help() //nolint:errcheck,gosec
				os.Exit(0)
			}
		}

		// bind flags to viper
		bindFlags(cmd, v)

		if err := v.Unmarshal(config); err != nil {
			return fmt.Errorf("unable to decode config, %w", err)
		}

		if err := checkConstraint(config.Version, version.Core()); err != nil {
			return err
		}

		// explicitly setting formatter to Config for non-root commands this
		// will effectively override formattter properties from config file
		// if 1) config file exists and 2) formatter is set and 3) explicitly
		// a subcommand was executed in the terminal
		if formatter != "root" {
			config.Formatter = formatter
		}

		// set the module root directory
		config.moduleRoot = args[0]

		// process and validate configuration
		return config.process()
	}
}

// RunEFunc returns actual 'cobra.Command#RunE' function for 'formatter' commands.
// This functions extract print.Settings and terraform.Options from generated and
// normalized Config and initializes required print.Format instance and executes it.
func RunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		settings, options := config.extract()
		options.Path = config.moduleRoot

		module, err := terraform.LoadWithOptions(options)
		if err != nil {
			return err
		}

		formatter, err := format.Factory(config.Formatter, settings)
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

		generator, err := formatter.Generate(module)
		if err != nil {
			return err
		}
		generator.Path(config.moduleRoot)

		content, err := generator.ExecuteTemplate(config.Content)
		if err != nil {
			return err
		}

		return writeContent(config, content)
	}
}

// bindFlags binds current command's changed flags to viper
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	sectionsCleared := false
	fs := cmd.Flags()
	fs.VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			return
		}

		switch f.Name {
		case "show", "hide":
			// If '--show' or '--hide' CLI flag is used, explicitly override and remove
			// all items from 'show' and 'hide' set in '.terraform-doc.yml'.
			if !sectionsCleared {
				v.Set("sections.show", []string{})
				v.Set("sections.hide", []string{})
				sectionsCleared = true
			}

			items, err := fs.GetStringSlice(f.Name)
			if err != nil {
				return
			}
			v.Set(flagMappings[f.Name], items)
		case "sort-by-required", "sort-by-type":
			v.Set("sort.by", flagMappings[f.Name])
		default:
			if _, ok := flagMappings[f.Name]; !ok {
				return
			}
			v.Set(flagMappings[f.Name], f.Value)
		}
	})
}

// checkConstraint validates if current version of terraform-docs being executed
// is valid against 'version' string provided in config file, and fail if the
// constraints is violated.
func checkConstraint(versionRange string, currentVersion string) error {
	if versionRange == "" {
		return nil
	}

	semver, err := goversion.NewSemver(currentVersion)
	if err != nil {
		return err
	}

	constraint, err := goversion.NewConstraint(versionRange)
	if err != nil || !constraint.Check(semver) {
		return fmt.Errorf("current version: %s, constraints: '%s'", semver, constraint)
	}

	return nil
}

// writeContent to a Writer. This can either be os.Stdout or specific
// file (e.g. README.md) if '--output-file' is provided.
func writeContent(config *Config, content string) error {
	var w io.Writer

	// writing to a file (either inject or replace)
	if config.Output.File != "" {
		w = &fileWriter{
			file: config.Output.File,
			dir:  config.moduleRoot,

			mode: config.Output.Mode,

			check: config.Output.Check,

			template: config.Output.Template,
			begin:    config.Output.beginComment,
			end:      config.Output.endComment,
		}
	} else {
		// writing to stdout
		w = &stdoutWriter{}
	}

	_, err := io.WriteString(w, content)

	return err
}
