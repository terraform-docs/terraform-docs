---
title: "pre-commit Hooks"
description: "How to use pre-commit hooks with terraform-docs"
menu:
  docs:
    parent: "how-to"
weight: 210
toc: false
---

Since `v0.12.0`

With [`pre-commit`], you can ensure your Terraform module documentation is kept
up-to-date each time you make a commit.

1. simply create or update a `.pre-commit-config.yaml`
in the root of your Git repo with at least the following content:

   ```yaml
   repos:
     - repo: https://github.com/terraform-docs/terraform-docs
       rev: "<VERSION, TAG, OR SHA TO USE>"             # e.g. "v0.11.2"
       hooks:
         - id: terraform-docs-go
           args: ["ARGS", "TO PASS", "INCLUDING PATH"]  # e.g. ["--output-file", "README.md", "./mymodule/path"]
   ```

   {{< alert type="info" >}}
   You can also include more than one entry under `hooks:` to update multiple docs.
   Just be sure to adjust the `args:` to pass the path you want terraform-docs to scan.
   {{< /alert >}}

1. install [`pre-commit`] and run `pre-commit` to activate the hooks.

1. make a Terraform change, `git add` and `git commit`.
pre-commit will regenerate your Terraform docs, after which you can
rerun `git add` and `git commit` to commit the code and doc changes together.

You can also regenerate the docs manually by running `pre-commit -a terraform-docs`.

### pre-commit via Docker

The pre-commit hook can also be run via Docker, for those who don't have Go installed.
Just use `id: terraform-docs-docker` in the previous example.

This will build the Docker image from the repo, which can be quite slow.
To download the pre-built image instead, change your `.pre-commit-config.yaml` to:

```yaml
repos:
  - repo: local
    hooks:
      - id: terraform-docs
        name: terraform-docs
        language: docker_image
        entry: quay.io/terraform-docs/terraform-docs:latest  # or, change latest to pin to a specific version
        args: ["ARGS", "TO PASS", "INCLUDING PATH"]          # e.g. ["--output-file", "README.md", "./mymodule/path"]
        pass_filenames: false
```

## Git Hook

A simple git hook (`.git/hooks/pre-commit`) added to your local terraform
repository can keep your Terraform module documentation up to date whenever you
make a commit. See also [git hooks] documentation.

```sh
#!/bin/sh

# Keep module docs up to date
for d in modules/*; do
  if terraform-docs md "$d" > "$d/README.md"; then
    git add "./$d/README.md"
  fi
done
```

{{< alert type="warning" >}}
This is very basic and highly simplified version of [pre-commit-terraform](https://github.com/antonbabenko/pre-commit-terraform).
Please refer to it for complete examples and guides.
{{< /alert >}}

[git hooks]: https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks
[`pre-commit`]: https://pre-commit.com/
