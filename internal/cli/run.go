package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/terraform-docs/terraform-docs/internal/format"
	"github.com/terraform-docs/terraform-docs/internal/module"
)

// list of flagset items which are explicitly changed from CLI
var changedfs = make(map[string]bool)

// PreRunEFunc returns actual 'cobra.Command#PreRunE' function
// for 'formatter' commands. This functions reads and normalizes
// flags and arguments passed through CLI execution.
func PreRunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		formatter := cmd.Annotations["command"]

		// root command must have an argument, otherwise we're going to show help
		if formatter == "root" && len(args) == 0 {
			cmd.Help() //nolint:errcheck
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
				cmd.Help() //nolint:errcheck
				os.Exit(0)
			}
		} else {
			// config file is found, we're now going to parse it
			if err := cfgreader.parse(); err != nil {
				return err
			}
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

		return nil
	}
}

// RunEFunc returns actual 'cobra.Command#RunE' function for
// 'formatter' commands. This functions extract print.Settings
// and module.Options from generated and normalized Config and
// initializes required print.Format instance and executes it.
func RunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		settings, options := config.extract()

		printer, err := format.Factory(config.Formatter, settings)
		if err != nil {
			return err
		}

		options.Path = args[0]

		tfmodule, err := module.LoadWithOptions(options)
		if err != nil {
			return err

		}

		output, err := printer.Print(tfmodule, settings)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	}
}
