package pairwise

import (
	"math"

	"github.com/gonum/matrix"
	"gonum.org/v1/gonum/mat"
)

type Manhattan struct{}

func NewManhattan() *Manhattan {
	return &Manhattan{}
}

// Distance computes the Manhattan distance, also known as L1 distance.
// == the sum of the absolute values of elements.
func (m *Manhattan) Distance(vectorX *mat.Dense, vectorY *mat.Dense) float64 {
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
