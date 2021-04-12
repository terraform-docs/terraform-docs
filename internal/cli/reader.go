/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type cfgreader struct {
	file      string
	config    *Config
	overrides Config
}

func (c *cfgreader) exist() (bool, error) {
	if c.file == "" {
		return false, fmt.Errorf("config file name is missing")
	}
	if info, err := os.Stat(c.file); os.IsNotExist(err) || info == nil || info.IsDir() {
		return false, err
	}
	return true, nil
}

func (c *cfgreader) parse() error { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	if ok, err := c.exist(); !ok {
		return err
	}

	content, err := ioutil.ReadFile(c.file)
	if err != nil {
		return err
	}

	c.overrides = *c.config
	if err := yaml.Unmarshal(content, c.config); err != nil {
		return err
	}

	mappings := map[string]struct {
		flag string
		from interface{}
		to   interface{}
	}{
		"header-from": {
			flag: "header-from",
			from: &c.overrides,
			to:   c.config,
		},
		"footer-from": {
			flag: "footer-from",
			from: &c.overrides,
			to:   c.config,
		},

		// sort
		"sort": {
			flag: "enabled",
			from: &c.overrides.Sort,
			to:   &c.config.Sort,
		},
		"sort-by": {
			flag: "by",
			from: &c.overrides.Sort,
			to:   &c.config.Sort,
		},
		"sort-by-required": {
			flag: "required",
			from: nil,
			to:   nil,
		},
		"sort-by-type": {
			flag: "type",
			from: nil,
			to:   nil,
		},

		// output
		"output-file": {
			flag: "file",
			from: &c.overrides.Output,
			to:   &c.config.Output,
		},
		"output-mode": {
			flag: "mode",
			from: &c.overrides.Output,
			to:   &c.config.Output,
		},
		"output-template": {
			flag: "template",
			from: &c.overrides.Output,
			to:   &c.config.Output,
		},

		// output-values
		"output-values": {
			flag: "enabled",
			from: &c.overrides.OutputValues,
			to:   &c.config.OutputValues,
		},
		"output-values-from": {
			flag: "from",
			from: &c.overrides.OutputValues,
			to:   &c.config.OutputValues,
		},

		// settings
		"anchor": {
			flag: "anchor",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"color": {
			flag: "color",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"default": {
			flag: "default",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"escape": {
			flag: "escape",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"indent": {
			flag: "indent",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"required": {
			flag: "required",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"sensitive": {
			flag: "sensitive",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
		"type": {
			flag: "type",
			from: &c.overrides.Settings,
			to:   &c.config.Settings,
		},
	}

	// If '--show' or '--hide' CLI flag is used, explicitly override and remove
	// all items from 'show' and 'hide' set in '.terraform-doc.yml'.
	if changedfs["show"] || changedfs["hide"] {
		c.config.Sections.Show = []string{}
		c.config.Sections.Hide = []string{}
	}

	for flag, enabled := range changedfs {
		if !enabled {
			continue
		}

		switch flag {
		case "show":
			c.overrideShow()
		case "hide":
			c.overrideHide()
		case "sort-by-required", "sort-by-type":
			c.config.Sort.By = mappings[flag].flag
		default:
			if _, ok := mappings[flag]; !ok {
				continue
			}
			if err := c.overrideValue(mappings[flag].flag, mappings[flag].to, mappings[flag].from); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *cfgreader) overrideValue(name string, to interface{}, from interface{}) error {
	if name == "" || name == "-" {
		return fmt.Errorf("tag name cannot be blank or empty")
	}
	toEl := reflect.ValueOf(to).Elem()
	field, err := c.findField(toEl, "yaml", name)
	if err != nil {
		return err
	}
	fromEl := reflect.ValueOf(from).Elem()
	toEl.FieldByName(field).Set(fromEl.FieldByName(field))
	return nil
}

func (c *cfgreader) overrideShow() {
	for _, item := range c.overrides.Sections.Show {
		if !contains(c.config.Sections.Show, item) {
			c.config.Sections.Show = append(c.config.Sections.Show, item)
		}
	}
}

func (c *cfgreader) overrideHide() {
	for _, item := range c.overrides.Sections.Hide {
		if !contains(c.config.Sections.Hide, item) {
			c.config.Sections.Hide = append(c.config.Sections.Hide, item)
		}
	}
}

func (c *cfgreader) findField(el reflect.Value, tag string, value string) (string, error) {
	for i := 0; i < el.NumField(); i++ {
		f := el.Type().Field(i)
		t := f.Tag.Get(tag)
		if t == "" || t == "-" || t != value {
			continue
		}
		return f.Name, nil
	}
	return "", fmt.Errorf("field with tag: '%s', value; '%s' not found or not readable", tag, value)
}
