terraform {
  required_version = "~> 0.11"
}

provider "sourcehut" {}

data "sourcehut_paste" "testpaste" {
  id = "0b010e5debe7a46fb5a85eeee053b5573873bbee"
}

data "sourcehut_blob" "testpasteblob" {
  id = "640ab2bae07bedc4c163f679a746f7ab7fb5d1fa"
}
