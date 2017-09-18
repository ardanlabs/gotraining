// Copyright 2016 The Internal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buffer

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
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
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

func test(t testing.TB, allocs, goroutines int) {
	ready := make(chan int, goroutines)
	run := make(chan int)
	done := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			a := rand.Perm(allocs)
			ready <- 1
			<-run
			for _, v := range a {
				p := Get(v)
				b := *p
				if g, e := len(b), v; g != e {
					t.Error(g, e)
					break
				}

				Put(p)
			}
			done <- 1
		}()
	}
	for i := 0; i < goroutines; i++ {
		<-ready
	}
	close(run)
	for i := 0; i < goroutines; i++ {
		<-done
	}
}

func Test2(t *testing.T) {
	test(t, 1<<15, 32)
}

func Benchmark1(b *testing.B) {
	const (
		allocs     = 1000
		goroutines = 100
	)
	for i := 0; i < b.N; i++ {
		test(b, allocs, goroutines)
	}
	b.SetBytes(goroutines * (allocs*allocs + allocs) / 2)
}
