package pairwise

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEuclidean(t *testing.T) {
	var vectorX, vectorY *mat64.Dense
	euclidean := NewEuclidean()

	Convey("Given two vectors", t, func() {
		vectorX = mat64.NewDense(3, 1, []float64{1, 2, 3})
		vectorY = mat64.NewDense(3, 1, []float64{2, 4, 5})

		Convey("When doing inner product", func() {
			result := euclidean.InnerProduct(vectorX, vectorY)

			Convey("The result should be 25", func() {
				So(result, ShouldEqual, 25)
			})
		})

		Convey("When calculating distance", func() {
			result := euclidean.Distance(vectorX, vectorY)

			Convey("The result should be 3", func() {
				So(result, ShouldEqual, 3)
			})

		})

	})
}
