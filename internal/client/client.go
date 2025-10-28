// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import (
	"fmt"
	"net/http"

	"git.sr.ht/~emersion/gqlclient"
)

// Client handles GraphQL API communication with sourcehut services
type Client struct {
	clients map[Service]*gqlclient.Client
	token   string
}

// NewClient creates a new sourcehut GraphQL API client
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	c := &Client{
		clients: make(map[Service]*gqlclient.Client),
		token:   token,
	}

	return c, nil
}

// getClient returns a GraphQL client for the specified service
func (c *Client) getClient(service Service) *gqlclient.Client {
	if client, exists := c.clients[service]; exists {
		return client
	}

	client := gqlclient.New(fmt.Sprintf("https://%s/query", service), &http.Client{
		Transport: &authedTransport{token: c.token},
	})
	c.clients[service] = client

	return client
}

// Git returns a GraphQL client for git.sr.ht
func (c *Client) Git() *gqlclient.Client {
	return c.getClient(GitService)
}

// Meta returns a GraphQL client for meta.sr.ht
func (c *Client) Meta() *gqlclient.Client {
	return c.getClient(MetaService)
}

// Paste returns a GraphQL client for paste.sr.ht
func (c *Client) Paste() *gqlclient.Client {
	return c.getClient(PasteService)
}

type authedTransport struct {
	token     string
	transport http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.transport == nil {
		t.transport = http.DefaultTransport
	}
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.transport.RoundTrip(req)
}
