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

	"github.com/terraform-docs/terraform-docs/internal/print"
)

// initializerFn returns a concrete implementation of an Engine.
type initializerFn func(*print.Settings) print.Engine

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

// Factory initializes and returns the concrete implementation of
// format.Engine based on the provided 'name', for example for name
// of 'json' it will return '*format.JSON' through 'format.NewJSON'
// function.
func Factory(name string, settings *print.Settings) (print.Engine, error) {
	fn, ok := initializers[name]
	if !ok {
		return nil, fmt.Errorf("formatter '%s' not found", name)
	}
	return fn(settings), nil
}
