// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"fmt"
	"os"

	"git.sr.ht/~wombelix/sourcehut-go"
	"git.sr.ht/~wombelix/sourcehut-go/git"
	"git.sr.ht/~wombelix/sourcehut-go/meta"
	"git.sr.ht/~wombelix/sourcehut-go/paste"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	// All config
	/* #nosec */
	tokenKey = "token"
	/* #nosec */
	tokenEnv = "SRHT_TOKEN"

	// Common key names
	idKey               = "id"
	createdKey          = "created"
	createdTimestampKey = "created_unix"
	userKey             = "user"
	canonicalUserKey    = "canonical_user"

	// Meta config
	metaURLKey = "meta_url"
	metaURLEnv = "SRHT_META_URL"
	metaURLDef = "https://meta.sr.ht/api"

	// Paste config
	pasteURLKey = "paste_url"
	pasteURLEnv = "SRHT_PASTE_URL"
	pasteURLDef = "https://paste.sr.ht/api"

	// Git config
	gitURLKey = "git_url"
	gitURLEnv = "SRHT_GIT_URL"
	gitURLDef = "https://git.sr.ht/api"
)

func provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			metaURLKey: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  metaURLDef,
				Description: fmt.Sprintf(
					`The URL to the SourceHut Meta API endpoint. It is required if using
					a private installation of SourceHut. The default is to use the cloud
					paste service. It can be provided via the %s environment variable.`,
					metaURLEnv),
			},
			pasteURLKey: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  pasteURLDef,
				Description: fmt.Sprintf(
					`The URL to the SourceHut Paste API endpoint. It is required if using
					a private installation of SourceHut. The default is to use the cloud
					paste service. It can be provided via the %s environment variable.`,
					pasteURLEnv),
			},
			gitURLKey: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  gitURLDef,
				Description: fmt.Sprintf(
					`The URL to the SourceHut Paste API endpoint. It is required if using
					a private installation of SourceHut. The default is to use the cloud
					git service. It can be provided via the %s environment variable.`,
					gitURLEnv),
			},
			tokenKey: {
				Type:     schema.TypeString,
				Optional: true,
				Description: fmt.Sprintf(
					`A SourceHut API personal access token. It is required to use most
					resources. It can be provided via the %s environment variable.`,
					tokenEnv),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			sshKeyName: resourceSSHKey(),
			pgpKeyName: resourcePGPKey(),
			repoName:   resourceRepo(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			pasteName: dataSourcePaste(),
			blobName:  dataSourceBlob(),
			userName:  dataSourceUser(),
			repoName:  dataSourceRepo(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	srhtClient := sourcehut.NewClient(
		sourcehut.Token(dataOrEnv(d, tokenKey, tokenEnv)),
		sourcehut.UserAgent("git.sr.ht/~wombelix/terraform-provider-sourcehut"),
	)

	pasteClient, err := paste.NewClient(
		paste.SrhtClient(srhtClient),
		paste.Base(dataOrEnv(d, pasteURLKey, pasteURLEnv)),
	)
	if err != nil {
		return nil, err
	}
	metaClient, err := meta.NewClient(
		meta.SrhtClient(srhtClient),
		meta.Base(dataOrEnv(d, metaURLKey, metaURLEnv)),
	)
	if err != nil {
		return nil, err
	}
	gitClient, err := git.NewClient(
		git.SrhtClient(srhtClient),
		git.Base(dataOrEnv(d, gitURLKey, gitURLEnv)),
	)
	if err != nil {
		return nil, err
	}

	return config{
		srhtClient:  srhtClient,
		metaClient:  metaClient,
		pasteClient: pasteClient,
		gitClient:   gitClient,
	}, nil
}

type config struct {
	srhtClient  sourcehut.Client
	metaClient  *meta.Client
	pasteClient *paste.Client
	gitClient   *git.Client
}

func dataOrEnv(d *schema.ResourceData, key, env string) string {
	var ret string
	if v, ok := d.Get(key).(string); ok {
		ret = v
	}
	if ret == "" {
		ret = os.Getenv(env)
	}
	return ret
}
