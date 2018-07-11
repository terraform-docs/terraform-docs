
  `terraform-docs(1)` &sdot; a quick utility to generate docs from terraform modules.

<img width="1284" alt="screen shot 2016-06-14 at 5 38 37 pm" src="https://cloud.githubusercontent.com/assets/1661587/16049202/1ad63c16-3257-11e6-9e2c-6bb83e684ba4.png">


## Features

  - View docs for inputs and outputs
  - Generate docs for inputs and outputs
    * (And AWS SSM Parameters)  
  - Generate JSON docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs
  
## Installation

  - `go get github.com/segmentio/terraform-docs`
  - [Binaries](https://github.com/segmentio/terraform-docs/releases)
  - `brew install terraform-docs` (on macOS)
  - docker 

## Usage

```bash
 Usage:
    terraform-docs  [--no-required]  [json | md | markdown]  <path>...
    terraform-docs  [-o=RESOURCE_NAME] [-a RESOURCE_ATTR] [json | md | markdown]  <path>...

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

    # Geneerate markdown tables of inputs and outputs, including amazon ssm parameters as outputs
    $ terraform-docs -o aws_ssm_parameter md ./my-module

    # Generate markdown tables of inputs and outputs, but don't print "Required" column
    $ terraform-docs --no-required md ./my-module

    # Generate markdown tables of inputs and outputs for the given module and ../config.tf
    $ terraform-docs md ./my-module ../config.tf



  Options:
    -h, --help                    show help information
    -o, --output-resource-name=RESOURCE_NAME  If you want to use any additional terraform resoruces as an output (e.g. aws_ssm_parameter,azurerm_key_vault_secret)
    -a, --output-resource-attr=RESOURCE_ATTR  If using an additional output resource, what attribute should be uesed to get the name [default: name]
```

### With Docker

If you use the docker file, e.g. run `docker build . -t terraform-docs` you can subsequently run with

`docker run -v $(pwd):/workspace terraform-docs` 
(the default uses markdown and `.`, i.e. the above is equavilent to )
`docker run -v $(pwd):/workspace terraform-docs md .` 

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
  value = "vpc-5c1f55fd"
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
      "Description": "The VPC ID.\n"
    }
  ]
}
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

```
## Output and Parameters

While using outputs is great within a given terraform "project" using outputs with remote states is dififcult at best.  To use the output from another terraform project
requires you to specify a bucket name, a bucket region, a key in the bucket to point at the state, and then the output name.  Where using something like AWS parameter store
allows you to just specify a name.   As such we currently allow [AWS SSM Parameters](https://www.terraform.io/docs/providers/aws/r/ssm_parameter.html) to be documented like outputs

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
