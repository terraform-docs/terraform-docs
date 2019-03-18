package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/terraform/terraform"
	"github.com/segmentio/terraform-docs/doc"
	"github.com/segmentio/terraform-docs/print"
	"github.com/tj/docopt"
)

var version = "dev"

const usage = `
Usage:
  terraform-docs [--inputs| --outputs] [--terraform-output] [--detailed] [--no-required] [--out-values=<file>] [--var-file=<file>...] [--color| --no-color] [json | yaml | hcl | md | markdown | xml] [<path>...]
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
  -i, --inputs             Render only inputs
  -o, --outputs            Render only outputs
  -t, --terraform-output   Render outputs in terraform output format ('terraform output -json')
  -d, --detailed           Render detailed value for <list> and <map>
  -c, --color              Force rendering of color even if the output is redirected or piped
  -C, --no-color           Do not use color to render the result
  -R, --no-required        Do not output "Required" column
  -O, --out-values=<file>  File used to get output values (result of 'terraform output -json' or 'terraform plan -out file')
  -v, --var-file=<file>... Files used to assign values to terraform variables (HCL format) 
  -h, --help               Show help information
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatal(err)
	}

	// If color options are specified, set the color mode (if nothing is specified, the color mode is determined by the library depending
	// if the output is redirected or piped or rendered to console)
	if args["--color"].(bool) {
		color.NoColor = false
	} else if args["--no-color"].(bool) {
		color.NoColor = true
	}

	var names []string
	paths := args["<path>"].([]string)
	if len(paths) == 0 {
		paths = []string{"."}
	}
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

	document := doc.Create(files)

	// Determine the render mode parameters
	var renderMode print.RenderMode
	if args["--inputs"].(bool) {
		renderMode |= print.RenderInputs
	}
	if args["--outputs"].(bool) {
		renderMode |= print.RenderOutputs
	}
	if args["--detailed"].(bool) {
		renderMode |= print.RenderDetailed
	}
	if renderMode&print.RenderAll == 0 {
		renderMode |= print.RenderAll
	}

	// Import the optionally supplied tfvars files
	vars := make(map[string]interface{})
	for _, file := range args["--var-file"].([]string) {
		if content, err := LoadHCL(file); err != nil {
			log.Fatal(err)
		} else {
			for k, v := range content {
				vars[k] = v
			}
		}
	}
	for i := range document.Inputs {
		i := &document.Inputs[i]
		if value, ok := vars[i.Name]; ok {
			if i.Default == nil {
				i.Default = &doc.Value{"", ""}
			}
			i.Default.Literal = fmt.Sprintf("%v", value)
		}
	}

	// Import the optionally supplied outputs file
	if outFile := args["--out-values"]; outFile != nil {
		outputs, err := LoadOutputs(outFile.(string))
		if err != nil {
			log.Fatal(err)
		}
		for i := range document.Outputs {
			o := &document.Outputs[i]
			if matched := outputs[o.Name]; matched != nil {
				o.Result = doc.Result{Sensitive: matched.Sensitive, Type: matched.Type, Value: matched.Value}
			}
		}
	}

	printRequired := !args["--no-required"].(bool)

	var out string

	switch {
	case args["markdown"].(bool) || args["md"].(bool):
		out, err = print.Markdown(document, renderMode, printRequired, args["--out-values"] != nil)
	case args["json"].(bool):
		if args["--terraform-output"].(bool) {
			out, err = print.TerraformOutput(document, renderMode)
		} else {
			out, err = print.JSON(document, renderMode)
		}
	case args["yaml"].(bool):
		out, err = print.YAML(document, renderMode)
	case args["hcl"].(bool):
		out, err = print.HCL(document, renderMode)
	case args["xml"].(bool):
		out, err = print.XML(document, renderMode)
	default:
		out, err = print.Pretty(document, renderMode)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

// LoadHCL loads hcl file into variable
func LoadHCL(filename string) (result map[string]interface{}, err error) {
	if content, err := ioutil.ReadFile(filename); err == nil {
		err = hcl.Unmarshal(content, &result)
	}
	return
}

// LoadOutputs loads output values coming either from a JSON file resulting of 'terraform output -json' or
// the outputs contained in the out file resulting of 'terraform plan -out <file>'
func LoadOutputs(filename string) (result map[string]*terraform.OutputState, err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return
	}
	plan, err := terraform.ReadPlan(reader)
	if err != nil {
		reader.Seek(0, 0)
		// The outFile may be a JSON file
		var content []byte
		if content, err = ioutil.ReadAll(reader); err == nil {
			err = json.Unmarshal(content, &result)
		}
		if err != nil {
			err = fmt.Errorf("The out-values file must be either a JSON file resulting from 'terraform output -json' or the out file produced by 'terraform plan -out <file>': %v", err)
			return
		}
	} else {
		result = plan.State.Modules[0].Outputs
	}
	return
}
