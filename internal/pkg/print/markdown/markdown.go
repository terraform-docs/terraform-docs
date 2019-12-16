package markdown

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// SanitizeName escapes underscore character which have special meaning in Markdown.
func SanitizeName(s string, settings *settings.Settings) string {
	if settings.EscapeMarkdown {
		// Escape underscore
		s = strings.Replace(s, "_", "\\_", -1)
	}

	return s
}

// SanitizeDescriptionForDocument converts description to suitable Markdown representation for a document. (including line-break, illegal characters, code blocks etc)
func SanitizeDescriptionForDocument(s string, settings *settings.Settings) string {
	// s = ConvertMultiLineText(s)
	// s = EscapeIllegalCharacters(s, settings)
	// return s
	// Isolate blocks of code. Dont escape anything inside them
	nextIsInCodeBlock := strings.HasPrefix(s, "```\n")
	segments := strings.Split(s, "\n```\n")
	buf := bytes.NewBufferString("")
	for i, segment := range segments {
		if !nextIsInCodeBlock {
			segment = ConvertMultiLineText(segment)
			segment = EscapeIllegalCharacters(segment, settings)
			if i > 0 && len(segment) > 0 {
				buf.WriteString("<br>```<br>")
			}
			buf.WriteString(segment)
			nextIsInCodeBlock = true
		} else {
			buf.WriteString("<br>```<br>")
			buf.WriteString(segment)
			buf.WriteString("<br>```")
			nextIsInCodeBlock = false
		}
	}
	return buf.String()
}

// SanitizeDescriptionForTable converts description to suitable Markdown representation for a table. (including line-break, illegal characters, code blocks etc)
func SanitizeDescriptionForTable(s string, settings *settings.Settings) string {
	// Isolate blocks of code. Dont escape anything inside them
	nextIsInCodeBlock := strings.HasPrefix(s, "```\n")
	segments := strings.Split(s, "```\n")
	buf := bytes.NewBufferString("")
	for _, segment := range segments {
		if !nextIsInCodeBlock {
			segment = ConvertMultiLineText(segment)
			segment = EscapeIllegalCharacters(segment, settings)
			buf.WriteString(segment)
			nextIsInCodeBlock = true
		} else {
			buf.WriteString("<code><pre>")
			buf.WriteString(strings.Replace(strings.Replace(segment, "\n", "<br>", -1), "\r", "", -1))
			buf.WriteString("</pre></code>")
			nextIsInCodeBlock = false
		}
	}

	return buf.String()
}

// ConvertMultiLineText converts a multi-line text into a suitable Markdown representation.
func ConvertMultiLineText(s string) string {

	// Convert double newlines to <br><br>.
	s = strings.Replace(
		strings.TrimSpace(s),
		"\n\n",
		"<br><br>",
		-1)

	// Convert space-space-newline to <br>
	s = strings.Replace(s, "  \n", "<br>", -1)

	// Convert single newline to space.
	return strings.Replace(s, "\n", " ", -1)
}

// EscapeIllegalCharacters escapes characters which have special meaning in Markdown into their corresponding literal.
func EscapeIllegalCharacters(s string, settings *settings.Settings) string {
	// Escape pipe
	s = strings.Replace(s, "|", "\\|", -1)

	if settings.EscapeMarkdown {
		// Escape underscore
		s = strings.Replace(s, "_", "\\_", -1)

		// Escape asterisk
		s = strings.Replace(s, "*", "\\*", -1)

		// Escape parenthesis
		s = strings.Replace(s, "(", "\\(", -1)
		s = strings.Replace(s, ")", "\\)", -1)

		// Escape brackets
		s = strings.Replace(s, "[", "\\[", -1)
		s = strings.Replace(s, "]", "\\]", -1)

		// Escape curly brackets
		s = strings.Replace(s, "{", "\\{", -1)
		s = strings.Replace(s, "}", "\\}", -1)
	}

	return s
}

// Sanitize cleans a Markdown document to soothe linters.
func Sanitize(markdown string) string {
	result := markdown

	// Remove trailing spaces from the end of lines
	result = regexp.MustCompile(` +(\r?\n)`).ReplaceAllString(result, "$1")
	result = regexp.MustCompile(` +$`).ReplaceAllLiteralString(result, "")

	// Remove multiple consecutive blank lines
	result = regexp.MustCompile(`(\r?\n){3,}`).ReplaceAllString(result, "$1$1")
	result = regexp.MustCompile(`(\r?\n){2,}$`).ReplaceAllString(result, "$1")

	return result
}

// GenerateIndentation generates indentation of Markdown headers
// with base level of provided 'settings.MarkdownIndent' plus any
// extra level needed for subsection (e.g. 'Required Inputs' which
// is a subsection of 'Inputs' section)
func GenerateIndentation(extra int, settings *settings.Settings) string {
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

func PrintCode(code string, language string) string {
	if strings.Contains(code, "\n") {
		return fmt.Sprintf("```%s\n%s\n```\n", language, code)
	}
	return fmt.Sprintf("`%s`", code)
}
