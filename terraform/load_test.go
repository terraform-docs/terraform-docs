/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/print"
)

func TestLoadModuleWithOptions(t *testing.T) {
	assert := assert.New(t)

	config := print.NewConfig()
	config.ModuleRoot = filepath.Join("testdata", "full-example")
	config.Sections.Header = true

	module, err := LoadWithOptions(config)

	assert.Nil(err)
	assert.Equal(true, module.HasHeader())
	assert.Equal(false, module.HasFooter())
	assert.Equal(true, module.HasInputs())
	assert.Equal(true, module.HasOutputs())
	assert.Equal(true, module.HasModuleCalls())
	assert.Equal(true, module.HasProviders())
	assert.Equal(true, module.HasRequirements())

	config.Sections.Header = false
	config.Sections.Footer = true
	config.FooterFrom = "doc.tf"

	module, err = LoadWithOptions(config)
	assert.Nil(err)
	assert.Equal(true, module.HasFooter())
	assert.Equal(false, module.HasHeader())
}

func TestLoadModule(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "load module from path",
			path:    "full-example",
			wantErr: false,
		},
		{
			name:    "load module from path",
			path:    "non-exist",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			_, err := loadModule(filepath.Join("testdata", tt.path))
			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
			}
		})
	}
}

func TestGetFileFormat(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected string
	}{
		{
			name:     "get file format",
			filename: "main.tf",
			expected: ".tf",
		},
		{
			name:     "get file format",
			filename: "main.file.tf",
			expected: ".tf",
		},
		{
			name:     "get file format",
			filename: "main_file.tf",
			expected: ".tf",
		},
		{
			name:     "get file format",
			filename: "main.file_tf",
			expected: ".file_tf",
		},
		{
			name:     "get file format",
			filename: "main_file_tf",
			expected: "",
		},
		{
			name:     "get file format",
			filename: "",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := getFileFormat(tt.filename)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestIsFileFormatSupported(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
		wantErr  bool
		errText  string
		section  string
	}{
		{
			name:     "is file format supported",
			filename: "main.adoc",
			expected: true,
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "is file format supported",
			filename: "main.md",
			expected: true,
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "is file format supported",
			filename: "main.tf",
			expected: true,
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "is file format supported",
			filename: "main.txt",
			expected: true,
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "is file format supported",
			filename: "main.doc",
			expected: false,
			wantErr:  true,
			errText:  "only .adoc, .md, .tf, and .txt formats are supported to read header from",
			section:  "header",
		},
		{
			name:     "is file format supported",
			filename: "",
			expected: false,
			wantErr:  true,
			errText:  "--header-from value is missing",
			section:  "header",
		}, {
			name:     "err message changes for footer",
			filename: "main.doc",
			expected: false,
			wantErr:  true,
			errText:  "only .adoc, .md, .tf, and .txt formats are supported to read footer from",
			section:  "footer",
		},
		{
			name:     "err message changes for footer",
			filename: "",
			expected: false,
			wantErr:  true,
			errText:  "--footer-from value is missing",
			section:  "footer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := isFileFormatSupported(tt.filename, tt.section)
			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errText, err.Error())
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, actual)
			}
		})
	}
}

func TestLoadHeader(t *testing.T) {
	tests := []struct {
		name         string
		testData     string
		showHeader   bool
		expectedData func() (string, error)
	}{
		{
			name:       "loadHeader should return a string from file",
			testData:   "full-example",
			showHeader: true,
			expectedData: func() (string, error) {
				path := filepath.Join("testdata", "expected", "full-example-mainTf-Header.golden")
				data, err := os.ReadFile(path)
				return string(data), err
			},
		},
		{
			name:       "loadHeader should return an empty string if not shown",
			testData:   "",
			showHeader: false,
			expectedData: func() (string, error) {
				return "", nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.ModuleRoot = filepath.Join("testdata", tt.testData)
			config.Sections.Header = tt.showHeader

			expected, err := tt.expectedData()
			assert.Nil(err)

			header, err := loadHeader(config)
			assert.Nil(err)
			assert.Equal(expected, header)
		})
	}
}

func TestLoadFooter(t *testing.T) {
	tests := []struct {
		name         string
		testData     string
		footerFile   string
		showFooter   bool
		expectedData func() (string, error)
	}{
		{
			name:       "loadFooter should return a string from file",
			testData:   "full-example",
			footerFile: "main.tf",
			showFooter: true,
			expectedData: func() (string, error) {
				path := filepath.Join("testdata", "expected", "full-example-mainTf-Header.golden")
				data, err := os.ReadFile(path)
				return string(data), err
			},
		},
		{
			name:       "loadHeader should return an empty string if not shown",
			testData:   "",
			footerFile: "",
			showFooter: false,
			expectedData: func() (string, error) {
				return "", nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.ModuleRoot = filepath.Join("testdata", tt.testData)
			config.Sections.Footer = tt.showFooter
			config.FooterFrom = tt.footerFile

			expected, err := tt.expectedData()
			assert.Nil(err)

			header, err := loadFooter(config)
			assert.Nil(err)
			assert.Equal(expected, header)
		})
	}
}

func TestLoadSections(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		file     string
		expected string
		wantErr  bool
		errText  string
		section  string
	}{
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "main.tf",
			expected: "Example of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n\nEven inline **formatting** in _here_ is possible.\nand some [link](https://domain.com/)",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "doc.tf",
			expected: "Custom Header:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "doc.md",
			expected: "# Custom Header\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "doc.adoc",
			expected: "= Custom Header\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "doc.txt",
			expected: "# Custom Header\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "no-inputs",
			file:     "main.tf",
			expected: "",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "non-existent.tf",
			expected: "",
			wantErr:  true,
			errText:  "stat testdata/full-example/non-existent.tf: no such file or directory",
			section:  "header",
		},
		{
			name:     "no error if header file is missing and is default 'main.tf'",
			path:     "inputs-lf",
			file:     "main.tf",
			expected: "",
			wantErr:  false,
			errText:  "",
			section:  "header",
		},
		{
			name:     "error if footer file is missing even if 'main.tf'",
			path:     "inputs-lf",
			file:     "main.tf",
			expected: "",
			wantErr:  true,
			errText:  "stat testdata/inputs-lf/main.tf: no such file or directory",
			section:  "footer",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "wrong-formate.docx",
			expected: "",
			wantErr:  true,
			errText:  "only .adoc, .md, .tf, and .txt formats are supported to read footer from",
			section:  "footer",
		},
		{
			name:     "load module header from path",
			path:     "full-example",
			file:     "",
			expected: "",
			wantErr:  true,
			errText:  "--header-from value is missing",
			section:  "header",
		},
		{
			name:     "load module header from path",
			path:     "empty-header",
			file:     "",
			expected: "",
			wantErr:  true,
			errText:  "--header-from value is missing",
			section:  "header",
		},
		{
			name:     "load module footer from path",
			path:     "non-exist",
			file:     "",
			expected: "",
			wantErr:  true,
			errText:  "--footer-from value is missing",
			section:  "footer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.ModuleRoot = filepath.Join("testdata", tt.path)

			actual, err := loadSection(config, tt.file, tt.section)
			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errText, err.Error())
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, actual)
			}
		})
	}
}

func TestLoadInputs(t *testing.T) {
	type expected struct {
		inputs    int
		requireds int
		optionals int
	}
	tests := []struct {
		name     string
		path     string
		expected expected
	}{
		{
			name: "load module inputs from path",
			path: "full-example",
			expected: expected{
				inputs:    7,
				requireds: 2,
				optionals: 5,
			},
		},
		{
			name: "load module inputs from path",
			path: "no-required-inputs",
			expected: expected{
				inputs:    6,
				requireds: 0,
				optionals: 6,
			},
		},
		{
			name: "load module inputs from path",
			path: "no-optional-inputs",
			expected: expected{
				inputs:    6,
				requireds: 6,
				optionals: 0,
			},
		},
		{
			name: "load module inputs from path",
			path: "no-inputs",
			expected: expected{
				inputs:    0,
				requireds: 0,
				optionals: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			module, _ := loadModule(filepath.Join("testdata", tt.path))
			inputs, requireds, optionals := loadInputs(module, config)

			assert.Equal(tt.expected.inputs, len(inputs))
			assert.Equal(tt.expected.requireds, len(requireds))
			assert.Equal(tt.expected.optionals, len(optionals))
		})
	}
}

func TestLoadModulecalls(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "load modulecalls from path",
			path:     "full-example",
			expected: 2,
		},
		{
			name:     "load modulecalls from path",
			path:     "no-modulecalls",
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			module, _ := loadModule(filepath.Join("testdata", tt.path))
			modulecalls := loadModulecalls(module, config)

			assert.Equal(tt.expected, len(modulecalls))
		})
	}
}

func TestLoadInputsLineEnding(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "load module inputs from file with lf line ending",
			path:     "inputs-lf",
			expected: "The quick brown fox jumps\nover the lazy dog\n",
		},
		{
			name:     "load module inputs from file with crlf line ending",
			path:     "inputs-crlf",
			expected: "The quick brown fox jumps\nover the lazy dog\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			module, _ := loadModule(filepath.Join("testdata", tt.path))
			inputs, _, _ := loadInputs(module, config)

			assert.Equal(1, len(inputs))
			assert.Equal(tt.expected, string(inputs[0].Description))
		})
	}
}

func TestLoadOutputs(t *testing.T) {
	type expected struct {
		outputs int
	}
	tests := []struct {
		name     string
		path     string
		expected expected
	}{
		{
			name: "load module outputs from path",
			path: "full-example",
			expected: expected{
				outputs: 3,
			},
		},
		{
			name: "load module outputs from path",
			path: "no-outputs",
			expected: expected{
				outputs: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			module, _ := loadModule(filepath.Join("testdata", tt.path))
			outputs, err := loadOutputs(module, config)

			assert.Nil(err)
			assert.Equal(tt.expected.outputs, len(outputs))

			for _, v := range outputs {
				assert.Equal(false, v.ShowValue)
			}
		})
	}
}

func TestLoadOutputsLineEnding(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "load module outputs from file with lf line ending",
			path:     "outputs-lf",
			expected: "The quick brown fox jumps\nover the lazy dog\n",
		},
		{
			name:     "load module outputs from file with crlf line ending",
			path:     "outputs-crlf",
			expected: "The quick brown fox jumps\nover the lazy dog\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			module, _ := loadModule(filepath.Join("testdata", tt.path))
			outputs, _ := loadOutputs(module, config)

			assert.Equal(1, len(outputs))
			assert.Equal(tt.expected, string(outputs[0].Description))
		})
	}
}

func TestLoadOutputsValues(t *testing.T) {
	type expected struct {
		outputs int
	}
	tests := []struct {
		name       string
		path       string
		outputPath string
		expected   expected
		wantErr    bool
	}{
		{
			name:       "load module outputs with values from path",
			path:       "full-example",
			outputPath: "output-values.json",
			expected: expected{
				outputs: 3,
			},
			wantErr: false,
		},
		{
			name:       "load module outputs with values from path",
			path:       "full-example",
			outputPath: "no-file.json",
			expected:   expected{},
			wantErr:    true,
		},
		{
			name:       "load module outputs with values from path",
			path:       "no-outputs",
			outputPath: "no-file.json",
			expected:   expected{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.OutputValues.Enabled = true
			config.OutputValues.From = filepath.Join("testdata", tt.path, tt.outputPath)

			module, _ := loadModule(filepath.Join("testdata", tt.path))
			outputs, err := loadOutputs(module, config)

			if tt.wantErr {
				assert.NotNil(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected.outputs, len(outputs))

				for _, v := range outputs {
					assert.Equal(true, v.ShowValue)
				}
			}
		})
	}
}

func TestLoadProviders(t *testing.T) {
	type expected struct {
		providers []string
	}
	tests := []struct {
		name     string
		path     string
		lockfile bool
		expected expected
	}{
		{
			name:     "load module providers from path",
			path:     "full-example",
			lockfile: false,
			expected: expected{
				providers: []string{"aws->= 2.15.0", "null-", "tls-"},
			},
		},
		{
			name:     "load module providers from path",
			path:     "with-lock-file",
			lockfile: true,
			expected: expected{
				providers: []string{"aws-3.42.0", "null-3.1.0", "tls-3.1.0"},
			},
		},
		{
			name:     "load module providers from path",
			path:     "with-lock-file",
			lockfile: false,
			expected: expected{
				providers: []string{"aws->= 2.15.0", "null-", "tls-"},
			},
		},
		{
			name: "load module providers from path",
			path: "no-providers",
			expected: expected{
				providers: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.ModuleRoot = filepath.Join("testdata", tt.path)
			config.Settings.LockFile = tt.lockfile

			module, _ := loadModule(filepath.Join("testdata", tt.path))
			providers := loadProviders(module, config)

			actual := []string{}

			for _, p := range providers {
				actual = append(actual, p.FullName()+"-"+string(p.Version))
			}
			sort.Strings(actual)

			assert.Equal(tt.expected.providers, actual)
		})
	}
}

func TestLoadComments(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		fileName   string
		lineNumber int
		expected   string
	}{
		{
			name:       "load resource comment from file",
			path:       "full-example",
			fileName:   "variables.tf",
			lineNumber: 2,
			expected:   "D description",
		},
		{
			name:       "load resource comment from file",
			path:       "full-example",
			fileName:   "variables.tf",
			lineNumber: 16,
			expected:   "A Description in multiple lines",
		},
		{
			name:       "load resource comment from file with wrong line number",
			path:       "full-example",
			fileName:   "variables.tf",
			lineNumber: 100,
			expected:   "",
		},
		{
			name:       "load resource comment from non existing file",
			path:       "full-example",
			fileName:   "non-exist.tf",
			lineNumber: 5,
			expected:   "",
		},
		{
			name:       "load resource comment from non existing file",
			path:       "non-exist",
			fileName:   "variables.tf",
			lineNumber: 5,
			expected:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := loadComments(filepath.Join("testdata", tt.path, tt.fileName), tt.lineNumber)
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestReadComments(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		fileName     string
		readComments bool
		expected     string
	}{
		{
			name:         "Validate description when 'ReadComments' is false",
			path:         "read-comments",
			fileName:     "variables.tf",
			readComments: false,
			expected:     "",
		},
		{
			name:         "Validate description when 'ReadComments' is true",
			path:         "read-comments",
			fileName:     "variables.tf",
			readComments: true,
			expected:     "B description",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			config := print.NewConfig()
			config.Settings.ReadComments = tt.readComments

			module, err := loadModule(filepath.Join("testdata", tt.path))

			assert.Nil(err)

			inputs, _, _ := loadInputs(module, config)
			assert.Equal(1, len(inputs))
			assert.Equal(tt.expected, string(inputs[0].Description))

			outputs, _ := loadOutputs(module, config)
			assert.Equal(1, len(outputs))
			assert.Equal(tt.expected, string(outputs[0].Description))
		})
	}
}

func TestSortItems(t *testing.T) {
	type expected struct {
		inputs    []string
		required  []string
		optional  []string
		outputs   []string
		providers []string
	}
	tests := []struct {
		name        string
		path        string
		sortenabled bool
		sorttype    string
		expected    expected
	}{
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: false,
			sorttype:    "",
			expected: expected{
				inputs:    []string{"D", "B", "E", "A", "C", "F", "G"},
				required:  []string{"A", "F"},
				optional:  []string{"D", "B", "E", "C", "G"},
				outputs:   []string{"C", "A", "B"},
				providers: []string{"tls", "aws", "null"},
			},
		},
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: false,
			sorttype:    print.SortName,
			expected: expected{
				inputs:    []string{"D", "B", "E", "A", "C", "F", "G"},
				required:  []string{"A", "F"},
				optional:  []string{"D", "B", "E", "C", "G"},
				outputs:   []string{"C", "A", "B"},
				providers: []string{"tls", "aws", "null"},
			},
		},
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: false,
			sorttype:    print.SortRequired,
			expected: expected{
				inputs:    []string{"D", "B", "E", "A", "C", "F", "G"},
				required:  []string{"A", "F"},
				optional:  []string{"D", "B", "E", "C", "G"},
				outputs:   []string{"C", "A", "B"},
				providers: []string{"tls", "aws", "null"},
			},
		},
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: true,
			sorttype:    print.SortName,
			expected: expected{
				inputs:    []string{"A", "B", "C", "D", "E", "F", "G"},
				required:  []string{"A", "F"},
				optional:  []string{"B", "C", "D", "E", "G"},
				outputs:   []string{"A", "B", "C"},
				providers: []string{"aws", "null", "tls"},
			},
		},
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: true,
			sorttype:    print.SortRequired,
			expected: expected{
				inputs:    []string{"A", "F", "B", "C", "D", "E", "G"},
				required:  []string{"A", "F"},
				optional:  []string{"B", "C", "D", "E", "G"},
				outputs:   []string{"A", "B", "C"},
				providers: []string{"aws", "null", "tls"},
			},
		},
		{
			name:        "sort module items",
			path:        "full-example",
			sortenabled: true,
			sorttype:    print.SortType,
			expected: expected{
				inputs:    []string{"A", "F", "G", "B", "C", "D", "E"},
				required:  []string{"A", "F"},
				optional:  []string{"G", "B", "C", "D", "E"},
				outputs:   []string{"A", "B", "C"},
				providers: []string{"aws", "null", "tls"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			path := filepath.Join("testdata", tt.path)

			config := print.NewConfig()
			config.ModuleRoot = path
			config.Sort.Enabled = tt.sortenabled
			config.Sort.By = tt.sorttype

			tfmodule, _ := loadModule(path)
			module, err := loadModuleItems(tfmodule, config)

			assert.Nil(err)
			sortItems(module, config)

			for i, v := range module.Inputs {
				assert.Equal(tt.expected.inputs[i], v.Name)
			}
			for i, v := range module.RequiredInputs {
				assert.Equal(tt.expected.required[i], v.Name)
			}
			for i, v := range module.OptionalInputs {
				assert.Equal(tt.expected.optional[i], v.Name)
			}
			for i, v := range module.Outputs {
				assert.Equal(tt.expected.outputs[i], v.Name)
			}
			for i, v := range module.Providers {
				assert.Equal(tt.expected.providers[i], v.Name)
			}
		})
	}
}
