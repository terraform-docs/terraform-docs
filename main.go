package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/tj/docopt"
)

const version = ""
const usage = `
  Usage:
    tf-docs <dir>
    tf-docs md <dir>
    tf-docs -h | --help

  Examples:

    # Generate a JSON of inputs and outputs
    $ tf-docs ./my-module

    # Generate markdown tables of inputs and outputs
    $ tf-docs md ./my-module

  Options:
    -h, --help     show help information

`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal(err)
	}

	dir := args["<dir>"].(string)
	names, err := filepath.Glob(fmt.Sprintf("%s/*.tf", dir))
	if err != nil {
		log.Fatal(err)
	}

	var vals []Value

	for _, name := range names {
		buf, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		f, err := hcl.ParseBytes(buf)
		if err != nil {
			log.Fatal(err)
		}

		vals = append(vals, values(f)...)
	}

	if args["md"].(bool) {
		markdown(vals)
		return
	}

	buf, err := json.MarshalIndent(vals, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}

type Value struct {
	Type        string
	Name        string
	Value       string
	Default     string
	Description string
}

func markdown(vals []Value) {
	var inputs bytes.Buffer
	var outputs bytes.Buffer

	inputs.WriteString("| Name | Description | Default | Required |\n")
	inputs.WriteString("|------|-------------|:-----:|:-----:|\n")
	outputs.WriteString("| Name | Description |\n")
	outputs.WriteString("|------|-------------|\n")

	for _, v := range vals {
		if v.Type == "input" {
			def := v.Default

			if def == "" {
				def = "-"
			} else {
				def = fmt.Sprintf("`%s`", def)
			}

			inputs.WriteString(fmt.Sprintf("| %s | %s | %s | %v |\n",
				v.Name,
				v.Description,
				def,
				humanize(v.Default == "")))
		} else {
			outputs.WriteString(fmt.Sprintf("| %s | %s |\n",
				v.Name,
				strings.TrimSpace(v.Description)))
		}
	}

	fmt.Println("## Inputs")
	fmt.Println(inputs.String())
	fmt.Println("## Outputs")
	fmt.Println(outputs.String())
}

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

func values(f *ast.File) (ret []Value) {
	list := f.Node.(*ast.ObjectList)

	for _, n := range list.Items {
		if is(n.Keys, "variable") {
			name := n.Keys[1].Token.Text
			items := n.Val.(*ast.ObjectType).List.Items

			ret = append(ret, Value{
				Type:        "input",
				Name:        clean(name),
				Description: clean(get(items, "description")),
				Default:     get(items, "default"),
			})
			continue
		}

		if is(n.Keys, "output") {
			name := n.Keys[1].Token.Text
			items := n.Val.(*ast.ObjectType).List.Items
			value := get(items, "value")
			desc := ""

			if c := n.LeadComment; c != nil {
				desc = comment(c.List)
			}

			ret = append(ret, Value{
				Type:        "output",
				Name:        clean(name),
				Description: desc,
				Value:       value,
			})
			continue
		}
	}

	return ret
}

func get(items []*ast.ObjectItem, key string) string {
	for _, item := range items {
		if is(item.Keys, key) {
			if lit, ok := item.Val.(*ast.LiteralType); ok {
				return lit.Token.Text
			}

			return ""
		}
	}

	return ""
}

func is(keys []*ast.ObjectKey, t string) bool {
	if len(keys) > 0 {
		return keys[0].Token.Text == t
	}

	return false
}

func comment(l []*ast.Comment) string {
	var line string
	var ret string

	for _, t := range l {
		line = strings.TrimSpace(t.Text)
		line = strings.TrimPrefix(line, "//")
		ret += strings.TrimSpace(line) + "\n"
	}

	return ret
}

func clean(s string) string {
	if len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return ""
}
