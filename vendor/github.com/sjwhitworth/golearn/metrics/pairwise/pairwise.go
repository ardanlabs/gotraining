// Package pairwise implements utilities to evaluate pairwise distances or inner product (via kernel).
package pairwise

import (
	"gonum.org/v1/gonum/mat"
)

type PairwiseDistanceFunc interface {
	Distance(vectorX *mat.Dense, vectorY *mat.Dense) float64
}
