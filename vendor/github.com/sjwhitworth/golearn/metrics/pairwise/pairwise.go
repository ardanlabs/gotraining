// Package pairwise implements utilities to evaluate pairwise distances or inner product (via kernel).
package pairwise

import (
	"github.com/gonum/matrix/mat64"
)

type PairwiseDistanceFunc interface {
	Distance(vectorX *mat64.Dense, vectorY *mat64.Dense) float64
}
