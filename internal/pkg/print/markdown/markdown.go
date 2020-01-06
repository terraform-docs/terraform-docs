package markdown

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

// Sanitize cleans a Markdown document to soothe linters.
func Sanitize(markdown string) string {
	result := markdown

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(` {2}(\r?\n)`).ReplaceAllString(result, "‡‡$1")

	// Remove trailing spaces from the end of lines
	result = regexp.MustCompile(` +(\r?\n)`).ReplaceAllString(result, "$1")
	result = regexp.MustCompile(` +$`).ReplaceAllLiteralString(result, "")

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(`‡‡(\r?\n)`).ReplaceAllString(result, "  $1")

	// Remove multiple consecutive blank lines
	result = regexp.MustCompile(`(\r?\n){3,}`).ReplaceAllString(result, "$1$1")
	result = regexp.MustCompile(`(\r?\n){2,}$`).ReplaceAllString(result, "$1")

	return result
}

// SanitizeName escapes underscore character which have special meaning in Markdown.
func SanitizeName(name string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		// Escape underscore
		name = strings.Replace(name, "_", "\\_", -1)
	}
	return name
}

// SanitizeItemForDocument converts passed 'string' to suitable Markdown representation
// for a document. (including line-break, illegal characters, code blocks etc)
func SanitizeItemForDocument(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"\n```",
		func(segment string) string {
			segment = ConvertMultiLineText(segment, false)
			segment = EscapeIllegalCharacters(segment, settings)
			segment = fmt.Sprintf("%s\n", segment)
			return segment
		},
		func(segment string) string {
			segment = fmt.Sprintf("\n```%s\n```", segment)
			return segment
		},
	)
	return strings.Replace(result, "<br>", "\n", -1)
}

// SanitizeItemForTable converts passed 'string' to suitable Markdown representation
// for a table. (including line-break, illegal characters, code blocks etc)
func SanitizeItemForTable(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```\n",
		func(segment string) string {
			segment = ConvertMultiLineText(segment, true)
			segment = EscapeIllegalCharacters(segment, settings)
			return segment
		},
		func(segment string) string {
			segment = fmt.Sprintf("<code><pre>%s</pre></code>", strings.Replace(strings.Replace(segment, "\n", "<br>", -1), "\r", "", -1))
			return segment
		},
	)
	return result
}

// ConvertMultiLineText converts a multi-line text into a suitable Markdown representation.
func ConvertMultiLineText(s string, convertDoubleSpaces bool) string {
	// Convert double newlines to <br><br>.
	s = strings.Replace(
		strings.TrimSpace(s),
		"\n\n",
		"<br><br>",
		-1,
	)

	if convertDoubleSpaces {
		// Convert space-space-newline to <br>
		s = strings.Replace(s, "  \n", "<br>", -1)

		// Convert single newline to space.
		s = strings.Replace(s, "\n", " ", -1)
	}

	return s
}

// EscapeIllegalCharacters escapes characters which have special meaning in Markdown into their corresponding literal.
func EscapeIllegalCharacters(s string, settings *print.Settings) string {
	// Escape pipe
	s = strings.Replace(s, "|", "\\|", -1)

	if settings.EscapeCharacters {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				// Escape underscore
				segment = strings.Replace(segment, "_", "\\_", -1)
				// Escape asterisk
				segment = strings.Replace(segment, "*", "\\*", -1)
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

// GenerateIndentation generates indentation of Markdown headers
// with base level of provided 'settings.MarkdownIndent' plus any
// extra level needed for subsection (e.g. 'Required Inputs' which
// is a subsection of 'Inputs' section)
func GenerateIndentation(extra int, settings *print.Settings) string {
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

// PrintFencedCodeBlock prints codes in fences, it automatically detects if
// the input 'code' contains '\n' it will use multi line fence, otherwise it
// wraps the 'code' inside single-tick block.
// If the fenced is multi-line it also appens an extra '\n` at the end and
// returns true accordingly, otherwise returns false for non-carriage return.
func PrintFencedCodeBlock(code string, language string) (string, bool) {
	if strings.Contains(code, "\n") {
		return fmt.Sprintf("\n\n```%s\n%s\n```\n", language, code), true
	}
	return fmt.Sprintf("`%s`", code), false
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
