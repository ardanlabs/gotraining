// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

const (
	// maxLen is the biggest slice/array len one can create on a 32/64b platform.
	maxLen = int64(int(^uint(0) >> 1))
)

var (
	sizeInt64   = binary.Size(int64(0))
	sizeFloat64 = binary.Size(float64(0))

	errTooBig    = errors.New("mat64: resulting data slice too big")
	errTooSmall  = errors.New("mat64: input slice too small")
	errBadBuffer = errors.New("mat64: data buffer size mismatch")
	errBadSize   = errors.New("mat64: invalid dimension")
)

// MarshalBinary encodes the receiver into a binary form and returns the result.
//
// Dense is little-endian encoded as follows:
//   0 -  7  number of rows    (int64)
//   8 - 15  number of columns (int64)
//  16 - ..  matrix data elements (float64)
//           [0,0] [0,1] ... [0,ncols-1]
//           [1,0] [1,1] ... [1,ncols-1]
//           ...
//           [nrows-1,0] ... [nrows-1,ncols-1]
func (m Dense) MarshalBinary() ([]byte, error) {
	bufLen := int64(m.mat.Rows)*int64(m.mat.Cols)*int64(sizeFloat64) + 2*int64(sizeInt64)
	if bufLen <= 0 {
		// bufLen is too big and has wrapped around.
		return nil, errTooBig
	}

	p := 0
	buf := make([]byte, bufLen)
	binary.LittleEndian.PutUint64(buf[p:p+sizeInt64], uint64(m.mat.Rows))
	p += sizeInt64
	binary.LittleEndian.PutUint64(buf[p:p+sizeInt64], uint64(m.mat.Cols))
	p += sizeInt64

	r, c := m.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			binary.LittleEndian.PutUint64(buf[p:p+sizeFloat64], math.Float64bits(m.at(i, j)))
			p += sizeFloat64
		}
	}

	return buf, nil
}

// MarshalBinaryTo encodes the receiver into a binary form and writes it into w.
// MarshalBinaryTo returns the number of bytes written into w and an error, if any.
//
// See MarshalBinary for the on-disk layout.
func (m Dense) MarshalBinaryTo(w io.Writer) (int, error) {
	var n int
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(m.mat.Rows))
	nn, err := w.Write(buf[:])
	n += nn
	if err != nil {
		return n, err
	}
	binary.LittleEndian.PutUint64(buf[:], uint64(m.mat.Cols))
	nn, err = w.Write(buf[:])
	n += nn
	if err != nil {
		return n, err
	}

	r, c := m.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			binary.LittleEndian.PutUint64(buf[:], math.Float64bits(m.at(i, j)))
			nn, err = w.Write(buf[:])
			n += nn
			if err != nil {
				return n, err
			}
		}
	}

	return n, nil
}

// UnmarshalBinary decodes the binary form into the receiver.
// It panics if the receiver is a non-zero Dense matrix.
//
// See MarshalBinary for the on-disk layout.
//
// Limited checks on the validity of the binary input are performed:
//  - matrix.ErrShape is returned if the number of rows or columns is negative,
//  - an error is returned if the resulting Dense matrix is too
//  big for the current architecture (e.g. a 16GB matrix written by a
//  64b application and read back from a 32b application.)
// UnmarshalBinary does not limit the size of the unmarshaled matrix, and so
// it should not be used on untrusted data.
func (m *Dense) UnmarshalBinary(data []byte) error {
	if !m.isZero() {
		panic("mat64: unmarshal into non-zero matrix")
	}

	if len(data) < 2*sizeInt64 {
		return errTooSmall
	}

	p := 0
	rows := int64(binary.LittleEndian.Uint64(data[p : p+sizeInt64]))
	p += sizeInt64
	cols := int64(binary.LittleEndian.Uint64(data[p : p+sizeInt64]))
	p += sizeInt64
	if rows < 0 || cols < 0 {
		return errBadSize
	}

	size := rows * cols
	if int(size) < 0 || size > maxLen {
		return errTooBig
	}

	if len(data) != int(size)*sizeFloat64+2*sizeInt64 {
		return errBadBuffer
	}

	m.reuseAs(int(rows), int(cols))
	for i := range m.mat.Data {
		m.mat.Data[i] = math.Float64frombits(binary.LittleEndian.Uint64(data[p : p+sizeFloat64]))
		p += sizeFloat64
	}

	return nil
}

// UnmarshalBinaryFrom decodes the binary form into the receiver and returns
// the number of bytes read and an error if any.
// It panics if the receiver is a non-zero Dense matrix.
//
// See MarshalBinary for the on-disk layout.
//
// Limited checks on the validity of the binary input are performed:
//  - matrix.ErrShape is returned if the number of rows or columns is negative,
//  - an error is returned if the resulting Dense matrix is too
//  big for the current architecture (e.g. a 16GB matrix written by a
//  64b application and read back from a 32b application.)
// UnmarshalBinary does not limit the size of the unmarshaled matrix, and so
// it should not be used on untrusted data.
func (m *Dense) UnmarshalBinaryFrom(r io.Reader) (int, error) {
	if !m.isZero() {
		panic("mat64: unmarshal into non-zero matrix")
	}

	var (
		n   int
		buf [8]byte
	)
	nn, err := readFull(r, buf[:])
	n += nn
	if err != nil {
		return n, err
	}
	rows := int64(binary.LittleEndian.Uint64(buf[:]))

	nn, err = readFull(r, buf[:])
	n += nn
	if err != nil {
		return n, err
	}
	cols := int64(binary.LittleEndian.Uint64(buf[:]))
	if rows < 0 || cols < 0 {
		return n, errBadSize
	}

	size := rows * cols
	if int(size) < 0 || size > maxLen {
		return n, errTooBig
	}

	m.reuseAs(int(rows), int(cols))
	for i := range m.mat.Data {
		nn, err = readFull(r, buf[:])
		n += nn
		if err != nil {
			return n, err
		}
		m.mat.Data[i] = math.Float64frombits(binary.LittleEndian.Uint64(buf[:]))
	}

	return n, nil
}

// MarshalBinary encodes the receiver into a binary form and returns the result.
//
// Vector is little-endian encoded as follows:
//   0 -  7  number of elements     (int64)
//   8 - ..  vector's data elements (float64)
func (v Vector) MarshalBinary() ([]byte, error) {
	bufLen := int64(sizeInt64) + int64(v.n)*int64(sizeFloat64)
	if bufLen <= 0 {
		// bufLen is too big and has wrapped around.
		return nil, errTooBig
	}

	p := 0
	buf := make([]byte, bufLen)
	binary.LittleEndian.PutUint64(buf[p:p+sizeInt64], uint64(v.n))
	p += sizeInt64

	for i := 0; i < v.n; i++ {
		binary.LittleEndian.PutUint64(buf[p:p+sizeFloat64], math.Float64bits(v.at(i)))
		p += sizeFloat64
	}

	return buf, nil
}

// MarshalBinaryTo encodes the receiver into a binary form, writes it to w and
// returns the number of bytes written and an error if any.
//
// See MarshalBainry for the on-disk format.
func (v Vector) MarshalBinaryTo(w io.Writer) (int, error) {
	var (
		n   int
		buf [8]byte
	)

	binary.LittleEndian.PutUint64(buf[:], uint64(v.n))
	nn, err := w.Write(buf[:])
	n += nn
	if err != nil {
		return n, err
	}

	for i := 0; i < v.n; i++ {
		binary.LittleEndian.PutUint64(buf[:], math.Float64bits(v.at(i)))
		nn, err = w.Write(buf[:])
		n += nn
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

// UnmarshalBinary decodes the binary form into the receiver.
// It panics if the receiver is a non-zero Vector.
//
// See MarshalBinary for the on-disk layout.
//
// Limited checks on the validity of the binary input are performed:
//  - matrix.ErrShape is returned if the number of rows is negative,
//  - an error is returned if the resulting Vector is too
//  big for the current architecture (e.g. a 16GB vector written by a
//  64b application and read back from a 32b application.)
// UnmarshalBinary does not limit the size of the unmarshaled vector, and so
// it should not be used on untrusted data.
func (v *Vector) UnmarshalBinary(data []byte) error {
	if !v.isZero() {
		panic("mat64: unmarshal into non-zero vector")
	}

	p := 0
	n := int64(binary.LittleEndian.Uint64(data[p : p+sizeInt64]))
	p += sizeInt64
	if n < 0 {
		return errBadSize
	}
	if n > maxLen {
		return errTooBig
	}
	if len(data) != int(n)*sizeFloat64+sizeInt64 {
		return errBadBuffer
	}

	v.reuseAs(int(n))
	for i := range v.mat.Data {
		v.mat.Data[i] = math.Float64frombits(binary.LittleEndian.Uint64(data[p : p+sizeFloat64]))
		p += sizeFloat64
	}

	return nil
}

// UnmarshalBinaryFrom decodes the binary form into the receiver, from the
// io.Reader and returns the number of bytes read and an error if any.
// It panics if the receiver is a non-zero Vector.
//
// See MarshalBinary for the on-disk layout.
// See UnmarshalBinary for the list of sanity checks performed on the input.
func (v *Vector) UnmarshalBinaryFrom(r io.Reader) (int, error) {
	if !v.isZero() {
		panic("mat64: unmarshal into non-zero vector")
	}

	var (
		n   int
		buf [8]byte
	)
	nn, err := readFull(r, buf[:])
	n += nn
	if err != nil {
		return n, err
	}
	sz := int64(binary.LittleEndian.Uint64(buf[:]))
	if sz < 0 {
		return n, errBadSize
	}
	if sz > maxLen {
		return n, errTooBig
	}

	v.reuseAs(int(sz))
	for i := range v.mat.Data {
		nn, err = readFull(r, buf[:])
		n += nn
		if err != nil {
			return n, err
		}
		v.mat.Data[i] = math.Float64frombits(binary.LittleEndian.Uint64(buf[:]))
	}

	if n != sizeInt64+int(sz)*sizeFloat64 {
		return n, io.ErrUnexpectedEOF
	}

	return n, nil
}

// readFull reads from r into buf until it has read len(buf).
// It returns the number of bytes copied and an error if fewer bytes were read.
// If an EOF happens after reading fewer than len(buf) bytes, io.ErrUnexpectedEOF is returned.
func readFull(r io.Reader, buf []byte) (int, error) {
	var n int
	var err error
	for n < len(buf) && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n == len(buf) {
		return n, nil
	}
	if err == io.EOF {
		return n, io.ErrUnexpectedEOF
	}
	return n, err
}
