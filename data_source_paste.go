// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// Datasource Name
	pasteName = "sourcehut_paste"
)

// dataSourcePaste returns a data source for getting information about a paste.
func dataSourcePaste() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePasteRead,

		Schema: map[string]*schema.Schema{
			idKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SHA1 hash of the paste.",
			},
			createdKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the paste was created in RFC3339 format.",
			},
			createdTimestampKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date on which the paste was created as a unix timestamp.",
			},
			userKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the user that owns the paste (eg. 'example').",
			},
			canonicalUserKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The canonical name of the user that owns the paste (eg. '~example').",
			},
		},
	}
}

func dataSourcePasteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	paste, err := config.pasteClient.Get(d.Get("id").(string))
	if err != nil {
		return err
	}

	d.SetId(paste.ID)
	err = d.Set(createdKey, paste.Created.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(createdTimestampKey, paste.Created.Unix())
	if err != nil {
		return err
	}
	err = d.Set(userKey, paste.User.Name)
	if err != nil {
		return err
	}
	return d.Set(canonicalUserKey, paste.User.CanonicalName)
}
