// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"net/http"
	"strconv"
	"time"

	"git.sr.ht/~samwhited/sourcehut-go/git"
	"github.com/hashicorp/terraform/helper/schema"
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
			Default:     "public",
			Description: `The visibility of the repository ("public", "unlisted", or "private").`,
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
	config := meta.(config)
	visibility := git.RepoVisibility(d.Get(visiKey).(string))
	repo, err := config.gitClient.NewRepo(
		d.Get(nameKey).(string),
		d.Get(descKey).(string),
		visibility,
	)
	if err != nil {
		return err
	}

	return setRepo(d, repo)
}

func resourceRepoRead(d *schema.ResourceData, meta interface{}) error {
	return repoRead(d, meta, false)
}

func repoRead(d *schema.ResourceData, meta interface{}, importing bool) error {
	config := meta.(config)

	name := d.Id()
	if !importing {
		name = d.Get(nameKey).(string)
	}

	repo, err := config.gitClient.Repo("", name)
	var statusCode int
	if e, ok := err.(interface {
		StatusCode() int
	}); ok {
		statusCode = e.StatusCode()
	}
	// If the error resulted from a 404, mark the resource as deleted.
	if statusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	return setRepo(d, repo)
}

func resourceRepoDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	return config.gitClient.DeleteRepo(d.Get(nameKey).(string))
}

func resourceRepoUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	oldName, newName := d.GetChange(nameKey)
	repo := &git.Repo{
		Description: d.Get(descKey).(string),
		Name:        newName.(string),
		Visibility:  git.RepoVisibility(d.Get(visiKey).(string)),
	}
	return config.gitClient.UpdateRepo(oldName.(string), repo)
}

func resourceRepoImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	err := repoRead(d, meta, true)
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func setRepo(d *schema.ResourceData, repo *git.Repo) error {
	d.SetId(strconv.FormatInt(repo.ID, 10))
	//err := d.Set(idKey, repo.ID)
	//if err != nil {
	//	return err
	//}
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
	err = d.Set(visiKey, string(repo.Visibility))
	if err != nil {
		return err
	}
	return d.Set(nameKey, repo.Name)
}
