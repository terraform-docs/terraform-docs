package types

import (
	"bytes"
	"encoding/xml"
	"go/types"
	"reflect"
	"sort"
	"strings"
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
}

// ValueOf returns actual value of a variable casted to 'Default' interface.
// This is done to be able to attach specific marshaller func to the type
// (if such a custom function was needed)
func ValueOf(v interface{}) Value {
	if v == nil {
		return new(Nil)
	}
	value := reflect.ValueOf(v)
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
type Nil types.Nil

// HasDefault return false for Nil, because there's no value set for the variable
func (n Nil) HasDefault() bool {
	return false
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

// String returns s as an actual string value
func (s String) String() string {
	return string(s)
}

// nolint
func (s String) underlying() string {
	return string(s)
}

// HasDefault indicates a Terraform variable has a default value set.
func (s String) HasDefault() bool {
	return true
}

// MarshalJSON custom marshal function which sets the value to literal `null` when empty
func (s String) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(s.String()) == 0 {
		buf.WriteString(`null`)
	} else {
		normalize := s.String()
		normalize = strings.Replace(normalize, "\n", "\\n", -1)
		normalize = strings.Replace(normalize, "\"", "\\\"", -1)
		buf.WriteString(`"` + normalize + `"`) // add double quation mark as json format required
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
	if len(s.String()) == 0 {
		return nil, nil
	}
	return s, nil
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

type xmllistentry struct {
	XMLName xml.Name
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
		e.Encode(xmllistentry{XMLName: xml.Name{Local: "item"}, Value: i}) //nolint: errcheck
	}
	return e.EncodeToken(start.End())
}

// Map represents a 'map' of values
type Map map[string]interface{}

// Underlying returns the underlying elements in the form of 'map[string]interface {}'
func (m Map) Underlying() map[string]interface{} {
	r := make(map[string]interface{}, 0)
	for k, e := range m {
		r[k] = e
	}
	return r
}

// HasDefault indicates a Terraform variable has a default value set.
func (m Map) HasDefault() bool {
	return true
}

type xmlmapentry struct {
	XMLName xml.Name
	Value   interface{} `xml:",chardata"`
}

type sortmapkeys []string

func (s sortmapkeys) Len() int           { return len(s) }
func (s sortmapkeys) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortmapkeys) Less(i, j int) bool { return s[i] < s[j] }

// MarshalXML custom marshal function which converts map to its literal
// XML representation. For example:
//
// m := Map{
//     "a": 1,
//     "b": 2,
//     "c": 3,
// }
//
// type foo struct {
//     Value Map `xml:"value"`
// }
//
// will get marshaled to:
//
// <foo>
//   <value>
//     <a>1</a>
//     <b>2</b>
//     <c>3</c>
//   </value>
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
		switch reflect.TypeOf(m[k]).Kind() {
		case reflect.Map:
			is := xml.StartElement{Name: xml.Name{Local: k}}
			Map(m[k].(map[string]interface{})).MarshalXML(e, is) //nolint: errcheck
		case reflect.Slice:
			is := xml.StartElement{Name: xml.Name{Local: k}}
			List(m[k].([]interface{})).MarshalXML(e, is) //nolint: errcheck
		default:
			e.Encode(xmlmapentry{XMLName: xml.Name{Local: k}, Value: m[k]}) //nolint: errcheck
		}
	}
	return e.EncodeToken(start.End())
}
