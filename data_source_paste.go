// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// Datasource Name
	pasteName = "sourcehut_paste"

	// Schema names
	userKey          = "user"
	canonicalUserKey = "canonical_user"
)

// dataSourcePaste returns a data source for getting information about a paste.
func dataSourcePaste() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePasteRead,

		Schema: map[string]*schema.Schema{
			idKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SHA1 hash of the paste.",
			},
			createdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the paste was created in RFC3339 format.",
			},
			createdTimestampKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date on which the paste was created as a unix timestamp.",
			},
			userKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the user that owns the paste (eg. 'example').",
			},
			canonicalUserKey: &schema.Schema{
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
	d.Set(createdKey, paste.Created.Format(time.RFC3339))
	d.Set(createdTimestampKey, paste.Created.Unix())
	d.Set(userKey, paste.User.Name)
	d.Set(canonicalUserKey, paste.User.CanonicalName)
	return nil
}
