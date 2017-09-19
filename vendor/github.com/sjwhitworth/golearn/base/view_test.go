package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInstancesViewRows(t *testing.T) {
	Convey("Given Iris", t, func() {
		instOrig, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldEqual, nil)
		Convey("Given a new row map containing only row 5", func() {
			rMap := make(map[int]int)
			rMap[0] = 5
			instView := NewInstancesViewFromRows(instOrig, rMap)
			Convey("The internal structure should be right...", func() {
				So(instView.rows[0], ShouldEqual, 5)
			})
			Convey("The reconstructed values should be correct...", func() {
				str := "5.4 3.9 1.7 0.4 Iris-setosa"
				row := instView.RowString(0)
				So(row, ShouldEqual, str)
			})
			Convey("And the size should be correct...", func() {
				width, height := instView.Size()
				So(width, ShouldEqual, 5)
				So(height, ShouldEqual, 150)
			})
		})
	})
}

func TestInstancesViewFromVisible(t *testing.T) {
	Convey("Given Iris", t, func() {
		instOrig, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldEqual, nil)
		Convey("Generate something that says every other row should be visible", func() {
			rowVisiblex1 := make([]int, 0)
			_, totalRows := instOrig.Size()
			for i := 0; i < totalRows; i += 2 {
				rowVisiblex1 = append(rowVisiblex1, i)
			}
			instViewx1 := NewInstancesViewFromVisible(instOrig, rowVisiblex1, instOrig.AllAttributes())
			for i, a := range rowVisiblex1 {
				rowStr1 := instViewx1.RowString(i)
				rowStr2 := instOrig.RowString(a)
				So(rowStr1, ShouldEqual, rowStr2)
			}
			Convey("And then generate something that says that every other row than that should be visible", func() {
				rowVisiblex2 := make([]int, 0)
				for i := 0; i < totalRows; i += 4 {
					rowVisiblex2 = append(rowVisiblex1, i)
				}
				instViewx2 := NewInstancesViewFromVisible(instOrig, rowVisiblex2, instOrig.AllAttributes())
				for i, a := range rowVisiblex2 {
					rowStr1 := instViewx2.RowString(i)
					rowStr2 := instOrig.RowString(a)
					So(rowStr1, ShouldEqual, rowStr2)
				}
			})
		})
	})
}

func TestInstancesViewAttrs(t *testing.T) {
	Convey("Given Iris", t, func() {
		instOrig, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldEqual, nil)
		Convey("Given a new Attribute vector with the last 4...", func() {
			cMap := instOrig.AllAttributes()[1:]
			instView := NewInstancesViewFromAttrs(instOrig, cMap)
			Convey("The size should be correct", func() {
				h, v := instView.Size()
				So(h, ShouldEqual, 4)
				_, vOrig := instOrig.Size()
				So(v, ShouldEqual, vOrig)
			})
			Convey("There should be 4 Attributes...", func() {
				attrs := instView.AllAttributes()
				So(len(attrs), ShouldEqual, 4)
			})
			Convey("There should be 4 Attributes with the right headers...", func() {
				attrs := instView.AllAttributes()
				So(attrs[0].GetName(), ShouldEqual, "Sepal width")
				So(attrs[1].GetName(), ShouldEqual, "Petal length")
				So(attrs[2].GetName(), ShouldEqual, "Petal width")
				So(attrs[3].GetName(), ShouldEqual, "Species")
			})
			Convey("There should be a class Attribute...", func() {
				attrs := instView.AllClassAttributes()
				So(len(attrs), ShouldEqual, 1)
			})
			Convey("The class Attribute should be preserved...", func() {
				attrs := instView.AllClassAttributes()
				So(attrs[0].GetName(), ShouldEqual, "Species")
			})
			Convey("Attempts to get the filtered Attribute should fail...", func() {
				_, err := instView.GetAttribute(instOrig.AllAttributes()[0])
				So(err, ShouldNotEqual, nil)
			})
			Convey("The filtered Attribute should not appear in the RowString", func() {
				str := "3.9 1.7 0.4 Iris-setosa"
				row := instView.RowString(5)
				So(row, ShouldEqual, str)
			})
			Convey("The filtered Attributes should all be the same type...", func() {
				attrs := instView.AllAttributes()
				_, ok1 := attrs[0].(*FloatAttribute)
				_, ok2 := attrs[1].(*FloatAttribute)
				_, ok3 := attrs[2].(*FloatAttribute)
				_, ok4 := attrs[3].(*CategoricalAttribute)
				So(ok1, ShouldEqual, true)
				So(ok2, ShouldEqual, true)
				So(ok3, ShouldEqual, true)
				So(ok4, ShouldEqual, true)
			})
			Convey("The InstancesView should match one prepared earlier...", func() {
				instRef, err := ParseCSVToInstances("../examples/datasets/iris_headers_subset.csv", true)
				So(err, ShouldBeNil)
				So(InstancesAreEqual(instRef, instView), ShouldBeTrue)
				Convey("And a DenseInstances conversion should too...", func() {
					instView2 := NewDenseCopy(instRef)
					So(InstancesAreEqual(instRef, instView2), ShouldBeTrue)
				})
			})
		})
	})
}
