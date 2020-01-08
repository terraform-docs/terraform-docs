resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

data "aws_caller_identity" "ident" {
  provider = "aws.ident"
}

terraform {
  required_providers {
    aws = ">= 2.15.0"
  }
}

resource "null_resource" "foo" {}
