/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"encoding/xml"
)

// Module represents a Terraform module. It consists of
type Module struct {
	XMLName xml.Name `json:"-" toml:"-" xml:"module" yaml:"-"`

	Header       string         `json:"header" toml:"header" xml:"header" yaml:"header"`
	Footer       string         `json:"footer" toml:"footer" xml:"footer" yaml:"footer"`
	Inputs       []*Input       `json:"inputs" toml:"inputs" xml:"inputs>input" yaml:"inputs"`
	ModuleCalls  []*ModuleCall  `json:"modules" toml:"modules" xml:"modules>module" yaml:"modules"`
	Outputs      []*Output      `json:"outputs" toml:"outputs" xml:"outputs>output" yaml:"outputs"`
	Providers    []*Provider    `json:"providers" toml:"providers" xml:"providers>provider" yaml:"providers"`
	Requirements []*Requirement `json:"requirements" toml:"requirements" xml:"requirements>requirement" yaml:"requirements"`
	Resources    []*Resource    `json:"resources" toml:"resources" xml:"resources>resource" yaml:"resources"`

	RequiredInputs []*Input `json:"-" toml:"-" xml:"-" yaml:"-"`
	OptionalInputs []*Input `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// HasHeader indicates if the module has header.
func (m *Module) HasHeader() bool {
	return len(m.Header) > 0
}

// HasFooter indicates if the module has footer.
func (m *Module) HasFooter() bool {
	return len(m.Footer) > 0
}

// HasInputs indicates if the module has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Inputs) > 0
}

// HasModuleCalls indicates if the module has modulecalls.
func (m *Module) HasModuleCalls() bool {
	return len(m.ModuleCalls) > 0
}

// HasOutputs indicates if the module has outputs.
func (m *Module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

// HasProviders indicates if the module has providers.
func (m *Module) HasProviders() bool {
	return len(m.Providers) > 0
}

// HasRequirements indicates if the module has requirements.
func (m *Module) HasRequirements() bool {
	return len(m.Requirements) > 0
}

// HasResources indicates if the module has resources.
func (m *Module) HasResources() bool {
	return len(m.Resources) > 0
}
