// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"context"
	"fmt"
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
	pasteID := d.Get("id").(string)

	// Get the paste first to identify its files
	paste, err := config.client.GetPaste(context.Background(), pasteID)
	if err != nil {
		return err
	}

	if len(paste.Files) == 0 {
		return fmt.Errorf("no files found in paste")
	}

	// Get the actual blob content
	blob, err := config.client.GetPasteBlob(context.Background(), pasteID, paste.Files[0].Hash)
	if err != nil {
		return err
	}

	d.SetId(blob.Hash)
	err = d.Set(createdKey, paste.Created.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(createdTimestampKey, paste.Created.Unix())
	if err != nil {
		return err
	}
	return d.Set(contentsKey, blob.Contents)
}
