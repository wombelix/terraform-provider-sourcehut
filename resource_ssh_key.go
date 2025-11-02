// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"git.sr.ht/~wombelix/terraform-provider-sourcehut/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		CreateContext: resourceSSHKeyCreate,
		ReadContext:   resourceSSHKeyRead,
		DeleteContext: resourceSSHKeyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceSSHKeyImport,
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
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date on which the key was last used as a unix timestamp.",
			},
		},
	}
}

func resourceSSHKeyImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	diags := resourceSSHKeyRead(ctx, d, m)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading sourcehut SSH key: %v", diags[0].Summary)
	}
	return []*schema.ResourceData{d}, nil
}

func resourceSSHKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid resource id: %s", d.Id()))
	}

	key, err := config.client.GetSSHKey(ctx, int(id))
	if err != nil {
		return diag.FromErr(err)
	}

	if key == nil {
		d.SetId("")
		return diags
	}

	user, err := config.client.GetCurrentUser(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setKey(d, key); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(userKey, user.Username); err != nil {
		return diag.FromErr(fmt.Errorf("error setting user key: %s", err))
	}

	if err := d.Set(canonicalUserKey, user.CanonicalName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting canonical user key: %s", err))
	}

	return diags
}

func resourceSSHKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*config)

	key, err := config.client.CreateSSHKey(ctx, d.Get(keyKey).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	user, err := config.client.GetCurrentUser(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setKey(d, key); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(userKey, user.Username); err != nil {
		return diag.FromErr(fmt.Errorf("error setting user key: %s", err))
	}

	if err := d.Set(canonicalUserKey, user.CanonicalName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting canonical user key: %s", err))
	}

	return diag.Diagnostics{}
}

func resourceSSHKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid resource id: %s", d.Id()))
	}

	if err := config.client.DeleteSSHKey(ctx, int(id)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setKey(d *schema.ResourceData, key *client.SSHKey) error {
	d.SetId(strconv.FormatInt(int64(key.ID), 10))

	if err := d.Set(createdKey, key.Created.Format(time.RFC3339)); err != nil {
		return fmt.Errorf("error setting created key: %s", err)
	}

	if err := d.Set(createdTimestampKey, key.Created.Unix()); err != nil {
		return fmt.Errorf("error setting created timestamp key: %s", err)
	}

	if err := d.Set(commentKey, key.Comment); err != nil {
		return fmt.Errorf("error setting comment key: %s", err)
	}

	if err := d.Set(fingerprintKey, key.Fingerprint); err != nil {
		return fmt.Errorf("error setting fingerprint key: %s", err)
	}

	if err := d.Set(lastUsedKey, key.LastUsed.Format(time.RFC3339)); err != nil {
		return fmt.Errorf("error setting last used key: %s", err)
	}

	if err := d.Set(lastUsedTimestampKey, key.LastUsed.Unix()); err != nil {
		return fmt.Errorf("error setting last used timestamp key: %s", err)
	}

	if err := d.Set(keyKey, key.Key); err != nil {
		return fmt.Errorf("error setting key: %s", err)
	}

	return nil
}
