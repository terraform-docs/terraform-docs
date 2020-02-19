package tfconf

// Module represents a Terraform module. It consists of
// - Header    ('header' json key):    Module header found in shape of multi line comments at the beginning of 'main.tf'
// - Inputs    ('inputs' json key):    List of input 'variables' extracted from the Terraform module .tf files
// - Outputs   ('outputs' json key):   List of 'outputs' extracted from Terraform module .tf files
// - Providers ('providers' json key): List of 'providers' extracted from resources used in Terraform module
type Module struct {
	Header    string      `json:"header" yaml:"header"`
	Inputs    []*Input    `json:"inputs" yaml:"inputs"`
	Outputs   []*Output   `json:"outputs" yaml:"outputs"`
	Providers []*Provider `json:"providers" yaml:"providers"`

	RequiredInputs []*Input `json:"-" yaml:"-"`
	OptionalInputs []*Input `json:"-" yaml:"-"`
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
