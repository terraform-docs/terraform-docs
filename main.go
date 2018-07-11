package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	docopt "github.com/docopt/docopt.go"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/segmentio/terraform-docs/doc"
	"github.com/segmentio/terraform-docs/print"
)

var version = "dev"

const usage = `
  Usage:
    terraform-docs  [--no-required]  [json | md | markdown]  <path>...
    terraform-docs  [-o=RESOURCE_NAME] [-a RESOURCE_ATTR] [json | md | markdown]  <path>...

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

    # Geneerate markdown tables of inputs and outputs, including amazon ssm parameters as outputs 
    $ terraform-docs -o aws_ssm_parameter md ./my-module
     
    # Generate markdown tables of inputs and outputs, but don't print "Required" column
    $ terraform-docs --no-required md ./my-module

    # Generate markdown tables of inputs and outputs for the given module and ../config.tf
    $ terraform-docs md ./my-module ../config.tf



  Options:
    -h, --help                    show help information
    -o, --output-resource-name=RESOURCE_NAME  If you want to use any additional terraform resoruces as an output (e.g. aws_ssm_parameter,azurerm_key_vault_secret) 
    -a, --output-resource-attr=RESOURCE_ATTR  If using an additional output resource, what attribute should be uesed to get the name [default: name]
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	paths := args["<path>"].([]string)
	for _, p := range paths {
		pi, err := os.Stat(p)
		if err != nil {
			log.Fatal(err)
		}

		if !pi.IsDir() {
			names = append(names, p)
			continue
		}

		files, err := filepath.Glob(fmt.Sprintf("%s/*.tf", p))
		if err != nil {
			log.Fatal(err)
		}

		names = append(names, files...)
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

	printRequired := !args["--no-required"].(bool)

	for k, v := range args {
		log.Printf("Got arg %s with val %v", k, v)
	}

	var docOpts doc.Opts
	if val, ok := args["--output-resource-name"]; ok && val != nil {
		docOpts = doc.Opts{ResourceOutput: val.(string),
			ResourceNameAttribute: args["--output-resource-attr"].(string)}
	}
	// if(args)
	// outputResource := args["--output-resource-name"].(string)
	// outputResourceName := args["--output-resource-attr"].(string)

	doc := doc.Create(docOpts, files)

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
