# Markdown

Generate Markdown of inputs and outputs.

## Usage

```text
Usage:
  terraform-docs markdown [PATH] [flags]
  terraform-docs markdown [command]

Aliases:
  markdown, md

Available Commands:
  document    Generate Markdown document of inputs and outputs
  table       Generate Markdown tables of inputs and outputs

Flags:
  -h, --help          help for markdown
      --indent int    indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --no-escape     do not escape special characters
      --no-required   do not show "Required" column or section

Global Flags:
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types

Use "terraform-docs markdown [command] --help" for more information about a command.
```

## Markdown Types

The following Markdown specific output formats are available:

- [Document](/docs/formats/markdown-document.md)
- [Table](/docs/formats/markdown-table.md)

Note that the default format, if not explicitly defined, is `tables`.
