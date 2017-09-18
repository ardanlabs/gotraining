package pairwise

import (
	"math"

	"github.com/gonum/matrix"
	"github.com/gonum/matrix/mat64"
)

type Manhattan struct{}

func NewManhattan() *Manhattan {
	return &Manhattan{}
}

// Distance computes the Manhattan distance, also known as L1 distance.
// == the sum of the absolute values of elements.
func (m *Manhattan) Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	r1, c1 := vectorX.Dims()
	r2, c2 := vectorY.Dims()
	if r1 != r2 || c1 != c2 {
		panic(matrix.ErrShape)
	}

	result := .0

	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result += math.Abs(vectorX.At(i, j) - vectorY.At(i, j))
		}
	}
	return result
}
