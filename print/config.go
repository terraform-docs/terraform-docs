/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Config represents all the available config options that can be accessed and
// passed through CLI.
type Config struct {
	File         string       `mapstructure:"-"`
	Formatter    string       `mapstructure:"formatter"`
	Version      string       `mapstructure:"version"`
	HeaderFrom   string       `mapstructure:"header-from"`
	FooterFrom   string       `mapstructure:"footer-from"`
	Recursive    recursive    `mapstructure:"recursive"`
	Content      string       `mapstructure:"content"`
	Sections     sections     `mapstructure:"sections"`
	Output       output       `mapstructure:"output"`
	OutputValues outputvalues `mapstructure:"output-values"`
	Sort         sort         `mapstructure:"sort"`
	Settings     settings     `mapstructure:"settings"`

	ModuleRoot string
}

// NewConfig returns new instances of Config with empty values.
func NewConfig() *Config {
	return &Config{
		HeaderFrom:   "main.tf",
		Recursive:    recursive{},
		Sections:     sections{},
		Output:       output{},
		OutputValues: outputvalues{},
		Sort:         sort{},
		Settings:     settings{},
	}
}

// DefaultConfig returns new instance of Config with default values set.
func DefaultConfig() *Config {
	return &Config{
		File:         "",
		Formatter:    "",
		Version:      "",
		HeaderFrom:   "main.tf",
		FooterFrom:   "",
		Recursive:    defaultRecursive(),
		Content:      "",
		Sections:     defaultSections(),
		Output:       defaultOutput(),
		OutputValues: defaultOutputValues(),
		Sort:         defaultSort(),
		Settings:     defaultSettings(),

		ModuleRoot: "",
	}
}

type recursive struct {
	Enabled     bool   `mapstructure:"enabled"`
	Path        string `mapstructure:"path"`
	IncludeMain bool   `mapstructure:"include-main"`
}

func defaultRecursive() recursive {
	return recursive{
		Enabled:     false,
		Path:        "modules",
		IncludeMain: true,
	}
}

func (r *recursive) validate() error {
	if r.Enabled && r.Path == "" {
		return fmt.Errorf("value of '--recursive-path' can't be empty")
	}
	return nil
}

const (
	sectionAll          = "all"
	sectionDataSources  = "data-sources"
	sectionFooter       = "footer"
	sectionHeader       = "header"
	sectionInputs       = "inputs"
	sectionModules      = "modules"
	sectionOutputs      = "outputs"
	sectionProviders    = "providers"
	sectionRequirements = "requirements"
	sectionResources    = "resources"
)

var allSections = []string{
	sectionAll,
	sectionDataSources,
	sectionFooter,
	sectionHeader,
	sectionInputs,
	sectionModules,
	sectionOutputs,
	sectionProviders,
	sectionRequirements,
	sectionResources,
}

// AllSections list.
var AllSections = strings.Join(allSections, ", ")

type sections struct {
	Show []string `mapstructure:"show"`
	Hide []string `mapstructure:"hide"`

	DataSources  bool
	Header       bool
	Footer       bool
	Inputs       bool
	ModuleCalls  bool
	Outputs      bool
	Providers    bool
	Requirements bool
	Resources    bool
}

func defaultSections() sections {
	return sections{
		Show: []string{},
		Hide: []string{},

		DataSources:  true,
		Header:       true,
		Footer:       false,
		Inputs:       true,
		ModuleCalls:  true,
		Outputs:      true,
		Providers:    true,
		Requirements: true,
		Resources:    true,
	}
}

func (s *sections) validate() error {
	if len(s.Show) > 0 && len(s.Hide) > 0 {
		return fmt.Errorf("'--show' and '--hide' can't be used together")
	}
	for _, item := range s.Show {
		if !contains(allSections, item) {
			return fmt.Errorf("'%s' is not a valid section", item)
		}
	}
	for _, item := range s.Hide {
		if !contains(allSections, item) {
			return fmt.Errorf("'%s' is not a valid section", item)
		}
	}
	return nil
}

func (s *sections) visibility(section string) bool {
	if len(s.Show) == 0 && len(s.Hide) == 0 {
		return true
	}
	for _, n := range s.Show {
		if n == sectionAll || n == section {
			return true
		}
	}
	for _, n := range s.Hide {
		if n == sectionAll || n == section {
			return false
		}
	}
	// hidden : if s.Show NOT empty AND s.Show does NOT contain section
	// visible: if s.Hide NOT empty AND s.Hide does NOT contain section
	return len(s.Hide) > 0
}

// Output modes.
const (
	OutputModeInject  = "inject"
	OutputModeReplace = "replace"
)

// Output template.
const (
	OutputBeginComment = "<!-- BEGIN_TF_DOCS -->"
	OutputContent      = "{{ .Content }}"
	OutputEndComment   = "<!-- END_TF_DOCS -->"
)

// Output to file template and modes.
var (
	OutputTemplate = fmt.Sprintf("%s\n%s\n%s", OutputBeginComment, OutputContent, OutputEndComment)
	OutputModes    = strings.Join([]string{OutputModeInject, OutputModeReplace}, ", ")
)

type output struct {
	File     string `mapstructure:"file"`
	Mode     string `mapstructure:"mode"`
	Template string `mapstructure:"template"`
	Check    bool

	BeginComment string
	EndComment   string
}

func defaultOutput() output {
	return output{
		File:     "",
		Mode:     OutputModeInject,
		Template: OutputTemplate,
		Check:    false,

		BeginComment: OutputBeginComment,
		EndComment:   OutputEndComment,
	}
}

func (o *output) validate() error {
	if o.File == "" {
		return nil
	}

	if o.Mode == "" {
		return fmt.Errorf("value of '--output-mode' can't be empty")
	}

	// Template is optional for mode 'replace'
	if o.Mode == OutputModeReplace && o.Template == "" {
		return nil
	}

	if o.Template == "" {
		return fmt.Errorf("value of '--output-template' can't be empty")
	}

	if !strings.Contains(o.Template, OutputContent) {
		return fmt.Errorf("value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)")
	}

	// No extra validation is needed for mode 'replace',
	// the followings only apply for every other modes.
	if o.Mode == OutputModeReplace {
		return nil
	}

	o.Template = strings.ReplaceAll(o.Template, "\\n", "\n")
	lines := strings.Split(o.Template, "\n")
	tests := []struct {
		condition  func() bool
		errMessage string
	}{
		{
			condition: func() bool {
				return len(lines) < 3
			},
			errMessage: "value of '--output-template' should contain at least 3 lines (begin comment, {{ .Content }}, and end comment)",
		},
		{
			condition: func() bool {
				return !isInlineComment(lines[0])
			},
			errMessage: "value of '--output-template' is missing begin comment",
		},
		{
			condition: func() bool {
				return !isInlineComment(lines[len(lines)-1])
			},
			errMessage: "value of '--output-template' is missing end comment",
		},
	}

	for _, t := range tests {
		if t.condition() {
			return fmt.Errorf(t.errMessage)
		}
	}

	o.BeginComment = strings.TrimSpace(lines[0])
	o.EndComment = strings.TrimSpace(lines[len(lines)-1])

	return nil
}

// Detect if a particular line is a Markdown comment.
//
// ref: https://www.jamestharpe.com/markdown-comments/
func isInlineComment(line string) bool {
	switch {
	// AsciiDoc specific
	case strings.HasPrefix(line, "//"):
		return true

	// Markdown specific
	default:
		cases := [][]string{
			{"<!--", "-->"},
			{"[]: # (", ")"},
			{"[]: # \"", "\""},
			{"[]: # '", "'"},
			{"[//]: # (", ")"},
			{"[comment]: # (", ")"},
		}
		for _, c := range cases {
			if strings.HasPrefix(line, c[0]) && strings.HasSuffix(line, c[1]) {
				return true
			}
		}
	}

	return false
}

type outputvalues struct {
	Enabled bool   `mapstructure:"enabled"`
	From    string `mapstructure:"from"`
}

func defaultOutputValues() outputvalues {
	return outputvalues{
		Enabled: false,
		From:    "",
	}
}

func (o *outputvalues) validate() error {
	if o.Enabled && o.From == "" {
		return fmt.Errorf("value of '--output-values-from' is missing")
	}
	return nil
}

// Sort types.
const (
	SortName     = "name"
	SortRequired = "required"
	SortType     = "type"
)

var allSorts = []string{
	SortName,
	SortRequired,
	SortType,
}

// SortTypes list.
var SortTypes = strings.Join(allSorts, ", ")

type sort struct {
	Enabled bool   `mapstructure:"enabled"`
	By      string `mapstructure:"by"`
}

func defaultSort() sort {
	return sort{
		Enabled: true,
		By:      SortName,
	}
}

func (s *sort) validate() error {
	if !contains(allSorts, s.By) {
		return fmt.Errorf("'%s' is not a valid sort type", s.By)
	}
	return nil
}

type settings struct {
	Anchor       bool `mapstructure:"anchor"`
	Color        bool `mapstructure:"color"`
	Default      bool `mapstructure:"default"`
	Description  bool `mapstructure:"description"`
	Escape       bool `mapstructure:"escape"`
	HideEmpty    bool `mapstructure:"hide-empty"`
	HTML         bool `mapstructure:"html"`
	Indent       int  `mapstructure:"indent"`
	LockFile     bool `mapstructure:"lockfile"`
	ReadComments bool `mapstructure:"read-comments"`
	Required     bool `mapstructure:"required"`
	Sensitive    bool `mapstructure:"sensitive"`
	Type         bool `mapstructure:"type"`
	Validation   bool `mapstructure:"validation"`
}

func defaultSettings() settings {
	return settings{
		Anchor:       true,
		Color:        true,
		Default:      true,
		Description:  false,
		Escape:       true,
		HideEmpty:    false,
		HTML:         true,
		Indent:       2,
		LockFile:     true,
		ReadComments: true,
		Required:     true,
		Sensitive:    true,
		Type:         true,
		Validation:   true,
	}
}

func (s *settings) validate() error {
	return nil
}

// Parse process config and set sections visibility.
func (c *Config) Parse() {
	// sections
	c.Sections.DataSources = c.Sections.visibility("data-sources")
	c.Sections.Header = c.Sections.visibility("header")
	c.Sections.Inputs = c.Sections.visibility("inputs")
	c.Sections.ModuleCalls = c.Sections.visibility("modules")
	c.Sections.Outputs = c.Sections.visibility("outputs")
	c.Sections.Providers = c.Sections.visibility("providers")
	c.Sections.Requirements = c.Sections.visibility("requirements")
	c.Sections.Resources = c.Sections.visibility("resources")

	// Footer section is optional and should only be enabled if --footer-from
	// is explicitly set, either via CLI or config file.
	if c.FooterFrom != "" {
		c.Sections.Footer = c.Sections.visibility("footer")
	}
}

// Validate provided Config and check for any misuse or misconfiguration.
func (c *Config) Validate() error {
	// formatter
	if c.Formatter == "" {
		return fmt.Errorf("value of 'formatter' can't be empty")
	}

	// header-from
	if c.HeaderFrom == "" {
		return fmt.Errorf("value of '--header-from' can't be empty")
	}

	// footer-from, not a 'default' section so can be empty
	if c.Sections.Footer && c.FooterFrom == "" {
		return fmt.Errorf("value of '--footer-from' can't be empty")
	}

	if c.FooterFrom == c.HeaderFrom {
		return fmt.Errorf("value of '--footer-from' can't equal value of '--header-from")
	}

	for _, fn := range [](func() error){
		c.Recursive.validate,
		c.Sections.validate,
		c.Output.validate,
		c.OutputValues.validate,
		c.Sort.validate,
		c.Settings.validate,
	} {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

// ReadConfig reads config file in `rootDir` with given `filename` and returns
// instance of Config. It returns error if config file not found or there is a
// problem with unmarshalling.
func ReadConfig(rootDir string, filename string) (*Config, error) {
	cfg := NewConfig()

	v := viper.New()
	v.SetConfigFile(path.Join(rootDir, filename))

	if err := v.ReadInConfig(); err != nil {
		var perr *os.PathError
		if errors.As(err, &perr) {
			return nil, fmt.Errorf("config file %s not found", filename)
		}

		var cerr viper.ConfigFileNotFoundError
		if !errors.As(err, &cerr) {
			return nil, err
		}
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config, %w", err)
	}

	cfg.ModuleRoot = rootDir

	// process and validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	cfg.Parse()

	return cfg, nil
}
