// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blas64

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/blas"
)

func newSymmetricFrom(a SymmetricCols) Symmetric {
	t := Symmetric{
		N:      a.N,
		Stride: a.N,
		Data:   make([]float64, a.N*a.N),
		Uplo:   a.Uplo,
	}
	t.From(a)
	return t
}

func (m Symmetric) n() int { return m.N }
func (m Symmetric) at(i, j int) float64 {
	if m.Uplo == blas.Lower && i < j && j < m.N {
		i, j = j, i
	}
	if m.Uplo == blas.Upper && i > j {
		i, j = j, i
	}
	return m.Data[i*m.Stride+j]
}
func (m Symmetric) uplo() blas.Uplo { return m.Uplo }

func newSymmetricColsFrom(a Symmetric) SymmetricCols {
	t := SymmetricCols{
		N:      a.N,
		Stride: a.N,
		Data:   make([]float64, a.N*a.N),
		Uplo:   a.Uplo,
	}
	t.From(a)
	return t
}

func (m SymmetricCols) n() int { return m.N }
func (m SymmetricCols) at(i, j int) float64 {
	if m.Uplo == blas.Lower && i < j {
		i, j = j, i
	}
	if m.Uplo == blas.Upper && i > j && i < m.N {
		i, j = j, i
	}
	return m.Data[i+j*m.Stride]
}
func (m SymmetricCols) uplo() blas.Uplo { return m.Uplo }

type symmetric interface {
	n() int
	at(i, j int) float64
	uplo() blas.Uplo
}

func sameSymmetric(a, b symmetric) bool {
	an := a.n()
	bn := b.n()
	if an != bn {
		return false
	}
	if a.uplo() != b.uplo() {
		return false
	}
	for i := 0; i < an; i++ {
		for j := 0; j < an; j++ {
			if a.at(i, j) != b.at(i, j) || math.IsNaN(a.at(i, j)) != math.IsNaN(b.at(i, j)) {
				return false
			}
		}
	}
	return true
}

var symmetricTests = []Symmetric{
	{N: 3, Stride: 3, Data: []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}},
	{N: 3, Stride: 5, Data: []float64{
		1, 2, 3, 0, 0,
		4, 5, 6, 0, 0,
		7, 8, 9, 0, 0,
	}},
}

func TestConvertSymmetric(t *testing.T) {
	for _, test := range symmetricTests {
		for _, uplo := range []blas.Uplo{blas.Upper, blas.Lower} {
			test.Uplo = uplo
			colmajor := newSymmetricColsFrom(test)
			if !sameSymmetric(colmajor, test) {
				t.Errorf("unexpected result for row major to col major conversion:\n\tgot: %#v\n\tfrom:%#v",
					colmajor, test)
			}
			rowmajor := newSymmetricFrom(colmajor)
			if !sameSymmetric(rowmajor, test) {
				t.Errorf("unexpected result for col major to row major conversion:\n\tgot: %#v\n\twant:%#v",
					rowmajor, test)
			}
		}
	}
}
func newSymmetricBandFrom(a SymmetricBandCols) SymmetricBand {
	t := SymmetricBand{
		N:      a.N,
		K:      a.K,
		Stride: a.K + 1,
		Data:   make([]float64, a.N*(a.K+1)),
		Uplo:   a.Uplo,
	}
	for i := range t.Data {
		t.Data[i] = math.NaN()
	}
	t.From(a)
	return t
}

func (m SymmetricBand) n() (n int) { return m.N }
func (m SymmetricBand) at(i, j int) float64 {
	b := Band{
		Rows: m.N, Cols: m.N,
		Stride: m.Stride,
		Data:   m.Data,
	}
	switch m.Uplo {
	default:
		panic("blas64: bad BLAS uplo")
	case blas.Upper:
		b.KU = m.K
		if i > j {
			i, j = j, i
		}
	case blas.Lower:
		b.KL = m.K
		if i < j {
			i, j = j, i
		}
	}
	return b.at(i, j)
}
func (m SymmetricBand) bandwidth() (k int) { return m.K }
func (m SymmetricBand) uplo() blas.Uplo    { return m.Uplo }

func newSymmetricBandColsFrom(a SymmetricBand) SymmetricBandCols {
	t := SymmetricBandCols{
		N:      a.N,
		K:      a.K,
		Stride: a.K + 1,
		Data:   make([]float64, a.N*(a.K+1)),
		Uplo:   a.Uplo,
	}
	for i := range t.Data {
		t.Data[i] = math.NaN()
	}
	t.From(a)
	return t
}

func (m SymmetricBandCols) n() (n int) { return m.N }
func (m SymmetricBandCols) at(i, j int) float64 {
	b := BandCols{
		Rows: m.N, Cols: m.N,
		Stride: m.Stride,
		Data:   m.Data,
	}
	switch m.Uplo {
	default:
		panic("blas64: bad BLAS uplo")
	case blas.Upper:
		b.KU = m.K
		if i > j {
			i, j = j, i
		}
	case blas.Lower:
		b.KL = m.K
		if i < j {
			i, j = j, i
		}
	}
	return b.at(i, j)
}
func (m SymmetricBandCols) bandwidth() (k int) { return m.K }
func (m SymmetricBandCols) uplo() blas.Uplo    { return m.Uplo }

type symmetricBand interface {
	n() (n int)
	at(i, j int) float64
	bandwidth() (k int)
	uplo() blas.Uplo
}

func sameSymmetricBand(a, b symmetricBand) bool {
	an := a.n()
	bn := b.n()
	if an != bn {
		return false
	}
	if a.uplo() != b.uplo() {
		return false
	}
	ak := a.bandwidth()
	bk := b.bandwidth()
	if ak != bk {
		return false
	}
	for i := 0; i < an; i++ {
		for j := 0; j < an; j++ {
			if a.at(i, j) != b.at(i, j) || math.IsNaN(a.at(i, j)) != math.IsNaN(b.at(i, j)) {
				return false
			}
		}
	}
	return true
}

var symmetricBandTests = []SymmetricBand{
	{N: 3, K: 0, Stride: 1, Uplo: blas.Upper, Data: []float64{
		1,
		2,
		3,
	}},
	{N: 3, K: 0, Stride: 1, Uplo: blas.Lower, Data: []float64{
		1,
		2,
		3,
	}},
	{N: 3, K: 1, Stride: 2, Uplo: blas.Upper, Data: []float64{
		1, 2,
		3, 4,
		5, -1,
	}},
	{N: 3, K: 1, Stride: 2, Uplo: blas.Lower, Data: []float64{
		-1, 1,
		2, 3,
		4, 5,
	}},
	{N: 3, K: 2, Stride: 3, Uplo: blas.Upper, Data: []float64{
		1, 2, 3,
		4, 5, -1,
		6, -2, -3,
	}},
	{N: 3, K: 2, Stride: 3, Uplo: blas.Lower, Data: []float64{
		-2, -1, 1,
		-3, 2, 4,
		3, 5, 6,
	}},

	{N: 3, K: 0, Stride: 5, Uplo: blas.Upper, Data: []float64{
		1, 0, 0, 0, 0,
		2, 0, 0, 0, 0,
		3, 0, 0, 0, 0,
	}},
	{N: 3, K: 0, Stride: 5, Uplo: blas.Lower, Data: []float64{
		1, 0, 0, 0, 0,
		2, 0, 0, 0, 0,
		3, 0, 0, 0, 0,
	}},
	{N: 3, K: 1, Stride: 5, Uplo: blas.Upper, Data: []float64{
		1, 2, 0, 0, 0,
		3, 4, 0, 0, 0,
		5, -1, 0, 0, 0,
	}},
	{N: 3, K: 1, Stride: 5, Uplo: blas.Lower, Data: []float64{
		-1, 1, 0, 0, 0,
		2, 3, 0, 0, 0,
		4, 5, 0, 0, 0,
	}},
	{N: 3, K: 2, Stride: 5, Uplo: blas.Upper, Data: []float64{
		1, 2, 3, 0, 0,
		4, 5, -1, 0, 0,
		6, -2, -3, 0, 0,
	}},
	{N: 3, K: 2, Stride: 5, Uplo: blas.Lower, Data: []float64{
		-2, -1, 1, 0, 0,
		-3, 2, 4, 0, 0,
		3, 5, 6, 0, 0,
	}},
}

func TestConvertSymBand(t *testing.T) {
	for _, test := range symmetricBandTests {
		colmajor := newSymmetricBandColsFrom(test)
		if !sameSymmetricBand(colmajor, test) {
			t.Errorf("unexpected result for row major to col major conversion:\n\tgot: %#v\n\tfrom:%#v",
				colmajor, test)
		}
		rowmajor := newSymmetricBandFrom(colmajor)
		if !sameSymmetricBand(rowmajor, test) {
			t.Errorf("unexpected result for col major to row major conversion:\n\tgot: %#v\n\twant:%#v",
				rowmajor, test)
		}
	}
}
