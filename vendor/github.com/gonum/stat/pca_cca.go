// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat

import (
	"errors"
	"math"

	"github.com/gonum/floats"
	"github.com/gonum/matrix"
	"github.com/gonum/matrix/mat64"
)

// PC is a type for computing and extracting the principal components of a
// matrix. The results of the principal components analysis are only valid
// if the call to PrincipalComponents was successful.
type PC struct {
	n, d    int
	weights []float64
	svd     *mat64.SVD
	ok      bool
}

// PrincipalComponents performs a weighted principal components analysis on the
// matrix of the input data which is represented as an n×d matrix a where each
// row is an observation and each column is a variable.
//
// PrincipalComponents centers the variables but does not scale the variance.
//
// The weights slice is used to weight the observations. If weights is nil, each
// weight is considered to have a value of one, otherwise the length of weights
// must match the number of observations or PrincipalComponents will panic.
//
// PrincipalComponents returns whether the analysis was successful.
func (c *PC) PrincipalComponents(a mat64.Matrix, weights []float64) (ok bool) {
	c.n, c.d = a.Dims()
	if weights != nil && len(weights) != c.n {
		panic("stat: len(weights) != observations")
	}

	c.svd, c.ok = svdFactorizeCentered(c.svd, a, weights)
	if c.ok {
		c.weights = append(c.weights[:0], weights...)
	}
	return c.ok
}

// Vectors returns the component direction vectors of a principal components
// analysis. The vectors are returned in the columns of a d×min(n, d) matrix.
// If dst is not nil it must either be zero-sized or be a d×min(n, d) matrix.
// dst will  be used as the destination for the direction vector data. If dst
// is nil, a new mat64.Dense is allocated for the destination.
func (c *PC) Vectors(dst *mat64.Dense) *mat64.Dense {
	if !c.ok {
		panic("stat: use of unsuccessful principal components analysis")
	}

	if dst == nil {
		dst = &mat64.Dense{}
	} else if d, n := dst.Dims(); (n != 0 || d != 0) && (d != c.d || n != min(c.n, c.d)) {
		panic(matrix.ErrShape)
	}
	dst.VFromSVD(c.svd)
	return dst
}

// Vars returns the column variances of the principal component scores,
// b * vecs, where b is a matrix with centered columns. Variances are returned
// in descending order.
// If dst is not nil it is used to store the variances and returned.
// Vars will panic if the receiver has not successfully performed a principal
// components analysis or dst is not nil and the length of dst is not min(n, d).
func (c *PC) Vars(dst []float64) []float64 {
	if !c.ok {
		panic("stat: use of unsuccessful principal components analysis")
	}
	if dst != nil && len(dst) != min(c.n, c.d) {
		panic("stat: length of slice does not match analysis")
	}

	dst = c.svd.Values(dst)
	var f float64
	if c.weights == nil {
		f = 1 / float64(c.n-1)
	} else {
		f = 1 / (floats.Sum(c.weights) - 1)
	}
	for i, v := range dst {
		dst[i] = f * v * v
	}
	return dst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// CC is a type for computing the canonical correlations of a pair of matrices.
// The results of the canonical correlation analysis are only valid
// if the call to CanonicalCorrelations was successful.
type CC struct {
	// n is the number of observations used to
	// construct the canonical correlations.
	n int

	// xd and yd are used for size checks.
	xd, yd int

	x, y, c *mat64.SVD
	ok      bool
}

// CanonicalCorrelations returns a CC which can provide the results of canonical
// correlation analysis of the input data x and y, columns of which should be
// interpretable as two sets of measurements on the same observations (rows).
// These observations are optionally weighted by weights.
//
// Canonical correlation analysis finds associations between two sets of
// variables on the same observations by finding linear combinations of the two
// sphered datasets that maximize the correlation between them.
//
// Some notation: let Xc and Yc denote the centered input data matrices x
// and y (column means subtracted from each column), let Sx and Sy denote the
// sample covariance matrices within x and y respectively, and let Sxy denote
// the covariance matrix between x and y. The sphered data can then be expressed
// as Xc * Sx^{-1/2} and Yc * Sy^{-1/2} respectively, and the correlation matrix
// between the sphered data is called the canonical correlation matrix,
// Sx^{-1/2} * Sxy * Sy^{-1/2}. In cases where S^{-1/2} is ambiguous for some
// covariance matrix S, S^{-1/2} is taken to be E * D^{-1/2} * E^T where S can
// be eigendecomposed as S = E * D * E^T.
//
// The canonical correlations are the correlations between the corresponding
// pairs of canonical variables and can be obtained with c.Corrs(). Canonical
// variables can be obtained by projecting the sphered data into the left and
// right eigenvectors of the canonical correlation matrix, and these
// eigenvectors can be obtained with c.Left(m, true) and c.Right(m, true)
// respectively. The canonical variables can also be obtained directly from the
// centered raw data by using the back-transformed eigenvectors which can be
// obtained with c.Left(m, false) and c.Right(m, false) respectively.
//
// The first pair of left and right eigenvectors of the canonical correlation
// matrix can be interpreted as directions into which the respective sphered
// data can be projected such that the correlation between the two projections
// is maximized. The second pair and onwards solve the same optimization but
// under the constraint that they are uncorrelated (orthogonal in sphered space)
// to previous projections.
//
// CanonicalCorrelations will panic if the inputs x and y do not have the same
// number of rows.
//
// The slice weights is used to weight the observations. If weights is nil, each
// weight is considered to have a value of one, otherwise the length of weights
// must match the number of observations (rows of both x and y) or
// CanonicalCorrelations will panic.
//
// More details can be found at
// https://en.wikipedia.org/wiki/Canonical_correlation
// or in Chapter 3 of
// Koch, Inge. Analysis of multivariate and high-dimensional data.
// Vol. 32. Cambridge University Press, 2013. ISBN: 9780521887939
func (c *CC) CanonicalCorrelations(x, y mat64.Matrix, weights []float64) error {
	var yn int
	c.n, c.xd = x.Dims()
	yn, c.yd = y.Dims()
	if c.n != yn {
		panic("stat: unequal number of observations")
	}
	if weights != nil && len(weights) != c.n {
		panic("stat: len(weights) != observations")
	}

	// Center and factorize x and y.
	c.x, c.ok = svdFactorizeCentered(c.x, x, weights)
	if !c.ok {
		return errors.New("stat: failed to factorize x")
	}
	c.y, c.ok = svdFactorizeCentered(c.y, y, weights)
	if !c.ok {
		return errors.New("stat: failed to factorize y")
	}
	var xu, xv, yu, yv mat64.Dense
	xu.UFromSVD(c.x)
	xv.VFromSVD(c.x)
	yu.UFromSVD(c.y)
	yv.VFromSVD(c.y)

	// Calculate and factorise the canonical correlation matrix.
	var ccor mat64.Dense
	ccor.Product(&xv, xu.T(), &yu, yv.T())
	if c.c == nil {
		c.c = &mat64.SVD{}
	}
	c.ok = c.c.Factorize(&ccor, matrix.SVDThin)
	if !c.ok {
		return errors.New("stat: failed to factorize ccor")
	}
	return nil
}

// Corrs returns the canonical correlations, using dst if it is not nil.
// If dst is not nil and len(dst) does not match the number of columns in
// the y input matrix, Corrs will panic.
func (c *CC) Corrs(dst []float64) []float64 {
	if !c.ok {
		panic("stat: canonical correlations missing or invalid")
	}

	if dst != nil && len(dst) != c.yd {
		panic("stat: length of destination does not match input dimension")
	}
	return c.c.Values(dst)
}

// Left returns the left eigenvectors of the canonical correlation matrix if
// spheredSpace is true. If spheredSpace is false it returns these eigenvectors
// back-transformed to the original data space.
// If dst is not nil it must either be zero-sized or be an xd×yd matrix where xd
// and yd are the number of variables in the input x and y matrices. dst will
// be used as the destination for the vector data. If dst is nil, a new
// mat64.Dense is allocated for the destination.
func (c *CC) Left(dst *mat64.Dense, spheredSpace bool) *mat64.Dense {
	if !c.ok || c.n < 2 {
		panic("stat: canonical correlations missing or invalid")
	}

	if dst == nil {
		dst = &mat64.Dense{}
	} else if d, n := dst.Dims(); (n != 0 || d != 0) && (n != c.yd || d != c.xd) {
		panic(matrix.ErrShape)
	}
	dst.UFromSVD(c.c)
	if spheredSpace {
		return dst
	}

	var xv mat64.Dense
	xs := c.x.Values(nil)
	xv.VFromSVD(c.x)

	scaleColsReciSqrt(&xv, xs)

	dst.Product(&xv, xv.T(), dst)
	dst.Scale(math.Sqrt(float64(c.n-1)), dst)
	return dst
}

// Right returns the right eigenvectors of the canonical correlation matrix if
// spheredSpace is true. If spheredSpace is false it returns these eigenvectors
// back-transformed to the original data space.
// If dst is not nil it must either be zero-sized or be an yd×yd matrix where yd
// is the number of variables in the input y matrix. dst will
// be used as the destination for the vector data. If dst is nil, a new
// mat64.Dense is allocated for the destination.
func (c *CC) Right(dst *mat64.Dense, spheredSpace bool) *mat64.Dense {
	if !c.ok || c.n < 2 {
		panic("stat: canonical correlations missing or invalid")
	}

	if dst == nil {
		dst = &mat64.Dense{}
	} else if d, n := dst.Dims(); (n != 0 || d != 0) && (n != c.yd || d != c.yd) {
		panic(matrix.ErrShape)
	}
	dst.VFromSVD(c.c)
	if spheredSpace {
		return dst
	}

	var yv mat64.Dense
	ys := c.y.Values(nil)
	yv.VFromSVD(c.y)

	scaleColsReciSqrt(&yv, ys)

	dst.Product(&yv, yv.T(), dst)
	dst.Scale(math.Sqrt(float64(c.n-1)), dst)
	return dst
}

func svdFactorizeCentered(work *mat64.SVD, m mat64.Matrix, weights []float64) (svd *mat64.SVD, ok bool) {
	n, d := m.Dims()
	centered := mat64.NewDense(n, d, nil)
	col := make([]float64, n)
	for j := 0; j < d; j++ {
		mat64.Col(col, j, m)
		floats.AddConst(-Mean(col, weights), col)
		centered.SetCol(j, col)
	}
	for i, w := range weights {
		floats.Scale(math.Sqrt(w), centered.RawRowView(i))
	}
	if work == nil {
		work = &mat64.SVD{}
	}
	ok = work.Factorize(centered, matrix.SVDThin)
	return work, ok
}

// scaleColsReciSqrt scales the columns of cols
// by the reciprocal square-root of vals.
func scaleColsReciSqrt(cols *mat64.Dense, vals []float64) {
	if cols == nil {
		panic("stat: input nil")
	}
	n, d := cols.Dims()
	if len(vals) != d {
		panic("stat: input length mismatch")
	}
	col := make([]float64, n)
	for j := 0; j < d; j++ {
		mat64.Col(col, j, cols)
		floats.Scale(math.Sqrt(1/vals[j]), col)
		cols.SetCol(j, col)
	}
}
