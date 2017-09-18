// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

func ExamplePrincipalComponents() {
	// iris is a truncated sample of the Fisher's Iris dataset.
	n := 10
	d := 4
	iris := mat64.NewDense(n, d, []float64{
		5.1, 3.5, 1.4, 0.2,
		4.9, 3.0, 1.4, 0.2,
		4.7, 3.2, 1.3, 0.2,
		4.6, 3.1, 1.5, 0.2,
		5.0, 3.6, 1.4, 0.2,
		5.4, 3.9, 1.7, 0.4,
		4.6, 3.4, 1.4, 0.3,
		5.0, 3.4, 1.5, 0.2,
		4.4, 2.9, 1.4, 0.2,
		4.9, 3.1, 1.5, 0.1,
	})

	// Calculate the principal component direction vectors
	// and variances.
	var pc stat.PC
	ok := pc.PrincipalComponents(iris, nil)
	if !ok {
		return
	}
	fmt.Printf("variances = %.4f\n\n", pc.Vars(nil))

	// Project the data onto the first 2 principal components.
	k := 2
	var proj mat64.Dense
	proj.Mul(iris, pc.Vectors(nil).Slice(0, d, 0, k))

	fmt.Printf("proj = %.4f", mat64.Formatted(&proj, mat64.Prefix("       ")))

	// Output:
	// variances = [0.1666 0.0207 0.0079 0.0019]
	//
	// proj = ⎡-6.1686   1.4659⎤
	//        ⎢-5.6767   1.6459⎥
	//        ⎢-5.6699   1.3642⎥
	//        ⎢-5.5643   1.3816⎥
	//        ⎢-6.1734   1.3309⎥
	//        ⎢-6.7278   1.4021⎥
	//        ⎢-5.7743   1.1498⎥
	//        ⎢-6.0466   1.4714⎥
	//        ⎢-5.2709   1.3570⎥
	//        ⎣-5.7533   1.6207⎦
}
