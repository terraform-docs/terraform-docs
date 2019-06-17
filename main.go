package main

import (
	"fmt"
	doc2 "github.com/segmentio/terraform-docs/internal/pkg/doc"
	"log"

	"github.com/docopt/docopt.go"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	markdown_document "github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	markdown_table "github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

var version = "dev"

const usage = `
  Usage:
    terraform-docs [--no-required] [--sort-variables-by-required] [--providers] [json | markdown | md] [document | table] <path>
    terraform-docs -h | --help

  Examples:

    # View variables and outputs
    $ terraform-docs ./my-module

    # Generate a JSON of variables and outputs
    $ terraform-docs json ./my-module

    # Generate Markdown tables of variables and outputs
    $ terraform-docs md ./my-module

    # Generate Markdown tables of variables and outputs
    $ terraform-docs md table ./my-module

    # Generate Markdown document of variables and outputs
    $ terraform-docs md document ./my-module

    # Generate Markdown tables of variables and outputs, but don't print "Required" column
    $ terraform-docs --no-required md ./my-module

  Options:
    -h, --help                       show help information
    --no-required                    omit "Required" column when generating Markdown
    --sort-variables-by-required     sort variables by name and prints required variables first
    --version                        print version
    --providers                      include Terraform provider info for the module

  Types of Markdown:
    document                         generate Markdown document of variables and outputs
    table                            generate Markdown tables of variables and outputs (default)

`

func main() {
	parser := &docopt.Parser{
		HelpHandler:   docopt.PrintHelpAndExit,
		OptionsFirst:  true,
		SkipHelpFlags: false,
	}

	args, err := parser.ParseArgs(usage, nil, version)
	if err != nil {
		log.Fatal(err)
	}

	path := args["<path>"].(string)

	module, diag := tfconfig.LoadModule(path)
	if diag != nil && diag.HasErrors() {
		log.Fatal(err)
	}

	var printSettings settings.Settings
	if !args["--no-required"].(bool) {
		printSettings.Add(settings.WithRequired)
	}

	if args["--sort-variables-by-required"].(bool) {
		printSettings.Add(settings.WithSortVariablesByRequired)
	}

	if args["--providers"].(bool) {
		printSettings.Add(settings.WithProviders)
	}

	doc, err := doc2.Create(module, printSettings)
	if err != nil {
		log.Fatal(err)
	}

	var out string

	switch {
	case args["markdown"].(bool), args["md"].(bool):
		if args["document"].(bool) {
			out, err = markdown_document.Print(doc, printSettings)
		} else {
			out, err = markdown_table.Print(doc, printSettings)
		}
	case args["json"].(bool):
		out, err = json.Print(doc)
	default:
		out, err = pretty.Print(doc, printSettings)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
