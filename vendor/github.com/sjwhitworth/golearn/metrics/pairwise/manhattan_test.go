package pairwise

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	. "github.com/smartystreets/goconvey/convey"
)

func TestManhattan(t *testing.T) {
	var vectorX, vectorY *mat64.Dense
	manhattan := NewManhattan()

	Convey("Given two vectors that are same", t, func() {
		vec := mat64.NewDense(7, 1, []float64{0, 1, -2, 3.4, 5, -6.7, 89})
		distance := manhattan.Distance(vec, vec)

		Convey("The result should be 0", func() {
			So(distance, ShouldEqual, 0)
		})
	})

	Convey("Given two vectors", t, func() {
		vectorX = mat64.NewDense(3, 1, []float64{2, 2, 3})
		vectorY = mat64.NewDense(3, 1, []float64{1, 4, 5})

		Convey("When calculating distance with column vectors", func() {
			result := manhattan.Distance(vectorX, vectorY)

			Convey("The result should be 5", func() {
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When calculating distance with row vectors", func() {
			vectorX.Copy(vectorX.T())
			vectorY.Copy(vectorY.T())
			result := manhattan.Distance(vectorX, vectorY)

			Convey("The result should be 5", func() {
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When calculating distance with different dimension matrices", func() {
			vectorX.Clone(vectorX.T())
			So(func() { manhattan.Distance(vectorX, vectorY) }, ShouldPanic)
		})

	})
}
