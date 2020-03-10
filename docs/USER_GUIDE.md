# User Guide

## Syntax, Usage, and Output Formats

Please refer to [docs/README.md](/docs/README.md) for guidance on output formats, execution syntax, CLI options, etc.

## Integrating With Your Terraform Repository

A simple git hook `.git/hooks/pre-commit` added to your local terraform repository can keep your Terraform module documentation up to date whenever you make a commit. See also [ git hooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) documentation.

```sh
#!/bin/sh

# Keep module docs up to date
for d in $(ls -1 modules)
do
  terraform-docs md modules/$d > modules/$d/README.md
  if [ $? -eq 0 ] ; then
    git add "./modules/$d/README.md"
  fi
done
```
