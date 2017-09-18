package pairwise

import (
	"math"

	"github.com/gonum/matrix/mat64"
)

type Cosine struct{}

func NewCosine() *Cosine {
	return &Cosine{}
}

// Dot computes dot value of vectorX and vectorY.
func (c *Cosine) Dot(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	subVector := mat64.NewDense(0, 0, nil)
	subVector.MulElem(vectorX, vectorY)
	result := mat64.Sum(subVector)

	return result
}

// Distance computes Cosine distance.
// It will return distance which represented as 1-cos() (ranged from 0 to 2).
func (c *Cosine) Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64 {
	dotXY := c.Dot(vectorX, vectorY)
	lengthX := math.Sqrt(c.Dot(vectorX, vectorX))
	lengthY := math.Sqrt(c.Dot(vectorY, vectorY))

	cos := dotXY / (lengthX * lengthY)

	return 1 - cos
}
