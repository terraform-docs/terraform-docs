resource "aws_instance" "foo" {
}

resource "aws_instance" "bar" {
  provider = notaws
}

resource "aws_instance" "baz" {
  provider = aws.aliased
}

resource "aws_instance" "deprecated_bar" {
  provider = "notaws"
}

resource "aws_instance" "deprecated_baz" {
  provider = "aws.aliased"
}
