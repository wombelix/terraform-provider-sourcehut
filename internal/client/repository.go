// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import (
	"context"
	"time"

	"git.sr.ht/~emersion/gqlclient"
)

// Repository represents a sourcehut git repository
type Repository struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Visibility  string    `json:"visibility"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Subject     string    `json:"subject,omitempty"`
}

// RepositoryInput represents the input parameters for repository operations
type RepositoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Visibility  string `json:"visibility"`
}

// CreateRepository creates a new git repository
func (c *Client) CreateRepository(ctx context.Context, input RepositoryInput) (*Repository, error) {
	op := gqlclient.NewOperation(`
		mutation CreateRepo($name: String!, $visibility: Visibility!, $description: String) {
			createRepository(
				name: $name,
				visibility: $visibility,
				description: $description
			) {
				id
				name
				description
				visibility
				created
				updated
				owner {
					canonicalName
				}
			}
		}
	`)

	op.Var("name", input.Name)
	op.Var("visibility", input.Visibility)
	if input.Description != "" {
		op.Var("description", input.Description)
	}

	var resp struct {
		CreateRepository Repository `json:"createRepository"`
	}

	if err := c.Git().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.CreateRepository, nil
}

// GetRepository retrieves a repository by owner and name
func (c *Client) GetRepository(ctx context.Context, name string) (*Repository, error) {
	op := gqlclient.NewOperation(`
		query GetRepo($name: String!) {
			me {
				repository(name: $name) {
					id
					name
					description
					visibility
					created
					updated
				}
			}
		}
	`)

	op.Var("name", name)

	var resp struct {
		Me struct {
			Repository Repository `json:"repository"`
		} `json:"me"`
	}

	if err := c.Git().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.Me.Repository, nil
}

// UpdateRepository updates an existing repository
func (c *Client) UpdateRepository(ctx context.Context, id int, input RepositoryInput) (*Repository, error) {
	op := gqlclient.NewOperation(`
		mutation UpdateRepo($id: Int!, $input: RepoInput!) {
			updateRepository(id: $id, input: $input) {
				id
				name
				description
				visibility
				created
				updated
			}
		}
	`)

	op.Var("id", id)
	op.Var("input", map[string]interface{}{
		"name":        input.Name,
		"description": input.Description,
		"visibility":  input.Visibility,
	})

	var resp struct {
		UpdateRepository Repository `json:"updateRepository"`
	}

	if err := c.Git().Execute(ctx, op, &resp); err != nil {
		return nil, err
	}

	return &resp.UpdateRepository, nil
}

// DeleteRepository deletes a repository by ID
func (c *Client) DeleteRepository(ctx context.Context, id int) error {
	op := gqlclient.NewOperation(`
		mutation DeleteRepo($id: Int!) {
			deleteRepository(id: $id) {
				id
			}
		}
	`)

	op.Var("id", id)

	return c.Git().Execute(ctx, op, nil)
}
