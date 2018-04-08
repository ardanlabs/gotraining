package pairwise

import (
	"math"

	"github.com/gonum/matrix"
	"gonum.org/v1/gonum/mat"
)

type Chebyshev struct{}

func NewChebyshev() *Chebyshev {
	return &Chebyshev{}
}

func (c *Chebyshev) Distance(vectorX *mat.Dense, vectorY *mat.Dense) float64 {
	r1, c1 := vectorX.Dims()
	r2, c2 := vectorY.Dims()
	if r1 != r2 || c1 != c2 {
		panic(matrix.ErrShape)
	}

	max := float64(0)

	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			max = math.Max(max, math.Abs(vectorX.At(i, j)-vectorY.At(i, j)))
		}
	}

	return max
}
