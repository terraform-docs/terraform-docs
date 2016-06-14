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

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("\n%s\n", d.Comment))
	}

	if len(d.Inputs) > 0 {
		buf.WriteString("\n")

		for _, i := range d.Inputs {
			format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
			desc := i.Description
			def := i.Default

			if def == "" {
				def = "required"
			}

			if desc == "" {
				desc = "-"
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
func Markdown(d *doc.Doc) (string, error) {
	var buf bytes.Buffer

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("%s\n", d.Comment))
	}

	if len(d.Inputs) > 0 {
		buf.WriteString("\n## Inputs\n\n")
		buf.WriteString("| Name | Description | Default | Required |\n")
		buf.WriteString("|------|-------------|:-----:|:-----:|\n")
	}

	for _, v := range d.Inputs {
		def := v.Default

		if def == "" {
			def = "-"
		} else {
			def = fmt.Sprintf("`%s`", def)
		}

		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %v |\n",
			v.Name,
			v.Description,
			def,
			humanize(v.Default == "")))
	}

	if len(d.Outputs) > 0 {
		buf.WriteString("\n## Outputs\n\n")
		buf.WriteString("| Name | Description |\n")
		buf.WriteString("|------|-------------|\n")
	}

	for _, v := range d.Outputs {
		buf.WriteString(fmt.Sprintf("| %s | %s |\n",
			v.Name,
			strings.TrimSpace(v.Description)))
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

// Humanize the given `v`.
func humanize(v interface{}) string {
	switch v.(type) {
	case bool:
		if v.(bool) {
			return "yes"
		}
		return "no"
	default:
		panic("unknown type")
	}
}
