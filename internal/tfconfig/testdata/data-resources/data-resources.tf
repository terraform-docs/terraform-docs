data "external" "foo" {
}

data "external" "bar" {
  provider = notexternal
}
