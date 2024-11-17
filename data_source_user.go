// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			userKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the authenticated user (eg. 'example').",
			},
			canonicalUserKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The canonical name of the authenticated user (eg. '~example').",
			},
			emailKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users email.",
			},
			urlKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users URL.",
			},
			locationKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users location.",
			},
			bioKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The users bio.",
			},
			pgpKeyKey: {
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
	err = d.Set(userKey, user.Name)
	if err != nil {
		return err
	}
	err = d.Set(canonicalUserKey, user.CanonicalName)
	if err != nil {
		return err
	}
	err = d.Set(emailKey, user.Email)
	if err != nil {
		return err
	}
	err = d.Set(urlKey, user.URL)
	if err != nil {
		return err
	}
	err = d.Set(locationKey, user.Location)
	if err != nil {
		return err
	}
	err = d.Set(bioKey, user.Bio)
	if err != nil {
		return err
	}
	return d.Set(pgpKeyKey, user.UsePGPKey)
}
