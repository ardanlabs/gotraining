package base

import (
	"github.com/gonum/matrix/mat64"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInlineMat64Creation(t *testing.T) {

	Convey("Given a literal array...", t, func() {
		mat := mat64.NewDense(4, 3, []float64{
			1, 0, 1,
			0, 1, 1,
			0, 0, 0,
			1, 1, 0,
		})
		inst := InstancesFromMat64(4, 3, mat)
		attrs := inst.AllAttributes()
		Convey("Attributes should be well-defined...", func() {
			So(len(attrs), ShouldEqual, 3)
		})

		Convey("No class variables set by default...", func() {
			classAttrs := inst.AllClassAttributes()
			So(len(classAttrs), ShouldEqual, 0)
		})

		Convey("Getting values should work...", func() {
			as, err := inst.GetAttribute(attrs[0])
			So(err, ShouldBeNil)
			valBytes := inst.Get(as, 3)
			val := UnpackBytesToFloat(valBytes)
			So(val, ShouldAlmostEqual, 1.0)
		})

	})

}
