package tfconf

import (
	"encoding/xml"
)

// Module represents a Terraform module. It consists of
//
// - Header       ('header' json key):    Module header found in shape of multi line comments at the beginning of 'main.tf'
// - Inputs       ('inputs' json key):    List of input 'variables' extracted from the Terraform module .tf files
// - Outputs      ('outputs' json key):   List of 'outputs' extracted from Terraform module .tf files
// - Providers    ('providers' json key): List of 'providers' extracted from resources used in Terraform module
// - Requirements ('header' json key):    List of 'requirements' extracted from the Terraform module .tf files
type Module struct {
	XMLName xml.Name `json:"-" xml:"module" yaml:"-"`

	Header       string         `json:"header" xml:"header" yaml:"header"`
	Inputs       []*Input       `json:"inputs" xml:"inputs>input" yaml:"inputs"`
	Outputs      []*Output      `json:"outputs" xml:"outputs>output" yaml:"outputs"`
	Providers    []*Provider    `json:"providers" xml:"providers>provider" yaml:"providers"`
	Requirements []*Requirement `json:"requirements" xml:"requirements>requirement" yaml:"requirements"`

	RequiredInputs []*Input `json:"-" xml:"-" yaml:"-"`
	OptionalInputs []*Input `json:"-" xml:"-" yaml:"-"`
}

// HasHeader indicates if the module has header.
func (m *Module) HasHeader() bool {
	return len(m.Header) > 0
}

// HasInputs indicates if the module has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Inputs) > 0
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
