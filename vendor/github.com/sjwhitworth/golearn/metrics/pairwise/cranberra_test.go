package pairwise

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCranberrra(t *testing.T) {
	var vectorX, vectorY *mat64.Dense
	cranberra := NewCranberra()

	Convey("Given two vectors that are same", t, func() {
		vec := mat64.NewDense(7, 1, []float64{0, 1, -2, 3.4, 5, -6.7, 89})
		distance := cranberra.Distance(vec, vec)

		Convey("The result should be 0", func() {
			So(distance, ShouldEqual, 0)
		})
	})

	Convey("Given two vectors", t, func() {
		vectorX = mat64.NewDense(5, 1, []float64{1, 2, 3, 4, 9})
		vectorY = mat64.NewDense(5, 1, []float64{-5, -6, 7, 4, 3})

		Convey("When calculating distance with two vectors", func() {
			result := cranberra.Distance(vectorX, vectorY)

			Convey("The result should be 2.9", func() {
				So(result, ShouldEqual, 2.9)
			})
		})

		Convey("When calculating distance with row vectors", func() {
			vectorX.Copy(vectorX.T())
			vectorY.Copy(vectorY.T())
			result := cranberra.Distance(vectorX, vectorY)

			Convey("The result should be 2.9", func() {
				So(result, ShouldEqual, 2.9)
			})
		})

		Convey("When calculating distance with different dimension matrices", func() {
			vectorX.Clone(vectorX.T())
			So(func() { cranberra.Distance(vectorX, vectorY) }, ShouldPanic)
		})

	})
}
