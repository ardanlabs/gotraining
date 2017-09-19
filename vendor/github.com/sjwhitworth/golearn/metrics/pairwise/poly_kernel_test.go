package pairwise

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPolyKernel(t *testing.T) {
	var vectorX, vectorY *mat64.Dense
	polyKernel := NewPolyKernel(3)

	Convey("Given two vectors", t, func() {
		vectorX = mat64.NewDense(3, 1, []float64{1, 2, 3})
		vectorY = mat64.NewDense(3, 1, []float64{2, 4, 5})

		Convey("When doing inner product", func() {
			result := polyKernel.InnerProduct(vectorX, vectorY)

			Convey("The result should be 17576", func() {
				So(result, ShouldEqual, 17576)
			})
		})

		Convey("When calculating distance", func() {
			result := polyKernel.Distance(vectorX, vectorY)

			Convey("The result should alomost equal 31.622776601683793", func() {
				So(result, ShouldAlmostEqual, 31.622776601683793)
			})

		})

	})
}
