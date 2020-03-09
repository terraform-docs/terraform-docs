/**
 * Example of 'foo_bar' module in `foo_bar.tf`.
 *
 * - list item 1
 * - list item 2
 *
 * Even inline **formatting** in _here_ is possible.
 * and some [link](https://domain.com/)
 */

resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

resource "null_resource" "foo" {}
