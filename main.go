package main

import (
	"fmt"
	"log"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/tj/docopt"
)

var version = "dev"

const usage = `
  Usage:
    terraform-docs [--no-required] [--no-sort | --sort-inputs-by-required] [--with-aggregate-type-defaults] [json | markdown | md] <path>...
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
    -h, --help                       show help information
    --no-required                    omit "Required" column when generating markdown
    --no-sort                        omit sorted rendering of inputs and ouputs
    --sort-inputs-by-required        sort inputs by name and prints required inputs first
    --with-aggregate-type-defaults   print default values of aggregate types
    --version                        print version

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
	case args["markdown"].(bool):
		out, err = markdown.Print(document, printSettings)
	case args["md"].(bool):
		out, err = markdown.Print(document, printSettings)
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
