// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"git.sr.ht/~samwhited/sourcehut-go"
	"git.sr.ht/~samwhited/sourcehut-go/paste"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// All config
	tokenKey = "token"
	tokenEnv = "SRHT_TOKEN"

	// Common key names
	idKey               = "id"
	createdKey          = "created"
	createdTimestampKey = "created_unix"

	// Paste config
	pasteURLKey = "paste_url"
	pasteURLEnv = "SRHT_PASTE_URL"
	pasteURLDef = "https://paste.sr.ht/api"
)

func provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
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
			tokenKey: {
				Type:     schema.TypeString,
				Optional: true,
				Description: fmt.Sprintf(
					`A SourceHut API personal access token. It is required to use most
					resources. It can be provided via the %s environment variable.`,
					tokenEnv),
			},
		},
		//ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			pasteName: dataSourcePaste(),
			blobName:  dataSourceBlob(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	srhtClient := sourcehut.NewClient(
		sourcehut.Token(dataOrEnv(d, tokenKey, tokenEnv)),
		sourcehut.UserAgent("git.sr.ht/~samwhited/terraform-provider-sourcehut"),
	)

	pasteClient, err := paste.NewClient(
		paste.SrhtClient(srhtClient),
		paste.Base(dataOrEnv(d, pasteURLKey, pasteURLEnv)),
	)
	if err != nil {
		return nil, err
	}

	return config{
		srhtClient:  srhtClient,
		pasteClient: pasteClient,
	}, nil
}

type config struct {
	srhtClient  sourcehut.Client
	pasteClient *paste.Client
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
