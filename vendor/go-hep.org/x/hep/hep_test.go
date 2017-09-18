// Copyright 2017 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hep

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGofmt(t *testing.T) {
	exe, err := exec.LookPath("goimports")
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			if e.Err == exec.ErrNotFound {
				exe, err = exec.LookPath("gofmt")
			}
		}
	}
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(exe, "-d", ".")
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf

	err = cmd.Run()
	if err != nil {
		t.Fatalf("error running %s:\n%s\n%v", exe, string(buf.Bytes()), err)
	}

	if len(buf.Bytes()) != 0 {
		t.Errorf("some files were not gofmt'ed:\n%s\n", string(buf.Bytes()))
	}
}
