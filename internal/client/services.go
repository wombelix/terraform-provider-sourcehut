// SPDX-FileCopyrightText: 2025 Dominik Wombacher <dominik@wombacher.cc>
//
// SPDX-License-Identifier: BSD-2-Clause

package client

// Service represents a sourcehut service endpoint
type Service string

const (
	// GitService represents git.sr.ht
	GitService Service = "git.sr.ht"
	// MetaService represents meta.sr.ht
	MetaService Service = "meta.sr.ht"
	// PasteService represents paste.sr.ht
	PasteService Service = "paste.sr.ht"
)
