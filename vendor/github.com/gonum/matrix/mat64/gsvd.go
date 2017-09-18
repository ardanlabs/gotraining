// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/lapack"
	"github.com/gonum/lapack/lapack64"
	"github.com/gonum/matrix"
)

// GSVD is a type for creating and using the Generalized Singular Value Decomposition
// (GSVD) of a matrix.
//
// The factorization is a linear transformation of the data sets from the given
// variable×sample spaces to reduced and diagonalized "eigenvariable"×"eigensample"
// spaces.
type GSVD struct {
	kind matrix.GSVDKind

	r, p, c, k, l int
	s1, s2        []float64
	a, b, u, v, q blas64.General

	work  []float64
	iwork []int
}

// Factorize computes the generalized singular value decomposition (GSVD) of the input
// the r×c matrix A and the p×c matrix B. The singular values of A and B are computed
// in all cases, while the singular vectors are optionally computed depending on the
// input kind.
//
// The full singular value decomposition (kind == GSVDU|GSVDV|GSVDQ) deconstructs A and B as
//  A = U * Σ₁ * [ 0 R ] * Q^T
//
//  B = V * Σ₂ * [ 0 R ] * Q^T
// where Σ₁ and Σ₂ are r×(k+l) and p×(k+l) diagonal matrices of singular values, and
// U, V and Q are r×r, p×p and c×c orthogonal matrices of singular vectors. k+l is the
// effective numerical rank of the matrix [ A^T B^T ]^T.
//
// It is frequently not necessary to compute the full GSVD. Computation time and
// storage costs can be reduced using the appropriate kind. Either only the singular
// values can be computed (kind == SVDNone), or in conjunction with specific singular
// vectors (kind bit set according to matrix.GSVDU, matrix.GSVDV and matrix.GSVDQ).
//
// Factorize returns whether the decomposition succeeded. If the decomposition
// failed, routines that require a successful factorization will panic.
func (gsvd *GSVD) Factorize(a, b Matrix, kind matrix.GSVDKind) (ok bool) {
	r, c := a.Dims()
	gsvd.r, gsvd.c = r, c
	p, c := b.Dims()
	gsvd.p = p
	if gsvd.c != c {
		panic(matrix.ErrShape)
	}
	var jobU, jobV, jobQ lapack.GSVDJob
	switch {
	default:
		panic("gsvd: bad input kind")
	case kind == matrix.GSVDNone:
		jobU = lapack.GSVDNone
		jobV = lapack.GSVDNone
		jobQ = lapack.GSVDNone
	case (matrix.GSVDU|matrix.GSVDV|matrix.GSVDQ)&kind != 0:
		if matrix.GSVDU&kind != 0 {
			jobU = lapack.GSVDU
			gsvd.u = blas64.General{
				Rows:   r,
				Cols:   r,
				Stride: r,
				Data:   use(gsvd.u.Data, r*r),
			}
		}
		if matrix.GSVDV&kind != 0 {
			jobV = lapack.GSVDV
			gsvd.v = blas64.General{
				Rows:   p,
				Cols:   p,
				Stride: p,
				Data:   use(gsvd.v.Data, p*p),
			}
		}
		if matrix.GSVDQ&kind != 0 {
			jobQ = lapack.GSVDQ
			gsvd.q = blas64.General{
				Rows:   c,
				Cols:   c,
				Stride: c,
				Data:   use(gsvd.q.Data, c*c),
			}
		}
	}

	// A and B are destroyed on call, so copy the matrices.
	aCopy := DenseCopyOf(a)
	bCopy := DenseCopyOf(b)

	gsvd.s1 = use(gsvd.s1, c)
	gsvd.s2 = use(gsvd.s2, c)

	gsvd.iwork = useInt(gsvd.iwork, c)

	gsvd.work = use(gsvd.work, 1)
	lapack64.Ggsvd3(jobU, jobV, jobQ, aCopy.mat, bCopy.mat, gsvd.s1, gsvd.s2, gsvd.u, gsvd.v, gsvd.q, gsvd.work, -1, gsvd.iwork)
	gsvd.work = use(gsvd.work, int(gsvd.work[0]))
	gsvd.k, gsvd.l, ok = lapack64.Ggsvd3(jobU, jobV, jobQ, aCopy.mat, bCopy.mat, gsvd.s1, gsvd.s2, gsvd.u, gsvd.v, gsvd.q, gsvd.work, len(gsvd.work), gsvd.iwork)
	if ok {
		gsvd.a = aCopy.mat
		gsvd.b = bCopy.mat
		gsvd.kind = kind
	}
	return ok
}

// Kind returns the matrix.GSVDKind of the decomposition. If no decomposition has been
// computed, Kind returns 0.
func (gsvd *GSVD) Kind() matrix.GSVDKind {
	return gsvd.kind
}

// Rank returns the k and l terms of the rank of [ A^T B^T ]^T.
func (gsvd *GSVD) Rank() (k, l int) {
	return gsvd.k, gsvd.l
}

// GeneralizedValues returns the generalized singular values of the factorized matrices.
// If the input slice is non-nil, the values will be stored in-place into the slice.
// In this case, the slice must have length min(r,c)-k, and GeneralizedValues will
// panic with matrix.ErrSliceLengthMismatch otherwise. If the input slice is nil,
// a new slice of the appropriate length will be allocated and returned.
//
// GeneralizedValues will panic if the receiver does not contain a successful factorization.
func (gsvd *GSVD) GeneralizedValues(v []float64) []float64 {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	c := gsvd.c
	k := gsvd.k
	d := min(r, c)
	if v == nil {
		v = make([]float64, d-k)
	}
	if len(v) != d-k {
		panic(matrix.ErrSliceLengthMismatch)
	}
	floats.DivTo(v, gsvd.s1[k:d], gsvd.s2[k:d])
	return v
}

// ValuesA returns the singular values of the factorized A matrix.
// If the input slice is non-nil, the values will be stored in-place into the slice.
// In this case, the slice must have length min(r,c)-k, and ValuesA will panic with
// matrix.ErrSliceLengthMismatch otherwise. If the input slice is nil,
// a new slice of the appropriate length will be allocated and returned.
//
// ValuesA will panic if the receiver does not contain a successful factorization.
func (gsvd *GSVD) ValuesA(s []float64) []float64 {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	c := gsvd.c
	k := gsvd.k
	d := min(r, c)
	if s == nil {
		s = make([]float64, d-k)
	}
	if len(s) != d-k {
		panic(matrix.ErrSliceLengthMismatch)
	}
	copy(s, gsvd.s1[k:min(r, c)])
	return s
}

// ValuesB returns the singular values of the factorized B matrix.
// If the input slice is non-nil, the values will be stored in-place into the slice.
// In this case, the slice must have length min(r,c)-k, and ValuesB will panic with
// matrix.ErrSliceLengthMismatch otherwise. If the input slice is nil,
// a new slice of the appropriate length will be allocated and returned.
//
// ValuesB will panic if the receiver does not contain a successful factorization.
func (gsvd *GSVD) ValuesB(s []float64) []float64 {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	c := gsvd.c
	k := gsvd.k
	d := min(r, c)
	if s == nil {
		s = make([]float64, d-k)
	}
	if len(s) != d-k {
		panic(matrix.ErrSliceLengthMismatch)
	}
	copy(s, gsvd.s2[k:d])
	return s
}

// ZeroRFromGSVD extracts the matrix [ 0 R ] from the singular value decomposition, storing
// the result in-place into the receiver. [ 0 R ] is size (k+l)×c.
func (m *Dense) ZeroRFromGSVD(gsvd *GSVD) {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	c := gsvd.c
	k := gsvd.k
	l := gsvd.l
	h := min(k+l, r)
	m.reuseAsZeroed(k+l, c)
	a := Dense{
		mat:     gsvd.a,
		capRows: r,
		capCols: c,
	}
	m.Slice(0, h, c-k-l, c).(*Dense).
		Copy(a.Slice(0, h, c-k-l, c))
	if r < k+l {
		b := Dense{
			mat:     gsvd.b,
			capRows: gsvd.p,
			capCols: c,
		}
		m.Slice(r, k+l, c+r-k-l, c).(*Dense).
			Copy(b.Slice(r-k, l, c+r-k-l, c))
	}
}

// SigmaAFromGSVD extracts the matrix Σ₁ from the singular value decomposition, storing
// the result in-place into the receiver. Σ₁ is size r×(k+l).
func (m *Dense) SigmaAFromGSVD(gsvd *GSVD) {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	k := gsvd.k
	l := gsvd.l
	m.reuseAsZeroed(r, k+l)
	for i := 0; i < k; i++ {
		m.set(i, i, 1)
	}
	for i := k; i < min(r, k+l); i++ {
		m.set(i, i, gsvd.s1[i])
	}
}

// SigmaBFromGSVD extracts the matrix Σ₂ from the singular value decomposition, storing
// the result in-place into the receiver. Σ₂ is size p×(k+l).
func (m *Dense) SigmaBFromGSVD(gsvd *GSVD) {
	if gsvd.kind == 0 {
		panic("gsvd: no decomposition computed")
	}
	r := gsvd.r
	p := gsvd.p
	k := gsvd.k
	l := gsvd.l
	m.reuseAsZeroed(p, k+l)
	for i := 0; i < min(l, r-k); i++ {
		m.set(i, i+k, gsvd.s2[k+i])
	}
	for i := r - k; i < l; i++ {
		m.set(i, i+k, 1)
	}
}

// UFromGSVD extracts the matrix U from the singular value decomposition, storing
// the result in-place into the receiver. U is size r×r.
func (m *Dense) UFromGSVD(gsvd *GSVD) {
	if gsvd.kind&matrix.GSVDU == 0 {
		panic("mat64: improper GSVD kind")
	}
	r := gsvd.u.Rows
	c := gsvd.u.Cols
	m.reuseAs(r, c)

	tmp := &Dense{
		mat:     gsvd.u,
		capRows: r,
		capCols: c,
	}
	m.Copy(tmp)
}

// VFromGSVD extracts the matrix V from the singular value decomposition, storing
// the result in-place into the receiver. V is size p×p.
func (m *Dense) VFromGSVD(gsvd *GSVD) {
	if gsvd.kind&matrix.GSVDV == 0 {
		panic("mat64: improper GSVD kind")
	}
	r := gsvd.v.Rows
	c := gsvd.v.Cols
	m.reuseAs(r, c)

	tmp := &Dense{
		mat:     gsvd.v,
		capRows: r,
		capCols: c,
	}
	m.Copy(tmp)
}

// QFromGSVD extracts the matrix Q from the singular value decomposition, storing
// the result in-place into the receiver. Q is size c×c.
func (m *Dense) QFromGSVD(gsvd *GSVD) {
	if gsvd.kind&matrix.GSVDQ == 0 {
		panic("mat64: improper GSVD kind")
	}
	r := gsvd.q.Rows
	c := gsvd.q.Cols
	m.reuseAs(r, c)

	tmp := &Dense{
		mat:     gsvd.q,
		capRows: r,
		capCols: c,
	}
	m.Copy(tmp)
}
