// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64_test

import (
	"fmt"
	"log"
	"math"

	"github.com/gonum/matrix"
	"github.com/gonum/matrix/mat64"
)

func ExampleGSVD() {
	// Perform a GSVD factorization on food production/consumption data for the
	// three years 1990, 2000 and 2014, for Africa and Latin America/Caribbean.
	//
	// See Lee et al. doi:10.1371/journal.pone.0030098 and
	// Alter at al. doi:10.1073/pnas.0530258100 for more details.
	var gsvd mat64.GSVD
	ok := gsvd.Factorize(FAO.Africa, FAO.LatinAmericaCaribbean, matrix.GSVDU|matrix.GSVDV|matrix.GSVDQ)
	if !ok {
		log.Fatal("GSVD factorization failed")
	}

	var u, v mat64.Dense
	u.UFromGSVD(&gsvd)
	v.VFromGSVD(&gsvd)

	s1 := gsvd.ValuesA(nil)
	s2 := gsvd.ValuesB(nil)

	fmt.Printf("Africa\n\ts1 = %.4f\n\n\tU = %.4f\n\n",
		s1, mat64.Formatted(&u, mat64.Prefix("\t    "), mat64.Excerpt(2)))
	fmt.Printf("Latin America/Caribbean\n\ts2 = %.4f\n\n\tV = %.4f\n",
		s2, mat64.Formatted(&v, mat64.Prefix("\t    "), mat64.Excerpt(2)))

	var zeroR, q mat64.Dense
	zeroR.ZeroRFromGSVD(&gsvd)
	q.QFromGSVD(&gsvd)
	q.Mul(&zeroR, &q)
	fmt.Printf("\nCommon basis vectors\n\n\tQ^T = %.4f\n",
		mat64.Formatted(q.T(), mat64.Prefix("\t      ")))

	// Calculate the antisymmetric angular distances for each eigenvariable.
	fmt.Println("\nSignificance:")
	for i := 0; i < 3; i++ {
		fmt.Printf("\teigenvar_%d: %+.4f\n", i, math.Atan(s1[i]/s2[i])-math.Pi/4)
	}

	// Output:
	//
	// Africa
	// 	s1 = [1.0000 0.9344 0.5118]
	//
	// 	U = Dims(21, 21)
	// 	    ⎡-0.0005   0.0142  ...  ...  -0.0060  -0.0055⎤
	// 	    ⎢-0.0010   0.0019             0.0071   0.0075⎥
	// 	     .
	// 	     .
	// 	     .
	// 	    ⎢-0.0007  -0.0024             0.9999  -0.0001⎥
	// 	    ⎣-0.0010  -0.0016  ...  ...  -0.0001   0.9999⎦
	//
	// Latin America/Caribbean
	// 	s2 = [0.0047 0.3563 0.8591]
	//
	// 	V = Dims(14, 14)
	// 	    ⎡ 0.1362   0.0008  ...  ...   0.0700   0.2636⎤
	// 	    ⎢ 0.1830  -0.0040             0.2908   0.7834⎥
	// 	     .
	// 	     .
	// 	     .
	// 	    ⎢-0.2598  -0.0324             0.9339  -0.2170⎥
	// 	    ⎣-0.8386   0.1494  ...  ...  -0.1639   0.4121⎦
	//
	// Common basis vectors
	//
	// 	Q^T = ⎡ -8172.4084   -4524.2933    4813.9616⎤
	// 	      ⎢ 22581.8020   12397.1070  -16364.8933⎥
	// 	      ⎣ -8910.8462  -10902.1488   15762.8719⎦
	//
	// Significance:
	// 	eigenvar_0: +0.7807
	// 	eigenvar_1: +0.4211
	// 	eigenvar_2: -0.2482
}
