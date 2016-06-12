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

func values(f *ast.File) (ret []Value) {
	list := f.Node.(*ast.ObjectList)

	for _, n := range list.Items {
		if is(n.Keys, "variable") {
			name := n.Keys[1].Token.Text
			items := n.Val.(*ast.ObjectType).List.Items

			ret = append(ret, Value{
				Type:        "variable",
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
			ret = append(ret, Value{
				Type:  "output",
				Name:  clean(name),
				Value: value,
			})
			continue
		}
	}

	return ret
}

func get(items []*ast.ObjectItem, key string) string {
	for _, item := range items {
		if is(item.Keys, key) {
			return item.Val.(*ast.LiteralType).Token.Text
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

func clean(s string) string {
	if len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return ""
}
