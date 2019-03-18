// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// Datasource Name
	userName = "sourcehut_user"

	// Keys
	emailKey    = "email"
	urlKey      = "url"
	locationKey = "location"
	bioKey      = "bio"
	pgpKeyKey   = "preferred_pgp_key"
)

// dataSourceUser returns a data source for getting information about the
// authenticated users account.
func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			userKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the authenticated user (eg. 'example').",
			},
			canonicalUserKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The canonical name of the authenticated user (eg. '~example').",
			},
			emailKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users email.",
			},
			urlKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users URL.",
			},
			locationKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users location.",
			},
			bioKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users bio.",
			},
			pgpKeyKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users preferred PGP key.",
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	user, err := config.metaClient.GetUser()
	if err != nil {
		return err
	}

	d.SetId(user.Name)
	d.Set(userKey, user.Name)
	d.Set(canonicalUserKey, user.CanonicalName)
	d.Set(emailKey, user.Email)
	d.Set(urlKey, user.URL)
	d.Set(locationKey, user.Location)
	d.Set(bioKey, user.Bio)
	d.Set(pgpKeyKey, user.UsePGPKey)
	return nil
}
