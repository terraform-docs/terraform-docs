/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package testutil

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// GetModule returns 'example' Module
func GetModule(config *print.Config) (*terraform.Module, error) {
	path, err := getExampleFolder(config.ModuleRoot)
	if err != nil {
		return nil, err
	}
	config.ModuleRoot = path

	if config.OutputValues.Enabled {
		config.OutputValues.From = filepath.Join(path, config.OutputValues.From)
	}

	tfmodule, err := terraform.LoadWithOptions(config)
	if err != nil {
		return nil, err
	}

	return tfmodule, nil
}

// GetExpected returns 'example' Module and expected Golden file content
func GetExpected(format, name string) (string, error) {
	path := filepath.Join(testDataPath(), format, name+".golden")
	bytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getExampleFolder(folder string) (string, error) {
	_, b, _, _ := runtime.Caller(0)
	var path string
	if folder != "" {
		path = filepath.Join(filepath.Dir(b), "..", "testutil", "testdata", folder)
	} else {
		path = filepath.Join(filepath.Dir(b), "..", "..", "examples")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}

func testDataPath() string {
	return "testdata"
}
