package print

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/getcloudnative/terraform-docs/doc"
)

// Pretty printer pretty prints a doc.
func Pretty(d *doc.Doc) (string, error) {
	var buf bytes.Buffer

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("\n%s\n", d.Comment))
	}

	if len(d.Inputs) > 0 {
		buf.WriteString("\n")

		for _, i := range d.Inputs {
			format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"

			desc := i.Description
			if desc == "" {
				desc = "-"
			}

			def, err := getDefaultValueFromInput(i)
			if err != nil {
				return "", err
			}

			buf.WriteString(fmt.Sprintf(format, i.Name, def, desc))
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

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("%s\n", d.Comment))
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
		var def string
		if v.Value() == "required" {
			def = "-"
		} else {
			value, err := getDefaultValueFromInput(v)
			if err != nil {
				return "", err
			}

			def = fmt.Sprintf("`%s`", value)
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

	return string(s), nil
}

func getDefaultValueFromInput(input doc.Input) (string, error) {
	var result string

	if input.Type != "string" {
		json, err := json.MarshalIndent(input.Value(), "", "")
		if err != nil {
			return "", err
		}

		result = strings.Replace(string(json), "\n", " ", -1)
	} else {
		result = input.Value().(string)
	}

	return result, nil
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
