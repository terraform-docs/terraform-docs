package terraform

import (
	"errors"

	"github.com/imdario/mergo"
)

// SortBy contains different sort criteria corresponding
// to available flags (e.g. name, required, etc)
type SortBy struct {
	Name     bool
	Required bool
	Type     bool
}

// Options contains required options to load a Module from path
type Options struct {
	Path             string
	ShowHeader       bool
	HeaderFromFile   string
	SortBy           *SortBy
	OutputValues     bool
	OutputValuesPath string
}

// NewOptions returns new instance of Options
func NewOptions() *Options {
	return &Options{
		Path:             "",
		ShowHeader:       true,
		HeaderFromFile:   "main.tf",
		SortBy:           &SortBy{Name: false, Required: false, Type: false},
		OutputValues:     false,
		OutputValuesPath: "",
	}
}

// With override options with existing Options
func (o *Options) With(override *Options) (*Options, error) {
	if override == nil {
		return nil, errors.New("cannot use nil as override value")
	}
	if err := mergo.Merge(o, *override); err != nil {
		return nil, err
	}
	return o, nil
}

// WithOverwrite override options with existing Options and overwrites non-empty
// items in destination
func (o *Options) WithOverwrite(override *Options) (*Options, error) {
	if override == nil {
		return nil, errors.New("cannot use nil as override value")
	}
	if err := mergo.MergeWithOverwrite(o, *override); err != nil {
		return nil, err
	}
	return o, nil
}
