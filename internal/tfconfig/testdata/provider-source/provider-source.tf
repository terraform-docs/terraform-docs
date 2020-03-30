terraform {
  required_providers {
    foo = "2.0.0"
    bat = {
      source  = "baz/bat"
      version = "1.0.0"
    }
  }
}
