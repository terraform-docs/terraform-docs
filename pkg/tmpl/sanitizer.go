package tmpl

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/segmentio/terraform-docs/pkg/print"
	"mvdan.cc/xurls/v2"
)

// sanitizeName escapes underscore character which have special meaning in Markdown.
func sanitizeName(name string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		// Escape underscore
		name = strings.Replace(name, "_", "\\_", -1)
	}
	return name
}

// sanitizeItemForDocument converts passed 'string' to suitable Markdown representation
// for a document. (including line-break, illegal characters, code blocks etc)
func sanitizeItemForDocument(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = convertMultiLineText(segment, false)
			segment = escapeIllegalCharacters(segment, settings)
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

// sanitizeItemForTable converts passed 'string' to suitable Markdown representation
// for a table. (including line-break, illegal characters, code blocks etc)
func sanitizeItemForTable(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = convertMultiLineText(segment, true)
			segment = escapeIllegalCharacters(segment, settings)
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

// convertMultiLineText converts a multi-line text into a suitable Markdown representation.
func convertMultiLineText(s string, isTable bool) string {
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
	s = regexp.MustCompile(`(\S*)(\r?\n)(\s*)(\w+)`).ReplaceAllString(s, "$1  $2$3$4")
	s = strings.Replace(s, "    \n", "  \n", -1)
	s = strings.Replace(s, "<br>  \n", "\n\n", -1)

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
func escapeIllegalCharacters(s string, settings *print.Settings) string {
	// Escape pipe
	if settings.EscapePipe {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				segment = strings.Replace(segment, "|", "\\|", -1)
				return segment
			},
			func(segment string) string {
				segment = fmt.Sprintf("`%s`", segment)
				return segment
			},
		)
	}

	if settings.EscapeCharacters {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				escape := func(char string) {
					segment = strings.Replace(segment, char+char, "‡‡", -1)
					segment = strings.Replace(segment, " "+char, " ‡", -1)
					segment = strings.Replace(segment, char+" ", "‡ ", -1)
					segment = strings.Replace(segment, char, "\\"+char, -1)
					segment = strings.Replace(segment, "‡", char, -1)
				}
				// Escape underscore
				escape("_")
				// Escape asterisk
				escape("*")
				return segment
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

// generateIndentation generates indentation of Markdown headers
// with base level of provided 'settings.MarkdownIndent' plus any
// extra level needed for subsection (e.g. 'Required Inputs' which
// is a subsection of 'Inputs' section)
func generateIndentation(extra int, settings *print.Settings) string {
	var base = settings.MarkdownIndent
	if base < 1 || base > 5 {
		base = 2
	}
	var indent string
	for i := 0; i < base+extra; i++ {
		indent += "#"
	}
	return indent
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
