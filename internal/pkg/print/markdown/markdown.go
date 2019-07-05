package markdown

import (
	"regexp"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// SanitizeName escapes underscore character which have special meaning in Markdown.
func SanitizeName(s string, settings settings.Settings) string {
	if settings.EscapeMarkdown {
		// Escape underscore
		s = strings.Replace(s, "_", "\\_", -1)
	}

	return s
}

// SanitizeDescription converts description to suitable Markdown representation. (including line-break, illegal characters, etc)
func SanitizeDescription(s string, settings settings.Settings) string {
	s = ConvertMultiLineText(s)
	s = EscapeIllegalCharacters(s, settings)

	return s
}

// ConvertMultiLineText converts a multi-line text into a suitable Markdown representation.
func ConvertMultiLineText(s string) string {
	// Convert double newlines to <br><br>.
	s = strings.Replace(
		strings.TrimSpace(s),
		"\n\n",
		"<br><br>",
		-1)

	// Convert single newline to space.
	return strings.Replace(s, "\n", " ", -1)
}

// EscapeIllegalCharacters escapes characters which have special meaning in Markdown into their corresponding literal.
func EscapeIllegalCharacters(s string, settings settings.Settings) string {
	// Escape pipe
	s = strings.Replace(s, "|", "\\|", -1)

	if settings.EscapeMarkdown {
		// Escape underscore
		s = strings.Replace(s, "_", "\\_", -1)

		// Escape asterisk
		s = strings.Replace(s, "*", "\\*", -1)

		// Escape paranthesis
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
