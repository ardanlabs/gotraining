package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLazySortDesc(t *testing.T) {
	Convey("Given data that's not already sorted descending", t, func() {
		unsorted, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldBeNil)

		as1 := ResolveAllAttributes(unsorted)
		So(isSortedDesc(unsorted, as1[0]), ShouldBeFalse)

		Convey("Given reference data that's alredy sorted descending", func() {
			sortedDescending, err := ParseCSVToInstances("../examples/datasets/iris_sorted_desc.csv", true)
			So(err, ShouldBeNil)

			as2 := ResolveAllAttributes(sortedDescending)
			So(isSortedDesc(sortedDescending, as2[0]), ShouldBeTrue)

			Convey("LazySorting Descending", func() {
				result, err := LazySort(unsorted, Descending, as1[0:len(as1)-1])
				So(err, ShouldBeNil)

				Convey("Result should be sorted descending", func() {
					So(isSortedDesc(result, as1[0]), ShouldBeTrue)
				})

				Convey("Result should match the reference", func() {
					So(InstancesAreEqual(sortedDescending, result), ShouldBeTrue)
				})
			})
		})
	})
}

func TestLazySortAsc(t *testing.T) {
	Convey("Given data that's not already sorted ascending", t, func() {
		unsorted, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldBeNil)

		as1 := ResolveAllAttributes(unsorted)
		So(isSortedAsc(unsorted, as1[0]), ShouldBeFalse)

		Convey("Given reference data that's alredy sorted ascending", func() {
			sortedAscending, err := ParseCSVToInstances("../examples/datasets/iris_sorted_asc.csv", true)
			So(err, ShouldBeNil)

			as2 := ResolveAllAttributes(sortedAscending)
			So(isSortedAsc(sortedAscending, as2[0]), ShouldBeTrue)

			Convey("LazySorting Ascending", func() {
				result, err := LazySort(unsorted, Ascending, as1[0:len(as1)-1])
				So(err, ShouldBeNil)

				Convey("Result should be sorted descending", func() {
					So(isSortedAsc(result, as1[0]), ShouldBeTrue)
				})

				Convey("Result should match the reference", func() {
					So(InstancesAreEqual(sortedAscending, result), ShouldBeTrue)
				})

				Convey("First element of Result should equal known value", func() {
					So(result.RowString(0), ShouldEqual, "4.3 3.0 1.1 0.1 Iris-setosa")
				})
			})
		})
	})
}
