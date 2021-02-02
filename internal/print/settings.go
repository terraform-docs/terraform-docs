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

	// EscapePipe escapes pipe character in Markdown
	//
	// default: true
	// scope: Markdown
	EscapePipe bool

	// IndentLevel control the indentation of AsciiDoc and Markdown headers [available: 1, 2, 3, 4, 5]
	//
	// default: 2
	// scope: Asciidoc, Markdown
	IndentLevel int

	// OutputValues extract and show Output values from Terraform module output
	//
	// default: false
	// scope: Global
	OutputValues bool

	// ShowColor print "colorized" version of result in the terminal
	//
	// default: true
	// scope: Pretty
	ShowColor bool

	// ShowHeader show "Header" module information
	//
	// default: true
	// scope: Global
	ShowHeader bool

	// ShowInputs show "Inputs" information
	//
	// default: true
	// scope: Global
	ShowInputs bool

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

	// ShowRequired show "Required" column when generating Markdown
	//
	// default: true
	// scope: Markdown
	ShowRequired bool

	// ShowSensitivity show "Sensitive" column when generating Markdown
	//
	// default: true
	// scope: Markdown
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

	// SortByName sorted rendering of inputs and outputs
	//
	// default: true
	// scope: Global
	SortByName bool

	// SortByRequired sort items (inputs, providers) by name and prints required ones first
	//
	// default: false
	// scope: Global
	SortByRequired bool

	// SortByType sort items (inputs, outputs) by type alphabetically
	//
	// default: false
	// scope: Global
	SortByType bool
}

// DefaultSettings returns new instance of Settings
func DefaultSettings() *Settings {
	return &Settings{
		EscapeCharacters: true,
		EscapePipe:       true,
		IndentLevel:      2,
		OutputValues:     false,
		ShowColor:        true,
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequired:     true,
		ShowSensitivity:  true,
		ShowRequirements: true,
		ShowResources:    true,
		SortByName:       true,
		SortByRequired:   false,
		SortByType:       false,
	}
}

// Convert internal Settings to its equivalent in plugin-sdk
func (s *Settings) Convert() *printsdk.Settings {
	return &printsdk.Settings{
		EscapeCharacters: s.EscapeCharacters,
		EscapePipe:       s.EscapePipe,
		IndentLevel:      s.IndentLevel,
		OutputValues:     s.OutputValues,
		ShowColor:        s.ShowColor,
		ShowHeader:       s.ShowHeader,
		ShowInputs:       s.ShowInputs,
		ShowOutputs:      s.ShowOutputs,
		ShowProviders:    s.ShowProviders,
		ShowRequired:     s.ShowRequired,
		ShowSensitivity:  s.ShowSensitivity,
		ShowRequirements: s.ShowRequirements,
		ShowResources:    s.ShowResources,
		SortByName:       s.SortByName,
		SortByRequired:   s.SortByRequired,
		SortByType:       s.SortByType,
	}
}
