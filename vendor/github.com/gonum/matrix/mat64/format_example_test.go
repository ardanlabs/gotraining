// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64_test

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func ExampleFormatted() {
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Create a matrix formatting value with a prefix and calculating each column
	// width individually...
	fa := mat64.Formatted(a, mat64.Prefix("    "), mat64.Squeeze())

	// and then print with and without zero value elements.
	fmt.Printf("with all values:\na = %v\n\n", fa)
	fmt.Printf("with only non-zero values:\na = % v\n\n", fa)

	// Modify the matrix...
	a.Set(0, 2, 0)

	// and print it without zero value elements.
	fmt.Printf("after modification with only non-zero values:\na = % v\n\n", fa)

	// Modify the matrix again...
	a.Set(0, 2, 123.456)

	// and print it using scientific notation for large exponents.
	fmt.Printf("after modification with scientific notation:\na = %.2g\n\n", fa)
	// See golang.org/pkg/fmt/ floating-point verbs for a comprehensive list.

	// Output:
	// with all values:
	// a = ⎡1  2  3⎤
	//     ⎢0  4  5⎥
	//     ⎣0  0  6⎦
	//
	// with only non-zero values:
	// a = ⎡1  2  3⎤
	//     ⎢.  4  5⎥
	//     ⎣.  .  6⎦
	//
	// after modification with only non-zero values:
	// a = ⎡1  2  .⎤
	//     ⎢.  4  5⎥
	//     ⎣.  .  6⎦
	//
	// after modification with scientific notation:
	// a = ⎡1  2  1.2e+02⎤
	//     ⎢0  4        5⎥
	//     ⎣0  0        6⎦
}

func ExampleExcerpt() {
	// Excerpt allows diagnostic display of very large
	// matrices and vectors.

	// The big matrix is too large to properly print...
	big := mat64.NewDense(100, 100, nil)
	for i := 0; i < 100; i++ {
		big.Set(i, i, 1)
	}

	// so only print corner excerpts of the matrix.
	fmt.Printf("excerpt big identity matrix: %v\n\n",
		mat64.Formatted(big, mat64.Prefix(" "), mat64.Excerpt(3)))

	// The long vector is also too large, ...
	long := mat64.NewVector(100, nil)
	for i := 0; i < 100; i++ {
		long.SetVec(i, float64(i))
	}

	// ... so print end excerpts of the vector,
	fmt.Printf("excerpt long column vector: %v\n\n",
		mat64.Formatted(long, mat64.Prefix(" "), mat64.Excerpt(3)))
	// or its transpose.
	fmt.Printf("excerpt long row vector: %v\n",
		mat64.Formatted(long.T(), mat64.Prefix(" "), mat64.Excerpt(3)))

	// Output:
	// excerpt big identity matrix: Dims(100, 100)
	//  ⎡1  0  0  ...  ...  0  0  0⎤
	//  ⎢0  1  0            0  0  0⎥
	//  ⎢0  0  1            0  0  0⎥
	//   .
	//   .
	//   .
	//  ⎢0  0  0            1  0  0⎥
	//  ⎢0  0  0            0  1  0⎥
	//  ⎣0  0  0  ...  ...  0  0  1⎦
	//
	// excerpt long column vector: Dims(100, 1)
	//  ⎡ 0⎤
	//  ⎢ 1⎥
	//  ⎢ 2⎥
	//   .
	//   .
	//   .
	//  ⎢97⎥
	//  ⎢98⎥
	//  ⎣99⎦
	//
	// excerpt long row vector: Dims(1, 100)
	//  [ 0   1   2  ...  ...  97  98  99]

}
