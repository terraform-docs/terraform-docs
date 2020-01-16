package tfconf

import (
	"bytes"
	"strings"
)

// String represents a 'string' value which
// JSON marshalled to `null` when empty
type String string

// MarshalJSON custom marshal function which
// sets the value to literal `null` when empty
func (s String) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(string(s)) == 0 {
		buf.WriteString(`null`)
	} else {
		normalize := string(s)
		normalize = strings.Replace(normalize, "\n", "\\n", -1)
		normalize = strings.Replace(normalize, "\"", "\\\"", -1)
		buf.WriteString(`"` + normalize + `"`) // add double quation mark as json format required
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON custom unmarshal function which
// sets the value to `""` when the json property
// is `null`
func (s *String) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*s = ""
		return nil
	}
	res := String(str)
	if len(res) >= 2 {
		res = res[1 : len(res)-1] // remove the wrapped qutation
	}
	*s = res
	return nil
}
