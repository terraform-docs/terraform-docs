# Although we never documented it as being good style, the legacy HCL parser
# allowed block labels to be unquoted if they were valid identifiers.
#
# Since there may be modules in the wild that rely on this, we support it
# through fallback to the legacy HCL parser, which should then be able
# to give us similar information as we'd get from the newer HCL parser, albeit
# with less accuracy in source locations and other such rough edges that
# the old parser implies.

terraform {
  required_version = ">= 0.11.0"

  backend "s3" {
    # This is ignored but included to make sure we do ignore it successfully
    foo = "bar"
  }
  ignored = 1
}

provider aws {
  version = "1.0.0"
  ignored = 1
}

provider noversion {
  ignored = 1
}

variable foo {
  description = "foo description"
  default     = "foo default"
  ignored = 1
}

output foo {
  description = "foo description"
  ignored = 1
}

resource null_resource foo {
  ignored = 1
  provider = "notnull.baz"
}

data external foo {
  ignored = 1
}

module foo {
  source = "foo/bar/baz"
  version = "1.2.3"
}
