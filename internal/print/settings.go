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
	"github.com/terraform-docs/terraform-docs/terraform"
)

// Settings represents all settings.
type Settings struct {
	// EscapeCharacters escapes special characters (such as _ * in Markdown and > < in JSON)
	//
	// default: true
	// scope: Markdown
	EscapeCharacters bool

	// IndentLevel control the indentation of headers [available: 1, 2, 3, 4, 5]
	//
	// default: 2
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues extract and show Output values from Terraform module output
	//
	// default: false
	// scope: Global
	OutputValues bool

	// ShowAnchor show html anchor
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowAnchor bool

	// ShowColor print "colorized" version of result in the terminal
	//
	// default: true
	// scope: Pretty
	ShowColor bool

	// ShowDatasources show the data sources on the "Resources" section
	//
	// default: true
	// scope: Global
	ShowDataSources bool

	// ShowDefault show "Default" column
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

	// ShowHTML use HTML tags (a, pre, br, ...)
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

	// ShowRequired show "Required" column
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" column
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

	// ShowType show "Type" column
	//
	// default: true
	// scope: Asciidoc, Markdown
	ShowType bool
}

// DefaultSettings returns new instance of Settings
func DefaultSettings() *Settings {
	return &Settings{
		EscapeCharacters: true,
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

// Convert internal Settings to its equivalent in plugin-sdk
func (s *Settings) Convert() *printsdk.Settings {
	return &printsdk.Settings{
		EscapeCharacters: s.EscapeCharacters,
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

// CopySections sets the sections that'll be printed
func CopySections(settings *Settings, src *terraform.Module, dest *terraform.Module) {
	if settings.ShowHeader {
		dest.Header = src.Header
	}
	if settings.ShowFooter {
		dest.Footer = src.Footer
	}
	if settings.ShowInputs {
		dest.Inputs = src.Inputs
	}
	if settings.ShowModuleCalls {
		dest.ModuleCalls = src.ModuleCalls
	}
	if settings.ShowOutputs {
		dest.Outputs = src.Outputs
	}
	if settings.ShowProviders {
		dest.Providers = src.Providers
	}
	if settings.ShowRequirements {
		dest.Requirements = src.Requirements
	}
	if settings.ShowResources || settings.ShowDataSources {
		dest.Resources = filterResourcesByMode(settings, src.Resources)
	}
}

// filterResourcesByMode returns the managed or data resources defined by the show argument
func filterResourcesByMode(settings *Settings, module []*terraform.Resource) []*terraform.Resource {
	resources := make([]*terraform.Resource, 0)
	for _, r := range module {
		if settings.ShowResources && r.Mode == "managed" {
			resources = append(resources, r)
		}
		if settings.ShowDataSources && r.Mode == "data" {
			resources = append(resources, r)
		}
	}
	return resources
}
