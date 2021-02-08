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
	sectionFooter       = "footer"
	sectionHeader       = "header"
	sectionInputs       = "inputs"
	sectionModules      = "modules"
	sectionOutputs      = "outputs"
	sectionProviders    = "providers"
	sectionRequirements = "requirements"
	sectionResources    = "resources"
)

var allSections = []string{sectionFooter, sectionHeader, sectionInputs, sectionModules, sectionOutputs, sectionProviders, sectionRequirements, sectionResources}

// AllSections list.
var AllSections = strings.Join(allSections, ", ")

type sections struct {
	Show    []string `yaml:"show"`
	Hide    []string `yaml:"hide"`
	ShowAll bool     `yaml:"show-all"`
	HideAll bool     `yaml:"hide-all"`

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
		Show:    []string{},
		Hide:    []string{},
		ShowAll: true,
		HideAll: false,

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
	for _, item := range s.Show {
		switch item {
		case allSections[0], allSections[1], allSections[2], allSections[3], allSections[4], allSections[5], allSections[6], allSections[7]:
		default:
			return fmt.Errorf("'%s' is not a valid section", item)
		}
	}
	for _, item := range s.Hide {
		switch item {
		case allSections[0], allSections[1], allSections[2], allSections[3], allSections[4], allSections[5], allSections[6], allSections[7]:
		default:
			return fmt.Errorf("'%s' is not a valid section", item)
		}
	}
	if s.ShowAll && s.HideAll {
		return fmt.Errorf("'--show-all' and '--hide-all' can't be used together")
	}
	if s.ShowAll && len(s.Show) != 0 {
		return fmt.Errorf("'--show-all' and '--show' can't be used together")
	}
	if s.HideAll && len(s.Hide) != 0 {
		return fmt.Errorf("'--hide-all' and '--hide' can't be used together")
	}
	return nil
}

func (s *sections) visibility(section string) bool {
	if s.ShowAll && !s.HideAll {
		for _, n := range s.Hide {
			if n == section {
				return false
			}
		}
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
	return false
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

func (o *output) validate() error {
	if o.File != "" {
		if o.Mode == "" {
			return fmt.Errorf("value of '--output-mode' can't be empty")
		}
		if o.Template == "" {
			return fmt.Errorf("value of '--output-template' can't be empty")
		}

		index := strings.Index(o.Template, outputContent)
		if index < 0 {
			return fmt.Errorf("value of '--output-template' doesn't have '{{ .Content }}' (note that spaces inside '{{ }}' are mandatory)")
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

type sortby struct {
	Required bool `name:"required"`
	Type     bool `name:"type"`
}
type sort struct {
	Enabled bool     `yaml:"enabled"`
	ByList  []string `yaml:"by"`
	By      sortby   `yaml:"-"`
}

func defaultSort() sort {
	return sort{
		Enabled: true,
		ByList:  []string{},
		By: sortby{
			Required: false,
			Type:     false,
		},
	}
}

func (s *sort) validate() error {
	types := []string{"required", "type"}
	for _, item := range s.ByList {
		switch item {
		case types[0], types[1]:
		default:
			return fmt.Errorf("'%s' is not a valid sort type", item)
		}
	}
	if s.By.Required && s.By.Type {
		return fmt.Errorf("'--sort-by-required' and '--sort-by-type' can't be used together")
	}
	return nil
}

type settings struct {
	Anchor    bool `yaml:"anchor"`
	Color     bool `yaml:"color"`
	Default   bool `yaml:"default"`
	Escape    bool `yaml:"escape"`
	Indent    int  `yaml:"indent"`
	Required  bool `yaml:"required"`
	Sensitive bool `yaml:"sensitive"`
	Type      bool `yaml:"type"`
}

func defaultSettings() settings {
	return settings{
		Anchor:    true,
		Color:     true,
		Default:   true,
		Escape:    true,
		Indent:    2,
		Required:  true,
		Sensitive: true,
		Type:      true,
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
	if c.Sections.HideAll && !changedfs["show-all"] {
		c.Sections.ShowAll = false
	}
	if !c.Sections.ShowAll && !changedfs["hide-all"] {
		c.Sections.HideAll = true
	}

	c.Sections.header = c.Sections.visibility("header")
	c.Sections.footer = c.Sections.visibility("footer")
	c.Sections.inputs = c.Sections.visibility("inputs")
	c.Sections.modulecalls = c.Sections.visibility("modules")
	c.Sections.outputs = c.Sections.visibility("outputs")
	c.Sections.providers = c.Sections.visibility("providers")
	c.Sections.requirements = c.Sections.visibility("requirements")
	c.Sections.resources = c.Sections.visibility("resources")

	// Footer section optional and should not cause error with --show-all
	if c.Sections.ShowAll && c.Sections.footer {
		c.Sections.footer = false
	}
}

// validate config and check for any misuse or misconfiguration
func (c *Config) validate() error {
	// formatter
	if c.Formatter == "" {
		return fmt.Errorf("value of 'formatter' can't be empty")
	}

	// header-from
	if c.HeaderFrom == "" {
		return fmt.Errorf("value of '--header-from' can't be empty")
	}

	// footer-from, not a 'default' section so can be empty even if show-all enabled
	if c.Sections.footer && !c.Sections.ShowAll && c.FooterFrom == "" {
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
	options.ShowHeader = settings.ShowHeader
	options.HeaderFromFile = c.HeaderFrom

	// footer-from
	options.ShowFooter = settings.ShowFooter
	options.FooterFromFile = c.FooterFrom

	// sections
	settings.ShowHeader = c.Sections.header
	settings.ShowFooter = c.Sections.footer
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
	options.SortBy.Name = c.Sort.Enabled
	options.SortBy.Required = c.Sort.Enabled && c.Sort.By.Required
	options.SortBy.Type = c.Sort.Enabled && c.Sort.By.Type

	// settings
	settings.EscapeCharacters = c.Settings.Escape
	settings.IndentLevel = c.Settings.Indent
	settings.ShowAnchor = c.Settings.Anchor
	settings.ShowColor = c.Settings.Color
	settings.ShowDefault = c.Settings.Default
	settings.ShowRequired = c.Settings.Required
	settings.ShowSensitivity = c.Settings.Sensitive
	settings.ShowType = c.Settings.Type

	return settings, options
}
