/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	printsdk "github.com/terraform-docs/plugin-sdk/print"
)

// Settings represents all settings.
type Settings struct {
	// EscapeCharacters escapes special characters (such as _ * in Markdown and > < in JSON)
	//
	// default: true
	// scope: Markdown
	EscapeCharacters bool

	// HideEmpty hide empty sections
	//
	// default: false
	// scope: Asciidoc, Markdown
	HideEmpty bool

	// IndentLevel control the indentation of headings [available: 1, 2, 3, 4, 5]
	//
	// default: 2
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues extract and show Output values from Terraform module output
	//
	// default: false
	// scope: Global
	OutputValues bool

	// ShowAnchor generate HTML anchor tag for elements
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowAnchor bool

	// ShowColor print "colorized" version of result in the terminal
	//
	// default: true
	// scope: Pretty
	ShowColor bool

	// ShowDataSources show the data sources on the "Resources" section
	//
	// default: true
	// scope: Global
	ShowDataSources bool

	// ShowDefault show "Default" as column (in table) or section (in document)
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowDefault bool

	// ShowDescription show "Descriptions" as comment on variables
	//
	// default: false
	// scope: tfvars hcl
	ShowDescription bool

	// ShowFooter show "Footer" module information
	//
	// default: false
	// scope: Global
	ShowFooter bool

	// ShowHeader show "Header" module information
	//
	// default: true
	// scope: Global
	ShowHeader bool

	// ShowHTML generate HTML tags (a, pre, br, ...) in the output
	//
	// default: true
	// scope: Markdown
	ShowHTML bool

	// ShowInputs show "Inputs" information
	//
	// default: true
	// scope: Global
	ShowInputs bool

	// ShowModuleCalls show "ModuleCalls" information
	//
	// default: true
	// scope: Global
	ShowModuleCalls bool

	// ShowOutputs show "Outputs" information
	//
	// default: true
	// scope: Global
	ShowOutputs bool

	// ShowProviders show "Providers" information
	//
	// default: true
	// scope: Global
	ShowProviders bool

	// ShowRequired show "Required" as column (in table) or section (in document)
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" as column (in table) or section (in document)
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowSensitivity bool

	// ShowRequirements show "Requirements" section
	//
	// default: true
	// scope: Global
	ShowRequirements bool

	// ShowResources show "Resources" section
	//
	// default: true
	// scope: Global
	ShowResources bool

	// ShowType show "Type" as column (in table) or section (in document)
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowType bool
}

// DefaultSettings returns new instance of Settings.
func DefaultSettings() *Settings {
	return &Settings{
		EscapeCharacters: true,
		HideEmpty:        false,
		IndentLevel:      2,
		OutputValues:     false,
		ShowAnchor:       true,
		ShowColor:        true,
		ShowDataSources:  true,
		ShowDefault:      true,
		ShowDescription:  false,
		ShowFooter:       false,
		ShowHeader:       true,
		ShowHTML:         true,
		ShowInputs:       true,
		ShowModuleCalls:  true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequired:     true,
		ShowSensitivity:  true,
		ShowRequirements: true,
		ShowResources:    true,
		ShowType:         true,
	}
}

// ToConfig converts Settings to Config.
func (s *Settings) ToConfig() *Config {
	config := NewConfig()

	config.Settings.Anchor = s.ShowAnchor
	config.Settings.Color = s.ShowColor
	config.Settings.Default = s.ShowDefault
	config.Settings.Description = s.ShowDescription
	config.Settings.Escape = s.EscapeCharacters
	config.Settings.HideEmpty = s.HideEmpty
	config.Settings.HTML = s.ShowHTML
	config.Settings.Indent = s.IndentLevel
	config.Settings.Required = s.ShowRequired
	config.Settings.Sensitive = s.ShowSensitivity
	config.Settings.Type = s.ShowType

	config.OutputValues.Enabled = s.OutputValues

	config.Sections.DataSources = s.ShowDataSources
	config.Sections.Footer = s.ShowFooter
	config.Sections.Header = s.ShowHeader
	config.Sections.Inputs = s.ShowInputs
	config.Sections.Outputs = s.ShowOutputs
	config.Sections.Modulecalls = s.ShowModuleCalls
	config.Sections.Providers = s.ShowProviders
	config.Sections.Requirements = s.ShowRequirements
	config.Sections.Resources = s.ShowResources

	return config
}

// Convert Settings to its equivalent in plugin-sdk.
func (s *Settings) Convert() *printsdk.Settings {
	return &printsdk.Settings{
		EscapeCharacters: s.EscapeCharacters,
		HideEmpty:        s.HideEmpty,
		IndentLevel:      s.IndentLevel,
		OutputValues:     s.OutputValues,
		ShowColor:        s.ShowColor,
		ShowDataSources:  s.ShowDataSources,
		ShowDefault:      s.ShowDefault,
		ShowDescription:  s.ShowDescription,
		ShowFooter:       s.ShowFooter,
		ShowHeader:       s.ShowHeader,
		ShowInputs:       s.ShowInputs,
		ShowOutputs:      s.ShowOutputs,
		ShowModuleCalls:  s.ShowModuleCalls,
		ShowProviders:    s.ShowProviders,
		ShowRequired:     s.ShowRequired,
		ShowSensitivity:  s.ShowSensitivity,
		ShowRequirements: s.ShowRequirements,
		ShowResources:    s.ShowResources,
		ShowType:         s.ShowType,
	}
}
