package pairwise

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

type PolyKernel struct {
	degree int
}

// NewPolyKernel returns a d-degree polynomial kernel
func NewPolyKernel(degree int) *PolyKernel {
	return &PolyKernel{degree: degree}
}

// InnerProduct computes the inner product through a kernel trick
// K(x, y) = (x^T y + 1)^d
func (p *PolyKernel) InnerProduct(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	subVectorX := vectorX.ColView(0)
	subVectorY := vectorY.ColView(0)
	result := mat64.Dot(subVectorX, subVectorY)
	result = math.Pow(result+1, float64(p.degree))

	return result
}

// Distance computes distance under the polynomial kernel (maybe not needed?)
func (p *PolyKernel) Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	subVector := mat64.NewDense(0, 0, nil)
	subVector.Sub(vectorX, vectorY)
	result := p.InnerProduct(subVector, subVector)

	return math.Sqrt(result)
}
