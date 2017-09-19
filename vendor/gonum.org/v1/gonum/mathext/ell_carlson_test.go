// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"math/rand"
	"testing"
)

// Testing EllipticF (and EllipticRF) using the addition theorems from http://dlmf.nist.gov/19.11.i
func TestEllipticF(t *testing.T) {
	const tol = 1.0e-14
	rnd := rand.New(rand.NewSource(1))

	// The following EllipticF(pi/3,m), m=0.1(0.1)0.9 was computed in Maxima 5.38.0 using Bigfloat arithmetic.
	vF := [...]float64{
		1.0631390181954904767742338285104637431858016483079,
		1.0803778062523490005579242592072579594037132891908,
		1.0991352230920430074586978843452269008747645822123,
		1.1196949183404746257742176145632376703505764745654,
		1.1424290580457772555013955266260457822322036529624,
		1.1678400583161860445148860686430780757517286094732,
		1.1966306515644649360767197589467723191317720122309,
		1.2298294422249382706933871574135731278765534034979,
		1.2690359140762658660446752406901433173504503955036,
	}
	phi := math.Pi / 3
	for m := 1; m <= 9; m++ {
		mf := float64(m) / 10
		delta := math.Abs(EllipticF(phi, mf) - vF[m-1])
		if delta > tol {
			t.Fatalf("EllipticF(pi/3,m) test fail for m=%v", mf)
		}
	}

	for test := 0; test < 100; test++ {
		alpha := rnd.Float64() * math.Pi / 4
		beta := rnd.Float64() * math.Pi / 4
		for mi := 0; mi < 9999; mi++ {
			m := float64(mi) / 10000
			Fa := EllipticF(alpha, m)
			Fb := EllipticF(beta, m)
			sina, cosa := math.Sincos(alpha)
			sinb, cosb := math.Sincos(beta)
			tan := (sina*math.Sqrt(1-m*sinb*sinb) + sinb*math.Sqrt(1-m*sina*sina)) / (cosa + cosb)
			gamma := 2 * math.Atan(tan)
			Fg := EllipticF(gamma, m)
			delta := math.Abs(Fa + Fb - Fg)
			if delta > tol {
				t.Fatalf("EllipticF test fail for m=%v, alpha=%v, beta=%v", m, alpha, beta)
			}
		}
	}
}

// Testing EllipticE (and EllipticRF, EllipticRD) using the addition theorems from http://dlmf.nist.gov/19.11.i
func TestEllipticE(t *testing.T) {
	const tol = 1.0e-14
	rnd := rand.New(rand.NewSource(1))

	// The following EllipticE(pi/3,m), m=0.1(0.1)0.9 was computed in Maxima 5.38.0 using Bigfloat arithmetic.
	vE := [...]float64{
		1.0316510822817691068014397636905610074934300946730,
		1.0156973658341766636288643556414001451527597364432,
		9.9929636467826398814855428365155224243586391115108e-1,
		9.8240033979859736941287149003648737502960015189033e-1,
		9.6495145764299257550956863602992167490195750321518e-1,
		9.4687829659158090935158610908054896203271861698355e-1,
		9.2809053417715769009517654522979827392794124845027e-1,
		9.0847044378047233264777277954768245721857017157916e-1,
		8.8785835036531301307661603341327881634688308777383e-1,
	}
	phi := math.Pi / 3
	for m := 1; m <= 9; m++ {
		mf := float64(m) / 10
		delta := math.Abs(EllipticE(phi, mf) - vE[m-1])
		if delta > tol {
			t.Fatalf("EllipticE(pi/3,m) test fail for m=%v", mf)
		}
	}

	for test := 0; test < 100; test++ {
		alpha := rnd.Float64() * math.Pi / 4
		beta := rnd.Float64() * math.Pi / 4
		for mi := 0; mi < 9999; mi++ {
			m := float64(mi) / 10000
			Ea := EllipticE(alpha, m)
			Eb := EllipticE(beta, m)
			sina, cosa := math.Sincos(alpha)
			sinb, cosb := math.Sincos(beta)
			tan := (sina*math.Sqrt(1-m*sinb*sinb) + sinb*math.Sqrt(1-m*sina*sina)) / (cosa + cosb)
			gamma := 2 * math.Atan(tan)
			Eg := EllipticE(gamma, m)
			delta := math.Abs(Ea + Eb - Eg - m*sina*sinb*math.Sin(gamma))
			if delta > tol {
				t.Fatalf("EllipticE test fail for m=%v, alpha=%v, beta=%v", m, alpha, beta)
			}
		}
	}
}
