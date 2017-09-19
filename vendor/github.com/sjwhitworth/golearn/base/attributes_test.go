package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloatAttributeSysVal(t *testing.T) {
	Convey("Given some float", t, func() {
		x := "1.21"
		attr := NewFloatAttribute("")
		Convey("When the float gets packed", func() {
			packed := attr.GetSysValFromString(x)
			Convey("And then unpacked", func() {
				unpacked := attr.GetStringFromSysVal(packed)
				Convey("The unpacked version should be the same", func() {
					So(unpacked, ShouldEqual, x)
				})
			})
		})
	})
}

func TestCategoricalAttributeVal(t *testing.T) {
	attr := NewCategoricalAttribute()
	Convey("Given some string", t, func() {
		x := "hello world!"
		Convey("When the string gets converted", func() {
			packed := attr.GetSysValFromString(x)
			Convey("And then unconverted", func() {
				unpacked := attr.GetStringFromSysVal(packed)
				Convey("The unpacked version should be the same", func() {
					So(unpacked, ShouldEqual, x)
				})
			})
		})
	})
	Convey("Given some second string", t, func() {
		x := "hello world 1!"
		Convey("When the string gets converted", func() {
			packed := attr.GetSysValFromString(x)
			So(packed[0], ShouldEqual, 0x1)
			Convey("And then unconverted", func() {
				unpacked := attr.GetStringFromSysVal(packed)
				Convey("The unpacked version should be the same", func() {
					So(unpacked, ShouldEqual, x)
				})
			})
		})
	})
}

func TestBinaryAttribute(t *testing.T) {
	attr := new(BinaryAttribute)
	Convey("Given some binary Attribute", t, func() {
		Convey("SetName, GetName should be equal", func() {
			attr.SetName("Hello")
			So(attr.GetName(), ShouldEqual, "Hello")
		})
		Convey("Non-zero values should equal 1", func() {
			sysVal := attr.GetSysValFromString("1")
			So(sysVal[0], ShouldEqual, 1)
		})
		Convey("Zero values should equal 0", func() {
			sysVal := attr.GetSysValFromString("0")
			So(sysVal[0], ShouldEqual, 0)
		})
	})
}
