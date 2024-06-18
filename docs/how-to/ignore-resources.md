---
title: "Ignore Resources to be Generated"
description: "How to ignore resources from generated output"
menu:
  docs:
    parent: "how-to"
weight: 204
toc: false
---

Since `v0.18.0`

Any type of resources can be ignored from the generated output by prepending them
with a comment `terraform-docs-ignore`. Supported type of Terraform resources to
get ignored are:

- `resource`
- `data`
- `module`
- `variable`
- `output`

{{< alert type="info" >}}
If a `resource` or `data` is ignored, their corresponding discovered provider
will also get ignored from "Providers" section.
{{< /alert>}}

Take the following example:

```hcl
##################################################################
# All of the following will be ignored from the generated output #
##################################################################

# terraform-docs-ignore
resource "foo_resource" "foo" {}

# This resource is going to get ignored from generated
# output by using the following known comment.
#
# terraform-docs-ignore
#
# The ignore keyword also doesn't have to be the first,
# last, or the only thing in a leading comment
resource "bar_resource" "bar" {}

# terraform-docs-ignore
data "foo_data_resource" "foo" {}

# terraform-docs-ignore
data "bar_data_resource" "bar" {}

// terraform-docs-ignore
module "foo" {
  source  = "foo"
  version = "x.x.x"
}

# terraform-docs-ignore
variable "foo" {
  default = "foo"
}

// terraform-docs-ignore
output "foo" {
  value = "foo"
}
```

{{< alert type="info" >}}
The ignore keyword (i.e. `terraform-docs-ignore`) doesn't have to be the first,
last, or only thing in a leading comment. As long as the keyword is present in
a comment, the following resource will get ignored.
{{< /alert>}}
