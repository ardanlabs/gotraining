// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"fmt"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

// symView is a helper for getting a View of a SymDense.
type symView struct {
	sym *mat64.SymDense

	i, j, r, c int
}

func (s symView) Dims() (r, c int) { return s.r, s.c }

func (s symView) At(i, j int) float64 {
	if i < 0 || s.r <= i {
		panic("i out of bounds")
	}
	if j < 0 || s.c <= j {
		panic("j out of bounds")
	}
	return s.sym.At(s.i+i, s.j+j)
}

func (s symView) T() mat64.Matrix { return mat64.Transpose{s} }

func ExampleCC() {
	// This example is directly analogous to Example 3.5 on page 87 of
	// Koch, Inge. Analysis of multivariate and high-dimensional data.
	// Vol. 32. Cambridge University Press, 2013. ISBN: 9780521887939

	// bostonData is the Boston Housing Data of Harrison and Rubinfeld (1978)
	n, _ := bostonData.Dims()
	var xd, yd = 7, 4
	// The variables (columns) of bostonData can be partitioned into two sets:
	// those that deal with environmental/social variables (xdata), and those
	// that contain information regarding the individual (ydata). Because the
	// variables can be naturally partitioned in this way, these data are
	// appropriate for canonical correlation analysis. The columns (variables)
	// of xdata are, in order:
	//  per capita crime rate by town,
	//  proportion of non-retail business acres per town,
	//  nitric oxide concentration (parts per 10 million),
	//  weighted distances to Boston employment centres,
	//  index of accessibility to radial highways,
	//  pupil-teacher ratio by town, and
	//  proportion of blacks by town.
	xdata := bostonData.Slice(0, n, 0, xd)

	// The columns (variables) of ydata are, in order:
	//  average number of rooms per dwelling,
	//  proportion of owner-occupied units built prior to 1940,
	//  full-value property-tax rate per $10000, and
	//  median value of owner-occupied homes in $1000s.
	ydata := bostonData.Slice(0, n, xd, xd+yd)

	// For comparison, calculate the correlation matrix for the original data.
	var cor mat64.SymDense
	stat.CorrelationMatrix(&cor, bostonData, nil)

	// Extract just those correlations that are between xdata and ydata.
	var corRaw = symView{sym: &cor, i: 0, j: xd, r: xd, c: yd}

	// Note that the strongest correlation between individual variables is 0.91
	// between the 5th variable of xdata (index of accessibility to radial
	// highways) and the 3rd variable of ydata (full-value property-tax rate per
	// $10000).
	fmt.Printf("corRaw = %.4f", mat64.Formatted(corRaw, mat64.Prefix("         ")))

	// Calculate the canonical correlations.
	var cc stat.CC
	err := cc.CanonicalCorrelations(xdata, ydata, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Unpack cc.
	ccors := cc.Corrs(nil)
	pVecs := cc.Left(nil, true)
	qVecs := cc.Right(nil, true)
	phiVs := cc.Left(nil, false)
	psiVs := cc.Right(nil, false)

	// Canonical Correlation Matrix, or the correlations between the sphered
	// data.
	var corSph mat64.Dense
	corSph.Clone(pVecs)
	col := make([]float64, xd)
	for j := 0; j < yd; j++ {
		mat64.Col(col, j, &corSph)
		floats.Scale(ccors[j], col)
		corSph.SetCol(j, col)
	}
	corSph.Product(&corSph, qVecs.T())
	fmt.Printf("\n\ncorSph = %.4f", mat64.Formatted(&corSph, mat64.Prefix("         ")))

	// Canonical Correlations. Note that the first canonical correlation is
	// 0.95, stronger than the greatest correlation in the original data, and
	// much stronger than the greatest correlation in the sphered data.
	fmt.Printf("\n\nccors = %.4f", ccors)

	// Left and right eigenvectors of the canonical correlation matrix.
	fmt.Printf("\n\npVecs = %.4f", mat64.Formatted(pVecs, mat64.Prefix("        ")))
	fmt.Printf("\n\nqVecs = %.4f", mat64.Formatted(qVecs, mat64.Prefix("        ")))

	// Canonical Correlation Transforms. These can be useful as they represent
	// the canonical variables as linear combinations of the original variables.
	fmt.Printf("\n\nphiVs = %.4f", mat64.Formatted(phiVs, mat64.Prefix("        ")))
	fmt.Printf("\n\npsiVs = %.4f", mat64.Formatted(psiVs, mat64.Prefix("        ")))

	// Output:
	// corRaw = ⎡-0.2192   0.3527   0.5828  -0.3883⎤
	//          ⎢-0.3917   0.6448   0.7208  -0.4837⎥
	//          ⎢-0.3022   0.7315   0.6680  -0.4273⎥
	//          ⎢ 0.2052  -0.7479  -0.5344   0.2499⎥
	//          ⎢-0.2098   0.4560   0.9102  -0.3816⎥
	//          ⎢-0.3555   0.2615   0.4609  -0.5078⎥
	//          ⎣ 0.1281  -0.2735  -0.4418   0.3335⎦
	//
	// corSph = ⎡ 0.0118   0.0525   0.2300  -0.1363⎤
	//          ⎢-0.1810   0.3213   0.3814  -0.1412⎥
	//          ⎢ 0.0166   0.2241   0.0104  -0.2235⎥
	//          ⎢ 0.0346  -0.5481  -0.0034  -0.1994⎥
	//          ⎢ 0.0303  -0.0956   0.7152   0.2039⎥
	//          ⎢-0.0298  -0.0022   0.0739  -0.3703⎥
	//          ⎣-0.1226  -0.0746  -0.3899   0.1541⎦
	//
	// ccors = [0.9451 0.6787 0.5714 0.2010]
	//
	// pVecs = ⎡-0.2574   0.0158   0.2122  -0.0946⎤
	//         ⎢-0.4837   0.3837   0.1474   0.6597⎥
	//         ⎢-0.0801   0.3494   0.3287  -0.2862⎥
	//         ⎢ 0.1278  -0.7337   0.4851   0.2248⎥
	//         ⎢-0.6969  -0.4342  -0.3603   0.0291⎥
	//         ⎢-0.0991   0.0503   0.6384   0.1022⎥
	//         ⎣ 0.4260   0.0323  -0.2290   0.6419⎦
	//
	// qVecs = ⎡ 0.0182  -0.1583  -0.0067  -0.9872⎤
	//         ⎢-0.2348   0.9483  -0.1462  -0.1554⎥
	//         ⎢-0.9701  -0.2406  -0.0252   0.0209⎥
	//         ⎣ 0.0593  -0.1330  -0.9889   0.0291⎦
	//
	// phiVs = ⎡-0.0027   0.0093   0.0490  -0.0155⎤
	//         ⎢-0.0429  -0.0242   0.0361   0.1839⎥
	//         ⎢-1.2248   5.6031   5.8094  -4.7927⎥
	//         ⎢-0.0044  -0.3424   0.4470   0.1150⎥
	//         ⎢-0.0742  -0.1193  -0.1116   0.0022⎥
	//         ⎢-0.0233   0.1046   0.3853  -0.0161⎥
	//         ⎣ 0.0001   0.0005  -0.0030   0.0082⎦
	//
	// psiVs = ⎡ 0.0302  -0.3002   0.0878  -1.9583⎤
	//         ⎢-0.0065   0.0392  -0.0118  -0.0061⎥
	//         ⎢-0.0052  -0.0046  -0.0023   0.0008⎥
	//         ⎣ 0.0020   0.0037  -0.1293   0.1038⎦
}
