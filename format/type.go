/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"fmt"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// Type represents an output format type (e.g. json, markdown table, yaml, etc).
type Type interface {
	Generate(*terraform.Module) error // generate the Terraform module

	Content() string // all the sections combined based on the underlying format

	Header() string       // header section based on the underlying format
	Footer() string       // footer section based on the underlying format
	Inputs() string       // inputs section based on the underlying format
	Modules() string      // modules section based on the underlying format
	Outputs() string      // outputs section based on the underlying format
	Providers() string    // providers section based on the underlying format
	Requirements() string // requirements section based on the underlying format
	Resources() string    // resources section based on the underlying format

	Render(tmpl string) (string, error)
}

// initializerFn returns a concrete implementation of an Engine.
type initializerFn func(*print.Config) Type

// initializers list of all registered engine initializer functions.
var initializers = make(map[string]initializerFn)

// register a formatter engine initializer function.
func register(e map[string]initializerFn) {
	if e == nil {
		return
	}
	for k, v := range e {
		initializers[k] = v
	}
}

// New initializes and returns the concrete implementation of
// format.Engine based on the provided 'name', for example for name
// of 'json' it will return '*format.JSON' through 'format.NewJSON'
// function.
func New(config *print.Config) (Type, error) {
	name := config.Formatter
	fn, ok := initializers[name]
	if !ok {
		return nil, fmt.Errorf("formatter '%s' not found", name)
	}
	return fn(config), nil
}
