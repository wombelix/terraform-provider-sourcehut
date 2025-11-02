// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import (
	"context"
	"fmt"
	"time"

	"git.sr.ht/~emersion/gqlclient"
)

// SSHKey represents a sourcehut SSH key
type SSHKey struct {
	ID          int       `json:"id"`
	Created     time.Time `json:"created"`
	LastUsed    time.Time `json:"lastUsed"`
	Key         string    `json:"key"`
	Fingerprint string    `json:"fingerprint"`
	Comment     string    `json:"comment"`
}

// PGPKey represents a sourcehut PGP key
type PGPKey struct {
	ID          int       `json:"id"`
	Created     time.Time `json:"created"`
	Key         string    `json:"key"`
	Fingerprint string    `json:"fingerprint"`
}

// CreateSSHKey creates a new SSH key
func (c *Client) CreateSSHKey(ctx context.Context, key string) (*SSHKey, error) {
	op := gqlclient.NewOperation(`
		mutation CreateSSHKey($key: String!) {
			createSSHKey(key: $key) {
				id
				created
				key
				fingerprint
				comment
			}
		}
	`)

	op.Var("key", key)

	var resp struct {
		CreateSSHKey SSHKey `json:"createSSHKey"`
	}

	if err := c.Meta().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.CreateSSHKey, nil
}

// GetSSHKey retrieves an SSH key by ID
func (c *Client) GetSSHKey(ctx context.Context, id int) (*SSHKey, error) {
	op := gqlclient.NewOperation(`
		query GetSSHKey($id: Int!) {
			me {
				sshKeys {
					results {
						id
						created
						lastUsed
						key
						fingerprint
						comment
					}
				}
			}
		}
	`)

	op.Var("id", id)

	var resp struct {
		Me struct {
			SSHKeys struct {
				Results []SSHKey `json:"results"`
			} `json:"sshKeys"`
		} `json:"me"`
	}

	if err := c.Meta().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	// Find the key with matching ID
	for _, key := range resp.Me.SSHKeys.Results {
		if key.ID == id {
			return &key, nil
		}
	}

	return nil, nil
}

// DeleteSSHKey deletes an SSH key by ID
func (c *Client) DeleteSSHKey(ctx context.Context, id int) error {
	op := gqlclient.NewOperation(`
		mutation DeleteSSHKey($id: Int!) {
			deleteSSHKey(id: $id) {
				id
			}
		}
	`)

	op.Var("id", id)

	return c.Meta().Execute(ctx, op, nil)
}

// CreatePGPKey creates a new PGP key
func (c *Client) CreatePGPKey(ctx context.Context, key string) (*PGPKey, error) {
	op := gqlclient.NewOperation(`
		mutation CreatePGPKey($key: String!) {
			createPGPKey(key: $key) {
				id
				created
				key
				fingerprint
			}
		}
	`)

	op.Var("key", key)

	var resp struct {
		CreatePGPKey PGPKey `json:"createPGPKey"`
	}

	if err := c.Meta().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.CreatePGPKey, nil
}

// GetPGPKey retrieves a PGP key by ID
func (c *Client) GetPGPKey(ctx context.Context, id int) (*PGPKey, error) {
	op := gqlclient.NewOperation(`
		query GetPGPKey($id: Int!) {
			me {
				pgpKeys {
					results {
						id
						created
						key
						fingerprint
					}
				}
			}
		}
	`)

	op.Var("id", id)

	var resp struct {
		Me struct {
			PGPKeys struct {
				Results []PGPKey `json:"results"`
			} `json:"pgpKeys"`
		} `json:"me"`
	}

	if err := c.Meta().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	// Find the key with matching ID
	for _, key := range resp.Me.PGPKeys.Results {
		if key.ID == id {
			result := key
			return &result, nil
		}
	}

	return nil, fmt.Errorf("PGP key with ID %d not found", id)
}

// DeletePGPKey deletes a PGP key by ID
func (c *Client) DeletePGPKey(ctx context.Context, id int) error {
	op := gqlclient.NewOperation(`
		mutation DeletePGPKey($id: Int!) {
			deletePGPKey(id: $id) {
				id
			}
		}
	`)

	op.Var("id", id)

	return c.Meta().Execute(ctx, op, nil)
}

// GetCurrentUser retrieves the authenticated user's profile
func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	op := gqlclient.NewOperation(`
		query GetCurrentUser {
			me {
				id
				username
				canonicalName
				created
				email
				url
				location
				bio
				pgpKeys {
					results {
						id
						key
						fingerprint
					}
				}
			}
		}
	`)

	var resp struct {
		Me User `json:"me"`
	}

	if err := c.Meta().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.Me, nil
}
