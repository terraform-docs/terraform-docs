/**
 * Usage:
 *
 * Example of 'foo_bar' module in `foo_bar.tf`.
 *
 * - list item 1
 * - list item 2
 *
 * Even inline **formatting** in _here_ is possible.
 * and some [link](https://domain.com/)
 *
 * * list item 3
 * * list item 4
 *
 * ```hcl
 * module "foo_bar" {
 *   source = "github.com/foo/bar"
 *
 *   id   = "1234567890"
 *   name = "baz"
 *
 *   zones = ["us-east-1", "us-west-1"]
 *
 *   tags = {
 *     Name         = "baz"
 *     Created-By   = "first.last@email.com"
 *     Date-Created = "20180101"
 *   }
 * }
 * ```
 *
 * Here is some trailing text after code block,
 * followed by another line of text.
 *
 * | Name | Description     |
 * |------|-----------------|
 * | Foo  | Foo description |
 * | Bar  | Bar description |
 */

terraform {
  required_version = ">= 0.12"
  required_providers {
    random = ">= 2.2.0"
    aws    = ">= 2.15.0"
  }
}

resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

data "aws_caller_identity" "ident" {
  provider = "aws.ident"
}

resource "null_resource" "foo" {}

module "foo" {
  source  = "bar"
  version = "1.2.3"
}

module "bar" {
  source  = "baz"
  version = "4.5.6"
}

module "baz" {
  source  = "baz"
  version = "4.5.6"
}