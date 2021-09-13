---
title: "header-from"
description: "header-from configuration"
menu:
  docs:
    parent: "configuration"
weight: 124
toc: true
---

Since `v0.10.0`

Relative path to a file to extract header for the generated output from. Supported
file formats are `.adoc`, `.md`, `.tf`, and `.txt`.

{{< alert type="info" >}}
The whole file content is being extracted as module header when extracting from
`.adoc`, `.md`, or `.txt`.
{{< /alert >}}

To extract header from `.tf` file you need to use following javascript, c, or java
like multi-line comment.

```tf
/**
 * # Main title
 *
 * Everything in this comment block will get extracted.
 *
 * You can put simple text or complete Markdown content
 * here. Subsequently if you want to render AsciiDoc format
 * you can put AsciiDoc compatible content in this comment
 * block.
 */

resource "foo" "bar" { ... }
```

{{< alert type="info" >}}
This comment must start at the immediate first line of the `.tf` file
before any `resource`, `variable`, `module`, etc.
{{< /alert >}}

{{< alert type="info" >}}
terraform-docs will never alter line-endings of extracted header text and will assume
whatever extracted is intended as is. It's up to you to apply any kind of Markdown
formatting to them (i.e. adding `<SPACE><SPACE>` at the end of lines for break, etc.)
{{< /alert >}}

## Options

Available options with their default values.

```yaml
header-from: main.tf
```

## Examples

Read `header.md` to extract header:

```yaml
header-from: header.md
```

Read `docs/.header.md` to extract header:

```yaml
header-from: "docs/.header.md"
```
