# SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
#
# SPDX-License-Identifier: BSD-2-Clause

terraform {
  required_version = ">= 1.8"
}

provider "sourcehut" {}

data "sourcehut_paste" "testpaste" {
  id = "0b010e5debe7a46fb5a85eeee053b5573873bbee"
}

data "sourcehut_blob" "testpasteblob" {
  id = "640ab2bae07bedc4c163f679a746f7ab7fb5d1fa"
}

data "sourcehut_user" "myuser" {}

resource "sourcehut_user_ssh_key" "laptop" {
  key = file("~/.ssh/id_ed25519.pub")
}

resource "sourcehut_user_pgp_key" "laptop" {
  key = file("pub.asc")
}
