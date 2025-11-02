// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.sr.ht/~wombelix/terraform-provider-sourcehut/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Resource Name
	pgpKeyName = "sourcehut_user_pgp_key"
)

func resourcePGPKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePGPKeyCreate,
		ReadContext:   resourcePGPKeyRead,
		DeleteContext: resourcePGPKeyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourcePGPKeyImport,
		},
		Schema: map[string]*schema.Schema{
			keyKey: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The armored PGP key.",
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
			fingerprintKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fingerprint of the key.",
			},
		},
	}
}

func resourcePGPKeyImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	diags := resourcePGPKeyRead(ctx, d, m)
	if diags.HasError() {
		return nil, fmt.Errorf("error reading sourcehut PGP key: %v", diags[0].Summary)
	}
	return []*schema.ResourceData{d}, nil
}

func resourcePGPKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config)

	key, err := config.client.GetPGPKey(context.Background(), 0) // TODO: implement GetPGPKeyByFingerprint
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}

	user, err := config.client.GetCurrentUser(context.Background())
	if err != nil {
		return diag.FromErr(err)
	}

	diags = resourcePGPKeyRefresh(key, d)
	if diags.HasError() {
		return diags
	}

	if err := d.Set(userKey, user.Username); err != nil {
		return diag.FromErr(fmt.Errorf("error setting user key: %s", err))
	}

	if err := d.Set(canonicalUserKey, user.CanonicalName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting canonical user key: %s", err))
	}

	return diags
}

func resourcePGPKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*config)

	key, err := config.client.CreatePGPKey(context.Background(), d.Get(keyKey).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	user, err := config.client.GetCurrentUser(context.Background())
	if err != nil {
		return diag.FromErr(err)
	}

	diags := resourcePGPKeyRefresh(key, d)
	if diags.HasError() {
		return diags
	}

	if err := d.Set(userKey, user.Username); err != nil {
		return diag.FromErr(fmt.Errorf("error setting user key: %s", err))
	}

	if err := d.Set(canonicalUserKey, user.CanonicalName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting canonical user key: %s", err))
	}

	return diags
}

func resourcePGPKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	config := m.(*config)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid resource id: %s", d.Id()))
	}

	err = config.client.DeletePGPKey(context.Background(), int(id))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePGPKeyRefresh(key *client.PGPKey, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId(strconv.FormatInt(int64(key.ID), 10))

	if err := d.Set(createdKey, key.Created.Format(time.RFC3339)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting created key: %s", err))
	}

	if err := d.Set(createdTimestampKey, key.Created.Unix()); err != nil {
		return diag.FromErr(fmt.Errorf("error setting created timestamp key: %s", err))
	}

	if err := d.Set(fingerprintKey, key.Fingerprint); err != nil {
		return diag.FromErr(fmt.Errorf("error setting fingerprint key: %s", err))
	}

	if err := d.Set(keyKey, key.Key); err != nil {
		return diag.FromErr(fmt.Errorf("error setting key: %s", err))
	}

	return diags
}
