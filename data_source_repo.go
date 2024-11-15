// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// dataSourceRepo returns a data source for getting information about a
// repository.
func dataSourceRepo() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRepoRead,
		Schema: repoSchema(),
	}
}
