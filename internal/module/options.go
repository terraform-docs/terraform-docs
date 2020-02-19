package module

import (
	"log"

	"github.com/imdario/mergo"
)

// SortBy contains different sort criteria corresponding
// to available flags (e.g. name, required, etc)
type SortBy struct {
	Name     bool
	Required bool
}

// Options contains required options to load a Module from path
type Options struct {
	Path             string
	SortBy           *SortBy
	OutputValues     bool
	OutputValuesPath string
}

// NewOptions returns new instance of Options
func NewOptions() *Options {
	return &Options{
		Path:             "",
		SortBy:           &SortBy{Name: false, Required: false},
		OutputValues:     false,
		OutputValuesPath: "",
	}
}

// With override options with existing Options
func (o *Options) With(override *Options) *Options {
	if err := mergo.Merge(o, override); err != nil {
		log.Fatal(err)
	}
	return o
}
