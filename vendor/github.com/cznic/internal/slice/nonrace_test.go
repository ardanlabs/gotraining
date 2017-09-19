// Copyright 2016 The Internal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !race

package slice

import (
	"testing"
)

// Does not work with race detector b/c things get purged from the pool.
func Test(t *testing.T) {
	a := [1 << 10]*[]byte{}
	m := map[*[]byte]struct{}{}
	pool := newBytes()
	for i := range a {
		p := pool.CGet(i).(*[]byte)
		if _, ok := m[p]; ok {
			t.Fatal(i)
		}

		a[i] = p
		m[p] = struct{}{}
		b := *p
		for j := range b {
			b[j] = 123
		}
	}
	for i := range a {
		pool.Put(a[i])
	}
	for i := range a {
		p := pool.CGet(i).(*[]byte)
		if _, ok := m[p]; !ok {
			t.Fatal(i)
		}

		delete(m, p)
		b := *p
		if g, e := len(b), i; g != e {
			t.Fatal(g, e)
		}

		for j, v := range b[:cap(b)] {
			if v != 0 {
				t.Fatal(i, j, v)
			}
		}
	}
}
