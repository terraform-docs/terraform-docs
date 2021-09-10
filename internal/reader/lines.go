/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

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
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return []string{}, nil
	}
	defer func() {
		_ = f.Close()
	}()
	return l.extract(f)
}

func (l *Lines) extract(r io.Reader) ([]string, error) { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	bf := bufio.NewReader(r)
	var lines = make([]string, 0)
	for lnum := 0; ; lnum++ {
		if l.LineNum != -1 && lnum >= l.LineNum-1 {
			break
		}
		line, err := bf.ReadString('\n')
		if errors.Is(err, io.EOF) && line == "" {
			switch lnum {
			case 0:
				return nil, errors.New("no lines in file")
			case 1:
				return nil, errors.New("only 1 line")
			default:
				if l.LineNum == -1 {
					break
				}
				return nil, fmt.Errorf("only %d lines", lnum)
			}
		}

		//nolint:gocritic
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
