/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"embed"
	"fmt"
	"io/fs"
	"regexp"
	"strings"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/template"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// sanitize cleans a Markdown document to soothe linters.
func sanitize(markdown string) string {
	result := markdown

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(` {2}(\r?\n)`).ReplaceAllString(result, "‡‡‡DOUBLESPACES‡‡‡$1")

	// Remove trailing spaces from the end of lines
	result = regexp.MustCompile(` +(\r?\n)`).ReplaceAllString(result, "$1")
	result = regexp.MustCompile(` +$`).ReplaceAllLiteralString(result, "")

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(`‡‡‡DOUBLESPACES‡‡‡(\r?\n)`).ReplaceAllString(result, "  $1")

	// Remove blank line with only double spaces in it
	result = regexp.MustCompile(`(\r?\n)  (\r?\n)`).ReplaceAllString(result, "$1")

	// Remove multiple consecutive blank lines
	result = regexp.MustCompile(`(\r?\n){3,}`).ReplaceAllString(result, "$1$1")
	result = regexp.MustCompile(`(\r?\n){2,}$`).ReplaceAllString(result, "")

	return result
}

// PrintFencedCodeBlock prints codes in fences, it automatically detects if
// the input 'code' contains '\n' it will use multi line fence, otherwise it
// wraps the 'code' inside single-tick block.
// If the fenced is multi-line it also appens an extra '\n` at the end and
// returns true accordingly, otherwise returns false for non-carriage return.
func PrintFencedCodeBlock(code string, language string) (string, bool) {
	if strings.Contains(code, "\n") {
		return fmt.Sprintf("\n\n```%s\n%s\n```\n", language, code), true
	}
	return fmt.Sprintf("`%s`", code), false
}

// PrintFencedAsciidocCodeBlock prints codes in fences, it automatically detects if
// the input 'code' contains '\n' it will use multi line fence, otherwise it
// wraps the 'code' inside single-tick block.
// If the fenced is multi-line it also appens an extra '\n` at the end and
// returns true accordingly, otherwise returns false for non-carriage return.
func PrintFencedAsciidocCodeBlock(code string, language string) (string, bool) {
	if strings.Contains(code, "\n") {
		return fmt.Sprintf("\n[source,%s]\n----\n%s\n----\n", language, code), true
	}
	return fmt.Sprintf("`%s`", code), false
}

// readTemplateItems reads all static formatter .tmpl files prefixed by specific string
// from an embed file system.
func readTemplateItems(efs embed.FS, prefix string) []*template.Item {
	items := make([]*template.Item, 0)

	files, err := fs.ReadDir(efs, "templates")
	if err != nil {
		return items
	}

	for _, f := range files {
		content, err := efs.ReadFile("templates/" + f.Name())
		if err != nil {
			continue
		}

		name := f.Name()
		name = strings.ReplaceAll(name, prefix, "")
		name = strings.ReplaceAll(name, "_", "")
		name = strings.ReplaceAll(name, ".tmpl", "")
		if name == "" {
			name = "all"
		}

		items = append(items, &template.Item{
			Name: name,
			Text: string(content),
		})
	}
	return items
}

// copySections sets the sections that'll be printed
func copySections(settings *print.Settings, src *terraform.Module) *terraform.Module {
	dest := &terraform.Module{
		Header:       "",
		Footer:       "",
		Inputs:       make([]*terraform.Input, 0),
		ModuleCalls:  make([]*terraform.ModuleCall, 0),
		Outputs:      make([]*terraform.Output, 0),
		Providers:    make([]*terraform.Provider, 0),
		Requirements: make([]*terraform.Requirement, 0),
		Resources:    make([]*terraform.Resource, 0),
	}

	if settings.ShowHeader {
		dest.Header = src.Header
	}
	if settings.ShowFooter {
		dest.Footer = src.Footer
	}
	if settings.ShowInputs {
		dest.Inputs = src.Inputs
	}
	if settings.ShowModuleCalls {
		dest.ModuleCalls = src.ModuleCalls
	}
	if settings.ShowOutputs {
		dest.Outputs = src.Outputs
	}
	if settings.ShowProviders {
		dest.Providers = src.Providers
	}
	if settings.ShowRequirements {
		dest.Requirements = src.Requirements
	}
	if settings.ShowResources || settings.ShowDataSources {
		dest.Resources = filterResourcesByMode(settings, src.Resources)
	}

	return dest
}

// filterResourcesByMode returns the managed or data resources defined by the show argument
func filterResourcesByMode(settings *print.Settings, module []*terraform.Resource) []*terraform.Resource {
	resources := make([]*terraform.Resource, 0)
	for _, r := range module {
		if settings.ShowResources && r.Mode == "managed" {
			resources = append(resources, r)
		}
		if settings.ShowDataSources && r.Mode == "data" {
			resources = append(resources, r)
		}
	}
	return resources
}
