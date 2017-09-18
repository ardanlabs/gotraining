// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "runtime"

func (A *DenseMatrix) Plus(B MatrixRO) (Matrix, error) {
	C := A.Copy()
	err := C.Add(B)
	return C, err
}
func (A *DenseMatrix) PlusDense(B *DenseMatrix) (*DenseMatrix, error) {
	C := A.Copy()
	err := C.AddDense(B)
	return C, err
}

func (A *DenseMatrix) Minus(B MatrixRO) (Matrix, error) {
	C := A.Copy()
	err := C.Subtract(B)
	return C, err
}

func (A *DenseMatrix) MinusDense(B *DenseMatrix) (*DenseMatrix, error) {
	C := A.Copy()
	err := C.SubtractDense(B)
	return C, err
}

func (A *DenseMatrix) Add(B MatrixRO) error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] += B.Get(i, j)
			index++
		}
	}

	return nil
}

func (A *DenseMatrix) AddDense(B *DenseMatrix) error {
	if A.cols != B.cols || A.rows != B.rows {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.elements[i*A.step+j] += B.elements[i*B.step+j]
		}
	}

	return nil
}

func (A *DenseMatrix) Subtract(B MatrixRO) error {
	if Bd, ok := B.(*DenseMatrix); ok {
		return A.SubtractDense(Bd)
	}

	if A.cols != B.Cols() || A.rows != B.Rows() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] -= B.Get(i, j)
			index++
		}
	}

	return nil
}

func (A *DenseMatrix) SubtractDense(B *DenseMatrix) error {

	if A.cols != B.cols || A.rows != B.rows {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		indexB := i * B.step

		for j := 0; j < A.cols; j++ {
			A.elements[indexA] -= B.elements[indexB]
			indexA++
			indexB++
		}
	}

	return nil
}

func (A *DenseMatrix) Times(B MatrixRO) (Matrix, error) {

	if Bd, ok := B.(*DenseMatrix); ok {
		return A.TimesDense(Bd)
	}

	if A.cols != B.Rows() {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.rows, B.Cols())

	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.elements[i*A.step+k] * B.Get(k, j)
			}
			C.elements[i*C.step+j] = sum
		}
	}

	return C, nil
}

type parJob struct {
	start, finish int
}

func parTimes1(A, B, C *DenseMatrix) {
	C = Zeros(A.rows, B.cols)

	mp := runtime.GOMAXPROCS(0)

	jobChan := make(chan box, 1+mp)

	go func() {
		rowCount := A.rows / mp
		for startRow := 0; startRow < A.rows; startRow += rowCount {
			start := startRow
			finish := startRow + rowCount
			if finish >= A.rows {
				finish = A.rows
			}
			jobChan <- parJob{start: start, finish: finish}
		}
		close(jobChan)
	}()

	wait := parFor(jobChan, func(iBox box) {
		job := iBox.(parJob)
		for i := job.start; i < job.finish; i++ {
			sums := C.elements[i*C.step : (i+1)*C.step]
			for k := 0; k < A.cols; k++ {
				for j := 0; j < B.cols; j++ {
					sums[j] += A.elements[i*A.step+k] * B.elements[k*B.step+j]
				}
			}
		}
	})
	wait()

	return
}

//this is an adaptation of code from a go-nuts post made by Dmitriy Vyukov
func parTimes2(A, B, C *DenseMatrix) {
	const threshold = 8

	currentGoroutineCount := 1
	maxGoroutines := runtime.GOMAXPROCS(0) + 2

	var aux func(sync chan bool, A, B, C *DenseMatrix, rs, re, cs, ce, ks, ke int)
	aux = func(sync chan bool, A, B, C *DenseMatrix, rs, re, cs, ce, ks, ke int) {
		dr := re - rs
		dc := ce - cs
		dk := ke - ks
		switch {
		case currentGoroutineCount < maxGoroutines && dr >= dc && dr >= dk && dr >= threshold:
			sync0 := make(chan bool, 1)
			rm := (rs + re) / 2
			currentGoroutineCount++
			go aux(sync0, A, B, C, rs, rm, cs, ce, ks, ke)
			aux(nil, A, B, C, rm, re, cs, ce, ks, ke)
			<-sync0
			currentGoroutineCount--
		case currentGoroutineCount < maxGoroutines && dc >= dk && dc >= dr && dc >= threshold:
			sync0 := make(chan bool, 1)
			cm := (cs + ce) / 2
			currentGoroutineCount++
			go aux(sync0, A, B, C, rs, re, cs, cm, ks, ke)
			aux(nil, A, B, C, rs, re, cm, ce, ks, ke)
			<-sync0
			currentGoroutineCount--
		case currentGoroutineCount < maxGoroutines && dk >= dc && dk >= dr && dk >= threshold:
			km := (ks + ke) / 2
			aux(nil, A, B, C, rs, re, cs, ce, ks, km)
			aux(nil, A, B, C, rs, re, cs, ce, km, ke)
		default:
			for row := rs; row < re; row++ {
				sums := C.elements[row*C.step : (row+1)*C.step]
				for k := ks; k < ke; k++ {
					for col := cs; col < ce; col++ {
						sums[col] += A.elements[row*A.step+k] * B.elements[k*B.step+col]
					}
				}
			}
		}
		if sync != nil {
			sync <- true
		}
	}

	aux(nil, A, B, C, 0, A.rows, 0, B.cols, 0, A.cols)

	return
}

var (
	WhichParMethod  = 2
	WhichSyncMethod = 1
)

func (A *DenseMatrix) TimesDense(B *DenseMatrix) (C *DenseMatrix, err error) {
	C = Zeros(A.rows, B.cols)
	err = A.TimesDenseFill(B, C)
	return
}
func (A *DenseMatrix) TimesDenseFill(B, C *DenseMatrix) (err error) {
	if C.rows != A.rows || C.cols != B.cols || A.cols != B.rows {
		err = ErrorDimensionMismatch
		return
	}
	if WhichParMethod > 0 && runtime.GOMAXPROCS(0) > 1 {
		switch WhichParMethod {
		case 1:
			parTimes1(A, B, C)
		case 2:
			parTimes2(A, B, C)
		}
	} else {
		switch {
		case A.cols > 100 && WhichSyncMethod == 2:
			transposeTimes(A, B, C)
		default:
			for i := 0; i < A.rows; i++ {
				sums := C.elements[i*C.step : (i+1)*C.step]
				for k, a := range A.elements[i*A.step : i*A.step + A.cols] {
					for j, b := range B.elements[k*B.step : k * B.step + B.cols] {
						sums[j] += a * b
					}
				}
			}
		}
	}

	return
}

func transposeTimes(A, B, C *DenseMatrix) {
	Bt := B.Transpose()

	Bcols := Bt.Arrays()

	for i := 0; i < A.rows; i++ {
		Arow := A.elements[i*A.step : i*A.step+A.cols]
		for j := 0; j < B.cols; j++ {
			Bcol := Bcols[j]
			for k := range Arow {
				C.elements[i*C.step+j] += Arow[k] * Bcol[k]
			}
		}
	}

	return
}

func (A *DenseMatrix) ElementMult(B MatrixRO) (Matrix, error) {
	C := A.Copy()
	err := C.ScaleMatrix(B)
	return C, err
}

func (A *DenseMatrix) ElementMultDense(B *DenseMatrix) (*DenseMatrix, error) {
	C := A.Copy()
	err := C.ScaleMatrixDense(B)
	return C, err
}

func (A *DenseMatrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] *= f
			index++
		}
	}
}

func (A *DenseMatrix) ScaleMatrix(B MatrixRO) error {
	if Bd, ok := B.(*DenseMatrix); ok {
		return A.ScaleMatrixDense(Bd)
	}

	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.Get(i, j)
			indexA++
		}
	}
	return nil
}

func (A *DenseMatrix) ScaleMatrixDense(B *DenseMatrix) error {
	if A.rows != B.rows || A.cols != B.cols {
		return ErrorDimensionMismatch
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		indexB := i * B.step
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.elements[indexB]
			indexA++
			indexB++
		}
	}
	return nil
}
