package print

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/doc"
)

// Pretty printer pretty prints a doc.
func Pretty(d *doc.Doc) (string, error) {
	var buf bytes.Buffer

	if len(d.Version) > 0 {
		format := "  \033[36mterraform.required_version\033[0m (%s)\n\n\n"
		buf.WriteString(fmt.Sprintf(format, d.Version))
	}

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("\n%s\n", d.Comment))
	}

	if len(d.Providers) > 0 {
		buf.WriteString("\n")

		for _, i := range d.Providers {
			format := "  \033[36mprovider.%s\033[0m\n  \033[90m%s\033[0m\n\n"
			buf.WriteString(fmt.Sprintf(format, i.Name, i.Documentation))
		}

		buf.WriteString("\n")
	}

	if len(d.Inputs) > 0 {
		buf.WriteString("\n")

		for _, i := range d.Inputs {
			format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
			desc := i.Description

			if desc == "" {
				desc = "-"
			}

			buf.WriteString(fmt.Sprintf(format, i.Name, i.Value(), desc))
		}

		buf.WriteString("\n")
	}

	if len(d.Outputs) > 0 {
		buf.WriteString("\n")

		for _, i := range d.Outputs {
			format := "  \033[36moutput.%s\033[0m\n  \033[90m%s\033[0m\n\n"
			s := fmt.Sprintf(format, i.Name, strings.TrimSpace(i.Description))
			buf.WriteString(s)
		}

		buf.WriteString("\n")
	}

	return buf.String(), nil
}

// Markdown prints the given doc as markdown.
func Markdown(d *doc.Doc, printRequired bool) (string, error) {
	var buf bytes.Buffer

	if len(d.Version) > 0 {
		buf.WriteString(fmt.Sprintf("Terraform required version %s\n", d.Version))
	}

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("%s\n", d.Comment))
	}

	if len(d.Providers) > 0 {
		buf.WriteString("\n## Providers\n\n")
		buf.WriteString("| Name | Documentation |\n")
		buf.WriteString("|------|---------------|\n")

		for _, i := range d.Providers {
			format := "| %s | [%s](%s) |\n"
			buf.WriteString(fmt.Sprintf(format, i.Name, i.Documentation, i.Documentation))
		}

		buf.WriteString("\n")
	}

	if len(d.Inputs) > 0 {
		buf.WriteString("\n## Inputs\n\n")
		buf.WriteString("| Name | Description | Type | Default |")

		if printRequired {
			buf.WriteString(" Required |\n")
		} else {
			buf.WriteString("\n")
		}

		buf.WriteString("|------|-------------|:----:|:-----:|")
		if printRequired {
			buf.WriteString(":-----:|\n")
		} else {
			buf.WriteString("\n")
		}
	}

	for _, v := range d.Inputs {
		def := v.Value()

		if def == "required" {
			def = "-"
		} else {
			def = fmt.Sprintf("`%s`", def)
		}

		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s |",
			v.Name,
			normalizeMarkdownDesc(v.Description),
			v.Type,
			normalizeMarkdownDesc(def)))

		if printRequired {
			buf.WriteString(fmt.Sprintf(" %v |\n",
				humanize(v.Default)))
		} else {
			buf.WriteString("\n")
		}
	}

	if len(d.Outputs) > 0 {
		buf.WriteString("\n## Outputs\n\n")
		buf.WriteString("| Name | Description |\n")
		buf.WriteString("|------|-------------|\n")
	}

	for _, v := range d.Outputs {
		buf.WriteString(fmt.Sprintf("| %s | %s |\n",
			v.Name,
			normalizeMarkdownDesc(v.Description)))
	}

	return buf.String(), nil
}

// JSON prints the given doc as json.
func JSON(d *doc.Doc) (string, error) {
	s, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}

	// <, >, and & are printed as code points by the json package.
	// The brackets are needed to pretty-print required_version.
	// Convert the brackets back into printable chars, limiting the
	// number of printed brackets to 1 each, which is enough to
	// prevent HTML injection (json's concern - why they encode).
	jsonString := strings.Replace(string(s), "\\u003c", "<", 1)
	jsonString = strings.Replace(jsonString, "\\u003e", ">", 1)

	return jsonString, nil
}

// Humanize the given `v`.
func humanize(def *doc.Value) string {
	if def == nil {
		return "yes"
	}

	return "no"
}

// normalizeMarkdownDesc fixes line breaks in descriptions for markdown:
//
//  * Double newlines are converted to <br><br>
//  * A second pass replaces all other newlines with spaces
func normalizeMarkdownDesc(s string) string {
	return strings.Replace(strings.Replace(strings.TrimSpace(s), "\n\n", "<br><br>", -1), "\n", " ", -1)
}
