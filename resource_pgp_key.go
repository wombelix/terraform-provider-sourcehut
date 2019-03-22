// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"strconv"
	"time"

	"git.sr.ht/~samwhited/sourcehut-go/meta"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// Resource Name
	pgpKeyName = "sourcehut_user_pgp_key"
)

func resourcePGPKey() *schema.Resource {
	return &schema.Resource{
		Create: resourcePGPKeyCreate,
		Read:   resourcePGPKeyRead,
		Delete: resourcePGPKeyDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePGPKeyImport,
		},
		Schema: map[string]*schema.Schema{
			keyKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The armored PGP key.",
			},
			createdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the key was authorized in RFC3339 format.",
			},
			createdTimestampKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date on which the key was authorized as a unix timestamp.",
			},
			userKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the user that owns the key (eg. 'example').",
			},
			canonicalUserKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The canonical name of the user that owns the key (eg. '~example').",
			},
			fingerprintKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fingerprint of the key.",
			},
		},
	}
}

func resourcePGPKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	err := resourcePGPKeyRead(d, meta)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourcePGPKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	key, err := config.metaClient.GetPGPKey(id)
	if err != nil {
		return err
	}

	return setPGPKey(d, key)
}

func resourcePGPKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	key, err := config.metaClient.NewPGPKey(d.Get(keyKey).(string))
	if err != nil {
		return err
	}

	return setPGPKey(d, key)
}

func resourcePGPKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	return config.metaClient.DeletePGPKey(id)
}

func setPGPKey(d *schema.ResourceData, key meta.PGPKey) error {
	d.SetId(strconv.FormatInt(key.ID, 10))
	err := d.Set(idKey, key.ID)
	if err != nil {
		return err
	}
	err = d.Set(createdKey, key.Authorized.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(createdTimestampKey, key.Authorized.Unix())
	if err != nil {
		return err
	}
	err = d.Set(userKey, key.Owner.Name)
	if err != nil {
		return err
	}
	err = d.Set(canonicalUserKey, key.Owner.CanonicalName)
	if err != nil {
		return err
	}
	err = d.Set(fingerprintKey, key.KeyID)
	if err != nil {
		return err
	}
	return d.Set(keyKey, key.Key)
}
