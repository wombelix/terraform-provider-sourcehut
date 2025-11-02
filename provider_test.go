// SPDX-FileCopyrightText: 2019 The SourceHut API Contributors
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"os"
	"testing"
)

func init() {
	// Set TF_SCHEMA_PANIC_ON_ERROR as a sanity check on tests.
	if err := os.Setenv("TF_SCHEMA_PANIC_ON_ERROR", "true"); err != nil {
		panic(err)
	}
}

func TestProvider(t *testing.T) {
	if err := provider().InternalValidate(); err != nil {
		t.Error(err)
	}
}
