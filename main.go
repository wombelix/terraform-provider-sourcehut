// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Build time configuration.
var (
	Version = "devel"
	Commit  = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider})
}
