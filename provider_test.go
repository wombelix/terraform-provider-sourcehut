// Copyright 2019 The SourceHut API Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"
)

func init() {
	// Set TF_SCHEMA_PANIC_ON_ERROR as a sanity check on tests.
	os.Setenv("TF_SCHEMA_PANIC_ON_ERROR", "true")
}

func TestProvider(t *testing.T) {
	if err := provider().InternalValidate(); err != nil {
		t.Error(err)
	}
}
