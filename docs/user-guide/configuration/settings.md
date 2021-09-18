---
title: "settings"
description: "settings configuration"
menu:
  docs:
    parent: "configuration"
weight: 128
toc: true
---

Since `v0.10.0`

General settings to control the behavior and generated output items.

## Options

Available options with their default values.

```yaml
settings:
  anchor: true
  color: true
  default: true
  description: false
  escape: true
  hide-empty: false
  html: true
  indent: 2
  lockfile: true
  read-comments: true
  required: true
  sensitive: true
  type: true
```

### anchor

> since: `v0.12.0`\
> scope: `asciidoc`, `markdown`

Generate HTML anchor tag for elements.

### color

> since: `v0.10.0`\
> scope: `pretty`

Print colorized version of result in the terminal.

### default

> since: `v0.12.0`\
> scope: `asciidoc`, `markdown`

Show "Default" value as column (in table format) or section (in document format).

### description

> since: `v0.13.0`\
> scope: `tfvars hcl`

Show "Descriptions" as comment on variables.

### escape

> since: `v0.10.0`\
> scope: `asciidoc`, `json`, `markdown`

Escape special characters (such as `_`, `*` in Markdown and `>`, `<` in JSON)

### hide-empty

> since: `v0.16.0`\
> scope: `asciidoc`, `markdown`

Hide empty sections.

### html

> since: `v0.14.0`\
> scope: `markdown`

Generate HTML tags (`a`, `pre`, `br`, ...) in the output.

### indent

> since: `v0.10.0`\
> scope: `asciidoc`, `markdown`

Indentation level of headings [available: 1, 2, 3, 4, 5].

### lockfile

> since: `v0.15.0`\
> scope: `global`

Read `.terraform.lock.hcl` to extract exact version of providers.

### read-comments

> since: `v0.16.0`\
> scope: `global`

Use comments from `tf` files for "Description" column (for inputs and outputs) when description is empty

### required

> since: `v0.10.0`\
> scope: `asciidoc`, `markdown`

Show "Required" as column (in table format) or section (in document format).

### sensitive

> since: `v0.10.0`\
> scope: `asciidoc`, `markdown`

Show "Sensitive" as column (in table format) or section (in document format).

### type

> since: `v0.12.0`\
> scope: `asciidoc`, `markdown`

Show "Type" as column (in table format) or section (in document format).

## Examples

Markdown linters rule [MD033] prohibits using raw HTML in markdown document,
the following can be used to appease it:

```yaml
settings:
  anchor: false
  html: false
```

If `.terraform.lock.hcl` is not checked in the repository, running terraform-docs
potentially will produce different providers version on each execution, to prevent
this you can disable it by:

```yaml
settings:
  lockfile: false
```

For simple modules the generated documentation contains a lot of sections that
simply say "no outputs", "no resources", etc. It is possible to hide these empty
sections manually, but if the module changes in the future, they explicitly have
to be enabled again. The following can be used to let terraform-docs automatically
hide empty sections:

```yaml
settings:
  hide-empty: true
```

[MD033]: https://github.com/markdownlint/markdownlint/blob/5329a84691ab0fbce873aa69bb5073a6f5f98bdb/docs/RULES.md#md033---inline-html
