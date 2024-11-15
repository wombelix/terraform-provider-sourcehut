// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

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
			keyKey: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The key in authorized_keys format.",
			},
			createdKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the key was authorized in RFC3339 format.",
			},
			createdTimestampKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date on which the key was authorized as a unix timestamp.",
			},
			userKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the user that owns the key (eg. 'example').",
			},
			canonicalUserKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The canonical name of the user that owns the key (eg. '~example').",
			},
			commentKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The comment on the key.",
			},
			fingerprintKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fingerprint of the key.",
			},
			lastUsedKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date on which the key was last used in RFC3339 format.",
			},
			lastUsedTimestampKey: {
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

	return setKey(d, key)
}

func resourceSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	key, err := config.metaClient.NewSSHKey(d.Get(keyKey).(string))
	if err != nil {
		return err
	}

	return setKey(d, key)
}

func resourceSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	return config.metaClient.DeleteSSHKey(id)
}

func setKey(d *schema.ResourceData, key meta.SSHKey) error {
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
	err = d.Set(commentKey, key.Comment)
	if err != nil {
		return err
	}
	err = d.Set(fingerprintKey, key.Fingerprint)
	if err != nil {
		return err
	}
	err = d.Set(lastUsedKey, key.LastUsed.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(lastUsedTimestampKey, key.LastUsed.Unix())
	if err != nil {
		return err
	}
	return d.Set(keyKey, key.Key)
}
