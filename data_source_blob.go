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
			contentsKey: &schema.Schema{
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
	d.Set(createdKey, blob.Created.Format(time.RFC3339))
	d.Set(createdTimestampKey, blob.Created.Unix())
	d.Set(contentsKey, blob.Contents)
	return nil
}
