package pairwise

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Euclidean struct{}

func NewEuclidean() *Euclidean {
	return &Euclidean{}
}

// InnerProduct computes a Eucledian inner product.
func (e *Euclidean) InnerProduct(vectorX *mat.Dense, vectorY *mat.Dense) float64 {
	subVector := mat.NewDense(0, 0, nil)
	subVector.MulElem(vectorX, vectorY)
	result := mat.Sum(subVector)

	return result
}

// Distance computes Euclidean distance (also known as L2 distance).
func (e *Euclidean) Distance(vectorX *mat.Dense, vectorY *mat.Dense) float64 {
	subVector := mat.NewDense(0, 0, nil)
	subVector.Sub(vectorX, vectorY)

	result := e.InnerProduct(subVector, subVector)

	return math.Sqrt(result)
}
