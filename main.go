package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/tj/docopt"
)

const version = ""
const usage = `
  Usage: tf-docs <dir>
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatal(err)
	}

	dir := args["<dir>"].(string)
	names, err := filepath.Glob(fmt.Sprintf("%s/*.tf", dir))
	if err != nil {
		log.Fatal(err)
	}

	var values []Value

	for _, name := range names {
		buf, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		f, err := hcl.ParseBytes(buf)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, inputs(f)...)
	}

	buf, err := json.MarshalIndent(values, "", "  ")
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

func inputs(f *ast.File) (ret []Value) {
	list := f.Node.(*ast.ObjectList)

	for _, n := range list.Items {
		if is(n.Keys, "variable") {
			name := n.Keys[1].Token.Text
			vals := n.Val.(*ast.ObjectType)

			var defaultText string
			var description string

			if items := vals.List.Items; len(items) > 0 {
				for _, item := range items {
					if is(item.Keys, "default") {
						defaultText = item.Val.(*ast.LiteralType).Token.Text
						break
					}

					if is(item.Keys, "description") {
						description = item.Val.(*ast.LiteralType).Token.Text
						break
					}

					if is(item.Keys, "value") {
						description = item.Val.(*ast.LiteralType).Token.Text
						break
					}
				}
			}

			ret = append(ret, Value{
				Type:        "variable",
				Name:        clean(name),
				Description: clean(description),
				Default:     defaultText,
			})
			continue
		}

		if is(n.Keys, "output") {
			name := n.Keys[1].Token.Text
			ret = append(ret, Value{
				Type: "output",
				Name: name[1 : len(name)-1],
			})
			continue
		}
	}

	return ret
}

func is(keys []*ast.ObjectKey, t string) bool {
	if len(keys) > 0 {
		return keys[0].Token.Text == t
	}

	return false
}

func clean(s string) string {
	if len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return ""
}
