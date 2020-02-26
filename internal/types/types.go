package types

import (
	"bytes"
	"go/types"
	"reflect"
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

// ValueOf returns actual value of a variable
// casted to 'Value' interface. This is done
// to be able to attach specific marshaller func
// to the type (if such a custom function was needed)
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

// TypeOf returns Terraform type of a value
// based on provided type by terraform-inspect
// or by looking the underlying type of the value
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

// Nil represents a 'nil' value which is
// marshaled to `null` when empty for JSON
// and YAML
type Nil types.Nil

// HasDefault return false for Nil, because there's no value set for the variable
func (n Nil) HasDefault() bool {
	return false
}

// MarshalJSON custom marshal function which
// sets the value to literal `null`
func (n Nil) MarshalJSON() ([]byte, error) {
	return []byte(`null`), nil
}

// MarshalYAML custom marshal function which
// sets the value to literal `null`
func (n Nil) MarshalYAML() (interface{}, error) {
	return nil, nil
}

// String represents a 'string' value which is
// marshaled to `null` when empty for JSON and
// YAML
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

// MarshalJSON custom marshal function which
// sets the value to literal `null` when empty
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

// MarshalYAML custom marshal function which
// sets the value to literal `null` when empty
func (s String) MarshalYAML() (interface{}, error) {
	if len(s.String()) == 0 {
		return nil, nil
	}
	return s, nil
}

// Empty represents an empty 'string' which is
// marshaled to `""` in JSON and YAML
type Empty string

// nolint
func (e Empty) underlying() string {
	return string(e)
}

// HasDefault indicates a Terraform variable has a default value set.
func (e Empty) HasDefault() bool {
	return true
}

// MarshalJSON custom marshal function which
// sets the value to `""`
func (e Empty) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

// Number represents a 'number' value which is
// marshaled to `null` when emty in JSON and YAML
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

// nolint
func (l List) underlying() []interface{} {
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

// Map represents a 'map' of values
type Map map[string]interface{}

// nolint
func (m Map) underlying() map[string]interface{} {
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
