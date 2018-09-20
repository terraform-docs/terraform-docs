package main

import (
	"fmt"
	"log"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/tj/docopt"
)

var version = "dev"

const usage = `
  Usage:
    terraform-docs [--no-required] [json | md | markdown] <path>...
    terraform-docs -h | --help

  Examples:

    # View inputs and outputs
    $ terraform-docs ./my-module

    # View inputs and outputs for variables.tf and outputs.tf only
    $ terraform-docs variables.tf outputs.tf

    # Generate a JSON of inputs and outputs
    $ terraform-docs json ./my-module

    # Generate markdown tables of inputs and outputs
    $ terraform-docs md ./my-module

    # Generate markdown tables of inputs and outputs, but don't print "Required" column
    $ terraform-docs --no-required md ./my-module

    # Generate markdown tables of inputs and outputs for the given module and ../config.tf
    $ terraform-docs md ./my-module ../config.tf

  Options:
    -h, --help     show help information
    --version      print version

`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal(err)
	}

	paths := args["<path>"].([]string)
	doc, err := doc.CreateFromPaths(paths)
	if err != nil {
		log.Fatal(err)
	}

	printRequired := !args["--no-required"].(bool)

	var out string

	switch {
	case args["markdown"].(bool):
		out, err = print.Markdown(doc, printRequired)
	case args["md"].(bool):
		out, err = print.Markdown(doc, printRequired)
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
