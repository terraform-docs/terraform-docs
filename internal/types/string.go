package types

import (
	"bytes"
	"strings"
)

// TFString represents a 'string' value which
// JSON marshaled to `null` when empty
type TFString string

// String returns s as an actual string value
func (s TFString) String() string {
	return string(s)
}

// MarshalJSON custom marshal function which
// sets the value to literal `null` when empty
func (s *TFString) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON custom unmarshal function which
// sets the value to `""` when the json property
// is `null`
func (s *TFString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*s = ""
		return nil
	}
	res := TFString(str)
	if len(res) >= 2 {
		res = res[1 : len(res)-1] // remove the wrapped qutation
	}
	*s = res
	return nil
}
