# `terraform-docs` Documentations

A utility to generate documentation from Terraform modules in various output formats.

## Usage

```text
Usage:
  terraform-docs [command]

Available Commands:
  completion  Generate autocomplete for terraform-docs
  help        Help about any command
  json        Generate a JSON of inputs and outputs
  markdown    Generate Markdown of inputs and outputs
  pretty      Generate a colorized pretty of inputs and outputs
  version     Print the version number of terraform-docs
  yaml        Generate a YAML of inputs and outputs

Flags:
  -h, --help                           help for terraform-docs
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --version                        version for terraform-docs
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types

Use "terraform-docs [command] --help" for more information about a command.
```

## Output Formats

The following output formats are available:

- [JSON](/docs/formats/json.md)
- [Markdown](/docs/formats/markdown.md)
  - [Document](/docs/formats/markdown-document.md)
  - [Table](/docs/formats/markdown-table.md)
- [Pretty](/docs/formats/pretty.md)
- [YAML](/docs/formats/yaml.md)

## Terraform Versions

Support for Terraform `v0.12.x` has been added in `terraform-docs` version `v0.8.0`. Note that you can still generate output of module configuration which is not compatible with Terraform v0.12 with terraform-docs `v0.8.0` and future releases.
