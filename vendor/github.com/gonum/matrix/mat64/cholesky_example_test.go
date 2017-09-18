// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64_test

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func ExampleCholesky() {
	// Construct a symmetric positive definite matrix.
	tmp := mat64.NewDense(4, 4, []float64{
		2, 6, 8, -4,
		1, 8, 7, -2,
		2, 2, 1, 7,
		8, -2, -2, 1,
	})
	var a mat64.SymDense
	a.SymOuterK(1, tmp)

	fmt.Printf("a = %0.4v\n", mat64.Formatted(&a, mat64.Prefix("    ")))

	// Compute the cholesky factorization.
	var chol mat64.Cholesky
	if ok := chol.Factorize(&a); !ok {
		fmt.Println("a matrix is not positive semi-definite.")
	}

	// Find the determinant.
	fmt.Printf("\nThe determinant of a is %0.4g\n\n", chol.Det())

	// Use the factorization to solve the system of equations a * x = b.
	b := mat64.NewVector(4, []float64{1, 2, 3, 4})
	var x mat64.Vector
	if err := x.SolveCholeskyVec(&chol, b); err != nil {
		fmt.Println("Matrix is near singular: ", err)
	}
	fmt.Println("Solve a * x = b")
	fmt.Printf("x = %0.4v\n", mat64.Formatted(&x, mat64.Prefix("    ")))

	// Extract the factorization and check that it equals the original matrix.
	var t mat64.TriDense
	t.LFromCholesky(&chol)
	var test mat64.Dense
	test.Mul(&t, t.T())
	fmt.Println()
	fmt.Printf("L * L^T = %0.4v\n", mat64.Formatted(&a, mat64.Prefix("          ")))

	// Output:
	// a = ⎡120  114   -4  -16⎤
	//     ⎢114  118   11  -24⎥
	//     ⎢ -4   11   58   17⎥
	//     ⎣-16  -24   17   73⎦
	//
	// The determinant of a is 1.543e+06
	//
	// Solve a * x = b
	// x = ⎡  -0.239⎤
	//     ⎢  0.2732⎥
	//     ⎢-0.04681⎥
	//     ⎣  0.1031⎦
	//
	// L * L^T = ⎡120  114   -4  -16⎤
	//           ⎢114  118   11  -24⎥
	//           ⎢ -4   11   58   17⎥
	//           ⎣-16  -24   17   73⎦
}

func ExampleCholeskySymRankOne() {
	a := mat64.NewSymDense(4, []float64{
		1, 1, 1, 1,
		0, 2, 3, 4,
		0, 0, 6, 10,
		0, 0, 0, 20,
	})
	fmt.Printf("A = %0.4v\n", mat64.Formatted(a, mat64.Prefix("    ")))

	// Compute the Cholesky factorization.
	var chol mat64.Cholesky
	if ok := chol.Factorize(a); !ok {
		fmt.Println("matrix a is not positive definite.")
	}

	x := mat64.NewVector(4, []float64{0, 0, 0, 1})
	fmt.Printf("\nx = %0.4v\n", mat64.Formatted(x, mat64.Prefix("    ")))

	// Rank-1 update the factorization.
	chol.SymRankOne(&chol, 1, x)
	// Rank-1 update the matrix a.
	a.SymRankOne(a, 1, x)

	var au mat64.SymDense
	au.FromCholesky(&chol)

	// Print the matrix that was updated directly.
	fmt.Printf("\nA' =        %0.4v\n", mat64.Formatted(a, mat64.Prefix("            ")))
	// Print the matrix recovered from the factorization.
	fmt.Printf("\nU'^T * U' = %0.4v\n", mat64.Formatted(&au, mat64.Prefix("            ")))

	// Output:
	// A = ⎡ 1   1   1   1⎤
	//     ⎢ 1   2   3   4⎥
	//     ⎢ 1   3   6  10⎥
	//     ⎣ 1   4  10  20⎦
	//
	// x = ⎡0⎤
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎣1⎦
	//
	// A' =        ⎡ 1   1   1   1⎤
	//             ⎢ 1   2   3   4⎥
	//             ⎢ 1   3   6  10⎥
	//             ⎣ 1   4  10  21⎦
	//
	// U'^T * U' = ⎡ 1   1   1   1⎤
	//             ⎢ 1   2   3   4⎥
	//             ⎢ 1   3   6  10⎥
	//             ⎣ 1   4  10  21⎦
}
