/**
 * Example of 'foo_bar' module in `foo_bar.tf`.
 *
 * - list item 1
 * - list item 2
 *
 * Even inline **formatting** in _here_ is possible.
 * and some [link](https://domain.com/)
 */

terraform {
  required_version = ">= 0.12"
  required_providers {
    aws = {
      version = ">= 2.15.0"
      configuration_aliases = [
        aws,
        aws.ident
      ]
    }
  }
}

resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

data "aws_caller_identity" "ident" {
    provider = "aws.ident"
}

# terraform-docs-ignore
data "aws_caller_identity" "ignored" {
  provider = "aws"
}

resource "null_resource" "foo" {}

# terraform-docs-ignore
resource "null_resource" "ignored" {}

module "foo" {
  source  = "bar"
  version = "1.2.3"
}

module "foobar" {
  source = "git@github.com:module/path?ref=v7.8.9"
}

locals {
  arn = provider::aws::arn_parse("arn:aws:iam::444455556666:role/example")
}

// terraform-docs-ignore
module "ignored" {
  source  = "baz"
  version = "1.2.3"
}
