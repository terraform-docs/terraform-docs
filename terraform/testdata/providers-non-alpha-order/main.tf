terraform {
  required_providers {
    tls = {
      version = ">= 4.0"
    }
  }
}

resource "tls_private_key" "key" {}

resource "aws_instance" "server" {}
