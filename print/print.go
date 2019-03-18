package print

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/coveo/gotemplate/hcl"
	"github.com/fatih/color"
	"github.com/segmentio/terraform-docs/doc"
	yaml "gopkg.in/yaml.v2"
)

// RenderMode represents the mode used to render the results
type RenderMode int

const (
	renderNone RenderMode = iota
	// RenderInputs = Only render Inputs
	RenderInputs
	// RenderOutputs = Only render Outputs
	RenderOutputs
	// RenderAll = render both Inputs & Outputs
	RenderAll
	// RenderDetailed = render content of lists & maps
	RenderDetailed
)

//var dimmed = color.New(color.FgHiWhite, color.Bold, color.Faint, color.Italic).Sprintf
var dimmed = color.New(color.FgHiBlack, color.Bold, color.Italic).Sprintf

func varStr(name string) string { return color.CyanString(fmt.Sprintf("var.%s", name)) }
func outStr(name string) string { return color.CyanString(fmt.Sprintf("output.%s", name)) }

func valStr(value interface{}) string {
	if value == nil {
		return ""
	}
	return color.WhiteString(fmt.Sprintf(" (%v)", value))
}

func desStr(description string) string {
	if description == "" {
		return ""
	}
	return fmt.Sprintf("\n%s", dimmed(description))
}

// Pretty printer pretty prints a doc.
func Pretty(d *doc.Doc, mode RenderMode) (string, error) {
	var buf bytes.Buffer

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("\n%s\n", d.Comment))
	}

	if mode&RenderInputs != 0 && len(d.Inputs) > 0 {
		for _, i := range d.Inputs {
			value := i.Value()
			if mode&RenderDetailed != 0 {
				hcl, err := hcl.Marshal(i.Default.Literal)
				if err != nil {
					return "", err
				}
				value = string(hcl)
			}
			buf.WriteString(fmt.Sprintf("%s%s%s\n\n", varStr(i.Name), valStr(value), desStr(i.Description)))
		}
	}

	if mode&RenderOutputs != 0 && len(d.Outputs) > 0 {
		for _, o := range d.Outputs {
			buf.WriteString(fmt.Sprintf("%s%s%s\n\n", outStr(o.Name), valStr(o), desStr(o.Description)))
		}
	}

	return buf.String(), nil
}

// Markdown prints the given doc as markdown.
func Markdown(d *doc.Doc, mode RenderMode, printRequired, printValues bool) (string, error) {
	var buf bytes.Buffer

	if len(d.Comment) > 0 {
		buf.WriteString(fmt.Sprintf("%s\n", d.Comment))
	}

	if mode&RenderInputs != 0 {
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
				def = fmt.Sprintf("`%s `", def)
			}

			buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s |", v.Name, normalizeMarkdownDesc(v.Description), v.Type, normalizeMarkdownDesc(def)))

			if printRequired {
				buf.WriteString(fmt.Sprintf(" %v |\n", humanize(v.Default == nil)))
			} else {
				buf.WriteString("\n")
			}
		}
	}

	if mode&RenderOutputs != 0 {
		if len(d.Outputs) > 0 {
			var ext, sep string
			if printValues {
				ext = " Value | Type | Sensitive |"
				sep = "-------|------|-----------|"
			}
			buf.WriteString("\n## Outputs\n\n")
			buf.WriteString(fmt.Sprintf("| Name | Description |%s\n", ext))
			buf.WriteString(fmt.Sprintf("|------|-------------|%s\n", sep))
		}

		for _, v := range d.Outputs {
			var val string
			if printValues {
				val = fmt.Sprintf(" `%v ` | %s | %s |", v.Result.Value, v.Result.Type, humanize(v.Result.Sensitive))
			}
			buf.WriteString(fmt.Sprintf("| %s | %s |%s\n", v.Name, normalizeMarkdownDesc(v.Description), val))
		}
	}

	return buf.String(), nil
}

// TerraformOutput prints the given doc as 'terraform output -json'.
func TerraformOutput(d *doc.Doc, mode RenderMode) (string, error) {
	jsonOutput := make(map[string]doc.Result)
	for i := range filter(*d, mode).Outputs {
		o := &d.Outputs[i]
		jsonOutput[o.Name] = o.Result
	}
	result, err := json.MarshalIndent(jsonOutput, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// JSON prints the given doc as json.
func JSON(d *doc.Doc, mode RenderMode) (string, error) {
	result, err := json.MarshalIndent(filter(*d, mode), "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// YAML prints the given doc as yaml.
func YAML(d *doc.Doc, mode RenderMode) (string, error) {
	result, err := yaml.Marshal(filter(*d, mode))
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// XML prints the given doc as xml.
func XML(d *doc.Doc, mode RenderMode) (string, error) {
	data := filter(*d, mode)

	for i := range data.Outputs {
		output := &data.Outputs[i]
		if output.Result.Value != nil {
			switch value := output.Result.Value.(type) {
			case map[string]interface{}:
				output.Result.Value, _ = hcl.Marshal(value)
			}
		}
	}
	result, err := xml.MarshalIndent(data, "", "  ")

	if err != nil {
		return "", err
	}

	return string(result), nil
}

// HCL prints the given doc as hcl.
func HCL(d *doc.Doc, mode RenderMode) (string, error) {
	result, err := hcl.MarshalIndent(filter(*d, mode), "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func filter(d doc.Doc, mode RenderMode) doc.Doc {
	switch mode & RenderAll {
	case RenderInputs:
		d.Outputs = nil
	case RenderOutputs:
		d.Inputs = nil
	}
	return d
}

// Humanize the given boolean value.
func humanize(value bool) string {
	return map[bool]string{
		true:  "yes",
		false: "no",
	}[value]
}

// normalizeMarkdownDesc fixes line breaks in descriptions for markdown:
//
//  * Double newlines are converted to <br><br>
//  * A second pass replaces all other newlines with spaces
func normalizeMarkdownDesc(s string) string {
	return strings.Replace(strings.Replace(strings.TrimSpace(s), "\n\n", "<br><br>", -1), "\n", " ", -1)
}
