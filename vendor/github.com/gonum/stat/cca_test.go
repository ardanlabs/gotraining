// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"testing"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

func TestCanonicalCorrelations(t *testing.T) {
tests:
	for i, test := range []struct {
		xdata     mat64.Matrix
		ydata     mat64.Matrix
		weights   []float64
		wantCorrs []float64
		wantpVecs *mat64.Dense
		wantqVecs *mat64.Dense
		wantphiVs *mat64.Dense
		wantpsiVs *mat64.Dense
		epsilon   float64
	}{
		// Test results verified using R.
		{ // Truncated iris data, Sepal vs Petal measurements.
			xdata: mat64.NewDense(10, 2, []float64{
				5.1, 3.5,
				4.9, 3.0,
				4.7, 3.2,
				4.6, 3.1,
				5.0, 3.6,
				5.4, 3.9,
				4.6, 3.4,
				5.0, 3.4,
				4.4, 2.9,
				4.9, 3.1,
			}),
			ydata: mat64.NewDense(10, 2, []float64{
				1.4, 0.2,
				1.4, 0.2,
				1.3, 0.2,
				1.5, 0.2,
				1.4, 0.2,
				1.7, 0.4,
				1.4, 0.3,
				1.5, 0.2,
				1.4, 0.2,
				1.5, 0.1,
			}),
			wantCorrs: []float64{0.7250624174504773, 0.5547679185730191},
			wantpVecs: mat64.NewDense(2, 2, []float64{
				0.0765914610875867, 0.9970625597666721,
				0.9970625597666721, -0.0765914610875868,
			}),
			wantqVecs: mat64.NewDense(2, 2, []float64{
				0.3075184850910837, 0.9515421069649439,
				0.9515421069649439, -0.3075184850910837,
			}),
			wantphiVs: mat64.NewDense(2, 2, []float64{
				-1.9794877596804641, 5.2016325219025124,
				4.5211829944066553, -2.7263663170835697,
			}),
			wantpsiVs: mat64.NewDense(2, 2, []float64{
				-0.0613084818030103, 10.8514169865438941,
				12.7209032660734298, -7.6793888180353775,
			}),
			epsilon: 1e-12,
		},
		// Test results compared to those results presented in examples by
		// Koch, Inge. Analysis of multivariate and high-dimensional data.
		// Vol. 32. Cambridge University Press, 2013. ISBN: 9780521887939
		{ // ASA Car Exposition Data of Ramos and Donoho (1983)
			// Displacement, Horsepower, Weight
			xdata: carData.Slice(0, 392, 0, 3),
			// Acceleration, MPG
			ydata:     carData.Slice(0, 392, 3, 5),
			wantCorrs: []float64{0.8782187384352336, 0.6328187219216761},
			wantpVecs: mat64.NewDense(3, 2, []float64{
				0.3218296374829181, 0.3947540257657075,
				0.4162807660635797, 0.7573719053303306,
				0.8503740401982725, -0.5201509936144236,
			}),
			wantqVecs: mat64.NewDense(2, 2, []float64{
				-0.5161984172278830, -0.8564690269072364,
				-0.8564690269072364, 0.5161984172278830,
			}),
			wantphiVs: mat64.NewDense(3, 2, []float64{
				0.0025033152994308, 0.0047795464118615,
				0.0201923608080173, 0.0409150208725958,
				-0.0000247374128745, -0.0026766435161875,
			}),
			wantpsiVs: mat64.NewDense(2, 2, []float64{
				-0.1666196759760772, -0.3637393866139658,
				-0.0915512109649727, 0.1077863777929168,
			}),
			epsilon: 1e-12,
		},
		// Test results compared to those results presented in examples by
		// Koch, Inge. Analysis of multivariate and high-dimensional data.
		// Vol. 32. Cambridge University Press, 2013. ISBN: 9780521887939
		{ // Boston Housing Data of Harrison and Rubinfeld (1978)
			// Per capita crime rate by town,
			// Proportion of non-retail business acres per town,
			// Nitric oxide concentration (parts per 10 million),
			// Weighted distances to Boston employment centres,
			// Index of accessibility to radial highways,
			// Pupil-teacher ratio by town, Proportion of blacks by town
			xdata: bostonData.Slice(0, 506, 0, 7),
			// Average number of rooms per dwelling,
			// Proportion of owner-occupied units built prior to 1940,
			// Full-value property-tax rate per $10000,
			// Median value of owner-occupied homes in $1000s
			ydata:     bostonData.Slice(0, 506, 7, 11),
			wantCorrs: []float64{0.9451239443886021, 0.6786622733370654, 0.5714338361583764, 0.2009739704710440},
			wantpVecs: mat64.NewDense(7, 4, []float64{
				-0.2574391924541903, 0.0158477516621194, 0.2122169934631024, -0.0945733803894706,
				-0.4836594430018478, 0.3837101908138468, 0.1474448317415911, 0.6597324886718275,
				-0.0800776365873296, 0.3493556742809252, 0.3287336458109373, -0.2862040444334655,
				0.1277586360386374, -0.7337427663667596, 0.4851134819037011, 0.2247964865970192,
				-0.6969432006136684, -0.4341748776002893, -0.3602872887636357, 0.0290661608626292,
				-0.0990903250057199, 0.0503411215453873, 0.6384330631742202, 0.1022367136218303,
				0.4260459963765036, 0.0323334351308141, -0.2289527516030810, 0.6419232947608805,
			}),
			wantqVecs: mat64.NewDense(4, 4, []float64{
				0.0181660502363264, -0.1583489460479038, -0.0066723577642883, -0.9871935400650649,
				-0.2347699045986119, 0.9483314614936594, -0.1462420505631345, -0.1554470767919033,
				-0.9700704038477141, -0.2406071741000039, -0.0251838984227037, 0.0209134074358349,
				0.0593000682318482, -0.1330460003097728, -0.9889057151969489, 0.0291161494720761,
			}),
			wantphiVs: mat64.NewDense(7, 4, []float64{
				-0.0027462234108197, 0.0093444513500898, 0.0489643932714296, -0.0154967189805819,
				-0.0428564455279537, -0.0241708702119420, 0.0360723472093996, 0.1838983230588095,
				-1.2248435648802380, 5.6030921364723980, 5.8094144583797025, -4.7926812190419676,
				-0.0043684825094649, -0.3424101164977618, 0.4469961215717917, 0.1150161814353696,
				-0.0741534069521954, -0.1193135794923700, -0.1115518305471460, 0.0021638758323088,
				-0.0233270323101624, 0.1046330818178399, 0.3853045975077387, -0.0160927870102877,
				0.0001293051387859, 0.0004540746921446, -0.0030296315865440, 0.0081895477974654,
			}),
			wantpsiVs: mat64.NewDense(4, 4, []float64{
				0.0301593362017375, -0.3002219289647127, 0.0878217377593682, -1.9583226531517062,
				-0.0065483104073892, 0.0392212086716247, -0.0117570776209991, -0.0061113064481860,
				-0.0052075523350125, -0.0045770200452960, -0.0022762313289592, 0.0008441873006821,
				0.0020111735096327, 0.0037352799829930, -0.1292578071621794, 0.1037709056329765,
			}),
			epsilon: 1e-12,
		},
	} {
		var cc stat.CC
		var corrs []float64
		var pVecs, qVecs *mat64.Dense
		var phiVs, psiVs *mat64.Dense
		for j := 0; j < 2; j++ {
			err := cc.CanonicalCorrelations(test.xdata, test.ydata, test.weights)
			if err != nil {
				t.Errorf("%d use %d: unexpected error: %v", i, j, err)
				continue tests
			}

			corrs = cc.Corrs(corrs)
			pVecs = cc.Left(pVecs, true)
			qVecs = cc.Right(qVecs, true)
			phiVs = cc.Left(phiVs, false)
			psiVs = cc.Right(psiVs, false)

			if !floats.EqualApprox(corrs, test.wantCorrs, test.epsilon) {
				t.Errorf("%d use %d: unexpected variance result got:%v, want:%v",
					i, j, corrs, test.wantCorrs)
			}
			if !mat64.EqualApprox(pVecs, test.wantpVecs, test.epsilon) {
				t.Errorf("%d use %d: unexpected CCA result got:\n%v\nwant:\n%v",
					i, j, mat64.Formatted(pVecs), mat64.Formatted(test.wantpVecs))
			}
			if !mat64.EqualApprox(qVecs, test.wantqVecs, test.epsilon) {
				t.Errorf("%d use %d: unexpected CCA result got:\n%v\nwant:\n%v",
					i, j, mat64.Formatted(qVecs), mat64.Formatted(test.wantqVecs))
			}
			if !mat64.EqualApprox(phiVs, test.wantphiVs, test.epsilon) {
				t.Errorf("%d use %d: unexpected CCA result got:\n%v\nwant:\n%v",
					i, j, mat64.Formatted(phiVs), mat64.Formatted(test.wantphiVs))
			}
			if !mat64.EqualApprox(psiVs, test.wantpsiVs, test.epsilon) {
				t.Errorf("%d use %d: unexpected CCA result got:\n%v\nwant:\n%v",
					i, j, mat64.Formatted(psiVs), mat64.Formatted(test.wantpsiVs))
			}
		}
	}
}
