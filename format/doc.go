/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

// Package format provides different, out of the box supported, output formats.
//
// Usage
//
// A specific format can be instantiated either for `format.Factory()` function or
// directly calling its function (e.g. `NewMarkdownTable`, etc)
//
//     formatter, err := format.Factory("markdown table", settings)
//     if err != nil {
//         return err
//     }
//
//     generator, err := formatter.Generate(tfmodule)
//     if err != nil {
//         return err
//     }
//
//     output, err := generator.ExecuteTemplate("")
//     if err != nil {
//         return err
//     }
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
//
package format
