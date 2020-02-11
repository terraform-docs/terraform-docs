## terraform-docs markdown

Generate Markdown of inputs and outputs

### Synopsis

Generate Markdown of inputs and outputs

```
terraform-docs markdown [PATH] [flags]
```

### Options

```
  -h, --help          help for markdown
      --indent int    indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --no-escape     do not escape special characters
      --no-required   do not show "Required" column or section
```

### Options inherited from parent commands

```
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

### SEE ALSO

* [terraform-docs markdown document](markdown-document.md)	 - Generate Markdown document of inputs and outputs
* [terraform-docs markdown table](markdown-table.md)	 - Generate Markdown tables of inputs and outputs
