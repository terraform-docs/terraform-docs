package types

import (
	"bytes"
	"fmt"
	"go/types"
	"strings"
)

// Default is a default value of an input
// it can be of several types:
// - Nil
// - String
// - Empty
// - Number
// - Bool
// - List
// - Map
type Default interface {
	HasDefault() bool
}

// ValueOf returns actual value of a variable
// casted to 'Default' interface. This is done
// to be able to attach specific marshaller func
// to the type (if such a custom function was needed)
func ValueOf(v interface{}) Default {
	if v == nil {
		return new(Nil)
	}
	switch i := v.(type) {
	case string:
		if i == "" {
			return Empty("")
		}
		return String(i)
	case float64:
		return Number(i)
	case float32:
		return Number(float64(i))
	case int64:
		return Number(float64(i))
	case int32:
		return Number(float64(i))
	case int16:
		return Number(float64(i))
	case int8:
		return Number(float64(i))
	case int:
		return Number(float64(i))
	case bool:
		return Bool(i)
	case []interface{}:
		return List(i)
	case map[string]interface{}:
		return Map(i)
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
		switch x := fmt.Sprintf("%T", v); x {
		case "string":
			return String("string")
		case "int", "int8", "int16", "int32", "int64", "float32", "float64":
			return String("number")
		case "bool":
			return String("bool")
		case "[]interface {}":
			return String("list")
		case "map[string]interface {}":
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

// HasDefault indicates a Terraform variable has a default value set.
func (n Number) HasDefault() bool {
	return true
}

// Bool represents a 'bool' value
type Bool bool

// HasDefault indicates a Terraform variable has a default value set.
func (b Bool) HasDefault() bool {
	return true
}

// List represents a 'list' of values
type List []interface{}

// HasDefault indicates a Terraform variable has a default value set.
func (l List) HasDefault() bool {
	return true
}

// Map represents a 'map' of values
type Map map[string]interface{}

// HasDefault indicates a Terraform variable has a default value set.
func (m Map) HasDefault() bool {
	return true
}
