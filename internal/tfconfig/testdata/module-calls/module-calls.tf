module "foo" {
  source = "foo/bar/baz"
  version = "1.0.2"

  unused = 2
}

module "bar" {
  source = "./child"

  unused = 1
}
