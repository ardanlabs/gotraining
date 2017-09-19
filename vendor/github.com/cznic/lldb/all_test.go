// Copyright 2014 The lldb Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lldb

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	testDbName = "_test.db"
	walName    = "_wal"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	_ = os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_ = os.Stderr.Sync()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

// ============================================================================

func now() time.Time { return time.Now() }

func hdump(b []byte) string {
	return hex.Dump(b)
}

func die() {
	os.Exit(1)
}

func stack() string {
	buf := make([]byte, 1<<16)
	return string(buf[:runtime.Stack(buf, false)])
}

func temp() (dir, name string) {
	dir, err := ioutil.TempDir("", "test-lldb-")
	if err != nil {
		panic(err)
	}

	return dir, filepath.Join(dir, "test.tmp")
}
