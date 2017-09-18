// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv_test

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func ExampleNormal() {
	// Create a normal distribution
	dist := distuv.Normal{
		Mu:    2,
		Sigma: 5,
	}

	data := make([]float64, 1e5)

	// Draw some random values from the standard normal distribution
	for i := range data {
		data[i] = dist.Rand()
	}

	mean, std := stat.MeanStdDev(data, nil)
	meanErr := stat.StdErr(std, float64(len(data)))

	fmt.Printf("mean= %1.1f ± %0.1v\n", mean, meanErr)

	// Output:
	// mean= 2.0 ± 0.02
}
