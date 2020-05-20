package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/pkg/print"
)

type sections struct {
	Hide []string

	header       bool
	inputs       bool
	outputs      bool
	providers    bool
	requirements bool
}

func (s *sections) contains(section string) bool {
	for _, item := range s.Hide {
		if item == section {
			return true
		}
	}
	return false
}

func (s *sections) validate(fs *pflag.FlagSet) error {
	sections := []string{"header", "inputs", "outputs", "providers", "requirements"}
	for _, h := range s.Hide {
		switch h {
		case sections[0], sections[1], sections[2], sections[3], sections[4]:
		default:
			return fmt.Errorf("'%s' is not a valid section", h)
		}
	}
	for _, section := range sections {
		if fs.Changed("no-"+section) && s.contains(section) {
			return fmt.Errorf("'--no-%s' and '--hide %s' cannot be used together", section, section)
		}
	}
	return nil
}

type outputvalues struct {
	Enabled bool
	From    string
}

func (o *outputvalues) validate() error {
	if o.Enabled && o.From == "" {
		return fmt.Errorf("value of '--output-values-from' cannot be empty")
	}
	return nil
}

type sortby struct {
	Required bool
	Type     bool
}
type sort struct {
	Enabled bool
	By      *sortby
}

func (s *sort) validate(fs *pflag.FlagSet) error {
	items := []string{"sort"}
	for _, item := range items {
		if fs.Changed("no-"+item) && fs.Changed(item) {
			return fmt.Errorf("'--no-%s' and '--%s' cannot be used together", item, item)
		}
	}
	return nil
}

type settings struct {
	Color     bool
	Escape    bool
	Indent    int
	Required  bool
	Sensitive bool
}

func (s *settings) validate(fs *pflag.FlagSet) error {
	items := []string{"escape", "color", "required", "sensitive"}
	for _, item := range items {
		if fs.Changed("no-"+item) && fs.Changed(item) {
			return fmt.Errorf("'--no-%s' and '--%s' cannot be used together", item, item)
		}
	}
	if fs.Changed("sort-by-required") && fs.Changed("sort-by-type") {
		return fmt.Errorf("'--sort-by-required' and '--sort-by-type' cannot be used together")
	}
	return nil
}

// Config represents all the available config options that can be accessed and passed through CLI
type Config struct {
	Formatter    string
	HeaderFrom   string
	Sections     *sections
	OutputValues *outputvalues
	Sort         *sort
	Settings     *settings
}

// DefaultConfig returns new instance of Config with default values set
func DefaultConfig() *Config {
	return &Config{
		Formatter:  "",
		HeaderFrom: "main.tf",
		Sections: &sections{
			Hide: []string{},

			header:       true,
			inputs:       true,
			outputs:      true,
			providers:    true,
			requirements: true,
		},
		OutputValues: &outputvalues{
			Enabled: false,
			From:    "",
		},
		Sort: &sort{
			Enabled: true,
			By: &sortby{
				Required: false,
				Type:     false,
			},
		},
		Settings: &settings{
			Color:     true,
			Escape:    true,
			Indent:    2,
			Required:  false,
			Sensitive: false,
		},
	}
}

// extract and build print.Settings and module.Options out of Config
func (c *Config) extract() (*print.Settings, *module.Options, error) {
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

	return settings, options, nil
}

// normalize provided Config and check for any misuse or misconfiguration
func normalize(command string, fs *pflag.FlagSet, config *Config) error {
	config.Formatter = strings.Replace(command, "terraform-docs ", "", -1)

	// header-from
	if fs.Changed("header-from") && config.HeaderFrom == "" {
		return fmt.Errorf("value of '--header-from' cannot be empty")
	}

	// sections
	if err := config.Sections.validate(fs); err != nil {
		return err
	}

	config.Sections.header = !(config.Sections.contains("header") || fs.Changed("no-header"))
	config.Sections.inputs = !(config.Sections.contains("inputs") || fs.Changed("no-inputs"))
	config.Sections.outputs = !(config.Sections.contains("outputs") || fs.Changed("no-outputs"))
	config.Sections.providers = !(config.Sections.contains("providers") || fs.Changed("no-providers"))
	config.Sections.requirements = !(config.Sections.contains("requirements") || fs.Changed("no-requirements"))

	// output values
	if err := config.OutputValues.validate(); err != nil {
		return err
	}

	// sort
	if err := config.Sort.validate(fs); err != nil {
		return err
	}

	if !fs.Changed("sort") {
		config.Sort.Enabled = !fs.Changed("no-sort")
	}

	// settings
	if err := config.Settings.validate(fs); err != nil {
		return err
	}

	if !fs.Changed("escape") {
		config.Settings.Escape = !fs.Changed("no-escape")
	}
	if !fs.Changed("color") {
		config.Settings.Color = !fs.Changed("no-color")
	}
	if !fs.Changed("required") {
		config.Settings.Required = !fs.Changed("no-required")
	}
	if !fs.Changed("sensitive") {
		config.Settings.Sensitive = !fs.Changed("no-sensitive")
	}

	return nil
}
