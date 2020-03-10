resource "tls_private_key" "baz" {}

data "aws_caller_identity" "current" {
  provider = "aws"
}

resource "null_resource" "foo" {}
