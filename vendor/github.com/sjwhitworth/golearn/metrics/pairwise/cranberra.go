package pairwise

import (
	"math"

	"github.com/gonum/matrix"
	"gonum.org/v1/gonum/mat"
)

type Cranberra struct{}

func NewCranberra() *Cranberra {
	return &Cranberra{}
}

func cranberraDistanceStep(num float64, denom float64) float64 {
	if num == .0 && denom == .0 {
		return .0
	}
	return num / denom
}

func (c *Cranberra) Distance(vectorX *mat.Dense, vectorY *mat.Dense) float64 {
	r1, c1 := vectorX.Dims()
	r2, c2 := vectorY.Dims()
	if r1 != r2 || c1 != c2 {
		panic(matrix.ErrShape)
	}

	sum := .0

	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			p1 := vectorX.At(i, j)
			p2 := vectorY.At(i, j)

			num := math.Abs(p1 - p2)
			denom := math.Abs(p1) + math.Abs(p2)

			sum += cranberraDistanceStep(num, denom)
		}
	}

	return sum
}
