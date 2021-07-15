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
	"strings"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// Mappings of CLI flags to Viper config
var flagMappings = map[string]string{
	"header-from": "header-from",
	"footer-from": "footer-from",

	"show": "sections.show",
	"hide": "sections.hide",

	"output-file":     "output.file",
	"output-mode":     "output.mode",
	"output-template": "output.template",

	"output-values":      "output-values.enabled",
	"output-values-from": "output-values.from",

	"sort":             "sort.enabled",
	"sort-by":          "sort.by",
	"sort-by-required": "required",
	"sort-by-type":     "type",

	"anchor":      "settings.anchor",
	"color":       "settings.color",
	"default":     "settings.default",
	"description": "settings.description",
	"escape":      "settings.escape",
	"indent":      "settings.indent",
	"required":    "settings.required",
	"sensitive":   "settings.sensitive",
	"type":        "settings.type",
}

// Config represents all the available config options that can be accessed and passed through CLI
type Config struct {
	File       string `mapstructure:"-"`
	Formatter  string `mapstructure:"formatter"`
	Version    string `mapstructure:"version"`
	HeaderFrom string `mapstructure:"header-from"`
	FooterFrom string `mapstructure:"footer-from"`
	// TOOD
	UseLockFile  bool         `mapstructure:"lockfile"`
	Content      string       `mapstructure:"content"`
	Sections     sections     `mapstructure:"sections"`
	Output       output       `mapstructure:"output"`
	OutputValues outputvalues `mapstructure:"output-values"`
	Sort         sort         `mapstructure:"sort"`
	Settings     settings     `mapstructure:"settings"`

	moduleRoot    string
	isFlagChanged func(string) bool
}

// DefaultConfig returns new instance of Config with default values set
func DefaultConfig() *Config {
	return &Config{
		File:         "",
		Formatter:    "",
		Version:      "",
		HeaderFrom:   "main.tf",
		FooterFrom:   "",
		UseLockFile:  true,
		Content:      "",
		Sections:     defaultSections(),
		Output:       defaultOutput(),
		OutputValues: defaultOutputValues(),
		Sort:         defaultSort(),
		Settings:     defaultSettings(),

		moduleRoot:    "",
		isFlagChanged: func(name string) bool { return false },
	}
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

	dataSources  bool
	header       bool
	footer       bool
	inputs       bool
	modulecalls  bool
	outputs      bool
	providers    bool
	requirements bool
	resources    bool
}

func defaultSections() sections {
	return sections{
		Show: []string{},
		Hide: []string{},

		dataSources:  false,
		header:       false,
		footer:       false,
		inputs:       false,
		modulecalls:  false,
		outputs:      false,
		providers:    false,
		requirements: false,
		resources:    false,
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

const (
	outputModeInject  = "inject"
	outputModeReplace = "replace"

	outputBeginComment = "<!-- BEGIN_TF_DOCS -->"
	outputContent      = "{{ .Content }}"
	outputEndComment   = "<!-- END_TF_DOCS -->"
)

// Output to file template and modes
var (
	OutputTemplate = fmt.Sprintf("%s\n%s\n%s", outputBeginComment, outputContent, outputEndComment)
	OutputModes    = strings.Join([]string{outputModeInject, outputModeReplace}, ", ")
)

type output struct {
	File     string `mapstructure:"file"`
	Mode     string `mapstructure:"mode"`
	Template string `mapstructure:"template"`

	beginComment string
	endComment   string
}

func defaultOutput() output {
	return output{
		File:     "",
		Mode:     outputModeInject,
		Template: OutputTemplate,

		beginComment: outputBeginComment,
		endComment:   outputEndComment,
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
	if o.Mode == outputModeReplace && o.Template == "" {
		return nil
	}

	if o.Template == "" {
		return fmt.Errorf("value of '--output-template' can't be empty")
	}

	if !strings.Contains(o.Template, outputContent) {
		return fmt.Errorf("value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)")
	}

	// No extra validation is needed for mode 'replace',
	// the followings only apply for every other modes.
	if o.Mode == outputModeReplace {
		return nil
	}

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

	o.beginComment = strings.TrimSpace(lines[0])
	o.endComment = strings.TrimSpace(lines[len(lines)-1])

	return nil
}

// Detect if a particular line is a Markdown comment
//
// ref: https://www.jamestharpe.com/markdown-comments/
func isInlineComment(line string) bool {
	switch {
	// Markdown specific
	case strings.HasPrefix(line, "<!--") && strings.HasSuffix(line, "-->"):
		return true
	case strings.HasPrefix(line, "[]: # ("):
		return true
	case strings.HasPrefix(line, "[]: # \""):
		return true
	case strings.HasPrefix(line, "[]: # '"):
		return true
	case strings.HasPrefix(line, "[//]: # ("):
		return true
	case strings.HasPrefix(line, "[comment]: # ("):
		return true

	// AsciiDoc specific
	case strings.HasPrefix(line, "//"):
		return true
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

const (
	sortName     = "name"
	sortRequired = "required"
	sortType     = "type"
)

var allSorts = []string{
	sortName,
	sortRequired,
	sortType,
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
		By:      sortName,
	}
}

func (s *sort) validate() error {
	if !contains(allSorts, s.By) {
		return fmt.Errorf("'%s' is not a valid sort type", s.By)
	}
	return nil
}

type settings struct {
	Anchor      bool `mapstructure:"anchor"`
	Color       bool `mapstructure:"color"`
	Default     bool `mapstructure:"default"`
	Description bool `mapstructure:"description"`
	Escape      bool `mapstructure:"escape"`
	HTML        bool `mapstructure:"html"`
	Indent      int  `mapstructure:"indent"`
	Required    bool `mapstructure:"required"`
	Sensitive   bool `mapstructure:"sensitive"`
	Type        bool `mapstructure:"type"`
}

func defaultSettings() settings {
	return settings{
		Anchor:      true,
		Color:       true,
		Default:     true,
		Description: false,
		Escape:      true,
		HTML:        true,
		Indent:      2,
		Required:    true,
		Sensitive:   true,
		Type:        true,
	}
}

func (s *settings) validate() error {
	return nil
}

// process provided Config and check for any misuse or misconfiguration
func (c *Config) process() error { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	// formatter
	if c.Formatter == "" {
		return fmt.Errorf("value of 'formatter' can't be empty")
	}

	// sections
	c.Sections.dataSources = c.Sections.visibility("data-sources")
	c.Sections.header = c.Sections.visibility("header")
	c.Sections.footer = c.Sections.visibility("footer")
	c.Sections.inputs = c.Sections.visibility("inputs")
	c.Sections.modulecalls = c.Sections.visibility("modules")
	c.Sections.outputs = c.Sections.visibility("outputs")
	c.Sections.providers = c.Sections.visibility("providers")
	c.Sections.requirements = c.Sections.visibility("requirements")
	c.Sections.resources = c.Sections.visibility("resources")

	// Footer section is optional and should only be enabled if --footer-from
	// is explicitly set, either via CLI or config file.
	if c.FooterFrom == "" && !c.isFlagChanged("footer-from") {
		c.Sections.footer = false
	}

	// header-from
	if c.HeaderFrom == "" {
		return fmt.Errorf("value of '--header-from' can't be empty")
	}

	// footer-from, not a 'default' section so can be empty
	if c.Sections.footer && c.FooterFrom == "" {
		return fmt.Errorf("value of '--footer-from' can't be empty")
	}

	if c.FooterFrom == c.HeaderFrom {
		return fmt.Errorf("value of '--footer-from' can't equal value of '--header-from")
	}

	for _, fn := range [](func() error){
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

// extract and build print.Settings and terraform.Options out of Config
func (c *Config) extract() (*print.Settings, *terraform.Options) {
	settings := print.DefaultSettings()
	options := terraform.NewOptions()

	// header-from
	settings.ShowHeader = c.Sections.header
	options.ShowHeader = settings.ShowHeader
	options.HeaderFromFile = c.HeaderFrom

	// footer-from
	settings.ShowFooter = c.Sections.footer
	options.ShowFooter = settings.ShowFooter
	options.FooterFromFile = c.FooterFrom

	// lock file
	options.UseLockFile = c.UseLockFile

	// sections
	settings.ShowDataSources = c.Sections.dataSources
	settings.ShowInputs = c.Sections.inputs
	settings.ShowModuleCalls = c.Sections.modulecalls
	settings.ShowOutputs = c.Sections.outputs
	settings.ShowProviders = c.Sections.providers
	settings.ShowRequirements = c.Sections.requirements
	settings.ShowResources = c.Sections.resources

	// output values
	settings.OutputValues = c.OutputValues.Enabled
	options.OutputValues = c.OutputValues.Enabled
	options.OutputValuesPath = c.OutputValues.From

	// sort
	options.SortBy.Name = c.Sort.Enabled && c.Sort.By == sortName
	options.SortBy.Required = c.Sort.Enabled && c.Sort.By == sortRequired
	options.SortBy.Type = c.Sort.Enabled && c.Sort.By == sortType

	// settings
	settings.EscapeCharacters = c.Settings.Escape
	settings.IndentLevel = c.Settings.Indent
	settings.ShowAnchor = c.Settings.Anchor
	settings.ShowDescription = c.Settings.Description
	settings.ShowColor = c.Settings.Color
	settings.ShowDefault = c.Settings.Default
	settings.ShowHTML = c.Settings.HTML
	settings.ShowRequired = c.Settings.Required
	settings.ShowSensitivity = c.Settings.Sensitive
	settings.ShowType = c.Settings.Type

	return settings, options
}
