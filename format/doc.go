/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

// Package format provides different, out of the box supported, output format types.
//
// # Usage
//
// A specific format can be instantiated either with `format.New()` function or
// directly calling its function (e.g. `NewMarkdownTable`, etc)
//
//	config := print.DefaultConfig()
//	config.Formatter = "markdown table"
//
//	formatter, err := format.New(config)
//	if err != nil {
//	    return err
//	}
//
//	err := formatter.Generate(tfmodule)
//	if err != nil {
//	    return err
//	}
//
//	output, err := formatter.Render"")
//	if err != nil {
//	    return err
//	}
//
// Note: if you don't intend to provide additional template for the generated
// content, or the target format doesn't provide templating (e.g. json, yaml,
// xml, or toml) you can use `Content()` function instead of `Render)`. Note
// that `Content()` returns all the sections combined with predefined order.
//
//	output := formatter.Content()
//
// Supported formats are:
//
// • `NewAsciidocDocument`
// • `NewAsciidocTable`
// • `NewJSON`
// • `NewMarkdownDocument`
// • `NewMarkdownTable`
// • `NewPretty`
// • `NewTfvarsHCL`
// • `NewTfvarsJSON`
// • `NewTOML`
// • `NewXML`
// • `NewYAML`
package format
