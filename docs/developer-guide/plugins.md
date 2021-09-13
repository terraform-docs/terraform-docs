---
title: "Plugins"
description: "terraform-docs plugin development guide"
menu:
  docs:
    parent: "developer-guide"
weight: 310
toc: false
---

If you want to add or change formatter, you need to write plugins. When changing
plugins, refer to the repository of each plugin and refer to how to build and
install.

If you want to create a new plugin, please refer to [tfdocs-format-template]. The
plugin can use [plugin-sdk] to communicate with the host process. You can create a
new repository from `Use this template`.

[tfdocs-format-template]: https://github.com/terraform-docs/tfdocs-format-template
[plugin-sdk]: https://github.com/terraform-docs/plugin-sdk
