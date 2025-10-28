// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

import "time"

// User represents a sourcehut user
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	CanonicalName string    `json:"canonicalName"`
	Created       time.Time `json:"created"`
	Email         string    `json:"email"`
	URL           string    `json:"url"`
	Location      string    `json:"location"`
	Bio           string    `json:"bio"`
}
