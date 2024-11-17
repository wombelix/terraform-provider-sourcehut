// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Datasource Name
	blobName = "sourcehut_blob"

	// Schema names
	contentsKey = "contents"
)

// dataSourceBlob returns a data source for getting information about a file
// (blob) in a paste.
func dataSourceBlob() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBlobRead,

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
			contentsKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The files contents as a UTF-8 encoded string.",
			},
		},
	}
}

func dataSourceBlobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	blob, err := config.pasteClient.GetBlob(d.Get("id").(string))
	if err != nil {
		return err
	}

	d.SetId(blob.ID)
	err = d.Set(createdKey, blob.Created.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(createdTimestampKey, blob.Created.Unix())
	if err != nil {
		return err
	}
	return d.Set(contentsKey, blob.Contents)
}
