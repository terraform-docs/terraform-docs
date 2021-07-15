---
title: "Configuration"
description: "terraform-docs configuration file, i.e. .terraform-docs.yml"
menu:
  docs:
    parent: "user-guide"
weight: 120
toc: true
---

The `terraform-docs` configuration is a yaml file. This is a convenient way to
share the configuation amongst teammates, CI, or other toolings. To do so you
can use `-c` or `--config` flag which accepts name of the config file.

Default name of this file is `.terraform-docs.yml`, and it will get picked up
(if existed) without needing to explicitly passing with config flag.

```bash
$ tree
.
├── main.tf
├── ...
├── ...
└── .terraform-docs.yml

$ terraform-docs .
```

Or you can use a config file with any arbitrary name:

```bash
$ tree
.
├── main.tf
├── ...
├── ...
└── .tfdocs-config.yml

$ terraform-docs -c .tfdocs-config.yml .
```

As of `v0.13.0`, the order for looking for config file is:

1. root of module directory
1. current directory
1. `$HOME/.tfdocs.d/`

if `.terraform-docs.yml` is found in any of the folders above, that will take
precedence and will override the other ones.

**Note:** Values passed directly as CLI flags will override all of the above.

## Options

Since `v0.10.0`

Below is a complete list of options that you can use with `terraform-docs`, with their
corresponding default values (if applicable).

```yaml
version: ""

formatter: <FORMATTER_NAME>

header-from: main.tf
footer-from: ""

sections:
  hide: []
  show: []

  hide-all: false # deprecated in v0.13.0, removed in v0.15.0
  show-all: true  # deprecated in v0.13.0, removed in v0.15.0

content: ""

output:
  file: ""
  mode: inject
  template: |-
    <!-- BEGIN_TF_DOCS -->
    {{ .Content }}
    <!-- END_TF_DOCS -->

output-values:
  enabled: false
  from: ""

sort:
  enabled: true
  by: name

settings:
  anchor: true
  color: true
  default: true
  description: false
  escape: true
  html: true
  indent: 2
  lockfile: true
  required: true
  sensitive: true
  type: true
```

**Note:** The following options cannot be used together:

- `sections.hide` and `sections.show`
- `sections.hide-all` and `sections.show-all`
- `sections.hide-all` and `sections.hide`
- `sections.show-all` and `sections.show`

**Note:** As of `v0.13.0`, `sections.hide-all` and `sections.show-all` are deprecated
in favor of explicit use of `sections.hide` and `sections.show`, and they are removed
as of `v0.15.0`.

## Version

Since `v0.13.0`

terraform-docs version constraints is almost identical to the syntax used by
Terraform. A version constraint is a string literal containing one or more condition,
which are separated by commas.

```yaml
version: ">= 0.13.0, < 1.0.0"
```

Each condition consists of an operator and a version number. A version number is
a series of numbers separated by dots (e.g. `0.13.0`). Note that version number
should not have leading `v` in it.

Valid operators are as follow:

- `=` (or no operator): allows for exact version number.
- `!=`: exclude an exact version number.
- `>`, `>=`, `<`, and `<=`: comparisons against a specific version.
- `~>`: only the rightmost version component to increment.

## Formatters

Since `v0.10.0`

The following options are supported out of the box by terraform-docs and can be
used for `FORMATTER_NAME`:

- `asciidoc` - [reference]({{< ref "asciidoc" >}})
- `asciidoc document` - [reference]({{< ref "asciidoc-document" >}})
- `asciidoc table` - [reference]({{< ref "asciidoc-table" >}})
- `json` - [reference]({{< ref "json" >}})
- `markdown` - [reference]({{< ref "markdown" >}})
- `markdown document` - [reference]({{< ref "markdown-document" >}})
- `markdown table` - [reference]({{< ref "markdown-table" >}})
- `pretty` - [reference]({{< ref "pretty" >}})
- `tfvars hcl` - [reference]({{< ref "tfvars-hcl" >}})
- `tfvars json` - [reference]({{< ref "tfvars-json" >}})
- `toml` - [reference]({{< ref "toml" >}})
- `xml` - [reference]({{< ref "xml" >}})
- `yaml` - [reference]({{< ref "yaml" >}})

**Note:** You need to pass name of a plugin as `formatter` in order to be able to
use the plugin. For example, if plugin binary file is called `tfdocs-format-foo`,
formatter name must be set to `foo`.

## Header From

Since `v0.10.0`

Relative path to a file to extract header for the generated output from. Supported
file formats are `.adoc`, `.md`, `.tf`, and `.txt`. Default value is `main.tf`.

## Footer From

Since `v0.12.0`

Relative path to a file to extract footer for the generated output from. Supported
file formats are `.adoc`, `.md`, `.tf`, and `.txt`. Default value is `""`.

## Sections

Since `v0.10.0`

The following options are supported and can be used for `sections.show` and
`sections.hide`:

- `all` (since `v0.15.0`)
- `data-sources` (since `v0.13.0`)
- `header`
- `footer` (since `v0.12.0`)
- `inputs`
- `modules` (since `v0.11.0`)
- `outputs`
- `providers`
- `requirements`
- `resources` (since `v0.11.0`)

**Note:** As of `v0.13.0`, `sections.hide-all` and `sections.show-all` are deprecated
in favor of explicit use of `sections.hide` and `sections.show`, and they are removed
as of `v0.15.0`.

## Content

Since `v0.14.0`

Generated content can be customized further away with `content` in configuration.
If the `content` is empty the default orders of section is used. `content` is a
Go template with following additional variables:

- `{{ .Header }}`
- `{{ .Footer }}`
- `{{ .Inputs }}`
- `{{ .Modules }}`
- `{{ .Outputs }}`
- `{{ .Providers }}`
- `{{ .Requirements }}`
- `{{ .Resources }}`

and

- `{{ include "relative/path/to/file" }}`

````yaml
content: |-
  Any arbitrary text can be placed anywhere in the content

  {{ .Header }}

  and even in between sections

  {{ .Providers }}

  and they don't even need to be in the default order

  {{ .Outputs }}

  {{ .Inputs }}

  and include any relative files

  {{ include "relative/path/to/file" }}

  or examples

  ```hcl
  {{ include "examples/foo/main.tf" }}
  ```
````

These variables are the generated output of individual sections in the selected
formatter. For example `{{ .Inputs }}` is Markdown Table representation of _inputs_
when formatter is set to `markdown table` and AsciiDoc Document representation
when formatter is set to `asciidoc document` and so on.

Compatible formats for customized content are:

- `asciidoc document`
- `asciidoc table`
- `markdown document`
- `markdown table`

**Note:** Sections visibility (i.e. `sections.show` and `sections.hide`) takes
precedence over the `content`. In the following example although `{{ .Providers }}`
is used it won't be rendered because `providers` is not set to be shown in
`sections.show`.

```yaml
sections:
  show:
    - header
    - inputs
    - outputs

content: |-
  {{ .Header }}

  Some more information can go here.

  {{ .Providers }}

  {{ .Inputs }}

  {{ .Outputs }}
```

## Output

Since `v0.12.0`

Insert generated output to file if `output.file` (or `--output-file string` CLI
flag) is not empty. Insersion behavior can be controlled by `output.mode` (or
`--output-mode string` CLI flag):

- `inject` (default)

  Partially replace the `output-file` with generated output. This will create
  the `output-file` if it doesn't exist. It will also append to `output-file`
  if it doesn't have surrounding comments.

- `replace`

  Completely replace the `output-file` with generated output. This will create
  the `output-file` if it doesn't exist.

The output generated by formatters (`markdown`, `asciidoc`, etc) will first be
inserted into a template, if provided, before getting saved into the file. This
template can be customized with `output.template` or `--output-template string`
CLI flag.

**Note:** `output.file` can be relative to module root or an absolute path.  
**Note:** `output.template` is optional for mode `replace`.

The default template value is:

```text
<!-- BEGIN_TF_DOCS -->
{{ .Content }}
<!-- END_TF_DOCS -->
```

This template consists of at least three lines (all of which are mandatory):

- begin comment
- `{{ .Content }}` slug
- end comment

You may change the wording of comment as you wish, but the comment must be present
in the template. Also note that `SPACE`s inside `{{ }}` are mandatory.

You may also add as many lines as you'd like before or after `{{ .Content }}` line.

**Note:** `{{ .Content }}` is mandatory if you want to customize template for mode
`replace`. For example if you wish to output to YAML file with trailing comment, the
following can be used:

```yaml
formatter: yaml

output:
  file: output.yaml
  mode: replace
  template: |-
    # Example trailing comments block which will be placed at the top of the
    # 'output.yaml' file.
    #
    # Note that there's no <!-- BEGIN_TF_DOCS --> and <!-- END_TF_DOCS -->
    # which will break the integrity yaml file.
    
    {{ .Content }}
```

### Template Comment

Markdown doesn't officially support inline commenting, there are multiple ways
to do it as a workaround, though. The following formats are supported as begin
and end comments of a template:

- `<!-- This is a comment -->`
- `[]: # (This is a comment)`
- `[]: # "This is a comment"`
- `[]: # 'This is a comment'`
- `[//]: # (This is a comment)`
- `[comment]: # (This is a comment)`

The following is also supported for AsciiDoc format:

- `// This is a comment`

The following can be used where HTML comments are not supported (e.g. Bitbucket
Cloud):

```yaml
output:
  file: README.md
  mode: inject
  template: |-
    [//]: # (BEGIN_TF_DOCS)
    {{ .Content }}

    [//]: # (END_TF_DOCS)
```

Note: The empty line before `[//]: # (END_TF_DOCS)` is mandatory in order for
Markdown engine to properly process the comment line after the paragraph.

## Sort

Since `v0.10.0`

To enable sorting of elements `sort.enabled` (or `--sort bool` CLI flag) can be
used. This will indicate sorting is enabled or not, but consecutively type of
sorting can also be specified with `sort.by` (or `--sort-by string` CLI flag).
The following sort types are supported:

- `name` (default): name of items
- `required`: by name of inputs AND show required ones first
- `type`: type of inputs

**Note:** As of `v0.13.0`, `sort.by` is converted from `list` to `string`.

```yaml
sort:
  enabled: true
  by: required   # this now only accepts string
```

The following error is an indicator that `.terraform-docs.yml` still uses
list for `sort.by`.

```text
Error: unable to decode config, 1 error(s) decoding:

* 'sort.by' expected type 'string', got unconvertible type '[]interface {}'
```
