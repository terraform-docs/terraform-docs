---
title: "footer-from"
description: "footer-from configuration"
menu:
  docs:
    parent: "configuration"
weight: 122
toc: true
---

Since `v0.12.0`

Relative path to a file to extract footer for the generated output from. Supported
file formats are `.adoc`, `.md`, `.tf`, `.tofu`, and `.txt`.

{{< alert type="info" >}}
The whole file content is being extracted as module footer when extracting from
`.adoc`, `.md`, or `.txt`.
{{< /alert >}}

To extract a footer from a `.tf` or `.tofu` file, use consecutive `#` comments:

```tf
# # Footer
#
# Everything in this comment block will get extracted.
#
# You can put simple text or complete Markdown content
# here. Subsequently if you want to render AsciiDoc format
# you can put AsciiDoc compatible content in this comment
# block.

resource "foo" "bar" { ... }
```

JavaScript-, C-, or Java-like multi-line comments are also supported:

```tf
/**
 * # Footer
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
This comment must start at the immediate first line of the `.tf` or `.tofu` file
before any `resource`, `variable`, `module`, etc.
{{< /alert >}}

{{< alert type="info" >}}
terraform-docs will never alter line-endings of extracted footer text and will assume
whatever extracted is intended as is. It's up to you to apply any kind of Markdown
formatting to them (i.e. adding `<SPACE><SPACE>` at the end of lines for break, etc.)
{{< /alert >}}

## Options

Available options with their default values.

```yaml
footer-from: ""
```

## Examples

Read `footer.md` to extract footer:

```yaml
footer-from: footer.md
```

Read `docs/.footer.md` to extract footer:

```yaml
footer-from: "docs/.footer.md"
```
