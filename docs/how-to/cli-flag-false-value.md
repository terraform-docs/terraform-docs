---
title: "CLI Flag 'false' value"
description: "How to use pass 'false' value to terraform-docs CLI flags"
menu:
  docs:
    parent: "how-to"
weight: 201
toc: false
---

Boolean flags can only take arguments via `--flag=[true|false]` or for short names
(if available) `-f=[true|false]`. You cannot use `--flag [true|false]` nor can you
use the shorthand `-f [true|false]` as it will result in the following error:

```text
Error: accepts 1 arg(s), received 2
```

Example:

```bash
# disable reading .terraform.lock.hcl
$ terraform-docs markdown --lockfile=false /path/to/module
```
