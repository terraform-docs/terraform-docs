terraform {
  required_providers {
    aws = {
      version = ">= 4.0"
    }
  }
}

resource "aws_s3_bucket" "first" {}

resource "tls_private_key" "key" {}

resource "aws_instance" "second" {}
