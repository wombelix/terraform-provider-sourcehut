// SPDX-FileCopyrightText: 2024 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import (
	"context"
	"fmt"
	"time"

	"git.sr.ht/~emersion/gqlclient"
)

// Paste represents a paste in the sourcehut API
type Paste struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
	Visibility string `json:"visibility"`
	User    *User     `json:"user"`
	Files   []File    `json:"files"`
}

// File represents a file in a paste
type File struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
	Contents string `json:"contents"`
}

// GetPaste retrieves metadata about a paste
func (c *Client) GetPaste(ctx context.Context, id string) (*Paste, error) {
	op := gqlclient.NewOperation(`
		query GetPaste($id: String!) {
			paste(id: $id) {
				id
				created
				user {
					username
					canonicalName
				}
				files {
					hash
				}
			}
		}
	`)
	op.Var("id", id)

	var resp struct {
		Paste *Paste `json:"paste"`
	}

	if err := c.Paste().Execute(ctx, op, &resp); err != nil {
		return nil, fmt.Errorf("failed to get paste: %w", err)
	}

	return resp.Paste, nil
}

// GetPasteBlob retrieves a specific file's content from a paste
func (c *Client) GetPasteBlob(ctx context.Context, id string, fileHash string) (*File, error) {
	op := gqlclient.NewOperation(`
		query GetPasteBlob($id: String!, $hash: String!) {
			paste(id: $id) {
				files(hash: $hash) {
					hash
					contents
				}
			}
		}
	`)
	op.Var("id", id)
	op.Var("hash", fileHash)

	var resp struct {
		Paste struct {
			Files []File `json:"files"`
		} `json:"paste"`
	}

	if err := c.Paste().Execute(ctx, op, &resp); err != nil {
		return nil, fmt.Errorf("failed to get paste blob: %w", err)
	}

	if len(resp.Paste.Files) == 0 {
		return nil, fmt.Errorf("file not found")
	}

	return &resp.Paste.Files[0], nil
}
func (c *Client) GetPastes(ctx context.Context) ([]Paste, error) {
	op := gqlclient.NewOperation(`
		query GetPastes {
			me {
				pastes {
					results {
						id
						created
						visibility
						files {
							filename
							hash
							contents
						}
					}
				}
			}
		}
	`)

	var resp struct {
		Me struct {
			Pastes struct {
				Results []Paste `json:"results"`
			} `json:"pastes"`
		} `json:"me"`
	}

	if err := c.Paste().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return resp.Me.Pastes.Results, nil
}
