# User Guide

## Terraform Versions

Support for Terraform `v0.12.x` has been added in `terraform-docs` version `v0.8.0`. Note that you can still generate output of module configuration which is not compatible with Terraform v0.12 with terraform-docs `v0.8.0` and future releases.

## Syntax, Usage, and Output Formats

Please refer to [Formats Guide](/docs/FORMATS_GUIDE.md) for guidance on output formats, execution syntax, CLI options, etc.

## Generate terraform.tfvars

You can generate `terraform.tfvars` in both `hcl` and `json` format by executing the following:

```bash
terraform-docs tfvars hcl /path/to/module

# or

terraform-docs tfvars json /path/to/module
```

Note that any required input variables will be empty, `""` in HCL and `null` in JSON format.

## Integrating With Your Terraform Repository

A simple git hook `.git/hooks/pre-commit` added to your local terraform repository can keep your Terraform module documentation up to date whenever you make a commit. See also [git hooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) documentation.

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
