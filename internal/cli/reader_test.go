/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		config   cfgreader
		expected bool
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "config file exists",
			config:   cfgreader{file: "testdata/sample-config.yaml"},
			expected: true,
			wantErr:  false,
			errMsg:   "",
		},
		{
			name:     "config file not found",
			config:   cfgreader{file: "testdata/noop.yaml"},
			expected: false,
			wantErr:  true,
			errMsg:   "",
		},
		{
			name:     "config file empty",
			config:   cfgreader{file: ""},
			expected: false,
			wantErr:  true,
			errMsg:   "config file name is missing",
		},
		{
			name:     "main argument is a file",
			config:   cfgreader{file: "testdata/sample-config.yaml/some-config.yaml"},
			expected: false,
			wantErr:  true,
			errMsg:   "stat testdata/sample-config.yaml/some-config.yaml: not a directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := tt.config.exist()
			if tt.wantErr {
				assert.NotNil(err)
				if tt.errMsg != "" {
					assert.Equal(tt.errMsg, err.Error())
				}
			} else {
				assert.Nil(err)
			}
			assert.Equal(tt.expected, actual)
		})
	}
}

func TestOverrideValue(t *testing.T) {
	config := DefaultConfig()
	override := DefaultConfig()
	tests := []struct {
		name       string
		tag        string
		to         func() interface{}
		from       func() interface{}
		overrideFn func()
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "override values of given field",
			tag:        "header-from",
			to:         func() interface{} { return config },
			from:       func() interface{} { return override },
			overrideFn: func() { override.HeaderFrom = "foo.txt" },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of given field",
			tag:        "enabled",
			to:         func() interface{} { return &config.Sort },
			from:       func() interface{} { return &override.Sort },
			overrideFn: func() { override.Sort.Enabled = false },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of given field",
			tag:        "color",
			to:         func() interface{} { return &config.Settings },
			from:       func() interface{} { return &override.Settings },
			overrideFn: func() { override.Settings.Color = false },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of given field",
			tag:        "mode",
			to:         func() interface{} { return &config.Output },
			from:       func() interface{} { return &override.Output },
			overrideFn: func() { override.Output.Mode = "replace" },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of unkwon field tag",
			tag:        "not-available",
			to:         func() interface{} { return config },
			from:       func() interface{} { return override },
			overrideFn: func() {},
			wantErr:    true,
			errMsg:     "field with tag: 'yaml', value; 'not-available' not found or not readable",
		},
		{
			name:       "override values of blank field tag",
			tag:        "-",
			to:         func() interface{} { return config },
			from:       func() interface{} { return override },
			overrideFn: func() {},
			wantErr:    true,
			errMsg:     "tag name cannot be blank or empty",
		},
		{
			name:       "override values of empty field tag",
			tag:        "",
			to:         func() interface{} { return config },
			from:       func() interface{} { return override },
			overrideFn: func() {},
			wantErr:    true,
			errMsg:     "tag name cannot be blank or empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			c := cfgreader{config: config}

			tt.overrideFn()

			if !tt.wantErr {
				// make sure before values are different
				assert.NotEqual(override, config)
			}

			// then override property 'from' to 'to'
			err := c.overrideValue(tt.tag, tt.to(), tt.from())

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
			}

			// then make sure values are the same now
			assert.Equal(override, config)
		})
	}
}

func TestOverrideShow(t *testing.T) {
	tests := []struct {
		name         string
		show         []string
		hide         []string
		showall      bool
		overrideShow []string
		expectedShow []string
		expectedHide []string
	}{
		{
			name:         "override section show",
			show:         []string{""},
			hide:         []string{"inputs", "outputs"},
			showall:      true,
			overrideShow: []string{"inputs"},
			expectedShow: []string{""},
			expectedHide: []string{"outputs"},
		},
		{
			name:         "override section show",
			show:         []string{"providers"},
			hide:         []string{"inputs"},
			showall:      true,
			overrideShow: []string{"outputs"},
			expectedShow: []string{"providers"},
			expectedHide: []string{"inputs"},
		},
		{
			name:         "override section show",
			show:         []string{"inputs"},
			hide:         []string{"providers"},
			showall:      false,
			overrideShow: []string{"outputs"},
			expectedShow: []string{"inputs", "outputs"},
			expectedHide: []string{"providers"},
		},
		{
			name:         "override section show",
			show:         []string{"inputs"},
			hide:         []string{"inputs"},
			showall:      false,
			overrideShow: []string{"inputs"},
			expectedShow: []string{"inputs"},
			expectedHide: []string{"inputs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			config := DefaultConfig()
			override := DefaultConfig()
			c := cfgreader{config: config, overrides: *override}

			c.config.Sections.Show = tt.show
			c.config.Sections.Hide = tt.hide
			c.config.Sections.ShowAll = tt.showall
			c.overrides.Sections.Show = tt.overrideShow

			c.overrideShow()

			assert.Equal(tt.expectedShow, c.config.Sections.Show)
			assert.Equal(tt.expectedHide, c.config.Sections.Hide)
		})
	}
}

func TestOverrideHide(t *testing.T) {
	tests := []struct {
		name         string
		show         []string
		hide         []string
		hideall      bool
		overrideHide []string
		expectedShow []string
		expectedHide []string
	}{
		{
			name:         "override section hide",
			show:         []string{"inputs", "outputs"},
			hide:         []string{""},
			hideall:      true,
			overrideHide: []string{"inputs"},
			expectedShow: []string{"outputs"},
			expectedHide: []string{""},
		},
		{
			name:         "override section hide",
			show:         []string{"inputs"},
			hide:         []string{"providers"},
			hideall:      true,
			overrideHide: []string{"outputs"},
			expectedShow: []string{"inputs"},
			expectedHide: []string{"providers"},
		},
		{
			name:         "override section hide",
			show:         []string{"providers"},
			hide:         []string{"inputs"},
			hideall:      false,
			overrideHide: []string{"outputs"},
			expectedShow: []string{"providers"},
			expectedHide: []string{"inputs", "outputs"},
		},
		{
			name:         "override section hide",
			show:         []string{"inputs"},
			hide:         []string{"inputs"},
			hideall:      false,
			overrideHide: []string{"inputs"},
			expectedShow: []string{"inputs"},
			expectedHide: []string{"inputs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			config := DefaultConfig()
			override := DefaultConfig()
			c := cfgreader{config: config, overrides: *override}

			c.config.Sections.Show = tt.show
			c.config.Sections.Hide = tt.hide
			c.config.Sections.HideAll = tt.hideall
			c.overrides.Sections.Hide = tt.overrideHide

			c.overrideHide()

			assert.Equal(tt.expectedShow, c.config.Sections.Show)
			assert.Equal(tt.expectedHide, c.config.Sections.Hide)
		})
	}
}

func TestUpdateSortTypes(t *testing.T) {
	tests := []struct {
		name       string
		appendFn   func(config *Config)
		expectedFn func(config *Config) bool
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "override values of given field",
			appendFn:   func(config *Config) { config.Sort.ByList = append(config.Sort.ByList, "required") },
			expectedFn: func(config *Config) bool { return config.Sort.By.Required },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of given field",
			appendFn:   func(config *Config) { config.Sort.ByList = append(config.Sort.ByList, "type") },
			expectedFn: func(config *Config) bool { return config.Sort.By.Type },
			wantErr:    false,
			errMsg:     "",
		},
		{
			name:       "override values of given field",
			appendFn:   func(config *Config) { config.Sort.ByList = append(config.Sort.ByList, "unknown") },
			expectedFn: func(config *Config) bool { return false },
			wantErr:    true,
			errMsg:     "field with tag: 'name', value; 'unknown' not found or not readable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			config := DefaultConfig()
			c := cfgreader{config: config}

			tt.appendFn(config)

			// make sure before values is false
			assert.Equal(false, tt.expectedFn(config))

			// then update sort types
			err := c.updateSortTypes()

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				// then make sure values is true
				assert.Nil(err)
				assert.Equal(true, tt.expectedFn(config))
			}
		})
	}
}

func TestFindField(t *testing.T) {
	type sample struct {
		A string `foo:"a"`
		B string `bar:"b"`
		C string `baz:"-"`
		D string `fizz:"-"`
		E string `buzz:""`
		F string
	}
	tests := []struct {
		name     string
		tag      string
		value    string
		expected string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "find field with given tag",
			tag:      "foo",
			value:    "a",
			expected: "A",
			wantErr:  false,
			errMsg:   "",
		},
		{
			name:     "find field with given tag",
			tag:      "bar",
			value:    "b",
			expected: "B",
			wantErr:  false,
			errMsg:   "",
		},
		{
			name:     "find field with tag none",
			tag:      "baz",
			value:    "-",
			expected: "",
			wantErr:  true,
			errMsg:   "field with tag: 'baz', value; '-' not found or not readable",
		},
		{
			name:     "find field with tag none",
			tag:      "fizz",
			value:    "-",
			expected: "",
			wantErr:  true,
			errMsg:   "field with tag: 'fizz', value; '-' not found or not readable",
		},
		{
			name:     "find field with tag empty",
			tag:      "buzz",
			value:    "",
			expected: "",
			wantErr:  true,
			errMsg:   "field with tag: 'buzz', value; '' not found or not readable",
		},
		{
			name:     "find field with tag unknown",
			tag:      "unknown",
			value:    "unknown",
			expected: "",
			wantErr:  true,
			errMsg:   "field with tag: 'unknown', value; 'unknown' not found or not readable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			c := cfgreader{}
			el := reflect.ValueOf(&sample{}).Elem()

			actual, err := c.findField(el, tt.tag, tt.value)

			if tt.wantErr {
				assert.NotNil(err)
				assert.Equal(tt.errMsg, err.Error())
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, actual)
			}
		})
	}
}
