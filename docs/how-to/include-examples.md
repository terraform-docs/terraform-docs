---
title: "Include Examples"
description: "How to include example in terraform-docs generated output"
menu:
  docs:
    parent: "how-to"
weight: 206
toc: false
---

Since `v0.14.0`

Example can be automatically included into README by using `content` in configuration
file. For example:

````bash
$ tree
.
├── examples
│   ├── example-1
│   │   ├── main.tf
│   └── example-2
│       └── main.tf
├── ...
├── main.tf
├── variables.tf
├── ...
└── .terraform-docs.yml
````

and

````yaml
# .terraform-docs.yml
content: |-
  {{ .Header }}

  ## Example

  ```hcl
  {{ include "examples/example-1/main.tf" }}
  ```

  {{ .Inputs }}

  {{ .Outputs }}
````
