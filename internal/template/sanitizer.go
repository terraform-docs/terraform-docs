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

	"mvdan.cc/xurls/v2"

	"github.com/terraform-docs/terraform-docs/internal/print"
)

// sanitizeName escapes underscore character which have special meaning in Markdown.
func sanitizeName(name string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		// Escape underscore
		name = strings.Replace(name, "_", "\\_", -1)
	}
	return name
}

// sanitizeHeader converts passed 'string' to suitable Markdown representation
// for a document. (including line-break, illegal characters, code blocks etc)
//
// IMPORTANT: sanitizeHeader will never change the line-endings and preserve them
// as they are provided by the users.
func sanitizeHeader(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = escapeIllegalCharacters(segment, settings, false)
			segment = convertMultiLineText(segment, false, true)
			segment = normalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
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

// sanitizeDocument converts passed 'string' to suitable Markdown or AsciiDoc
// representation for a document. (including line-break, illegal characters,
// code blocks etc)
func sanitizeDocument(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = escapeIllegalCharacters(segment, settings, false)
			segment = convertMultiLineText(segment, false, false)
			segment = normalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
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

// sanitizeMarkdownTable converts passed 'string' to suitable Markdown representation
// for a table. (including line-break, illegal characters, code blocks etc)
func sanitizeMarkdownTable(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = escapeIllegalCharacters(segment, settings, true)
			segment = convertMultiLineText(segment, true, false)
			segment = normalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
			segment = strings.TrimSpace(segment)
			segment = strings.Replace(segment, "\n", "<br>", -1)
			segment = strings.Replace(segment, "\r", "", -1)
			segment = fmt.Sprintf("<pre>%s</pre>", segment)
			return segment
		},
	)
	return result
}

// sanitizeAsciidocTable converts passed 'string' to suitable AsciiDoc representation
// for a table. (including line-break, illegal characters, code blocks etc)
func sanitizeAsciidocTable(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = escapeIllegalCharacters(segment, settings, true)
			segment = normalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
			segment = strings.TrimSpace(segment)
			segment = fmt.Sprintf("[source]\n----\n%s\n----", segment)
			return segment
		},
	)
	return result
}

// convertMultiLineText converts a multi-line text into a suitable Markdown representation.
func convertMultiLineText(s string, isTable bool, isHeader bool) string {
	if isTable {
		s = strings.TrimSpace(s)
	}

	// Convert double newlines to <br><br>.
	s = strings.Replace(s, "\n\n", "<br><br>", -1)

	// Convert line-break on a non-empty line followed by another line
	// starting with "alphanumeric" word into space-space-newline
	// which is a know convention of Markdown for multi-lines paragprah.
	// This doesn't apply on a markdown list for example, because all the
	// consecutive lines start with hyphen which is a special character.
	if !isHeader {
		s = regexp.MustCompile(`(\S*)(\r?\n)(\s*)(\w+)`).ReplaceAllString(s, "$1  $2$3$4")
		s = strings.Replace(s, "    \n", "  \n", -1)
		s = strings.Replace(s, "<br>  \n", "\n\n", -1)
	}

	if isTable {
		// Convert space-space-newline to <br>
		s = strings.Replace(s, "  \n", "<br>", -1)

		// Convert single newline to <br>.
		s = strings.Replace(s, "\n", "<br>", -1)
	} else {
		s = strings.Replace(s, "<br>", "\n", -1)
	}

	return s
}

// escapeIllegalCharacters escapes characters which have special meaning in Markdown into their corresponding literal.
func escapeIllegalCharacters(s string, settings *print.Settings, escapePipe bool) string {
	// Escape pipe (only for 'markdown table' or 'asciidoc table')
	if escapePipe {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				return strings.Replace(segment, "|", "\\|", -1)
			},
			func(segment string) string {
				return fmt.Sprintf("`%s`", segment)
			},
		)
	}

	if settings.EscapeCharacters {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				return executePerLine(segment, func(line string) string {
					escape := func(char string) {
						c := strings.Replace(char, "*", "\\*", -1)
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
						for _, c := range cases {
							r := regexp.MustCompile(c.pattern)
							m := r.FindAllStringSubmatch(line, -1)
							i := r.FindAllStringSubmatchIndex(line, -1)
							for j := range m {
								for _, k := range c.index {
									line = line[:i[j][k*2]] + strings.Replace(m[j][k], char, "‡‡‡DONTESCAPE‡‡‡", -1) + line[i[j][(k*2)+1]:]
								}
							}
						}
						line = strings.Replace(line, char, "\\"+char, -1)
						line = strings.Replace(line, "‡‡‡DONTESCAPE‡‡‡", char, -1)
					}
					escape("_") // Escape underscore
					return line
				})
			},
			func(segment string) string {
				segment = fmt.Sprintf("`%s`", segment)
				return segment
			},
		)
	}

	return s
}

// normalizeURLs runs after escape function and normalizes URL back
// to the original state. For example any underscore in the URL which
// got escaped by 'EscapeIllegalCharacters' will be reverted back.
func normalizeURLs(s string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		if urls := xurls.Strict().FindAllString(s, -1); len(urls) > 0 {
			for _, url := range urls {
				normalized := strings.Replace(url, "\\", "", -1)
				s = strings.Replace(s, url, normalized, -1)
			}
		}
	}
	return s
}

func processSegments(s string, prefix string, normalFn func(segment string) string, codeFn func(segment string) string) string {
	// Isolate blocks of code. Dont escape anything inside them
	nextIsInCodeBlock := strings.HasPrefix(s, prefix)
	segments := strings.Split(s, prefix)
	buffer := bytes.NewBufferString("")
	for _, segment := range segments {
		if len(segment) == 0 {
			continue
		}
		if !nextIsInCodeBlock {
			segment = normalFn(segment)
		} else {
			segment = codeFn(segment)
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
