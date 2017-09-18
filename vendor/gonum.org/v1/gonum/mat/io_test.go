// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"bytes"
	"encoding"
	"io"
	"io/ioutil"
	"math"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

var (
	_ encoding.BinaryMarshaler   = (*Dense)(nil)
	_ encoding.BinaryUnmarshaler = (*Dense)(nil)
	_ encoding.BinaryMarshaler   = (*VecDense)(nil)
	_ encoding.BinaryUnmarshaler = (*VecDense)(nil)
)

var denseData = []struct {
	raw  []byte
	want *Dense
	eq   func(got, want Matrix) bool
}{
	{
		raw:  []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
		want: NewDense(0, 0, []float64{}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x02\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@"),
		want: NewDense(2, 2, []float64{1, 2, 3, 4}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x02\x00\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		want: NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		want: NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@"),
		want: NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x02\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@"),
		want: NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).Slice(0, 2, 0, 2).(*Dense),
		eq:   Equal,
	},
	{
		raw:  []byte("\x02\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@"),
		want: NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).Slice(1, 3, 1, 3).(*Dense),
		eq:   Equal,
	},
	{
		raw:  []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@"),
		want: NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).Slice(0, 3, 1, 3).(*Dense),
		eq:   Equal,
	},
	{
		raw:  []byte("\x01\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0\xff\x00\x00\x00\x00\x00\x00\xf0\u007f\x01\x00\x00\x00\x00\x00\xf8\u007f"),
		want: NewDense(1, 4, []float64{0, math.Inf(-1), math.Inf(+1), math.NaN()}),
		eq: func(got, want Matrix) bool {
			for _, v := range []bool{
				got.At(0, 0) == 0,
				math.IsInf(got.At(0, 1), -1),
				math.IsInf(got.At(0, 2), +1),
				math.IsNaN(got.At(0, 3)),
			} {
				if !v {
					return false
				}
			}
			return true
		},
	},
}

func TestDenseMarshal(t *testing.T) {
	for i, test := range denseData {
		buf, err := test.want.MarshalBinary()
		if err != nil {
			t.Errorf("error encoding test-%d: %v\n", i, err)
			continue
		}

		nrows, ncols := test.want.Dims()
		sz := nrows*ncols*sizeFloat64 + 2*sizeInt64
		if len(buf) != sz {
			t.Errorf("encoded size test-%d: want=%d got=%d\n", i, sz, len(buf))
		}

		if !bytes.Equal(buf, test.raw) {
			t.Errorf("error encoding test-%d: bytes mismatch.\n got=%q\nwant=%q\n",
				i,
				string(buf),
				string(test.raw),
			)
			continue
		}
	}
}

func TestDenseMarshalTo(t *testing.T) {
	for i, test := range denseData {
		buf := new(bytes.Buffer)
		n, err := test.want.MarshalBinaryTo(buf)
		if err != nil {
			t.Errorf("error encoding test-%d: %v\n", i, err)
			continue
		}

		nrows, ncols := test.want.Dims()
		sz := nrows*ncols*sizeFloat64 + 2*sizeInt64
		if n != sz {
			t.Errorf("encoded size test-%d: want=%d got=%d\n", i, sz, n)
		}

		if !bytes.Equal(buf.Bytes(), test.raw) {
			t.Errorf("error encoding test-%d: bytes mismatch.\n got=%q\nwant=%q\n",
				i,
				string(buf.Bytes()),
				string(test.raw),
			)
			continue
		}
	}
}

func TestDenseUnmarshal(t *testing.T) {
	for i, test := range denseData {
		var v Dense
		err := v.UnmarshalBinary(test.raw)
		if err != nil {
			t.Errorf("error decoding test-%d: %v\n", i, err)
			continue
		}
		if !test.eq(&v, test.want) {
			t.Errorf("error decoding test-%d: values differ.\n got=%v\nwant=%v\n",
				i,
				&v,
				test.want,
			)
		}
	}
}

func TestDenseUnmarshalFrom(t *testing.T) {
	for i, test := range denseData {
		var v Dense
		buf := bytes.NewReader(test.raw)
		n, err := v.UnmarshalBinaryFrom(buf)
		if err != nil {
			t.Errorf("error decoding test-%d: %v\n", i, err)
			continue
		}
		if n != len(test.raw) {
			t.Errorf("error decoding test-%d: lengths differ.\n got=%d\nwant=%d\n",
				i, n, len(test.raw),
			)
		}
		if !test.eq(&v, test.want) {
			t.Errorf("error decoding test-%d: values differ.\n got=%v\nwant=%v\n",
				i,
				&v,
				test.want,
			)
		}
	}
}

func TestDenseUnmarshalFromError(t *testing.T) {
	test := denseData[1]
	for i, tt := range []struct {
		beg int
		end int
	}{
		{
			beg: 0,
			end: len(test.raw) - 1,
		},
		{
			beg: 0,
			end: len(test.raw) - sizeFloat64,
		},
		{
			beg: 0,
			end: 0,
		},
		{
			beg: 0,
			end: 1,
		},
		{
			beg: 0,
			end: sizeInt64,
		},
		{
			beg: 0,
			end: sizeInt64 - 1,
		},
		{
			beg: 0,
			end: sizeInt64 + 1,
		},
		{
			beg: 0,
			end: 2*sizeInt64 - 1,
		},
		{
			beg: 0,
			end: 2 * sizeInt64,
		},
		{
			beg: 0,
			end: 2*sizeInt64 + 1,
		},
		{
			beg: 0,
			end: 2*sizeInt64 + sizeFloat64 - 1,
		},
		{
			beg: 0,
			end: 2*sizeInt64 + sizeFloat64,
		},
		{
			beg: 0,
			end: 2*sizeInt64 + sizeFloat64 + 1,
		},
	} {
		buf := bytes.NewReader(test.raw[tt.beg:tt.end])
		var m Dense
		_, err := m.UnmarshalBinaryFrom(buf)
		if err != io.ErrUnexpectedEOF {
			t.Errorf("test #%d: error decoding. got=%v. want=%v\n", i, err, io.ErrUnexpectedEOF)
		}
	}
}

func TestDenseIORoundTrip(t *testing.T) {
	for i, test := range denseData {
		buf, err := test.want.MarshalBinary()
		if err != nil {
			t.Errorf("error encoding test #%d: %v\n", i, err)
		}

		var got Dense
		err = got.UnmarshalBinary(buf)
		if err != nil {
			t.Errorf("error decoding test #%d: %v\n", i, err)
		}

		if !test.eq(&got, test.want) {
			t.Errorf("r/w test #%d failed\n got=%#v\nwant=%#v\n", i, &got, test.want)
		}

		wbuf := new(bytes.Buffer)
		_, err = test.want.MarshalBinaryTo(wbuf)
		if err != nil {
			t.Errorf("error encoding test #%d: %v\n", i, err)
		}

		if !bytes.Equal(buf, wbuf.Bytes()) {
			t.Errorf("r/w test #%d encoding via MarshalBinary and MarshalBinaryTo differ:\nwith-stream: %q\n  no-stream: %q\n",
				i, wbuf.Bytes(), buf,
			)
		}

		var wgot Dense
		_, err = wgot.UnmarshalBinaryFrom(wbuf)
		if err != nil {
			t.Errorf("error decoding test #%d: %v\n", i, err)
		}

		if !test.eq(&wgot, test.want) {
			t.Errorf("r/w test #%d failed\n got=%#v\nwant=%#v\n", i, &wgot, test.want)
		}
	}
}

var vectorData = []struct {
	raw  []byte
	want *VecDense
	eq   func(got, want Matrix) bool
}{
	{
		raw:  []byte("\x00\x00\x00\x00\x00\x00\x00\x00"),
		want: NewVecDense(0, []float64{}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@"),
		want: NewVecDense(4, []float64{1, 2, 3, 4}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		want: NewVecDense(6, []float64{1, 2, 3, 4, 5, 6}),
		eq:   Equal,
	},
	{
		raw:  []byte("\t\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@"),
		want: NewVecDense(9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		eq:   Equal,
	},
	{
		raw:  []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@"),
		want: NewVecDense(9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).SliceVec(0, 3),
		eq:   Equal,
	},
	{
		raw:  []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@"),
		want: NewVecDense(9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).SliceVec(1, 4),
		eq:   Equal,
	},
	{
		raw:  []byte("\b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		want: NewVecDense(9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}).SliceVec(0, 8),
		eq:   Equal,
	},
	{
		raw: []byte("\x03\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\b@\x00\x00\x00\x00\x00\x00\x18@"),
		want: &VecDense{
			mat: blas64.Vector{
				Data: []float64{0, 1, 2, 3, 4, 5, 6},
				Inc:  3,
			},
			n: 3,
		},
		eq: Equal,
	},
	{
		raw:  []byte("\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0\xff\x00\x00\x00\x00\x00\x00\xf0\u007f\x01\x00\x00\x00\x00\x00\xf8\u007f"),
		want: NewVecDense(4, []float64{0, math.Inf(-1), math.Inf(+1), math.NaN()}),
		eq: func(got, want Matrix) bool {
			for _, v := range []bool{
				got.At(0, 0) == 0,
				math.IsInf(got.At(1, 0), -1),
				math.IsInf(got.At(2, 0), +1),
				math.IsNaN(got.At(3, 0)),
			} {
				if !v {
					return false
				}
			}
			return true
		},
	},
}

func TestVecDenseMarshal(t *testing.T) {
	for i, test := range vectorData {
		buf, err := test.want.MarshalBinary()
		if err != nil {
			t.Errorf("error encoding test-%d: %v\n", i, err)
			continue
		}

		nrows, ncols := test.want.Dims()
		sz := nrows*ncols*sizeFloat64 + sizeInt64
		if len(buf) != sz {
			t.Errorf("encoded size test-%d: want=%d got=%d\n", i, sz, len(buf))
		}

		if !bytes.Equal(buf, test.raw) {
			t.Errorf("error encoding test-%d: bytes mismatch.\n got=%q\nwant=%q\n",
				i,
				string(buf),
				string(test.raw),
			)
			continue
		}
	}
}

func TestVecDenseMarshalTo(t *testing.T) {
	for i, test := range vectorData {
		buf := new(bytes.Buffer)
		n, err := test.want.MarshalBinaryTo(buf)
		if err != nil {
			t.Errorf("error encoding test-%d: %v\n", i, err)
			continue
		}

		nrows, ncols := test.want.Dims()
		sz := nrows*ncols*sizeFloat64 + sizeInt64
		if n != sz {
			t.Errorf("encoded size test-%d: want=%d got=%d\n", i, sz, n)
		}

		if !bytes.Equal(buf.Bytes(), test.raw) {
			t.Errorf("error encoding test-%d: bytes mismatch.\n got=%q\nwant=%q\n",
				i,
				string(buf.Bytes()),
				string(test.raw),
			)
			continue
		}
	}
}

func TestVecDenseUnmarshal(t *testing.T) {
	for i, test := range vectorData {
		var v VecDense
		err := v.UnmarshalBinary(test.raw)
		if err != nil {
			t.Errorf("error decoding test-%d: %v\n", i, err)
			continue
		}
		if !test.eq(&v, test.want) {
			t.Errorf("error decoding test-%d: values differ.\n got=%v\nwant=%v\n",
				i,
				&v,
				test.want,
			)
		}
	}
}

func TestVecDenseUnmarshalFrom(t *testing.T) {
	for i, test := range vectorData {
		var v VecDense
		buf := bytes.NewReader(test.raw)
		n, err := v.UnmarshalBinaryFrom(buf)
		if err != nil {
			t.Errorf("error decoding test-%d: %v\n", i, err)
			continue
		}
		if n != len(test.raw) {
			t.Errorf("error decoding test-%d: lengths differ.\n got=%d\nwant=%d\n",
				i,
				n,
				len(test.raw),
			)
		}
		if !test.eq(&v, test.want) {
			t.Errorf("error decoding test-%d: values differ.\n got=%v\nwant=%v\n",
				i,
				&v,
				test.want,
			)
		}
	}
}

func TestVecDenseUnmarshalFromError(t *testing.T) {
	test := vectorData[1]
	for i, tt := range []struct {
		beg int
		end int
	}{
		{
			beg: 0,
			end: len(test.raw) - 1,
		},
		{
			beg: 0,
			end: len(test.raw) - sizeFloat64,
		},
		{
			beg: 0,
			end: 0,
		},
		{
			beg: 0,
			end: 1,
		},
		{
			beg: 0,
			end: sizeInt64,
		},
		{
			beg: 0,
			end: sizeInt64 - 1,
		},
		{
			beg: 0,
			end: sizeInt64 + 1,
		},
		{
			beg: 0,
			end: sizeInt64 + sizeFloat64 - 1,
		},
		{
			beg: 0,
			end: sizeInt64 + sizeFloat64,
		},
		{
			beg: 0,
			end: sizeInt64 + sizeFloat64 + 1,
		},
	} {
		buf := bytes.NewReader(test.raw[tt.beg:tt.end])
		var v VecDense
		_, err := v.UnmarshalBinaryFrom(buf)
		if err != io.ErrUnexpectedEOF {
			t.Errorf("test #%d: error decoding. got=%v. want=%v\n", i, err, io.ErrUnexpectedEOF)
		}
	}
}

func TestVecDenseIORoundTrip(t *testing.T) {
	for i, test := range vectorData {
		buf, err := test.want.MarshalBinary()
		if err != nil {
			t.Errorf("error encoding test #%d: %v\n", i, err)
		}

		var got VecDense
		err = got.UnmarshalBinary(buf)
		if err != nil {
			t.Errorf("error decoding test #%d: %v\n", i, err)
		}
		if !test.eq(&got, test.want) {
			t.Errorf("r/w test #%d failed\n got=%#v\nwant=%#v\n", i, &got, test.want)
		}

		wbuf := new(bytes.Buffer)
		_, err = test.want.MarshalBinaryTo(wbuf)
		if err != nil {
			t.Errorf("error encoding test #%d: %v\n", i, err)
		}

		if !bytes.Equal(buf, wbuf.Bytes()) {
			t.Errorf("test #%d encoding via MarshalBinary and MarshalBinaryTo differ:\nwith-stream: %q\n  no-stream: %q\n",
				i, wbuf.Bytes(), buf,
			)
		}

		var wgot VecDense
		_, err = wgot.UnmarshalBinaryFrom(wbuf)
		if err != nil {
			t.Errorf("error decoding test #%d: %v\n", i, err)
		}

		if !test.eq(&wgot, test.want) {
			t.Errorf("r/w test #%d failed\n got=%#v\nwant=%#v\n", i, &wgot, test.want)
		}
	}
}

func BenchmarkMarshalDense10(b *testing.B)    { marshalBinaryBenchDense(b, 10) }
func BenchmarkMarshalDense100(b *testing.B)   { marshalBinaryBenchDense(b, 100) }
func BenchmarkMarshalDense1000(b *testing.B)  { marshalBinaryBenchDense(b, 1000) }
func BenchmarkMarshalDense10000(b *testing.B) { marshalBinaryBenchDense(b, 10000) }

func marshalBinaryBenchDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	m := NewDense(1, size, data)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.MarshalBinary()
	}
}

func BenchmarkUnmarshalDense10(b *testing.B)    { unmarshalBinaryBenchDense(b, 10) }
func BenchmarkUnmarshalDense100(b *testing.B)   { unmarshalBinaryBenchDense(b, 100) }
func BenchmarkUnmarshalDense1000(b *testing.B)  { unmarshalBinaryBenchDense(b, 1000) }
func BenchmarkUnmarshalDense10000(b *testing.B) { unmarshalBinaryBenchDense(b, 10000) }

func unmarshalBinaryBenchDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	buf, err := NewDense(1, size, data).MarshalBinary()
	if err != nil {
		b.Fatalf("error creating binary buffer (size=%d): %v\n", size, err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var m Dense
		m.UnmarshalBinary(buf)
	}
}

func BenchmarkMarshalToDense10(b *testing.B)    { marshalBinaryToBenchDense(b, 10) }
func BenchmarkMarshalToDense100(b *testing.B)   { marshalBinaryToBenchDense(b, 100) }
func BenchmarkMarshalToDense1000(b *testing.B)  { marshalBinaryToBenchDense(b, 1000) }
func BenchmarkMarshalToDense10000(b *testing.B) { marshalBinaryToBenchDense(b, 10000) }

func marshalBinaryToBenchDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	m := NewDense(1, size, data)
	w := ioutil.Discard
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m.MarshalBinaryTo(w)
	}
}

type readerTest struct {
	buf []byte
	pos int
}

func (r *readerTest) Read(data []byte) (int, error) {
	n := copy(data, r.buf[r.pos:r.pos+len(data)])
	r.pos += n
	return n, nil
}

func (r *readerTest) reset() {
	r.pos = 0
}

func BenchmarkUnmarshalFromDense10(b *testing.B)    { unmarshalBinaryFromBenchDense(b, 10) }
func BenchmarkUnmarshalFromDense100(b *testing.B)   { unmarshalBinaryFromBenchDense(b, 100) }
func BenchmarkUnmarshalFromDense1000(b *testing.B)  { unmarshalBinaryFromBenchDense(b, 1000) }
func BenchmarkUnmarshalFromDense10000(b *testing.B) { unmarshalBinaryFromBenchDense(b, 10000) }

func unmarshalBinaryFromBenchDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	buf, err := NewDense(1, size, data).MarshalBinary()
	if err != nil {
		b.Fatalf("error creating binary buffer (size=%d): %v\n", size, err)
	}
	r := &readerTest{buf: buf}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var m Dense
		m.UnmarshalBinaryFrom(r)
		r.reset()
	}
}

func BenchmarkMarshalVecDense10(b *testing.B)    { marshalBinaryBenchVecDense(b, 10) }
func BenchmarkMarshalVecDense100(b *testing.B)   { marshalBinaryBenchVecDense(b, 100) }
func BenchmarkMarshalVecDense1000(b *testing.B)  { marshalBinaryBenchVecDense(b, 1000) }
func BenchmarkMarshalVecDense10000(b *testing.B) { marshalBinaryBenchVecDense(b, 10000) }

func marshalBinaryBenchVecDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	vec := NewVecDense(size, data)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		vec.MarshalBinary()
	}
}

func BenchmarkUnmarshalVecDense10(b *testing.B)    { unmarshalBinaryBenchVecDense(b, 10) }
func BenchmarkUnmarshalVecDense100(b *testing.B)   { unmarshalBinaryBenchVecDense(b, 100) }
func BenchmarkUnmarshalVecDense1000(b *testing.B)  { unmarshalBinaryBenchVecDense(b, 1000) }
func BenchmarkUnmarshalVecDense10000(b *testing.B) { unmarshalBinaryBenchVecDense(b, 10000) }

func unmarshalBinaryBenchVecDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	buf, err := NewVecDense(size, data).MarshalBinary()
	if err != nil {
		b.Fatalf("error creating binary buffer (size=%d): %v\n", size, err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var vec VecDense
		vec.UnmarshalBinary(buf)
	}
}

func BenchmarkMarshalToVecDense10(b *testing.B)    { marshalBinaryToBenchVecDense(b, 10) }
func BenchmarkMarshalToVecDense100(b *testing.B)   { marshalBinaryToBenchVecDense(b, 100) }
func BenchmarkMarshalToVecDense1000(b *testing.B)  { marshalBinaryToBenchVecDense(b, 1000) }
func BenchmarkMarshalToVecDense10000(b *testing.B) { marshalBinaryToBenchVecDense(b, 10000) }

func marshalBinaryToBenchVecDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	vec := NewVecDense(size, data)
	w := ioutil.Discard
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		vec.MarshalBinaryTo(w)
	}
}

func BenchmarkUnmarshalFromVecDense10(b *testing.B)    { unmarshalBinaryFromBenchVecDense(b, 10) }
func BenchmarkUnmarshalFromVecDense100(b *testing.B)   { unmarshalBinaryFromBenchVecDense(b, 100) }
func BenchmarkUnmarshalFromVecDense1000(b *testing.B)  { unmarshalBinaryFromBenchVecDense(b, 1000) }
func BenchmarkUnmarshalFromVecDense10000(b *testing.B) { unmarshalBinaryFromBenchVecDense(b, 10000) }

func unmarshalBinaryFromBenchVecDense(b *testing.B, size int) {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	buf, err := NewVecDense(size, data).MarshalBinary()
	if err != nil {
		b.Fatalf("error creating binary buffer (size=%d): %v\n", size, err)
	}
	r := &readerTest{buf: buf}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var vec VecDense
		vec.UnmarshalBinaryFrom(r)
		r.reset()
	}
}
