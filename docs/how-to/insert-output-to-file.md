---
title: "Insert Output To File"
description: "How to insert generated terraform-docs output to file"
menu:
  docs:
    parent: "how-to"
weight: 205
toc: false
---

Since `v0.12.0`

Generated output can be insterted directly into the file. There are two modes of
insersion: `inject` (default) or `replace`. Take a look at [output] configuration
for all the details.

```bash
terraform-docs markdown table --output-file README.md --output-mode inject /path/to/module
```

{{< alert type="info" >}}
`--output-file` can be relative to module path or an absolute path in filesystem.
{{< /alert>}}

```bash
$ pwd
/path/to/module

$ tree .
.
├── docs
│   └── README.md
├── ...
└── main.tf

# this works, relative path
$ terraform-docs markdown table --output-file ./docs/README.md .

# so does this, absolute path
$ terraform-docs markdown table --output-file /path/to/module/docs/README.md .
```

[output]: {{< ref "output" >}}
