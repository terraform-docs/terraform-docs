/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

// Package print provides configuration, and a Generator.
//
// # Configuration
//
// `print.Config` is the data structure representation for `.terraform-docs.yml`
// which will be read and extracted upon execution of terraform-docs cli. On the
// other hand it can be used directly if you are using terraform-docs as a library.
//
// This will return an instance of `Config` with default values set:
//
//	config := print.DefaultConfig()
//
// Alternatively this will return an empty instance of `Config`:
//
//	config := print.NewConfig()
//
// # Generator
//
// `Generator` is an abstract implementation of `format.Type`. It doesn't implement
// `Generate(*terraform.Module) error` function. It is used directly by different
// format types, i.e. each format extends `Generator` and provides its implementation
// of `Generate` function.
//
// Generator holds a reference to all the sections (e.g. header, footer, inputs, etc)
// and also it renders all of them, in a predefined order, in `Content()`.
//
// It also provides `Render(string)` function to process and render the template to generate
// the final output content. Following variables and functions are available:
//
// • `{{ .Header }}`
// • `{{ .Footer }}`
// • `{{ .Inputs }}`
// • `{{ .Modules }}`
// • `{{ .Outputs }}`
// • `{{ .Providers }}`
// • `{{ .Requirements }}`
// • `{{ .Resources }}`
// • `{{ include "path/fo/file" }}`
package print
