package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHighDimensionalInstancesLoad(t *testing.T) {
	Convey("Given a high-dimensional dataset...", t, func() {
		_, err := ParseCSVToInstances("../examples/datasets/mnist_train.csv", true)
		So(err, ShouldEqual, nil)
	})
}
func TestHighDimensionalInstancesLoad2(t *testing.T) {
	Convey("Given a high-dimensional dataset...", t, func() {
		// Create the class Attribute
		classAttrs := make(map[int]Attribute)
		classAttrs[0] = NewCategoricalAttribute()
		classAttrs[0].SetName("Number")
		// Setup the class Attribute to be in its own group
		classAttrGroups := make(map[string]string)
		classAttrGroups["Number"] = "ClassGroup"
		// The rest can go in a default group
		attrGroups := make(map[string]string)

		_, err := ParseCSVToInstancesWithAttributeGroups(
			"../examples/datasets/mnist_train.csv",
			attrGroups,
			classAttrGroups,
			classAttrs,
			true,
		)
		So(err, ShouldEqual, nil)
	})
}
