/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"testing"
)

func TestSectionsValidate(t *testing.T) {
	type fields struct {
		Show         []string
		Hide         []string
		ShowAll      bool
		HideAll      bool
		header       bool
		inputs       bool
		modulecalls  bool
		outputs      bool
		providers    bool
		requirements bool
		resources    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "show all and hide all at once",
			fields: fields{
				ShowAll: true,
				HideAll: true,
			},
			wantErr: true,
		},
		{
			name: "show unknown section",
			fields: fields{
				Show: []string{"test"},
			},
			wantErr: true,
		},
		{
			name: "hide unknown section",
			fields: fields{
				Hide: []string{"test"},
			},
			wantErr: true,
		},
		{
			name: "show allowed sections section",
			fields: fields{
				Show: []string{"header", "inputs", "modules", "outputs", "providers", "requirements", "resources"},
			},
			wantErr: false,
		},
		{
			name: "hide allowed sections section",
			fields: fields{
				Hide: []string{"header", "inputs", "modules", "outputs", "providers", "requirements", "resources"},
			},
			wantErr: false,
		},
		{
			name: "show-all and explicit show",
			fields: fields{
				ShowAll: true,
				Show:    []string{"header"},
			},
			wantErr: true,
		},
		{
			name: "show-all and explicit hide",
			fields: fields{
				HideAll: true,
				Hide:    []string{"header"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sections{
				Show:         tt.fields.Show,
				Hide:         tt.fields.Hide,
				ShowAll:      tt.fields.ShowAll,
				HideAll:      tt.fields.HideAll,
				header:       tt.fields.header,
				inputs:       tt.fields.inputs,
				modulecalls:  tt.fields.modulecalls,
				outputs:      tt.fields.outputs,
				providers:    tt.fields.providers,
				requirements: tt.fields.requirements,
				resources:    tt.fields.resources,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("sections.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSettingsValidate(t *testing.T) {
	type fields struct {
		Color     bool
		Escape    bool
		Indent    int
		Required  bool
		Sensitive bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &settings{
				Color:     tt.fields.Color,
				Escape:    tt.fields.Escape,
				Indent:    tt.fields.Indent,
				Required:  tt.fields.Required,
				Sensitive: tt.fields.Sensitive,
			}
			if err := s.validate(); (err != nil) != tt.wantErr {
				t.Errorf("settings.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigProcess(t *testing.T) {
	type fields struct {
		File         string
		Formatter    string
		HeaderFrom   string
		Sections     sections
		OutputValues outputvalues
		Sort         sort
		Settings     settings
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "show all sections",
			fields: fields{
				Sections: sections{
					ShowAll:      true,
					header:       true,
					inputs:       true,
					modulecalls:  true,
					outputs:      true,
					providers:    true,
					requirements: true,
					resources:    true,
				},
			},
		},
		{
			name: "hide all",
			fields: fields{
				Sections: sections{
					HideAll:      true,
					header:       false,
					inputs:       false,
					modulecalls:  false,
					outputs:      false,
					providers:    false,
					requirements: false,
					resources:    false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				File:         tt.fields.File,
				Formatter:    tt.fields.Formatter,
				HeaderFrom:   tt.fields.HeaderFrom,
				Sections:     tt.fields.Sections,
				OutputValues: tt.fields.OutputValues,
				Sort:         tt.fields.Sort,
				Settings:     tt.fields.Settings,
			}
			c.process()
			if c.Sections.header != tt.fields.Sections.header {
				t.Errorf("Config.process() header = %v, should be %v", c.Sections.header, tt.fields.Sections.header)
			}
			if c.Sections.inputs != tt.fields.Sections.inputs {
				t.Errorf("Config.process() inputs = %v, should be %v", c.Sections.inputs, tt.fields.Sections.inputs)
			}
			if c.Sections.modulecalls != tt.fields.Sections.modulecalls {
				t.Errorf("Config.process() modulecalls = %v, should be %v", c.Sections.modulecalls, tt.fields.Sections.modulecalls)
			}
			if c.Sections.outputs != tt.fields.Sections.outputs {
				t.Errorf("Config.process() outputs = %v, should be %v", c.Sections.outputs, tt.fields.Sections.outputs)
			}
			if c.Sections.providers != tt.fields.Sections.providers {
				t.Errorf("Config.process() providers = %v, should be %v", c.Sections.providers, tt.fields.Sections.providers)
			}
			if c.Sections.requirements != tt.fields.Sections.requirements {
				t.Errorf("Config.process() requirements = %v, should be %v", c.Sections.requirements, tt.fields.Sections.requirements)
			}
			if c.Sections.resources != tt.fields.Sections.resources {
				t.Errorf("Config.process() resources = %v, should be %v", c.Sections.resources, tt.fields.Sections.resources)
			}
		})
	}
}

func TestConfigValidate(t *testing.T) {
	type fields struct {
		File         string
		Formatter    string
		HeaderFrom   string
		Sections     sections
		OutputValues outputvalues
		Sort         sort
		Settings     settings
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "known good config",
			fields: fields{
				Formatter:    "md",
				HeaderFrom:   "test.md",
				Sections:     defaultSections(),
				OutputValues: defaultOutputValues(),
				Sort:         defaultSort(),
				Settings:     defaultSettings(),
			},
			wantErr: false,
		},
		{
			name: "empty formatter",
			fields: fields{
				Formatter:    "",
				HeaderFrom:   "test.md",
				Sections:     defaultSections(),
				OutputValues: defaultOutputValues(),
				Sort:         defaultSort(),
				Settings:     defaultSettings(),
			},
			wantErr: true,
		},
		{
			name: "empty HeaderFrom",
			fields: fields{
				Formatter:    "md",
				HeaderFrom:   "",
				Sections:     defaultSections(),
				OutputValues: defaultOutputValues(),
				Sort:         defaultSort(),
				Settings:     defaultSettings(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				File:         tt.fields.File,
				Formatter:    tt.fields.Formatter,
				HeaderFrom:   tt.fields.HeaderFrom,
				Sections:     tt.fields.Sections,
				OutputValues: tt.fields.OutputValues,
				Sort:         tt.fields.Sort,
				Settings:     tt.fields.Settings,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
