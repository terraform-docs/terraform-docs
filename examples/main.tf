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
    foo = {
      source  = "https://registry.acme.com/foo"
      version = ">= 1.0"
    }
  }
}

// this description for tls_private_key.baz
// which can be multiline.
resource "tls_private_key" "baz" {}
resource "foo_resource" "baz" {}

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

# This resource is going to get ignored from generated
# output by using the following known comment.
# terraform-docs-ignore
# And the ignore keyword also doesn't have to be the first,
# last, or only thing in a leading comment
resource "null_resource" "ignored" {}

module "bar" {
  source  = "baz"
  version = "4.5.6"
}

# another type of description for module foo
module "foo" {
  source  = "bar"
  version = "1.2.3"
}

module "baz" {
  source  = "baz"
  version = "4.5.6"
}

module "foobar" {
  source = "git@github.com:module/path?ref=v7.8.9"
}

// terraform-docs-ignore
module "ignored" {
  source  = "foobaz"
  version = "7.8.9"
}
