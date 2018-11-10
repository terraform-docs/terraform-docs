package main

import (
	"fmt"
	"log"
	"strings"

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
    terraform-docs [--no-required] [--no-sort | --sort-inputs-by-required] [--with-aggregate-type-defaults] [--follow-modules] [json | markdown | md] <path>...
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
    --follow-modules                 follow modules in stacks
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

	// construct the final output from multiple sources
	var out strings.Builder
	// most functions have double output (string, err)
	// which can not be used as input directly
	var tempstring string

	// get the main output (formatted)
	tempstring, err = doPrint(args, document, printSettings)

	if err != nil {
		log.Fatal(err)
	}
	// add the formatted document to the output string
	out.WriteString(tempstring)

	// done with the standard stuff, modules follow
	// no chance to use JSON as the logic of the program had to be changed
	if args["--follow-modules"].(bool) && !args["json"].(bool) && document.HasModules() {

		for _, module := range document.Modules {
			paths := []string{module.Source}

			document, err := doc.CreateFromPaths(paths)
			if err != nil {
				log.Fatal(err)
			}

			tempstring, err = doPrint(args, document, printSettings)

			// print the Module name as header
			switch {
			case args["markdown"].(bool):
				out.WriteString(fmt.Sprintf("\n----\n# Module: %s\n\n", module.Name))
			case args["md"].(bool):
				out.WriteString(fmt.Sprintf("\n----\n# Module: %s\n\n", module.Name))
			default:
				format := "\n\033[4m\033[1mModule:\033[21m %s\033[0m\n\n"
				out.WriteString(fmt.Sprintf(format, module.Name))
			}

			out.WriteString(tempstring)
		}
	}

	// finally print the result
	fmt.Println(out.String())

}

// helper function to save code on switch()
func doPrint(args map[string]interface{}, document *doc.Doc, printSettings settings.Settings) (string, error) {
	switch {
	case args["markdown"].(bool):
		return markdown.Print(document, printSettings)
	case args["md"].(bool):
		return markdown.Print(document, printSettings)
	case args["json"].(bool):
		return json.Print(document, printSettings)
	default:
		return pretty.Print(document, printSettings)
	}
}
