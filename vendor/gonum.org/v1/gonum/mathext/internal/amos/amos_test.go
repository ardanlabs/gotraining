// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amos

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"gonum.org/v1/gonum/floats"
)

type input struct {
	x    []float64
	is   []int
	kode int
	id   int
	yr   []float64
	yi   []float64
	n    int
	tol  float64
}

func randnum(rnd *rand.Rand) float64 {
	r := 2e2 // Fortran has infinite loop if this is set higher than 2e3
	if rnd.Float64() > 0.99 {
		return 0
	}
	return rnd.Float64()*r - r/2
}

func randInput(rnd *rand.Rand) input {
	x := make([]float64, 8)
	for j := range x {
		x[j] = randnum(rnd)
	}
	is := make([]int, 3)
	for j := range is {
		is[j] = rand.Intn(1000)
	}
	kode := rand.Intn(2) + 1
	id := rand.Intn(2)
	n := rand.Intn(5) + 1
	yr := make([]float64, n+1)
	yi := make([]float64, n+1)
	for j := range yr {
		yr[j] = randnum(rnd)
		yi[j] = randnum(rnd)
	}
	tol := 1e-14

	return input{
		x, is, kode, id, yr, yi, n, tol,
	}
}

const nInputs = 100000

func TestAiry(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zairytest(t, in.x, in.kode, in.id)
	}
}

func TestZacai(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zacaitest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZbknu(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zbknutest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZasyi(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zasyitest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZseri(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zseritest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZmlri(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zmlritest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi, in.kode)
	}
}

func TestZkscl(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zkscltest(t, in.x, in.is, in.tol, in.n, in.yr, in.yi)
	}
}

func TestZuchk(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zuchktest(t, in.x, in.is, in.tol)
	}
}

func TestZs1s2(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < nInputs; i++ {
		in := randInput(rnd)
		zs1s2test(t, in.x, in.is)
	}
}

func zs1s2test(t *testing.T, x []float64, is []int) {

	type data struct {
		ZRR, ZRI, S1R, S1I, S2R, S2I float64
		NZ                           int
		ASCLE, ALIM                  float64
		IUF                          int
	}

	input := data{
		x[0], x[1], x[2], x[3], x[4], x[5],
		is[0],
		x[6], x[7],
		is[1],
	}

	zr := complex(input.ZRR, input.ZRI)
	s1 := complex(input.S1R, input.S1I)
	s2 := complex(input.S2R, input.S2I)
	impl := func(input data) data {
		s1, s2, nz, iuf := Zs1s2(zr, s1, s2, input.ASCLE, input.ALIM, input.IUF)
		zrr := real(zr)
		zri := imag(zr)
		s1r := real(s1)
		s1i := imag(s1)
		s2r := real(s2)
		s2i := imag(s2)
		alim := input.ALIM
		ascle := input.ASCLE
		return data{zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf}
	}

	comp := func(input data) data {
		zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf :=
			zs1s2Orig(input.ZRR, input.ZRI, input.S1R, input.S1I, input.S2R, input.S2I, input.NZ, input.ASCLE, input.ALIM, input.IUF)
		return data{zrr, zri, s1r, s1i, s2r, s2i, nz, ascle, alim, iuf}
	}

	oi := impl(input)
	oc := comp(input)

	sameF64(t, "zs1s2 zrr", oc.ZRR, oi.ZRR)
	sameF64(t, "zs1s2 zri", oc.ZRI, oi.ZRI)
	sameF64(t, "zs1s2 s1r", oc.S1R, oi.S1R)
	sameF64(t, "zs1s2 s1i", oc.S1I, oi.S1I)
	sameF64(t, "zs1s2 s2r", oc.S2R, oi.S2R)
	sameF64(t, "zs1s2 s2i", oc.S2I, oi.S2I)
	sameF64(t, "zs1s2 ascle", oc.ASCLE, oi.ASCLE)
	sameF64(t, "zs1s2 alim", oc.ALIM, oi.ALIM)
	sameInt(t, "iuf", oc.IUF, oi.IUF)
	sameInt(t, "nz", oc.NZ, oi.NZ)
}

func zuchktest(t *testing.T, x []float64, is []int, tol float64) {
	YR := x[0]
	YI := x[1]
	NZ := is[0]
	ASCLE := x[2]
	TOL := tol

	YRfort, YIfort, NZfort, ASCLEfort, TOLfort := zuchkOrig(YR, YI, NZ, ASCLE, TOL)
	y := complex(YR, YI)
	NZamos := Zuchk(y, ASCLE, TOL)
	YRamos := real(y)
	YIamos := imag(y)
	ASCLEamos := ASCLE
	TOLamos := TOL

	sameF64(t, "zuchk yr", YRfort, YRamos)
	sameF64(t, "zuchk yi", YIfort, YIamos)
	sameInt(t, "zuchk nz", NZfort, NZamos)
	sameF64(t, "zuchk ascle", ASCLEfort, ASCLEamos)
	sameF64(t, "zuchk tol", TOLfort, TOLamos)
}

func zkscltest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64) {
	ZRR := x[0]
	ZRI := x[1]
	FNU := x[2]
	NZ := is[1]
	ELIM := x[3]
	ASCLE := x[4]
	RZR := x[6]
	RZI := x[7]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRRfort, ZRIfort, FNUfort, Nfort, YRfort, YIfort, NZfort, RZRfort, RZIfort, ASCLEfort, TOLfort, ELIMfort :=
		zksclOrig(ZRR, ZRI, FNU, n, yrfort, yifort, NZ, RZR, RZI, ASCLE, tol, ELIM)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRRamos, ZRIamos, FNUamos, Namos, YRamos, YIamos, NZamos, RZRamos, RZIamos, ASCLEamos, TOLamos, ELIMamos :=
		Zkscl(ZRR, ZRI, FNU, n, yramos, yiamos, NZ, RZR, RZI, ASCLE, tol, ELIM)

	sameF64(t, "zkscl zrr", ZRRfort, ZRRamos)
	sameF64(t, "zkscl zri", ZRIfort, ZRIamos)
	sameF64(t, "zkscl fnu", FNUfort, FNUamos)
	sameInt(t, "zkscl n", Nfort, Namos)
	sameInt(t, "zkscl nz", NZfort, NZamos)
	sameF64(t, "zkscl rzr", RZRfort, RZRamos)
	sameF64(t, "zkscl rzi", RZIfort, RZIamos)
	sameF64(t, "zkscl ascle", ASCLEfort, ASCLEamos)
	sameF64(t, "zkscl tol", TOLfort, TOLamos)
	sameF64(t, "zkscl elim", ELIMfort, ELIMamos)

	sameF64SApprox(t, "zkscl yr", YRfort, YRamos, 1e-14)
	sameF64SApprox(t, "zkscl yi", YIfort, YIamos, 1e-14)
}

func zmlritest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort :=
		zmlriOrig(ZR, ZI, FNU, KODE, n, yrfort, yifort, NZ, tol)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos :=
		Zmlri(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol)

	sameF64(t, "zmlri zr", ZRfort, ZRamos)
	sameF64(t, "zmlri zi", ZIfort, ZIamos)
	sameF64(t, "zmlri fnu", FNUfort, FNUamos)
	sameInt(t, "zmlri kode", KODEfort, KODEamos)
	sameInt(t, "zmlri n", Nfort, Namos)
	sameInt(t, "zmlri nz", NZfort, NZamos)
	sameF64(t, "zmlri tol", TOLfort, TOLamos)

	sameF64S(t, "zmlri yr", YRfort, YRamos)
	sameF64S(t, "zmlri yi", YIfort, YIamos)
}

func zseritest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		zseriOrig(ZR, ZI, FNU, KODE, n, yrfort, yifort, NZ, tol, ELIM, ALIM)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	y := make([]complex128, len(yramos))
	for i, v := range yramos {
		y[i] = complex(v, yiamos[i])
	}
	z := complex(ZR, ZI)

	NZamos := Zseri(z, FNU, KODE, n, y[1:], tol, ELIM, ALIM)

	ZRamos := real(z)
	ZIamos := imag(z)
	FNUamos := FNU
	KODEamos := KODE
	Namos := n
	TOLamos := tol
	ELIMamos := ELIM
	ALIMamos := ALIM
	YRamos := make([]float64, len(y))
	YIamos := make([]float64, len(y))
	for i, v := range y {
		YRamos[i] = real(v)
		YIamos[i] = imag(v)
	}

	sameF64(t, "zseri zr", ZRfort, ZRamos)
	sameF64(t, "zseri zi", ZIfort, ZIamos)
	sameF64(t, "zseri fnu", FNUfort, FNUamos)
	sameInt(t, "zseri kode", KODEfort, KODEamos)
	sameInt(t, "zseri n", Nfort, Namos)
	sameInt(t, "zseri nz", NZfort, NZamos)
	sameF64(t, "zseri tol", TOLfort, TOLamos)
	sameF64(t, "zseri elim", ELIMfort, ELIMamos)
	sameF64(t, "zseri elim", ALIMfort, ALIMamos)

	sameF64SApprox(t, "zseri yr", YRfort, YRamos, 1e-10)
	sameF64SApprox(t, "zseri yi", YIfort, YIamos, 1e-10)
}

func zasyitest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		zasyiOrig(ZR, ZI, FNU, KODE, n, yrfort, yifort, NZ, RL, tol, ELIM, ALIM)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		Zasyi(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	sameF64(t, "zasyi zr", ZRfort, ZRamos)
	sameF64(t, "zasyi zr", ZIfort, ZIamos)
	sameF64(t, "zasyi fnu", FNUfort, FNUamos)
	sameInt(t, "zasyi kode", KODEfort, KODEamos)
	sameInt(t, "zasyi n", Nfort, Namos)
	sameInt(t, "zasyi nz", NZfort, NZamos)
	sameF64(t, "zasyi rl", RLfort, RLamos)
	sameF64(t, "zasyi tol", TOLfort, TOLamos)
	sameF64(t, "zasyi elim", ELIMfort, ELIMamos)
	sameF64(t, "zasyi alim", ALIMfort, ALIMamos)

	sameF64SApprox(t, "zasyi yr", YRfort, YRamos, 1e-12)
	sameF64SApprox(t, "zasyi yi", YIfort, YIamos, 1e-12)
}

func zbknutest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		zbknuOrig(ZR, ZI, FNU, KODE, n, yrfort, yifort, NZ, tol, ELIM, ALIM)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos, ELIMamos, ALIMamos :=
		Zbknu(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol, ELIM, ALIM)

	sameF64(t, "zbknu zr", ZRfort, ZRamos)
	sameF64(t, "zbknu zr", ZIfort, ZIamos)
	sameF64(t, "zbknu fnu", FNUfort, FNUamos)
	sameInt(t, "zbknu kode", KODEfort, KODEamos)
	sameInt(t, "zbknu n", Nfort, Namos)
	sameInt(t, "zbknu nz", NZfort, NZamos)
	sameF64(t, "zbknu tol", TOLfort, TOLamos)
	sameF64(t, "zbknu elim", ELIMfort, ELIMamos)
	sameF64(t, "zbknu alim", ALIMfort, ALIMamos)

	sameF64SApprox(t, "zbknu yr", YRfort, YRamos, 1e-12)
	sameF64SApprox(t, "zbknu yi", YIfort, YIamos, 1e-12)
}

func zairytest(t *testing.T, x []float64, kode, id int) {
	ZR := x[0]
	ZI := x[1]
	KODE := kode
	ID := id

	AIRfort, AIIfort, NZfort := zairyOrig(ZR, ZI, ID, KODE)
	AIRamos, AIIamos, NZamos := Zairy(ZR, ZI, ID, KODE)

	sameF64Approx(t, "zairy air", AIRfort, AIRamos, 1e-12)
	sameF64Approx(t, "zairy aii", AIIfort, AIIamos, 1e-12)
	sameInt(t, "zairy nz", NZfort, NZamos)
}

func zacaitest(t *testing.T, x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	MR := is[2]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, MRfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		zacaiOrig(ZR, ZI, FNU, KODE, MR, n, yrfort, yifort, NZ, RL, tol, ELIM, ALIM)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, MRamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		Zacai(ZR, ZI, FNU, KODE, MR, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	sameF64(t, "zacai zr", ZRfort, ZRamos)
	sameF64(t, "zacai zi", ZIfort, ZIamos)
	sameF64(t, "zacai fnu", FNUfort, FNUamos)
	sameInt(t, "zacai kode", KODEfort, KODEamos)
	sameInt(t, "zacai mr", MRfort, MRamos)
	sameInt(t, "zacai n", Nfort, Namos)
	sameInt(t, "zacai nz", NZfort, NZamos)
	sameF64(t, "zacai rl", RLfort, RLamos)
	sameF64(t, "zacai tol", TOLfort, TOLamos)
	sameF64(t, "zacai elim", ELIMfort, ELIMamos)
	sameF64(t, "zacai elim", ALIMfort, ALIMamos)

	sameF64SApprox(t, "zacai yr", YRfort, YRamos, 1e-12)
	sameF64SApprox(t, "zacai yi", YIfort, YIamos, 1e-12)
}

func sameF64(t *testing.T, str string, c, native float64) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if c == native {
		return
	}
	cb := math.Float64bits(c)
	nb := math.Float64bits(native)
	t.Errorf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n", str, c, native, cb, nb)
}

func sameF64Approx(t *testing.T, str string, c, native, tol float64) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if floats.EqualWithinAbsOrRel(c, native, tol, tol) {
		return
	}
	// Have a much looser tolerance for correctness when the values are large.
	// Floating point noise makes the relative tolerance difference greater for
	// higher values.
	if c > 1e200 && floats.EqualWithinAbsOrRel(c, native, 10, 10) {
		return
	}
	cb := math.Float64bits(c)
	nb := math.Float64bits(native)
	t.Errorf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n", str, c, native, cb, nb)
}

func sameInt(t *testing.T, str string, c, native int) {
	if c != native {
		t.Errorf("Case %s: Int mismatch. c = %v, native = %v.", str, c, native)
	}
}

func sameF64S(t *testing.T, str string, c, native []float64) {
	if len(c) != len(native) {
		panic(str)
	}
	for i, v := range c {
		sameF64(t, str+"_idx_"+strconv.Itoa(i), v, native[i])
	}
}

func sameF64SApprox(t *testing.T, str string, c, native []float64, tol float64) {
	if len(c) != len(native) {
		panic(str)
	}
	for i, v := range c {
		sameF64Approx(t, str+"_idx_"+strconv.Itoa(i), v, native[i], tol)
	}
}
