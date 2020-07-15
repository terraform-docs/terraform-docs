package testutil

import (
	"github.com/imdario/mergo"

	"github.com/terraform-docs/terraform-docs/pkg/print"
)

// TestSettings respresents the Settings instance for tests
type TestSettings struct {
	full *print.Settings
}

// Settings returns TestSettings instance with predefined set of print.Settings
func Settings() *TestSettings {
	shared := &print.Settings{
		EscapePipe: true,
	}
	return &TestSettings{
		full: shared,
	}
}

// With appends provided 'override' print.Settings to TestSettings
func (s *TestSettings) With(override *print.Settings) *TestSettings {
	if err := mergo.Merge(override, s.full); err == nil {
		s.full = override
	}
	return s
}

// WithColor appends predefined 'ShowColor: true' to TestSettings
func (s *TestSettings) WithColor() *TestSettings {
	color := &print.Settings{
		ShowColor: true,
	}
	if err := mergo.Merge(color, s.full); err == nil {
		s.full = color
	}
	return s
}

// WithSections appends predefined show all sections ShowHeader, ShowProviders, ShowInputs, ShowOutputs to TestSettings
func (s *TestSettings) WithSections() *TestSettings {
	sections := &print.Settings{
		ShowHeader:       true,
		ShowInputs:       true,
		ShowOutputs:      true,
		ShowProviders:    true,
		ShowRequirements: true,
	}
	if err := mergo.Merge(sections, s.full); err == nil {
		s.full = sections
	}
	return s
}

// Build builds and returns print.Settings based on the provided overrides in TestSettings
func (s *TestSettings) Build() *print.Settings {
	return s.full
}
