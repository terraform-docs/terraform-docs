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

const (
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
	Show []string `yaml:"show"`
	Hide []string `yaml:"hide"`

	dataSources  bool `yaml:"-"`
	header       bool `yaml:"-"`
	footer       bool `yaml:"-"`
	inputs       bool `yaml:"-"`
	modulecalls  bool `yaml:"-"`
	outputs      bool `yaml:"-"`
	providers    bool `yaml:"-"`
	requirements bool `yaml:"-"`
	resources    bool `yaml:"-"`
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
		if n == section {
			return true
		}
	}
	for _, n := range s.Hide {
		if n == section {
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
	File     string `yaml:"file"`
	Mode     string `yaml:"mode"`
	Template string `yaml:"template"`

	BeginComment string `yaml:"-"`
	EndComment   string `yaml:"-"`
}

func defaultOutput() output {
	return output{
		File:     "",
		Mode:     outputModeInject,
		Template: OutputTemplate,

		BeginComment: outputBeginComment,
		EndComment:   outputEndComment,
	}
}

func (o *output) validate() error { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	if o.File != "" {
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

		if index := strings.Index(o.Template, outputContent); index < 0 {
			return fmt.Errorf("value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)")
		}

		// No extra validation is needed for mode 'replace',
		// the followings only apply for every other modes.
		if o.Mode == outputModeReplace {
			return nil
		}

		lines := strings.Split(o.Template, "\n")
		if len(lines) < 3 {
			return fmt.Errorf("value of '--output-template' should contain at least 3 lines (begin comment, {{ .Content }}, and end comment)")
		}

		if !strings.Contains(lines[0], "<!--") || !strings.Contains(lines[0], "-->") {
			return fmt.Errorf("value of '--output-template' is missing begin comment")
		}
		o.BeginComment = strings.TrimSpace(lines[0])

		if !strings.Contains(lines[len(lines)-1], "<!--") || !strings.Contains(lines[len(lines)-1], "-->") {
			return fmt.Errorf("value of '--output-template' is missing end comment")
		}
		o.EndComment = strings.TrimSpace(lines[len(lines)-1])
	}
	return nil
}

type outputvalues struct {
	Enabled bool   `yaml:"enabled"`
	From    string `yaml:"from"`
}

func defaultOutputValues() outputvalues {
	return outputvalues{
		Enabled: false,
		From:    "",
	}
}

func (o *outputvalues) validate() error {
	if o.Enabled && o.From == "" {
		if changedfs["output-values-from"] {
			return fmt.Errorf("value of '--output-values-from' can't be empty")
		}
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

type sortby struct {
	Name     bool `name:"name"`
	Required bool `name:"required"`
	Type     bool `name:"type"`
}
type sort struct {
	Enabled  bool   `yaml:"enabled"`
	By       string `yaml:"by"`
	Criteria sortby `yaml:"-"`
}

func defaultSort() sort {
	return sort{
		Enabled: true,
		By:      sortName,
		Criteria: sortby{
			Name:     true,
			Required: false,
			Type:     false,
		},
	}
}

func (s *sort) validate() error {
	if !contains(allSorts, s.By) {
		return fmt.Errorf("'%s' is not a valid sort type", s.By)
	}
	if s.Criteria.Required && s.Criteria.Type {
		return fmt.Errorf("'--sort-by-required' and '--sort-by-type' can't be used together")
	}
	return nil
}

type settings struct {
	Anchor      bool `yaml:"anchor"`
	Color       bool `yaml:"color"`
	Default     bool `yaml:"default"`
	Escape      bool `yaml:"escape"`
	Indent      int  `yaml:"indent"`
	Required    bool `yaml:"required"`
	Sensitive   bool `yaml:"sensitive"`
	Type        bool `yaml:"type"`
	Description bool `yaml:"description"`
}

func defaultSettings() settings {
	return settings{
		Anchor:      true,
		Color:       true,
		Default:     true,
		Escape:      true,
		Indent:      2,
		Required:    true,
		Sensitive:   true,
		Type:        true,
		Description: false,
	}
}

func (s *settings) validate() error {
	return nil
}

// Config represents all the available config options that can be accessed and passed through CLI
type Config struct {
	BaseDir      string       `yaml:"-"`
	File         string       `yaml:"-"`
	Formatter    string       `yaml:"formatter"`
	HeaderFrom   string       `yaml:"header-from"`
	FooterFrom   string       `yaml:"footer-from"`
	Sections     sections     `yaml:"sections"`
	Output       output       `yaml:"output"`
	OutputValues outputvalues `yaml:"output-values"`
	Sort         sort         `yaml:"sort"`
	Settings     settings     `yaml:"settings"`
}

// DefaultConfig returns new instance of Config with default values set
func DefaultConfig() *Config {
	return &Config{
		BaseDir:      "",
		File:         "",
		Formatter:    "",
		HeaderFrom:   "main.tf",
		FooterFrom:   "",
		Sections:     defaultSections(),
		Output:       defaultOutput(),
		OutputValues: defaultOutputValues(),
		Sort:         defaultSort(),
		Settings:     defaultSettings(),
	}
}

// process provided Config
func (c *Config) process() {
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
	if c.FooterFrom == "" && !changedfs["footer-from"] {
		c.Sections.footer = false
	}

	// Enable specified sort criteria
	c.Sort.Criteria.Name = c.Sort.Enabled && c.Sort.By == sortName
	c.Sort.Criteria.Required = c.Sort.Enabled && c.Sort.By == sortRequired
	c.Sort.Criteria.Type = c.Sort.Enabled && c.Sort.By == sortType
}

// validate config and check for any misuse or misconfiguration
func (c *Config) validate() error { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	// formatter
	if c.Formatter == "" {
		return fmt.Errorf("value of 'formatter' can't be empty")
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

	// sections
	if err := c.Sections.validate(); err != nil {
		return err
	}

	// output
	if err := c.Output.validate(); err != nil {
		return err
	}

	// output values
	if err := c.OutputValues.validate(); err != nil {
		return err
	}

	// sort
	if err := c.Sort.validate(); err != nil {
		return err
	}

	// settings
	if err := c.Settings.validate(); err != nil {
		return err
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
	options.SortBy.Name = c.Sort.Enabled && c.Sort.Criteria.Name
	options.SortBy.Required = c.Sort.Enabled && c.Sort.Criteria.Required
	options.SortBy.Type = c.Sort.Enabled && c.Sort.Criteria.Type

	// settings
	settings.EscapeCharacters = c.Settings.Escape
	settings.IndentLevel = c.Settings.Indent
	settings.ShowAnchor = c.Settings.Anchor
	settings.ShowDescription = c.Settings.Description
	settings.ShowColor = c.Settings.Color
	settings.ShowDefault = c.Settings.Default
	settings.ShowRequired = c.Settings.Required
	settings.ShowSensitivity = c.Settings.Sensitive
	settings.ShowType = c.Settings.Type

	return settings, options
}
