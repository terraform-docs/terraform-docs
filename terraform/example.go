/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"unicode"

	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/print"
)

// Requirement represents a requirement for Terraform module.
type Example struct {
	Name    string       `json:"name" toml:"name" xml:"name" yaml:"name"`
	Content types.String `json:"content" toml:"content" xml:"content" yaml:"content"`
}

// Creating a struct that handles the actual getting of example content
// and allowing load.go to handle the logic of loading that into a
// module object
type ExampleLoader struct {
	config   *print.Config
	examples []*Example
}

func NewExampleLoader(config *print.Config) *ExampleLoader {
	loader := &ExampleLoader{
		config,
		[]*Example{},
	}
	return loader
}

func (loader *ExampleLoader) SearchFolder() error {
	path := filepath.Join(loader.config.ModuleRoot, loader.config.ExamplesFrom)
	if !FileExists(path) {
		return fmt.Errorf("specified examples folder '%s' does not exist", path)
	}
	possibleExamples := GetChildFolders(path)
	for _, f := range possibleExamples {
		if HasChildMainTf(f) {
			maintfPath := filepath.Join(f, "main.tf")
			name := filepath.Base(f)
			if len(loader.config.Examples.Include) > 0 && !contains(loader.config.Examples.Include, name) {
				continue
			}
			if len(loader.config.Examples.Exclude) > 0 && contains(loader.config.Examples.Exclude, name) {
				continue
			}
			buf, err := os.ReadFile(maintfPath)
			if err != nil {
				return err
			}
			content := string(buf)
			if !isOnlyWhitespace(content) {
				loader.examples = append(loader.examples, &Example{name, types.String(content)})
				if len(loader.examples) == loader.config.Examples.Limit { // if limit is reached,
					return nil
				}
			}
		}
	}
	return nil
}

func (loader *ExampleLoader) GetFile(relativePath string) error {
	path := filepath.Join(loader.config.ModuleRoot, relativePath)
	if !FileExists(path) {
		return fmt.Errorf("specified example file '%s' does not exist", path)
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)
	if !isOnlyWhitespace(content) {
		exampleName := filepath.Base(filepath.Dir(path))
		loader.examples = append(loader.examples, &Example{exampleName, types.String(content)})
	}
	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetChildFolders(path string) []string {
	var childFolders []string
	filepath.WalkDir(path, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			childFolders = append(childFolders, path)
		}
		return nil
	})

	return childFolders[1:]
}

func HasChildMainTf(path string) bool {
	var childFiles []string
	filepath.WalkDir(path, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !file.IsDir() {
			childFiles = append(childFiles, path)
		}
		return nil
	})

	for _, fileName := range childFiles {
		if filepath.Base(fileName) == "main.tf" {
			return true
		}
	}
	return false
}

func isOnlyWhitespace(input string) bool {
	for _, c := range input {
		if !unicode.IsSpace(c) { //if
			return false
		}
	}
	return true
}

func contains(set []string, str string) bool {
	for _, i := range set {
		if i == str {
			return true
		}
	}
	return false
}
