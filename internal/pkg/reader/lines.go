package reader

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// Lines represents line reader in a given 'FileName' immediately
// before the given 'LineNum'. Extraction happens when 'Condition'
// is met and being processed by 'Parser' function.
type Lines struct {
	FileName  string
	LineNum   int // value -1 means scan the whole file and break after finding what we were looking for
	Condition func(line string) bool
	Parser    func(line string) (string, bool)
}

// Extract extracts lines in given file and based on the provided
// condition. returns empty if nothing found.
func (l *Lines) Extract() ([]string, error) {
	f, err := os.Open(l.FileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	bf := bufio.NewReader(f)
	var lines = make([]string, 0)

	for lnum := 0; ; lnum++ {
		if l.LineNum != -1 && lnum >= l.LineNum {
			break
		}
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
		if l.Condition(line) {
			if extracted, capture := l.Parser(line); capture {
				lines = append(lines, extracted)
			}
		} else if l.LineNum == -1 {
			break
		} else {
			lines = nil
		}
	}

	return lines, nil
}
