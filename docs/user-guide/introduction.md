---
title: "Introduction"
description: "Generate documentation from Terraform modules in various output formats."
menu:
  docs:
    parent: "user-guide"
weight: 100
toc: true
---

`terraform-docs` is a utility to generate documentation from Terraform modules in
various output formats.

{{< img-simple src="teaser.png" >}}

## Configuration

You can also have consistent execution through a `.terraform-docs.yml` file.

Once you set it up and configured it, every time you or your teammates want to
regenerate documentation (manually, through a pre-commit hook, or as part
of a CI pipeline) all you need to do is run `terraform-docs /module/path`.

{{< img-simple src="config.png" >}}

Read all about [Configuration].

## Formats

One of the most popular format is [markdown table], which is a very good fit for
generating README of module.

{{< img-simple src="markdown-table.png" >}}

which produces:

{{< img-simple src="markdown-table-output.png" >}}

Read all about available [formats].

[Configuration]: {{< ref "configuration" >}}
[markdown table]: {{< ref "markdown-table" >}}
[formats]: {{< ref "terraform-docs" >}}
