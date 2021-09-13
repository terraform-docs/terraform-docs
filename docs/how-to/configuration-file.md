---
title: "Configuration File"
description: "How to use terraform-docs configuration file"
menu:
  docs:
    parent: "how-to"
weight: 202
toc: false
---

Since `v0.10.0`

Configuration can be loaded with `-c, --config string`. Take a look at [configuration]
page for all the details.

```bash
$ pwd
/path/to/parent/folder

$ tree
.
├── module-a
│   └── main.tf
├── module-b
│   └── main.tf
├── ...
└── .terraform-docs.yml

# executing from parent
$ terraform-docs -c .terraform-docs.yml module-a/

# executing from child
$ cd module-a/
$ terraform-docs -c ../.terraform-docs.yml .

# or an absolute path
$ terraform-docs -c /path/to/parent/folder/.terraform-docs.yml .
```

{{< alert type="info" >}}
As of `v0.13.0`, `--config` flag accepts both relative and absolute paths.
{{< /alert >}}

[configuration]: {{< ref "configuration" >}}
