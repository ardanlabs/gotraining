package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func isSortedAsc(inst FixedDataGrid, attr AttributeSpec) bool {
	valPrev := 0.0
	_, rows := inst.Size()
	for i := 0; i < rows; i++ {
		cur := UnpackBytesToFloat(inst.Get(attr, i))
		if i > 0 {
			if valPrev > cur {
				return false
			}
		}
		valPrev = cur
	}
	return true
}

func isSortedDesc(inst FixedDataGrid, attr AttributeSpec) bool {
	valPrev := 0.0
	_, rows := inst.Size()
	for i := 0; i < rows; i++ {
		cur := UnpackBytesToFloat(inst.Get(attr, i))
		if i > 0 {
			if valPrev < cur {
				return false
			}
		}
		valPrev = cur
	}
	return true
}

func TestSortDesc(t *testing.T) {
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

			Convey("Sorting Descending", func() {
				result, err := Sort(unsorted, Descending, as1[0:len(as1)-1])
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

func TestSortAsc(t *testing.T) {
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

			Convey("Sorting Ascending", func() {
				result, err := Sort(unsorted, Ascending, as1[0:len(as1)-1])
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
