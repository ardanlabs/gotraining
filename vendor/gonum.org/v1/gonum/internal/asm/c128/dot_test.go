// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c128

import (
	"fmt"
	"testing"
)

var dotTests = []struct {
	x, y               []complex128
	wantu, wantc       complex128
	wantuRev, wantcRev complex128
	n                  int
}{
	{
		x:     []complex128{},
		y:     []complex128{},
		n:     0,
		wantu: 0, wantc: 0,
		wantuRev: 0, wantcRev: 0,
	},
	{
		x:     []complex128{1 + 1i},
		y:     []complex128{1 + 1i},
		n:     1,
		wantu: 0 + 2i, wantc: 2,
		wantuRev: 0 + 2i, wantcRev: 2,
	},
	{
		x:     []complex128{1 + 2i},
		y:     []complex128{1 + 1i},
		n:     1,
		wantu: -1 + 3i, wantc: 3 - 1i,
		wantuRev: -1 + 3i, wantcRev: 3 - 1i,
	},
	{
		x:     []complex128{1 + 2i, 3 + 4i, 5 + 6i, 7 + 8i, 9 + 10i, 11 + 12i, 13 + 14i, 15 + 16i, 17 + 18i, 19 + 20i},
		y:     []complex128{1 + 2i, 3 + 4i, 5 + 6i, 7 + 8i, 9 + 10i, 11 + 12i, 13 + 14i, 15 + 16i, 17 + 18i, 19 + 20i},
		n:     10,
		wantu: -210 + 2860i, wantc: 2870 + 0i,
		wantuRev: -210 + 1540i, wantcRev: 1550 + 0i,
	},
	{
		x:     []complex128{1 + 1i, 1 + 1i, 1 + 2i, 1 + 1i, 1 + 1i, 1 + 1i, 1 + 3i, 1 + 1i, 1 + 1i, 1 + 4i},
		y:     []complex128{1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i},
		n:     10,
		wantu: -22 + 36i, wantc: 42 + 4i,
		wantuRev: -22 + 36i, wantcRev: 42 + 4i,
	},
	{
		x:     []complex128{1 + 1i, 1 + 1i, 2 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, 2 + 1i},
		y:     []complex128{1 + 2i, 1 + 2i, 1 + 3i, 1 + 2i, 1 + 3i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i},
		n:     10,
		wantu: -10 + 37i, wantc: 34 + 17i,
		wantuRev: -10 + 36i, wantcRev: 34 + 16i,
	},
	{
		x:     []complex128{1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, complex(inf, 1), 1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i, 1 + 1i},
		y:     []complex128{1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i, 1 + 2i},
		n:     10,
		wantu: complex(inf, inf), wantc: complex(inf, inf),
		wantuRev: complex(inf, inf), wantcRev: complex(inf, inf),
	},
}

func TestDotcUnitary(t *testing.T) {
	const gd = 1 + 5i
	for i, test := range dotTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, align.x, align.y)
			xgLn, ygLn := 4+align.x, 4+align.y
			xg, yg := guardVector(test.x, gd, xgLn), guardVector(test.y, gd, ygLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			res := DotcUnitary(x, y)
			if !same(res, test.wantc) {
				t.Errorf(msgVal, prefix, i, res, test.wantc)
			}
			if !isValidGuard(xg, gd, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, gd, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
		}
	}
}

func TestDotcInc(t *testing.T) {
	const gd, gdLn = 2 + 5i, 4
	for i, test := range dotTests {
		for _, inc := range newIncSet(1, 2, 5, 10, -1, -2, -5, -10) {
			xg, yg := guardIncVector(test.x, gd, inc.x, gdLn), guardIncVector(test.y, gd, inc.y, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
			want := test.wantc
			var ix, iy int
			if inc.x < 0 {
				ix, want = -inc.x*(test.n-1), test.wantcRev
			}
			if inc.y < 0 {
				iy, want = -inc.y*(test.n-1), test.wantcRev
			}
			prefix := fmt.Sprintf("Test %v (x:%v y:%v) (ix:%v iy:%v)", i, inc.x, inc.y, ix, iy)
			res := DotcInc(x, y, uintptr(test.n), uintptr(inc.x), uintptr(inc.y), uintptr(ix), uintptr(iy))
			if inc.x*inc.y > 0 {
				want = test.wantc
			}
			if !same(res, want) {
				t.Errorf(msgVal, prefix, i, res, want)
				t.Error(x, y)
			}
			checkValidIncGuard(t, xg, gd, inc.x, gdLn)
			checkValidIncGuard(t, yg, gd, inc.y, gdLn)
		}
	}
}

func TestDotuUnitary(t *testing.T) {
	const gd = 1 + 5i
	for i, test := range dotTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, align.x, align.y)
			xgLn, ygLn := 4+align.x, 4+align.y
			xg, yg := guardVector(test.x, gd, xgLn), guardVector(test.y, gd, ygLn)
			x, y := xg[xgLn:len(xg)-xgLn], yg[ygLn:len(yg)-ygLn]
			res := DotuUnitary(x, y)
			if !same(res, test.wantu) {
				t.Errorf(msgVal, prefix, i, res, test.wantu)
			}
			if !isValidGuard(xg, gd, xgLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:xgLn], xg[len(xg)-xgLn:])
			}
			if !isValidGuard(yg, gd, ygLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:ygLn], yg[len(yg)-ygLn:])
			}
		}
	}
}

func TestDotuInc(t *testing.T) {
	const gd, gdLn = 1 + 5i, 4
	for i, test := range dotTests {
		for _, inc := range newIncSet(1, 2, 5, 10, -1, -2, -5, -10) {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v)", i, inc.x, inc.y)
			xg, yg := guardIncVector(test.x, gd, inc.x, gdLn), guardIncVector(test.y, gd, inc.y, gdLn)
			x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
			want := test.wantc
			var ix, iy int
			if inc.x < 0 {
				ix, want = -inc.x*(test.n-1), test.wantuRev
			}
			if inc.y < 0 {
				iy, want = -inc.y*(test.n-1), test.wantuRev
			}
			res := DotuInc(x, y, uintptr(test.n), uintptr(inc.x), uintptr(inc.y), uintptr(ix), uintptr(iy))
			if inc.x*inc.y > 0 {
				want = test.wantu
			}
			if !same(res, want) {
				t.Errorf(msgVal, prefix, i, res, want)
			}
			checkValidIncGuard(t, xg, gd, inc.x, gdLn)
			checkValidIncGuard(t, yg, gd, inc.y, gdLn)
		}
	}
}
