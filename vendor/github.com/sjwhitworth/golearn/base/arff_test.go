package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestParseARFFGetRows(t *testing.T) {
	Convey("Getting the number of rows for a ARFF file", t, func() {
		Convey("With a valid file path", func() {
			numNonHeaderRows := 150
			lineCount, err := ParseARFFGetRows("../examples/datasets/iris.arff")
			So(err, ShouldBeNil)
			So(lineCount, ShouldEqual, numNonHeaderRows)
		})
	})
}

func TestParseARFFGetAttributes(t *testing.T) {
	Convey("Getting the attributes in the headers of a CSV file", t, func() {
		attributes := ParseARFFGetAttributes("../examples/datasets/iris.arff")
		sepalLengthAttribute := attributes[0]
		sepalWidthAttribute := attributes[1]
		petalLengthAttribute := attributes[2]
		petalWidthAttribute := attributes[3]
		speciesAttribute := attributes[4]

		Convey("It gets the correct types for the headers based on the column values", func() {
			_, ok1 := sepalLengthAttribute.(*FloatAttribute)
			_, ok2 := sepalWidthAttribute.(*FloatAttribute)
			_, ok3 := petalLengthAttribute.(*FloatAttribute)
			_, ok4 := petalWidthAttribute.(*FloatAttribute)
			sA, ok5 := speciesAttribute.(*CategoricalAttribute)
			So(ok1, ShouldBeTrue)
			So(ok2, ShouldBeTrue)
			So(ok3, ShouldBeTrue)
			So(ok4, ShouldBeTrue)
			So(ok5, ShouldBeTrue)
			So(sA.GetValues(), ShouldResemble, []string{"Iris-setosa", "Iris-versicolor", "Iris-virginica"})
		})
	})
}

func TestParseARFF1(t *testing.T) {
	Convey("Should just be able to load in an ARFF...", t, func() {
		inst, err := ParseDenseARFFToInstances("../examples/datasets/iris.arff")
		So(err, ShouldBeNil)
		So(inst, ShouldNotBeNil)
		So(inst.RowString(0), ShouldEqual, "5.1 3.5 1.4 0.2 Iris-setosa")
		So(inst.RowString(50), ShouldEqual, "7.0 3.2 4.7 1.4 Iris-versicolor")
		So(inst.RowString(100), ShouldEqual, "6.3 3.3 6.0 2.5 Iris-virginica")
	})
}

func TestParseARFF2(t *testing.T) {
	Convey("Loading the weather dataset...", t, func() {
		inst, err := ParseDenseARFFToInstances("../examples/datasets/weather.arff")
		So(err, ShouldBeNil)

		Convey("Attributes should be right...", func() {
			So(GetAttributeByName(inst, "outlook"), ShouldNotBeNil)
			So(GetAttributeByName(inst, "temperature"), ShouldNotBeNil)
			So(GetAttributeByName(inst, "humidity"), ShouldNotBeNil)
			So(GetAttributeByName(inst, "windy"), ShouldNotBeNil)
			So(GetAttributeByName(inst, "play"), ShouldNotBeNil)
			Convey("outlook attribute values should match reference...", func() {
				outlookAttr := GetAttributeByName(inst, "outlook").(*CategoricalAttribute)
				So(outlookAttr.GetValues(), ShouldResemble, []string{"sunny", "overcast", "rainy"})
			})
			Convey("windy values should match reference...", func() {
				windyAttr := GetAttributeByName(inst, "windy").(*CategoricalAttribute)
				So(windyAttr.GetValues(), ShouldResemble, []string{"TRUE", "FALSE"})
			})
			Convey("play values should match reference...", func() {
				playAttr := GetAttributeByName(inst, "play").(*CategoricalAttribute)
				So(playAttr.GetValues(), ShouldResemble, []string{"yes", "no"})
			})

		})

	})
}

func TestSerializeToARFF(t *testing.T) {
	Convey("Loading the weather dataset...", t, func() {
		inst, err := ParseDenseARFFToInstances("../examples/datasets/weather.arff")
		So(err, ShouldBeNil)
		Convey("Saving back should suceed...", func() {
			attrs := ParseARFFGetAttributes("../examples/datasets/weather.arff")
			f, err := ioutil.TempFile("", "inst")
			So(err, ShouldBeNil)
			err = SerializeInstancesToDenseARFFWithAttributes(inst, attrs, f.Name(), "weather")
			So(err, ShouldBeNil)
			Convey("Reading the file back should be lossless...", func() {
				inst2, err := ParseDenseARFFToInstances(f.Name())
				So(err, ShouldBeNil)
				So(InstancesAreEqual(inst, inst2), ShouldBeTrue)
			})
			Convey("The file should be exactly the same as the original...", func() {
				ref, err := ioutil.ReadFile("../examples/datasets/weather.arff")
				So(err, ShouldBeNil)
				gen, err := ioutil.ReadFile(f.Name())
				So(err, ShouldBeNil)
				So(string(gen), ShouldEqual, string(ref))
			})
		})
	})
}
