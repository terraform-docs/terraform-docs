package tfconf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func readComment(filename string, lineNum int) string {
	lines, err := readLines(filename, lineNum)
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(lines, " ")
}

// Reading leading comments, i.e. lines start with # or //
// immediately before specific line number in given file
func readLines(filename string, lineNum int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	bf := bufio.NewReader(f)
	var lines = make([]string, 0)

	for lnum := 0; lnum < lineNum; lnum++ {
		line, err := bf.ReadString('\n')
		if err == io.EOF {
			switch lnum {
			case 0:
				return nil, errors.New("no lines in file")
			case 1:
				return nil, errors.New("only 1 line")
			default:
				return nil, fmt.Errorf("only %d lines", lnum)
			}
		}
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "#")
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			lines = append(lines, line)
		} else {
			lines = nil
		}
	}

	return lines, nil
}
