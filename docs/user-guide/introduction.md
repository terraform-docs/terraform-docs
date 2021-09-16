---
title: "Introduction"
description: "Generate documentation from Terraform modules in various output formats"
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

Read all about [configuration].

## Formats

One of the most popular format is [markdown table], which is a very good fit for
generating README of module.

{{< img-simple src="markdown-table.png" >}}

which produces:

{{< img-simple src="markdown-table-output.png" >}}

Read all about available [formats].

## Compatibility

terraform-docs compatiblity matrix with Terraform can be found below:

<table class="table pure-table">
  <thead>
    <tr>
      <th>terraform-docs</th>
      <th>Terraform</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>&gt;= 0.13</code></td>
      <td><code>&gt;= 0.15</code></td>
    </tr>
    <tr>
      <td><code>&gt;= 0.8, &lt; 0.13</code></td>
      <td><code>&gt;= 0.12, &lt; 0.15</code></td>
    </tr>
    <tr>
      <td><code>&lt; 0.8</code></td>
      <td><code>&lt; 0.12</code></td>
    </tr>
  </tbody>
</table>

[configuration]: {{< ref "configuration" >}}
[formats]: {{< ref "terraform-docs" >}}
[markdown table]: {{< ref "markdown-table" >}}
