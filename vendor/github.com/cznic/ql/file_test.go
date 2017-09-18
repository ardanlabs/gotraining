// Copyright (c) 2014 ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ql

import (
	"os"
	"testing"
)

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func TestWALRemoval(t *testing.T) {
	const tempDBName = "./_test_was_removal.db"
	wName := walName(tempDBName)
	defer os.Remove(tempDBName)
	defer os.Remove(wName)

	db, err := OpenFile(tempDBName, &Options{CanCreate: true})
	if err != nil {
		t.Fatalf("Cannot open db %s: %s\n", tempDBName, err)
	}
	db.Close()
	if !fileExists(wName) {
		t.Fatalf("Expect WAL file %s to exist but it doesn't", wName)
	}

	db, err = OpenFile(tempDBName, &Options{CanCreate: true, RemoveEmptyWAL: true})
	if err != nil {
		t.Fatalf("Cannot open db %s: %s\n", tempDBName, err)
	}
	db.Close()
	if fileExists(wName) {
		t.Fatalf("Expect WAL file %s to be removed but it still exists", wName)
	}
}
