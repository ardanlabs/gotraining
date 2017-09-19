// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix"
)

func TestDenseOverlaps(t *testing.T) {
	type view struct {
		i, j, r, c int
		*Dense
	}

	rnd := rand.New(rand.NewSource(1))

	for r := 1; r < 20; r++ {
		for c := 1; c < 20; c++ {
			m := NewDense(r, c, nil)
			panicked, message := panics(func() { m.checkOverlap(m.RawMatrix()) })
			if !panicked {
				t.Error("expected matrix overlap with self")
			}
			if message != regionIdentity {
				t.Errorf("unexpected panic message for self overlap: got: %q want: %q", message, regionIdentity)
			}

			for i := 0; i < 1000; i++ {
				var views [2]view
				for k := range views {
					if r > 1 {
						views[k].i = rnd.Intn(r - 1)
						views[k].r = rnd.Intn(r-views[k].i-1) + 1
					} else {
						views[k].r = 1
					}
					if c > 1 {
						views[k].j = rnd.Intn(c - 1)
						views[k].c = rnd.Intn(c-views[k].j-1) + 1
					} else {
						views[k].c = 1
					}
					views[k].Dense = m.Slice(views[k].i, views[k].i+views[k].r, views[k].j, views[k].j+views[k].c).(*Dense)

					panicked, _ = panics(func() { m.checkOverlap(views[k].RawMatrix()) })
					if !panicked {
						t.Errorf("expected matrix (%d×%d) overlap with view {rows=%d:%d, cols=%d:%d}",
							r, c, views[k].i, views[k].i+views[k].r, views[k].j, views[k].j+views[k].c)
					}
					panicked, _ = panics(func() { views[k].checkOverlap(m.RawMatrix()) })
					if !panicked {
						t.Errorf("expected view {rows=%d:%d, cols=%d:%d} overlap with parent (%d×%d)",
							views[k].i, views[k].i+views[k].r, views[k].j, views[k].j+views[k].c, r, c)
					}
				}

				overlapRows := intervalsOverlap(
					interval{views[0].i, views[0].i + views[0].r},
					interval{views[1].i, views[1].i + views[1].r},
				)
				overlapCols := intervalsOverlap(
					interval{views[0].j, views[0].j + views[0].c},
					interval{views[1].j, views[1].j + views[1].c},
				)
				want := overlapRows && overlapCols

				for k, v := range views {
					w := views[1-k]
					got, _ := panics(func() { v.checkOverlap(w.RawMatrix()) })
					if got != want {
						t.Errorf("unexpected result for overlap test for {rows=%d:%d, cols=%d:%d} with {rows=%d:%d, cols=%d:%d}: got: %t want: %t",
							v.i, v.i+v.r, v.j, v.j+v.c,
							w.i, w.i+w.r, w.j, w.j+w.c,
							got, want)
					}
				}
			}
		}
	}
}

func TestTriDenseOverlaps(t *testing.T) {
	type view struct {
		i, j, n int
		*TriDense
	}

	rnd := rand.New(rand.NewSource(1))

	for _, parentKind := range []matrix.TriKind{matrix.Upper, matrix.Lower} {
		for n := 1; n < 20; n++ {
			data := make([]float64, n*n)
			for i := range data {
				data[i] = float64(i + 1)
			}
			m := NewDense(n, n, data)
			mt := denseAsTriDense(m, parentKind)
			panicked, message := panics(func() { mt.checkOverlap(mt.RawTriangular()) })
			if !panicked {
				t.Error("expected matrix overlap with self")
			}
			if message != regionIdentity {
				t.Errorf("unexpected panic message for self overlap: got: %q want: %q", message, regionIdentity)
			}

			for i := 0; i < 1000; i++ {
				var views [2]view
				for k := range views {
					if n > 1 {
						views[k].i = rnd.Intn(n - 1)
						views[k].j = rnd.Intn(n - 1)
						views[k].n = rnd.Intn(n-max(views[k].i, views[k].j)-1) + 1
					} else {
						views[k].n = 1
					}
					viewKind := []matrix.TriKind{matrix.Upper, matrix.Lower}[rnd.Intn(2)]
					views[k].TriDense = denseAsTriDense(
						m.Slice(views[k].i, views[k].i+views[k].n, views[k].j, views[k].j+views[k].n).(*Dense),
						viewKind)

					wantPanick := overlapsParentTriangle(views[k].i, views[k].j, views[k].n, parentKind, viewKind)

					panicked, _ = panics(func() { mt.checkOverlap(views[k].RawTriangular()) })
					if panicked != wantPanick {
						t.Errorf("unexpected (%d×%d)%s overlap with view {rows=%d:%d, cols=%d:%d}%s got:%t want:%t\n% v\n\n% v\n",
							n, n, kindString(parentKind),
							views[k].i, views[k].i+views[k].n, views[k].j, views[k].j+views[k].n, kindString(viewKind),
							panicked, wantPanick,
							Formatted(mt), Formatted(views[k].TriDense))
					}
					panicked, _ = panics(func() { views[k].checkOverlap(mt.RawTriangular()) })
					if panicked != wantPanick {
						t.Errorf("unexpected {rows=%d:%d, cols=%d:%d}%s overlap with parent (%d×%d)%s got:%t want:%t\n% v\n\n% v\n",
							views[k].i, views[k].i+views[k].n, views[k].j, views[k].j+views[k].n, kindString(viewKind),
							n, n, kindString(parentKind),
							panicked, wantPanick,
							Formatted(views[k].TriDense), Formatted(mt))
					}
				}

				want := overlapSiblingTriangles(
					views[0].i, views[0].j, views[0].n, views[0].mat.Uplo == blas.Upper,
					views[1].i, views[1].j, views[1].n, views[1].mat.Uplo == blas.Upper,
				)

				for k, v := range views {
					w := views[1-k]
					got, _ := panics(func() { v.checkOverlap(w.RawTriangular()) })
					if got != want {
						t.Errorf("unexpected result for overlap test for {rows=%d:%d, cols=%d:%d}%s with {rows=%d:%d, cols=%d:%d}%s: got:%t want:%t\n% v\n\n% v\n",
							v.i, v.i+v.n, v.j, v.j+v.n, kindString(v.mat.Uplo == blas.Upper),
							w.i, w.i+w.n, w.j, w.j+w.n, kindString(w.mat.Uplo == blas.Upper),
							got, want,
							Formatted(v.TriDense), Formatted(w.TriDense))
					}
				}
			}
		}
	}
}

type interval struct{ from, to int }

func intervalsOverlap(a, b interval) bool {
	return a.to > b.from && b.to > a.from
}

func overlapsParentTriangle(i, j, n int, parent, view matrix.TriKind) bool {
	switch parent {
	case matrix.Upper:
		if i <= j {
			return true
		}
		if view == matrix.Upper {
			return i < j+n
		}

	case matrix.Lower:
		if i >= j {
			return true
		}
		if view == matrix.Lower {
			return i+n > j
		}
	}

	return false
}

func overlapSiblingTriangles(ai, aj, an int, aKind matrix.TriKind, bi, bj, bn int, bKind matrix.TriKind) bool {
	for i := max(ai, bi); i < min(ai+an, bi+bn); i++ {
		var a, b interval

		if aKind == matrix.Upper {
			a = interval{from: aj - ai + i, to: aj + an}
		} else {
			a = interval{from: aj, to: aj - ai + i + 1}
		}

		if bKind == matrix.Upper {
			b = interval{from: bj - bi + i, to: bj + bn}
		} else {
			b = interval{from: bj, to: bj - bi + i + 1}
		}

		if intervalsOverlap(a, b) {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func kindString(k matrix.TriKind) string {
	if k == matrix.Upper {
		return "U"
	}
	return "L"
}

// See https://github.com/gonum/matrix/issues/359 for details.
func TestIssue359(t *testing.T) {
	for xi := 0; xi < 2; xi++ {
		for xj := 0; xj < 2; xj++ {
			for yi := 0; yi < 2; yi++ {
				for yj := 0; yj < 2; yj++ {
					a := NewDense(3, 3, []float64{
						1, 2, 3,
						4, 5, 6,
						7, 8, 9,
					})
					x := a.Slice(xi, xi+2, xj, xj+2).(*Dense)
					y := a.Slice(yi, yi+2, yj, yj+2).(*Dense)

					panicked, _ := panics(func() { x.checkOverlap(y.mat) })
					if !panicked {
						t.Errorf("expected panic for aliased with offsets x(%d,%d) y(%d,%d):\nx:\n%v\ny:\n%v",
							xi, xj, yi, yj, Formatted(x), Formatted(y),
						)
					}
				}
			}
		}
	}
}

// denseAsTriDense returns a triangular matrix derived from the
// square matrix m, with the orientation specified by kind.
func denseAsTriDense(m *Dense, kind matrix.TriKind) *TriDense {
	r, c := m.Dims()
	if r != c {
		panic(matrix.ErrShape)
	}
	n := r
	uplo := blas.Lower
	if kind == matrix.Upper {
		uplo = blas.Upper
	}
	return &TriDense{
		mat: blas64.Triangular{
			N:      n,
			Stride: m.mat.Stride,
			Data:   m.mat.Data,
			Uplo:   uplo,
			Diag:   blas.NonUnit,
		},
		cap: n,
	}
}
