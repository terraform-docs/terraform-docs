package tfconf

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/reader"
)

func readHeader(path string) string {
	filename := filepath.Join(path, "main.tf")
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	lines := reader.Lines{
		FileName: filename,
		LineNum:  -1,
		Condition: func(line string) bool {
			line = strings.TrimSpace(line)
			return strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") || strings.HasPrefix(line, "*/")
		},
		Parser: func(line string) (string, bool) {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*/") {
				return "", false
			}
			if line == "*" {
				return "", true
			}
			line = strings.TrimPrefix(line, "* ")
			return line, true
		},
	}
	header, err := lines.Extract()
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(header, "\n")
}
