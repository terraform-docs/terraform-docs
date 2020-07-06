package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/terraform-docs/terraform-docs/internal/format"
	"github.com/terraform-docs/terraform-docs/internal/module"
)

// PreRunEFunc returns actual 'cobra.Command#PreRunE' function
// for 'formatter' commands. This functions reads and normalizes
// flags and arguments passed through CLI execution.
func PreRunEFunc(config *Config) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			changedfs[f.Name] = f.Changed
		})

		config.normalize(cmd.CommandPath())

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
