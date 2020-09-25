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

func (c *cfgreader) parse() error {
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

	if c.config.Sections.HideAll && !changedfs["show-all"] {
		c.config.Sections.ShowAll = false
	}
	if !c.config.Sections.ShowAll && !changedfs["hide-all"] {
		c.config.Sections.HideAll = true
	}

	for flag, enabled := range changedfs {
		if !enabled {
			continue
		}

		switch flag {
		case "header-from":
			if err := c.overrideValue(flag, c.config, &c.overrides); err != nil {
				return err
			}
		case "show":
			c.overrideShow()
		case "hide":
			c.overrideHide()
		case "sort":
			if err := c.overrideValue("enabled", &c.config.Sort, &c.overrides.Sort); err != nil {
				return err
			}
		case "sort-by-required", "sort-by-type":
			mapping := map[string]string{"sort-by-required": "required", "sort-by-type": "type"}
			if !contains(c.config.Sort.ByList, mapping[flag]) {
				c.config.Sort.ByList = append(c.config.Sort.ByList, mapping[flag])
			}
			el := reflect.ValueOf(&c.overrides.Sort.By).Elem()
			field, err := c.findField(el, "name", mapping[flag])
			if err != nil {
				return err
			}
			if !el.FieldByName(field).Bool() {
				c.config.Sort.ByList = remove(c.config.Sort.ByList, mapping[flag])
			}
		case "output-values", "output-values-from":
			mapping := map[string]string{"output-values": "enabled", "output-values-from": "from"}
			if err := c.overrideValue(mapping[flag], &c.config.OutputValues, &c.overrides.OutputValues); err != nil {
				return err
			}
		case "color", "escape", "indent", "required", "sensitive":
			if err := c.overrideValue(flag, &c.config.Settings, &c.overrides.Settings); err != nil {
				return err
			}
		}
	}

	if err := c.updateSortTypes(); err != nil {
		return err
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
		if c.config.Sections.ShowAll {
			if contains(c.config.Sections.Hide, item) {
				c.config.Sections.Hide = remove(c.config.Sections.Hide, item)
				c.config.Sections.Show = remove(c.config.Sections.Show, item)
			}
		} else {
			if !contains(c.config.Sections.Show, item) {
				c.config.Sections.Show = append(c.config.Sections.Show, item)
			}
		}
	}
}

func (c *cfgreader) overrideHide() {
	for _, item := range c.overrides.Sections.Hide {
		if c.config.Sections.HideAll {
			if contains(c.config.Sections.Show, item) {
				c.config.Sections.Show = remove(c.config.Sections.Show, item)
				c.config.Sections.Hide = remove(c.config.Sections.Hide, item)
			}
		} else {
			if !contains(c.config.Sections.Hide, item) {
				c.config.Sections.Hide = append(c.config.Sections.Hide, item)
			}
		}
	}
}

func (c *cfgreader) updateSortTypes() error {
	for _, item := range c.config.Sort.ByList {
		el := reflect.ValueOf(&c.config.Sort.By).Elem()
		field, err := c.findField(el, "name", item)
		if err != nil {
			return err
		}
		el.FieldByName(field).Set(reflect.ValueOf(true))
	}
	return nil
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
