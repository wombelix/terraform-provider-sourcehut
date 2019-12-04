// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

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
