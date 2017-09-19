package pairwise

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

type RBFKernel struct {
	gamma float64
}

// NewRBFKernel returns a representation of a Radial Basis Function Kernel
func NewRBFKernel(gamma float64) *RBFKernel {
	return &RBFKernel{gamma: gamma}
}

// InnerProduct computes the inner product through a kernel trick
// K(x, y) = exp(-gamma * ||x - y||^2)
func (r *RBFKernel) InnerProduct(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	euclidean := NewEuclidean()
	distance := euclidean.Distance(vectorX, vectorY)

	result := math.Exp(-r.gamma * math.Pow(distance, 2))

	return result
}
