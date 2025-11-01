// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"git.sr.ht/~emersion/gqlclient"
	"git.sr.ht/~wombelix/terraform-provider-sourcehut/internal/client"
)

const (
	// Resource Name
	repoName = "sourcehut_repository"

	// Schema keys
	nameKey    = "name"
	descKey    = "description"
	visiKey    = "visibility"
	subjectKey = "subject"
)

// repoSchema returns a schema that is used by both the repo resource and the
// repo datasource.
func repoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		nameKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the repository.",
		},
		descKey: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A description of the repository.",
		},
		visiKey: {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "PUBLIC",
			Description: `The visibility of the repository ("public", "unlisted", or "private").`,
			StateFunc: func(v interface{}) string {
				return strings.ToUpper(v.(string))
			},
		},
		createdKey: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date on which the repo was created in RFC3339 format.",
		},
		createdTimestampKey: {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The date on which the repo was created as a unix timestamp.",
		},
		subjectKey: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The message subject.",
		},
	}
}

func resourceRepo() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepoCreate,
		Read:   resourceRepoRead,
		Delete: resourceRepoDelete,
		Update: resourceRepoUpdate,

		Importer: &schema.ResourceImporter{
			State: resourceRepoImport,
		},
		Schema: repoSchema(),
	}
}

func resourceRepoCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config)
	input := client.RepositoryInput{
		Name:        d.Get(nameKey).(string),
		Description: d.Get(descKey).(string),
		Visibility:  d.Get(visiKey).(string),
	}

	repo, err := config.client.CreateRepository(context.Background(), input)
	if err != nil {
		return err
	}

	return setRepo(d, repo)
}

func resourceRepoRead(d *schema.ResourceData, meta interface{}) error {
	return repoRead(d, meta, false)
}

func repoRead(d *schema.ResourceData, meta interface{}, importing bool) error {
	config := meta.(*config)
	ctx := context.Background()

	name := d.Id()
	if !importing {
		name = d.Get(nameKey).(string)
	}

	repo, err := config.client.GetRepository(ctx, name)
	if err != nil {
		if httpErr, ok := err.(*gqlclient.HTTPError); ok && httpErr.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return err
	}

	return setRepo(d, repo)
}

func resourceRepoDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config)
	id, _ := strconv.Atoi(d.Id())
	return config.client.DeleteRepository(context.Background(), id)
}

func resourceRepoUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config)
	id, _ := strconv.Atoi(d.Id())

	input := client.RepositoryInput{
		Name:        d.Get(nameKey).(string),
		Description: d.Get(descKey).(string),
		Visibility:  d.Get(visiKey).(string),
	}

	_, err := config.client.UpdateRepository(context.Background(), id, input)
	return err
}

func resourceRepoImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	err := repoRead(d, meta, true)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func setRepo(d *schema.ResourceData, repo *client.Repository) error {
	d.SetId(strconv.Itoa(repo.ID))
	err := d.Set(createdKey, repo.Created.Format(time.RFC3339))
	if err != nil {
		return err
	}
	err = d.Set(createdTimestampKey, repo.Created.Unix())
	if err != nil {
		return err
	}
	err = d.Set(subjectKey, repo.Subject)
	if err != nil {
		return err
	}
	err = d.Set(descKey, repo.Description)
	if err != nil {
		return err
	}
	err = d.Set(visiKey, repo.Visibility)
	if err != nil {
		return err
	}
	return d.Set(nameKey, repo.Name)
}
