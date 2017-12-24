
  `terraform-docs(1)` &sdot; a quick utility to generate docs from terraform modules.

<img width="1284" alt="screen shot 2016-06-14 at 5 38 37 pm" src="https://cloud.githubusercontent.com/assets/1661587/16049202/1ad63c16-3257-11e6-9e2c-6bb83e684ba4.png">


## Features

  - View docs for inputs and outputs
  - Generate docs for inputs and outputs
  - Generate JSON, YAML or HCL docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs

## Installation

  - `go get github.com/segmentio/terraform-docs`
  - [Binaries](https://github.com/segmentio/terraform-docs/releases)
  - `brew install terraform-docs` (on macOS)

## Usage

```bash
Usage:
  terraform-docs [--inputs| --outputs] [--detailed] [--no-required] [--out-values=<file>] [--var-file=<file>...] [--color| --no-color] [json | yaml | hcl | md | markdown | xml] [<path>...]
  terraform-docs -h | --help

Examples:
  # View inputs and outputs
  $ terraform-docs ./my-module

  # View inputs and outputs for variables.tf and outputs.tf only
  $ terraform-docs variables.tf outputs.tf

  # Generate a JSON of inputs and outputs
  $ terraform-docs json ./my-module

  # Generate markdown tables of inputs and outputs
  $ terraform-docs md ./my-module

  # Generate markdown tables of inputs and outputs, but don't print "Required" column
  $ terraform-docs --no-required md ./my-module

  # Generate markdown tables of inputs and outputs for the given module and ../config.tf
  $ terraform-docs md ./my-module ../config.tf

Options:
  -i, --inputs             Render only inputs
  -o, --outputs            Render only outputs
  -d, --detailed           Render detailed value for <list> and <map>
  -c, --color              Force rendering of color even if the output is redirected or piped
  -C, --no-color           Do not use color to render the result
  -R, --no-required        Do not output "Required" column
  -O, --out-values=<file>  File used to get output values (result of 'terraform output -json' or 'terraform plan -out file')
  -v, --var-file=<file>... Files used to assign values to terraform variables (HCL format)
  -h, --help               Show help information
```

## Example

Given a simple module at `./_example`:

```tf
/**
 * This module has a variable and an output.  This text here will be output before any inputs or outputs!
 */

variable "subnet_ids" {
  description = "a comma-separated list of subnet IDs"
}

// The VPC ID.
output "vpc_id" {
  value     = "vpc-5c1f55fd"
  sensitive = true
}

// This comment will be ignored
output "vpc_cidr" {
  value       = "10.1.0.0/24"
  description = "The VPC CIDR"
}

```

To view docs:

```bash
$ terraform-docs _example
```

To output JSON docs:

```bash
$ terraform-docs json _example
{
  "Comment": "This module has a variable and an output.  This text here will be output before any inputs or outputs!\n",
  "Inputs": [
    {
      "Name": "subnet_ids",
      "Description": "a comma-separated list of subnet IDs",
      "Default": ""
    }
  ],
  "Outputs": [
    {
      "Name": "vpc_id",
      "Description": "The VPC ID."
    },
    {
      "Name": "vpc_cidr",
      "Description": "The VPC CIDR"
    }
  ]
}
```

To output YAML outputs:

```bash
$ terraform-docs yaml _example --outputs
- name: vpc_id
  description: The VPC ID.
```

To output markdown docs:

```bash
$ terraform-docs md _example
This module has a variable and an output.  This text here will be output before any inputs or outputs!


## Inputs

| Name | Description | Default | Required |
|------|-------------|:-----:|:-----:|
| subnet_ids | a comma-separated list of subnet IDs | - | yes |

## Outputs

| Name | Description |
|------|-------------|
| vpc_id | The VPC ID. |
| vpc_cidr | The VPC CIDR |

```

To output markdown docs with resulting output:

```bash
$ terraform plan -out current_plan
$ terraform-docs --out-values current_plan md _example
This module has a variable and an output.  This text here will be output before any inputs or outputs!


## Inputs

| Name | Description | Default | Required |
|------|-------------|:-----:|:-----:|
| subnet_ids | a comma-separated list of subnet IDs | - | yes |

## Outputs

| Name | Description | Value | Type | Sensitive |
|------|-------------|-------|------|-----------|
| vpc_id | The VPC ID. | vpc-5c1f55fd | string | true |
| vpc_cidr | The VPC CIDR | 10.1.0.0/24 | string | false |

```
## License

MIT License

Copyright (c) 2017 Segment, Inc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
