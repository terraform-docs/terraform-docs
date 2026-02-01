provider "aws" {
  for_each = toset(["us-east-1"])
  alias    = "main"
}

data "aws_caller_identity" "current" {
  provider = aws.main["us-east-1"]
}
