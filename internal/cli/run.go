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
	"path/filepath"

	goversion "github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/rquadling/terraform-docs/format"
	"github.com/rquadling/terraform-docs/internal/plugin"
	"github.com/rquadling/terraform-docs/internal/version"
	pluginsdk "github.com/rquadling/terraform-docs/plugin"
	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// Runtime represents the execution runtime for CLI.
type Runtime struct {
	rootDir string

	formatter string
	config    *print.Config

	cmd           *cobra.Command
	isFlagChanged func(string) bool
}

// NewRuntime returns new instance of Runtime. If `config` is not provided
// default config will be used.
func NewRuntime(config *print.Config) *Runtime {
	if config == nil {
		config = print.DefaultConfig()
	}
	return &Runtime{config: config}
}

// PreRunEFunc is the 'cobra.Command#PreRunE' function for 'formatter'
// commands. This function reads and normalizes flags and arguments passed
// through CLI execution.
func (r *Runtime) PreRunEFunc(cmd *cobra.Command, args []string) error {
	r.formatter = cmd.Annotations["command"]

	// root command must have an argument, otherwise we're going to show help
	if r.formatter == "root" && len(args) == 0 {
		cmd.Help() //nolint:errcheck,gosec
		os.Exit(0)
	}

	r.isFlagChanged = cmd.Flags().Changed
	r.rootDir = args[0]
	r.cmd = cmd

	// this can only happen in one way: terraform-docs -c "" /path/to/module
	if r.config.File == "" {
		return fmt.Errorf("value of '--config' can't be empty")
	}

	// read config file and override them with corresponding flags
	v := viper.New()

	if err := r.readConfig(v, r.config.File, ""); err != nil {
		return err
	}

	// and override them with corresponding flags
	if err := r.unmarshalConfig(v, r.config); err != nil {
		return err
	}

	return checkConstraint(r.config.Version, version.Core())
}

type module struct {
	rootDir string
	config  *print.Config
}

// RunEFunc is the 'cobra.Command#RunE' function for 'formatter' commands. It attempts
// to discover submodules, on `--recursive` flag, and generates the content for them
// as well as the root module.
func (r *Runtime) RunEFunc(cmd *cobra.Command, args []string) error { //nolint:gocyclo
	modules := []module{}

	if !r.config.Recursive.Enabled || r.config.Recursive.IncludeMain {
		modules = append(modules, module{r.rootDir, r.config})
	}

	// Generating content recursively is only allowed when `config.Output.File`
	// is set. Otherwise it would be impossible to distinguish where output of
	// one module ends and the other begin, if content is output to stdout.
	if r.config.Recursive.Enabled && r.config.Recursive.Path != "" {
		items, err := r.findSubmodules()
		if err != nil {
			return err
		}

		modules = append(modules, items...)
	}

	for _, module := range modules {
		cfg := r.config

		// If submodules contains its own configuration file, use that instead
		if module.config != nil {
			cfg = module.config
		}

		// set the module root directory
		cfg.ModuleRoot = module.rootDir

		// process and validate configuration
		if err := cfg.Validate(); err != nil {
			return err
		}

		if r.config.Recursive.Enabled && cfg.Output.File == "" {
			return fmt.Errorf("value of '--output-file' cannot be empty with '--recursive'")
		}

		if err := generateContent(cfg); err != nil {
			return err
		}
	}

	return nil
}

// readConfig attempts to read config file, either default `.terraform-docs.yml`
// or provided file with `-c, --config` flag. It will then attempt to override
// them with corresponding flags (if set).
func (r *Runtime) readConfig(v *viper.Viper, file string, submoduleDir string) error {
	if r.isFlagChanged("config") {
		v.SetConfigFile(file)
	} else {
		v.SetConfigName(".terraform-docs")
		v.SetConfigType("yml")
	}

	if submoduleDir != "" {
		v.AddConfigPath(submoduleDir)              // first look at submodule root
		v.AddConfigPath(submoduleDir + "/.config") // then .config/ folder at submodule root
	}

	v.AddConfigPath(r.rootDir)              // first look at module root
	v.AddConfigPath(r.rootDir + "/.config") // then .config/ folder at module root
	v.AddConfigPath(".")                    // then current directory
	v.AddConfigPath(".config")              // then .config/ folder at current directory
	v.AddConfigPath("$HOME/.tfdocs.d")      // and finally $HOME/.tfdocs.d/

	if err := v.ReadInConfig(); err != nil {
		var perr *os.PathError
		if errors.As(err, &perr) {
			return fmt.Errorf("config file %s not found", file)
		}

		var cerr viper.ConfigFileNotFoundError
		if !errors.As(err, &cerr) {
			return err
		}

		// config is not provided, only show error for root command
		if r.formatter == "root" {
			r.cmd.Help() //nolint:errcheck,gosec
			os.Exit(0)
		}
	}

	return nil
}

func (r *Runtime) unmarshalConfig(v *viper.Viper, config *print.Config) error {
	r.bindFlags(v)

	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("unable to decode config, %w", err)
	}

	// explicitly setting formatter to Config for non-root commands this
	// will effectively override formattter properties from config file
	// if 1) config file exists and 2) formatter is set and 3) explicitly
	// a subcommand was executed in the terminal
	if r.formatter != "root" {
		config.Formatter = r.formatter
	}

	config.Parse()

	return nil
}

// bindFlags binds current command's changed flags to viper.
func (r *Runtime) bindFlags(v *viper.Viper) {
	sectionsCleared := false
	fs := r.cmd.Flags()
	fs.VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			return
		}

		switch f.Name {
		case "show", "hide":
			// If '--show' or '--hide' CLI flag is used, explicitly override and remove
			// all items from 'show' and 'hide' set in '.terraform-docs.yml'.
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

func (r *Runtime) mergeConfig(v *viper.Viper) (*print.Config, error) {
	copy := *r.config
	merged := &copy

	if v.IsSet("sections.show") || v.IsSet("sections.hide") {
		merged.Sections.Show = []string{}
		merged.Sections.Hide = []string{}
	}

	if err := r.unmarshalConfig(v, merged); err != nil {
		return nil, err
	}

	return merged, nil
}

// findSubmodules generates list of submodules in `rootDir/RecursivePath` if
// `--recursive` flag is set. This keeps track of `.terraform-docs.yml` in any
// of the submodules (if exists) to override the root configuration.
func (r *Runtime) findSubmodules() ([]module, error) {
	dir := filepath.Join(r.rootDir, r.config.Recursive.Path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	info, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	modules := []module{}

	for _, file := range info {
		if !file.IsDir() {
			continue
		}

		var cfg *print.Config

		path := filepath.Join(dir, file.Name())
		cfgfile := filepath.Join(path, r.config.File)

		if _, err := os.Stat(cfgfile); !os.IsNotExist(err) {
			v := viper.New()

			if err = r.readConfig(v, cfgfile, path); err != nil {
				return nil, err
			}

			if cfg, err = r.mergeConfig(v); err != nil {
				return nil, err
			}
		}

		modules = append(modules, module{rootDir: path, config: cfg})
	}

	return modules, nil
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

// generateContent extracts print.Settings and terraform.Options from normalized
// Config and generates the output content for the module (and submodules if available)
// and write the result to the output (either stdout or a file).
func generateContent(config *print.Config) error {
	module, err := terraform.LoadWithOptions(config)
	if err != nil {
		return err
	}

	formatter, err := format.New(config)

	// formatter is unknown, this might mean that the intended formatter is
	// coming from a plugin. We are going to attempt to find a plugin with
	// that name and generate the content with it or error out if not found.
	if err != nil {
		plugins, perr := plugin.Discover()
		if perr != nil {
			return fmt.Errorf("formatter '%s' not found", config.Formatter)
		}

		client, found := plugins.Get(config.Formatter)
		if !found {
			return fmt.Errorf("formatter '%s' not found", config.Formatter)
		}

		content, cerr := client.Execute(&pluginsdk.ExecuteArgs{
			Module: module,
			Config: config,
		})
		if cerr != nil {
			return cerr
		}

		return writeContent(config, content)
	}

	err = formatter.Generate(module)
	if err != nil {
		return err
	}

	content, err := formatter.Render(config.Content)
	if err != nil {
		return err
	}

	return writeContent(config, content)
}

// writeContent to a Writer. This can either be os.Stdout or specific
// file (e.g. README.md) if '--output-file' is provided.
func writeContent(config *print.Config, content string) error {
	var w io.Writer

	// writing to a file (either inject or replace)
	if config.Output.File != "" {
		w = &fileWriter{
			file: config.Output.File,
			dir:  config.ModuleRoot,

			mode: config.Output.Mode,

			check: config.Output.Check,

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
