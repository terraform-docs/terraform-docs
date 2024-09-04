/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package template

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"mvdan.cc/xurls/v2"
)

// SanitizeName escapes underscore character which have special meaning in
// Markdown.
func SanitizeName(name string, escape bool) string {
	if escape {
		// Escape underscore
		name = strings.ReplaceAll(name, "_", "\\_")
	}
	return name
}

// SanitizeSection converts passed 'string' to suitable Markdown or AsciiDoc
// representation for a document. (including line-break, illegal characters,
// code blocks etc). This is in particular being used for header and footer.
//
// IMPORTANT: SanitizeSection will never change the line-endings and preserve
// them as they are provided by the users.
func SanitizeSection(s string, escape bool, html bool) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string, first bool, last bool) string {
			segment = EscapeCharacters(segment, escape, false)
			segment = ConvertMultiLineText(segment, false, true, html)
			segment = NormalizeURLs(segment, escape)
			return segment
		},
		func(segment string, first bool, last bool) string {
			lastbreak := ""
			if !strings.HasSuffix(segment, "\n") {
				lastbreak = "\n"
			}

			// Adjust indention and linebreak for indented codeblock
			// https://github.com/terraform-docs/terraform-docs/issues/521
			lastindent := ""
			lines := strings.Split(segment, "\n")
			if len(strings.TrimSpace(lines[len(lines)-1])) == 0 {
				lastbreak = ""
			}

			segment = fmt.Sprintf("```%s%s%s```", segment, lastindent, lastbreak)
			return segment
		},
	)
	return result
}

// SanitizeDocument converts passed 'string' to suitable Markdown or AsciiDoc
// representation for a document. (including line-break, illegal characters,
// code blocks etc).
func SanitizeDocument(s string, escape bool, html bool) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string, first bool, last bool) string {
			segment = EscapeCharacters(segment, escape, false)
			segment = ConvertMultiLineText(segment, false, false, html)
			segment = NormalizeURLs(segment, escape)
			return segment
		},
		func(segment string, first bool, last bool) string {
			lastbreak := ""
			if !strings.HasSuffix(segment, "\n") {
				lastbreak = "\n"
			}
			segment = fmt.Sprintf("```%s%s```", segment, lastbreak)
			return segment
		},
	)
	return result
}

// SanitizeMarkdownTable converts passed 'string' to suitable Markdown representation
// for a table. (including line-break, illegal characters, code blocks etc).
func SanitizeMarkdownTable(s string, escape bool, html bool) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string, first bool, last bool) string {
			segment = EscapeCharacters(segment, escape, true)
			segment = ConvertMultiLineText(segment, true, false, html)
			segment = NormalizeURLs(segment, escape)
			return segment
		},
		func(segment string, first bool, last bool) string {
			linebreak := "<br/>"
			codestart := "<pre>"
			codeend := "</pre>"

			segment = strings.TrimSpace(segment)

			if !html {
				linebreak = ""
				codestart = " ```"
				codeend = "``` "

				if first {
					codestart = codestart[1:]
				}
				if last {
					codeend = codeend[:3]
				}

				segment = ConvertOneLineCodeBlock(segment)
			}

			segment = strings.ReplaceAll(segment, "\n", linebreak)
			segment = strings.ReplaceAll(segment, "\r", "")
			segment = fmt.Sprintf("%s%s%s", codestart, segment, codeend)
			return segment
		},
	)
	return result
}

// SanitizeAsciidocTable converts passed 'string' to suitable AsciiDoc representation
// for a table. (including line-break, illegal characters, code blocks etc).
func SanitizeAsciidocTable(s string, escape bool, html bool) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string, first bool, last bool) string {
			segment = EscapeCharacters(segment, escape, true)
			segment = NormalizeURLs(segment, escape)
			return segment
		},
		func(segment string, first bool, last bool) string {
			segment = strings.TrimSpace(segment)
			segment = fmt.Sprintf("[source]\n----\n%s\n----", segment)
			return segment
		},
	)
	return result
}

// ConvertMultiLineText converts a multi-line text into a suitable Markdown representation.
func ConvertMultiLineText(s string, isTable bool, isHeader bool, showHTML bool) string {
	if isTable {
		s = strings.TrimSpace(s)
	}

	// Convert line-break on a non-empty line followed by another line
	// starting with "alphanumeric" word into space-space-newline
	// which is a know convention of Markdown for multi-lines paragprah.
	// This doesn't apply on a markdown list for example, because all the
	// consecutive lines start with hyphen which is a special character.
	if !isHeader {
		s = regexp.MustCompile(`(\S*)(\r?\n)(\s*)(\w+)`).ReplaceAllString(s, "$1  $2$3$4")
		s = strings.ReplaceAll(s, "    \n", "  \n")
		s = strings.ReplaceAll(s, "  \n\n", "\n\n")
		s = strings.ReplaceAll(s, "\n  \n", "\n\n")
	}

	if !isTable {
		return s
	}

	// representation of line break. <br/> if showHTML is true, <space> if false.
	linebreak := " "

	if showHTML {
		linebreak = "<br/>"
	}

	// Convert space-space-newline to 'linebreak'.
	s = strings.ReplaceAll(s, "  \n", linebreak)

	// Convert single newline to 'linebreak'.
	return strings.ReplaceAll(s, "\n", linebreak)
}

// ConvertOneLineCodeBlock converts a multi-line code block into a one-liner.
// Line breaks are replaced with single space.
func ConvertOneLineCodeBlock(s string) string {
	splitted := strings.Split(s, "\n")
	result := []string{}
	for _, segment := range splitted {
		if len(strings.TrimSpace(segment)) == 0 {
			continue
		}
		segment = regexp.MustCompile(`(\s*)=(\s*)`).ReplaceAllString(segment, " = ")
		segment = strings.TrimLeftFunc(segment, unicode.IsSpace)
		result = append(result, segment)
	}
	return strings.Join(result, " ")
}

// EscapeCharacters escapes characters which have special meaning in Markdown into
// their corresponding literal.
func EscapeCharacters(s string, escape bool, escapePipe bool) string {
	// Escape pipe (only for 'markdown table' or 'asciidoc table')
	if escapePipe {
		s = processSegments(
			s,
			"`",
			func(segment string, first bool, last bool) string {
				return strings.ReplaceAll(segment, "|", "\\|")
			},
			func(segment string, first bool, last bool) string {
				return fmt.Sprintf("`%s`", segment)
			},
		)
	}

	if escape {
		s = processSegments(
			s,
			"`",
			func(segment string, first bool, last bool) string {
				return executePerLine(segment, func(line string) string {
					escape := func(char string) {
						c := strings.ReplaceAll(char, "*", "\\*")
						cases := []struct {
							pattern string
							index   []int
						}{
							{
								pattern: `^(\s*)(` + c + `+)(\s+)(.*)`,
								index:   []int{2},
							},
							{
								pattern: `(\s+)(` + c + `+)([^\t\n\f\r ` + c + `])(.*)([^\t\n\f\r ` + c + `])(` + c + `+)(\s+)`,
								index:   []int{6, 2},
							},
						}
						for i := range cases {
							c := cases[i]
							r := regexp.MustCompile(c.pattern)
							m := r.FindAllStringSubmatch(line, -1)
							i := r.FindAllStringSubmatchIndex(line, -1)
							for j := range m {
								for _, k := range c.index {
									line = line[:i[j][k*2]] + strings.ReplaceAll(m[j][k], char, "‡‡‡DONTESCAPE‡‡‡") + line[i[j][(k*2)+1]:]
								}
							}
						}
						line = strings.ReplaceAll(line, char, "\\"+char)
						line = strings.ReplaceAll(line, "‡‡‡DONTESCAPE‡‡‡", char)
					}
					escape("_") // Escape underscore
					return line
				})
			},
			func(segment string, first bool, last bool) string {
				segment = fmt.Sprintf("`%s`", segment)
				return segment
			},
		)
	}

	return s
}

// NormalizeURLs runs after escape function and normalizes URL back to the original
// state. For example any underscore in the URL which got escaped by 'EscapeCharacters'
// will be reverted back.
func NormalizeURLs(s string, escape bool) string {
	if escape {
		if urls := xurls.Strict().FindAllString(s, -1); len(urls) > 0 {
			for _, url := range urls {
				normalized := strings.ReplaceAll(url, "\\", "")
				s = strings.ReplaceAll(s, url, normalized)
			}
		}
	}
	return s
}

type segmentCallbackFn func(string, bool, bool) string

func processSegments(s string, prefix string, normalFn segmentCallbackFn, codeFn segmentCallbackFn) string {
	// Isolate blocks of code. Dont escape anything inside them
	nextIsInCodeBlock := strings.HasPrefix(s, prefix)
	segments := strings.Split(s, prefix)
	buffer := bytes.NewBufferString("")
	for i, segment := range segments {
		if len(segment) == 0 {
			continue
		}

		first := i == 0 || len(strings.TrimSpace(segments[i-1])) == 0
		last := i == len(segments)-1 || len(strings.TrimSpace(segments[i+1])) == 0

		if !nextIsInCodeBlock {
			segment = normalFn(segment, first, last)
		} else {
			segment = codeFn(segment, first, last)
		}
		buffer.WriteString(segment)
		nextIsInCodeBlock = !nextIsInCodeBlock
	}
	return buffer.String()
}

func executePerLine(s string, fn func(string) string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = fn(l)
	}
	return strings.Join(lines, "\n")
}
