# `terraform-docs`

A utility to generate documentation from Terraform modules in various output formats

## Synopsis

A utility to generate documentation from Terraform modules in various output formats

## Options

```
  -h, --help                           help for terraform-docs
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --output-values                  inject output values into outputs
      --output-values-from string      inject output values from file into outputs
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

## Output Formats

* [terraform-docs json](/docs/formats/json.md)	 - Generate JSON of inputs and outputs
* [terraform-docs markdown](/docs/formats/markdown.md)	 - Generate Markdown of inputs and outputs
  * [terraform-docs markdown document](/docs/formats/markdown-document.md)	 - Generate Markdown document of inputs and outputs
  * [terraform-docs markdown table](/docs/formats/markdown-table.md)	 - Generate Markdown tables of inputs and outputs
* [terraform-docs pretty](/docs/formats/pretty.md)	 - Generate colorized pretty of inputs and outputs
* [terraform-docs xml](/docs/formats/xml.md)	 - Generate XML of inputs and outputs
* [terraform-docs yaml](/docs/formats/yaml.md)	 - Generate YAML of inputs and outputs

## Terraform Versions

Support for Terraform `v0.12.x` has been added in `terraform-docs` version `v0.8.0`. Note that you can still generate output of module configuration which is not compatible with Terraform v0.12 with terraform-docs `v0.8.0` and future releases.
