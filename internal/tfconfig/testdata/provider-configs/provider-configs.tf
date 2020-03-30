provider "foo" {
}

provider "bar" {
  version = "1.0.0"
}

# Ensure that an implied dependency doesn't overwrite the explicit dependency
# on version 1.0.0.
resource "bar_bar" "bar" {
}

terraform {
  required_providers {
    baz = "2.0.0"
    bar = "1.1.0"
  }
}
