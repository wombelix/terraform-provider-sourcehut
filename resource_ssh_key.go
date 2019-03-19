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
	sshKeyName = "sourcehut_user_ssh_key"

	// Schema keys
	commentKey           = "comment"
	fingerprintKey       = "fingerprint"
	keyKey               = "key"
	lastUsedKey          = "last_used"
	lastUsedTimestampKey = "last_used_timestamp"
)

func resourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSHKeyCreate,
		Read:   resourceSSHKeyRead,
		Delete: resourceSSHKeyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSSHKeyImport,
		},
		Schema: map[string]*schema.Schema{
			keyKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The key in authorized_keys format.",
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
			commentKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The comment on the key.",
			},
			fingerprintKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fingerprint of the key.",
			},
			lastUsedKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the key was last used in RFC3339 format.",
			},
			lastUsedTimestampKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the key was last used as a unix timestamp.",
			},
		},
	}
}

func resourceSSHKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	err := resourceSSHKeyRead(d, meta)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourceSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	key, err := config.metaClient.GetSSHKey(id)
	if err != nil {
		return err
	}

	setKey(d, key)
	return nil
}

func resourceSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	key, err := config.metaClient.NewSSHKey(d.Get(keyKey).(string))
	if err != nil {
		return err
	}

	setKey(d, key)
	return nil
}

func resourceSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	return config.metaClient.DeleteSSHKey(id)
}

func setKey(d *schema.ResourceData, key meta.SSHKey) {
	d.SetId(strconv.FormatInt(key.ID, 10))
	d.Set(idKey, key.ID)
	d.Set(createdKey, key.Authorized.Format(time.RFC3339))
	d.Set(createdTimestampKey, key.Authorized.Unix())
	d.Set(userKey, key.Owner.Name)
	d.Set(canonicalUserKey, key.Owner.CanonicalName)
	d.Set(commentKey, key.Comment)
	d.Set(fingerprintKey, key.Fingerprint)
	d.Set(lastUsedKey, key.LastUsed.Format(time.RFC3339))
	d.Set(lastUsedTimestampKey, key.LastUsed.Unix())
	d.Set(keyKey, key.Key)
}
