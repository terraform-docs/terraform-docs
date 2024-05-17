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
    aws = ">= 2.15.0"
  }
}

resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

resource "null_resource" "foo" {}

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
