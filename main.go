package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt.go"
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	markdown_document "github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	markdown_table "github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

var version = "dev"

const usage = `
  Usage:
    terraform-docs [--no-required] [--no-sort | --sort-inputs-by-required] [--with-aggregate-type-defaults] [json | markdown | md] [document | table] <path>...
    terraform-docs -h | --help

  Examples:

    # View inputs and outputs
    $ terraform-docs ./my-module

    # View inputs and outputs for variables.tf and outputs.tf only
    $ terraform-docs variables.tf outputs.tf

    # Generate a JSON of inputs and outputs
    $ terraform-docs json ./my-module

    # Generate Markdown tables of inputs and outputs
    $ terraform-docs md ./my-module

    # Generate Markdown tables of inputs and outputs
    $ terraform-docs md table ./my-module

    # Generate Markdown document of inputs and outputs
    $ terraform-docs md document ./my-module

    # Generate Markdown tables of inputs and outputs, but don't print "Required" column
    $ terraform-docs --no-required md ./my-module

    # Generate Markdown tables of inputs and outputs for the given module and ../config.tf
    $ terraform-docs md ./my-module ../config.tf

  Options:
    -h, --help                       show help information
    --no-required                    omit "Required" column when generating Markdown
    --no-sort                        omit sorted rendering of inputs and ouputs
    --sort-inputs-by-required        sort inputs by name and prints required inputs first
    --with-aggregate-type-defaults   print default values of aggregate types
    --version                        print version

  Types of Markdown:
    document                         generate Markdown document of inputs and outputs
    table                            generate Markdown tables of inputs and outputs (default)

`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal(err)
	}

	paths := args["<path>"].([]string)

	document, err := doc.CreateFromPaths(paths)
	if err != nil {
		log.Fatal(err)
	}

	var printSettings settings.Settings
	if !args["--no-required"].(bool) {
		printSettings.Add(print.WithRequired)
	}

	if !args["--no-sort"].(bool) {
		printSettings.Add(print.WithSortByName)
	}

	if args["--sort-inputs-by-required"].(bool) {
		printSettings.Add(print.WithSortInputsByRequired)
	}

	if args["--with-aggregate-type-defaults"].(bool) {
		printSettings.Add(print.WithAggregateTypeDefaults)
	}

	var out string

	switch {
	case args["markdown"].(bool), args["md"].(bool):
		if args["document"].(bool) {
			out, err = markdown_document.Print(document, printSettings)
		} else {
			out, err = markdown_table.Print(document, printSettings)
		}
	case args["json"].(bool):
		out, err = json.Print(document, printSettings)
	default:
		out, err = pretty.Print(document, printSettings)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
