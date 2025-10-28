// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.sr.ht/~emersion/gqlclient"
)

func TestCreateRepository(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var req struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}

		// Verify variables
		name := req.Variables["name"].(string)
		visibility := req.Variables["visibility"].(string)
		description := req.Variables["description"].(string)

		// Set proper content type
		w.Header().Set("Content-Type", "application/json")

		// Return mock response
		resp := map[string]interface{}{
			"data": map[string]interface{}{
				"createRepository": map[string]interface{}{
					"id":          1,
					"name":        name,
					"description": description,
					"visibility":  visibility,
					"created":     "2025-10-28T12:00:00Z",
					"updated":     "2025-10-28T12:00:00Z",
				},
			},
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	// Create client pointing to test server
	c := &Client{
		clients: map[Service]*gqlclient.Client{
			GitService: gqlclient.New(server.URL, &http.Client{
				Transport: &authedTransport{token: "test-token"},
			}),
		},
		token: "test-token",
	}

	// Test repository creation
	input := RepositoryInput{
		Name:        "test-repo",
		Description: "Test repository",
		Visibility:  "public",
	}

	repo, err := c.CreateRepository(context.Background(), input)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	// Verify response
	if repo.Name != input.Name {
		t.Errorf("Expected repo name %s, got %s", input.Name, repo.Name)
	}
	if repo.Description != input.Description {
		t.Errorf("Expected repo description %s, got %s", input.Description, repo.Description)
	}
	if repo.Visibility != input.Visibility {
		t.Errorf("Expected repo visibility %s, got %s", input.Visibility, repo.Visibility)
	}
}
