/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package types

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"sort"
)

// Value is a default value of an input or output.
// it can be of several types:
//
// - Nil
// - String
// - Empty
// - Number
// - Bool
// - List
// - Map
type Value interface {
	HasDefault() bool
	Length() int
	Raw() interface{}
}

// ValueOf returns actual value of a variable casted to 'Default' interface.
// This is done to be able to attach specific marshaller func to the type
// (if such a custom function was needed)
func ValueOf(v interface{}) Value {
	if v == nil {
		return new(Nil)
	}
	value := reflect.ValueOf(v)

	// We don't really care about all the other kinds.
	//
	//nolint:exhaustive
	switch value.Kind() {
	case reflect.String:
		if value.IsZero() {
			return Empty("")
		}
		return String(value.String())
	case reflect.Float32, reflect.Float64:
		return Number(value.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Number(float64(value.Int()))
	case reflect.Bool:
		return Bool(value.Bool())
	case reflect.Slice:
		return List(value.Interface().([]interface{}))
	case reflect.Map:
		return Map(value.Interface().(map[string]interface{}))
	}
	return new(Nil)
}

// TypeOf returns Terraform type of a value based on provided type by
// terraform-inspect or by looking the underlying type of the value
func TypeOf(t string, v interface{}) String {
	if t != "" {
		return String(t)
	}
	if v != nil {
		// We don't really care about all the other kinds.
		//
		//nolint:exhaustive
		switch reflect.ValueOf(v).Kind() {
		case reflect.String:
			return String("string")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			return String("number")
		case reflect.Bool:
			return String("bool")
		case reflect.Slice:
			return String("list")
		case reflect.Map:
			return String("map")
		}
	}
	return String("any")
}

// Nil represents a 'nil' value which is marshaled to `null` when empty for JSON and YAML
type Nil struct{}

// HasDefault return false for Nil, because there's no value set for the variable
func (n Nil) HasDefault() bool {
	return false
}

// Length returns the length of underlying item
func (n Nil) Length() int {
	return 0
}

// Raw underlying value of this type.
func (n Nil) Raw() interface{} {
	return nil
}

// MarshalJSON custom marshal function which sets the value to literal `null`
func (n Nil) MarshalJSON() ([]byte, error) {
	return []byte(`null`), nil
}

// MarshalXML custom marshal function which adds property 'xsi:nil="true"' to a tag
// of a 'nil' item
func (n Nil) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xsi:nil"}, Value: "true"})
	return e.EncodeElement(``, start)
}

// MarshalYAML custom marshal function which sets the value to literal `null`
func (n Nil) MarshalYAML() (interface{}, error) {
	return nil, nil
}

// String represents a 'string' value which is marshaled to `null` when empty for JSON and YAML
type String string

// nolint
func (s String) underlying() string {
	return string(s)
}

// HasDefault indicates a Terraform variable has a default value set.
func (s String) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (s String) Length() int {
	return len(s.underlying())
}

// Raw underlying value of this type.
func (s String) Raw() interface{} {
	return s.underlying()
}

// MarshalJSON custom marshal function which sets the value to literal `null` when empty
func (s String) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(string(s)) == 0 {
		buf.WriteString(`null`)
	} else {
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(string(s)); err != nil {
			return nil, err
		}
		buf.Truncate(buf.Len() - 1) // The json encoder adds a newline, this is not configurable
	}
	return buf.Bytes(), nil
}

// MarshalXML custom marshal function which adds property 'xsi:nil="true"' to a tag
// if the underlying item is 'nil'
func (s String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if string(s) == "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xsi:nil"}, Value: "true"})
		return e.EncodeElement(``, start)
	}
	return e.EncodeElement(string(s), start)
}

// MarshalYAML custom marshal function which sets the value to literal `null` when empty
func (s String) MarshalYAML() (interface{}, error) {
	if len(string(s)) == 0 || string(s) == `""` {
		return nil, nil
	}
	return string(s), nil
}

// Empty represents an empty 'string' which is marshaled to `""` in JSON and YAML
type Empty string

// nolint
func (e Empty) underlying() string {
	return string(e)
}

// HasDefault indicates a Terraform variable has a default value set.
func (e Empty) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (e Empty) Length() int {
	return len(e.underlying())
}

// Raw underlying value of this type.
func (e Empty) Raw() interface{} {
	return e.underlying()
}

// MarshalJSON custom marshal function which sets the value to `""`
func (e Empty) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

// Number represents a 'number' value which is marshaled to `null` when emty in JSON and YAML
type Number float64

// nolint
func (n Number) underlying() float64 {
	return float64(n)
}

// HasDefault indicates a Terraform variable has a default value set.
func (n Number) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (n Number) Length() int {
	return 0
}

// Raw underlying value of this type.
func (n Number) Raw() interface{} {
	return n.underlying()
}

// Bool represents a 'bool' value
type Bool bool

// nolint
func (b Bool) underlying() bool {
	return bool(b)
}

// HasDefault indicates a Terraform variable has a default value set.
func (b Bool) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (b Bool) Length() int {
	return 0
}

// Raw underlying value of this type.
func (b Bool) Raw() interface{} {
	return b.underlying()
}

// List represents a 'list' of values
type List []interface{}

// Underlying returns the underlying elements in the form of '[]interface {}'
func (l List) Underlying() []interface{} {
	r := make([]interface{}, 0)
	for _, i := range l {
		r = append(r, i)
	}
	return r
}

// HasDefault indicates a Terraform variable has a default value set.
func (l List) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (l List) Length() int {
	return len(l)
}

// Raw underlying value of this type.
func (l List) Raw() interface{} {
	return l.Underlying()
}

type xmllistentry struct {
	XMLName xml.Name    `xml:"item"`
	Value   interface{} `xml:",chardata"`
}

// MarshalXML custom marshal function which wraps list items in '<default></default>'
// tag and each items of the list will be wrapped in a '<item></item>' tag
func (l List) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(l) == 0 {
		return e.EncodeElement(``, start)
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for _, i := range l {
		e.Encode(xmllistentry{XMLName: xml.Name{Local: "item"}, Value: i}) //nolint:errcheck,gosec
	}
	return e.EncodeToken(start.End())
}

// Map represents a 'map' of values
type Map map[string]interface{}

// Underlying returns the underlying elements in the form of 'map[string]interface {}'
func (m Map) Underlying() map[string]interface{} {
	r := make(map[string]interface{})
	for k, e := range m {
		r[k] = e
	}
	return r
}

// Raw underlying value of this type.
func (m Map) Raw() interface{} {
	return m.Underlying()
}

// HasDefault indicates a Terraform variable has a default value set.
func (m Map) HasDefault() bool {
	return true
}

// Length returns the length of underlying item
func (m Map) Length() int {
	return len(m)
}

type xmlmapentry struct {
	XMLName xml.Name    `xml:","`
	Value   interface{} `xml:",chardata"`
}

type sortmapkeys []string

func (s sortmapkeys) Len() int           { return len(s) }
func (s sortmapkeys) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortmapkeys) Less(i, j int) bool { return s[i] < s[j] }

// MarshalXML custom marshal function which converts map to its literal
// XML representation. For example:
//
//	m := Map{
//	    "a": 1,
//	    "b": 2,
//	    "c": 3,
//	}
//
//	type foo struct {
//	    Value Map `xml:"value"`
//	}
//
// will get marshaled to:
//
// <foo>
//
//	<value>
//	  <a>1</a>
//	  <b>2</b>
//	  <c>3</c>
//	</value>
//
// </foo>
func (m Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return e.EncodeElement(``, start)
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sortmapkeys(keys))
	for _, k := range keys {
		// We don't really care about all the other kinds.
		//
		//nolint:exhaustive
		switch reflect.TypeOf(m[k]).Kind() {
		case reflect.Map:
			is := xml.StartElement{Name: xml.Name{Local: k}}
			Map(m[k].(map[string]interface{})).MarshalXML(e, is) //nolint:errcheck,gosec
		case reflect.Slice:
			is := xml.StartElement{Name: xml.Name{Local: k}}
			List(m[k].([]interface{})).MarshalXML(e, is) //nolint:errcheck,gosec
		default:
			e.Encode(xmlmapentry{XMLName: xml.Name{Local: k}, Value: m[k]}) //nolint:errcheck,gosec
		}
	}
	return e.EncodeToken(start.End())
}
