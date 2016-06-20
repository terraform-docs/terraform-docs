package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/segmentio/terraform-docs/doc"
	"github.com/segmentio/terraform-docs/print"
	"github.com/tj/docopt"
)

var version = ""

const usage = `
  Usage:
    terraform-docs <dir>
    terraform-docs json <dir>
    terraform-docs markdown <dir>
    terraform-docs md <dir>
    terraform-docs -h | --help

  Examples:

    # View inputs and outputs
    $ teraform-docs ./my-module

    # Generate a JSON of inputs and outputs
    $ teraform-docs json ./my-module

    # Generate markdown tables of inputs and outputs
    $ teraform-docs md ./my-module

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

	files := make(map[string]*ast.File, len(names))

	for _, name := range names {
		buf, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		f, err := hcl.ParseBytes(buf)
		if err != nil {
			log.Fatal(err)
		}

		files[name] = f
	}

	doc := doc.Create(files)

	var out string

	switch {
	case args["markdown"].(bool):
		out, err = print.Markdown(doc)
	case args["md"].(bool):
		out, err = print.Markdown(doc)
	case args["json"].(bool):
		out, err = print.JSON(doc)
	default:
		out, err = print.Pretty(doc)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
