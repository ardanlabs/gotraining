package kdtree

import (
	"testing"

	"github.com/sjwhitworth/golearn/metrics/pairwise"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKdtree(t *testing.T) {
	kd := New()

	Convey("Given a kdtree", t, func() {
		data := [][]float64{{2, 3}, {5, 4}, {4, 7}, {8, 1}, {7, 2}, {9, 6}}
		kd.Build(data)
		euclidean := pairwise.NewEuclidean()

		Convey("When k is 3 with euclidean", func() {
			result, _, _ := kd.Search(3, euclidean, []float64{7, 3})

			Convey("The result[0] should be 4", func() {
				So(result[0], ShouldEqual, 4)
			})
			Convey("The result[1] should be 3", func() {
				So(result[1], ShouldEqual, 3)
			})
			Convey("The result[2] should be 1", func() {
				So(result[2], ShouldEqual, 1)
			})
		})

		Convey("When k is 2 with euclidean", func() {
			result, _, _ := kd.Search(2, euclidean, []float64{7, 3})

			Convey("The result[0] should be 4", func() {
				So(result[0], ShouldEqual, 4)
			})
			Convey("The result[1] should be 1", func() {
				So(result[1], ShouldEqual, 1)
			})
		})

	})
}
