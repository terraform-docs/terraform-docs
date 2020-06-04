package cli

import (
	"fmt"

	"github.com/terraform-docs/terraform-docs/internal/module"
	"github.com/terraform-docs/terraform-docs/pkg/print"
)

type _sections struct {
	NoHeader       bool
	NoInputs       bool
	NoOutputs      bool
	NoProviders    bool
	NoRequirements bool
}
type sections struct {
	Show       []string  `yaml:"show"`
	Hide       []string  `yaml:"hide"`
	ShowAll    bool      `yaml:"show-all"`
	HideAll    bool      `yaml:"hide-all"`
	Deprecated _sections `yaml:"-"`

	header       bool `yaml:"-"`
	inputs       bool `yaml:"-"`
	outputs      bool `yaml:"-"`
	providers    bool `yaml:"-"`
	requirements bool `yaml:"-"`
}

func defaultSections() sections {
	return sections{
		Show:    []string{},
		Hide:    []string{},
		ShowAll: true,
		HideAll: false,
		Deprecated: _sections{
			NoHeader:       false,
			NoInputs:       false,
			NoOutputs:      false,
			NoProviders:    false,
			NoRequirements: false,
		},

		header:       false,
		inputs:       false,
		outputs:      false,
		providers:    false,
		requirements: false,
	}
}

func (s *sections) validate() error {
	items := []string{"header", "inputs", "outputs", "providers", "requirements"}
	for _, item := range s.Show {
		switch item {
		case items[0], items[1], items[2], items[3], items[4]:
		default:
			return fmt.Errorf("'%s' is not a valid section", item)
		}
	}
	for _, item := range s.Hide {
		switch item {
		case items[0], items[1], items[2], items[3], items[4]:
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
	for _, section := range items {
		if changedfs["no-"+section] && contains(s.Hide, section) {
			return fmt.Errorf("'--no-%s' and '--hide %s' can't be used together", section, section)
		}
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
	Required bool
	Type     bool
}
type _sort struct {
	NoSort bool
}
type sort struct {
	Enabled    bool     `yaml:"enabled"`
	ByList     []string `yaml:"by"`
	By         sortby   `yaml:"-"`
	Deprecated _sort    `yaml:"-"`
}

func defaultSort() sort {
	return sort{
		Enabled: true,
		ByList:  []string{},
		By: sortby{
			Required: false,
			Type:     false,
		},
		Deprecated: _sort{
			NoSort: false,
		},
	}
}

func (s *sort) validate() error {
	items := []string{"sort"}
	for _, item := range items {
		if changedfs[item] && changedfs["no-"+item] {
			return fmt.Errorf("'--%s' and '--no-%s' can't be used together", item, item)
		}
	}
	if s.By.Required && s.By.Type {
		return fmt.Errorf("'--sort-by-required' and '--sort-by-type' can't be used together")
	}
	return nil
}

type _settings struct {
	NoColor     bool
	NoEscape    bool
	NoRequired  bool
	NoSensitive bool
}
type settings struct {
	Color      bool      `yaml:"color"`
	Escape     bool      `yaml:"escape"`
	Indent     int       `yaml:"indent"`
	Required   bool      `yaml:"required"`
	Sensitive  bool      `yaml:"sensitive"`
	Deprecated _settings `yaml:"-"`
}

func defaultSettings() settings {
	return settings{
		Color:     true,
		Escape:    true,
		Indent:    2,
		Required:  true,
		Sensitive: true,
		Deprecated: _settings{
			NoColor:     false,
			NoEscape:    false,
			NoRequired:  false,
			NoSensitive: false,
		},
	}
}

func (s *settings) validate() error {
	items := []string{"escape", "color", "required", "sensitive"}
	for _, item := range items {
		if changedfs[item] && changedfs["no-"+item] {
			return fmt.Errorf("'--%s' and '--no-%s' can't be used together", item, item)
		}
	}
	return nil
}

// Config represents all the available config options that can be accessed and passed through CLI
type Config struct {
	File         string       `yaml:"-"`
	Formatter    string       `yaml:"formatter"`
	HeaderFrom   string       `yaml:"header-from"`
	Sections     sections     `yaml:"sections"`
	OutputValues outputvalues `yaml:"output-values"`
	Sort         sort         `yaml:"sort"`
	Settings     settings     `yaml:"settings"`
}

// DefaultConfig returns new instance of Config with default values set
func DefaultConfig() *Config {
	return &Config{
		File:         "",
		Formatter:    "",
		HeaderFrom:   "main.tf",
		Sections:     defaultSections(),
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
	c.Sections.inputs = c.Sections.visibility("inputs")
	c.Sections.outputs = c.Sections.visibility("outputs")
	c.Sections.providers = c.Sections.visibility("providers")
	c.Sections.requirements = c.Sections.visibility("requirements")

	// sort
	if !changedfs["sort"] && changedfs["no-sort"] {
		c.Sort.Enabled = !c.Sort.Deprecated.NoSort
	}

	// settings
	if !changedfs["escape"] && changedfs["no-escape"] {
		c.Settings.Escape = !c.Settings.Deprecated.NoEscape
	}
	if !changedfs["color"] && changedfs["no-color"] {
		c.Settings.Color = !c.Settings.Deprecated.NoColor
	}
	if !changedfs["required"] && changedfs["no-required"] {
		c.Settings.Required = !c.Settings.Deprecated.NoRequired
	}
	if !changedfs["sensitive"] && changedfs["no-sensitive"] {
		c.Settings.Sensitive = !c.Settings.Deprecated.NoSensitive
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

	// sections
	if err := c.Sections.validate(); err != nil {
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

// extract and build print.Settings and module.Options out of Config
func (c *Config) extract() (*print.Settings, *module.Options) {
	settings := print.NewSettings()
	options := module.NewOptions()

	// header-from
	options.HeaderFromFile = c.HeaderFrom

	// sections
	settings.ShowHeader = c.Sections.header
	settings.ShowInputs = c.Sections.inputs
	settings.ShowOutputs = c.Sections.outputs
	settings.ShowProviders = c.Sections.providers
	settings.ShowRequirements = c.Sections.requirements
	options.ShowHeader = settings.ShowHeader

	// output values
	settings.OutputValues = c.OutputValues.Enabled
	options.OutputValues = c.OutputValues.Enabled
	options.OutputValuesPath = c.OutputValues.From

	// sort
	settings.SortByName = c.Sort.Enabled
	settings.SortByRequired = c.Sort.Enabled && c.Sort.By.Required
	settings.SortByType = c.Sort.Enabled && c.Sort.By.Type
	options.SortBy.Name = settings.SortByName
	options.SortBy.Required = settings.SortByRequired
	options.SortBy.Type = settings.SortByType

	// settings
	settings.EscapeCharacters = c.Settings.Escape
	settings.IndentLevel = c.Settings.Indent
	settings.ShowColor = c.Settings.Color
	settings.ShowRequired = c.Settings.Required
	settings.ShowSensitivity = c.Settings.Sensitive

	return settings, options
}

func contains(list []string, name string) bool {
	for _, i := range list {
		if i == name {
			return true
		}
	}
	return false
}
