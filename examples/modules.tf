module "module-without-description" {
  source       = "../path/to/module/1"
}

module "module-with-description" {
  source       = "../path/to/module/2"
  description  = "Direct description"
}

# Description in comment
module "module-with-description-in-comment" {
  source       = "../path/to/module/3"
}

# Description in comment
module "module-with-both-descriptions" {
  source       = "../path/to/module/4"
  description  = "Direct description"
}
