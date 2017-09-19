package base

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCSVGetRows(t *testing.T) {
	Convey("Getting the number of rows for a CSV file", t, func() {
		Convey("With a valid file path", func() {
			numNonHeaderRows := 150

			Convey("When the CSV file doesn't have a header row", func() {
				lineCount, err := ParseCSVGetRows("../examples/datasets/iris.csv")
				So(err, ShouldBeNil)

				Convey("It counts the correct number of rows", func() {
					So(lineCount, ShouldEqual, numNonHeaderRows)
				})
			})

			Convey("When the CSV file has a header row", func() {
				lineCount, err := ParseCSVGetRows("../examples/datasets/iris_headers.csv")
				So(err, ShouldBeNil)

				Convey("It counts the correct number of rows, *including* the header row", func() {
					So(lineCount, ShouldEqual, numNonHeaderRows+1)
				})
			})
		})

		Convey("With a path to a non-existent file", func() {
			_, err := ParseCSVGetRows("../examples/datasets/non-existent.csv")

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestParseCSVGetAttributes(t *testing.T) {
	Convey("Getting the attributes in the headers of a CSV file", t, func() {
		attributes := ParseCSVGetAttributes("../examples/datasets/iris_headers.csv", true)
		sepalLengthAttribute := attributes[0]
		speciesAttribute := attributes[4]

		Convey("It gets the correct types for the headers based on the column values", func() {
			So(sepalLengthAttribute.GetType(), ShouldEqual, Float64Type)
			So(speciesAttribute.GetType(), ShouldEqual, CategoricalType)
		})

		Convey("It gets the correct attribute names", func() {
			So(sepalLengthAttribute.GetName(), ShouldEqual, "Sepal length")
			So(speciesAttribute.GetName(), ShouldEqual, "Species")
		})
	})
}

func TestParseCSVSniffAttributeTypes(t *testing.T) {
	Convey("Getting just the attribute types for the columns in the CSV", t, func() {
		attributes := ParseCSVSniffAttributeTypes("../examples/datasets/iris_headers.csv", true)

		Convey("It gets the correct types", func() {
			So(attributes[0].GetType(), ShouldEqual, Float64Type)
			So(attributes[1].GetType(), ShouldEqual, Float64Type)
			So(attributes[2].GetType(), ShouldEqual, Float64Type)
			So(attributes[3].GetType(), ShouldEqual, Float64Type)
			So(attributes[4].GetType(), ShouldEqual, CategoricalType)
		})
	})
}

func TestParseCSVSniffAttributeNames(t *testing.T) {
	Convey("Getting just the attribute name for the columns in the CSV", t, func() {
		attributeNames := ParseCSVSniffAttributeNames("../examples/datasets/iris_headers.csv", true)

		Convey("It gets the correct names", func() {
			So(attributeNames[0], ShouldEqual, "Sepal length")
			So(attributeNames[1], ShouldEqual, "Sepal width")
			So(attributeNames[2], ShouldEqual, "Petal length")
			So(attributeNames[3], ShouldEqual, "Petal width")
			So(attributeNames[4], ShouldEqual, "Species")
		})
	})
}

func TestParseCSVToInstances(t *testing.T) {
	Convey("Parsing a CSV file to Instances", t, func() {
		Convey("Given a path to a reasonable CSV file", func() {
			instances, err := ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
			So(err, ShouldBeNil)

			Convey("Should parse the rows correctly", func() {
				So(instances.RowString(0), ShouldEqual, "5.1 3.5 1.4 0.2 Iris-setosa")
				So(instances.RowString(50), ShouldEqual, "7.0 3.2 4.7 1.4 Iris-versicolor")
				So(instances.RowString(100), ShouldEqual, "6.3 3.3 6.0 2.5 Iris-virginica")
			})
		})

		Convey("Given a path to another reasonable CSV file", func() {
			_, err := ParseCSVToInstances("../examples/datasets/c45-numeric.csv", true)
			So(err, ShouldBeNil)
		})

		Convey("Given a path to a non-existent file", func() {
			_, err := ParseCSVToInstances("../examples/datasets/non-existent.csv", true)

			Convey("It should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Given a path to a CSV file with awkward data", func() { // what's so awkward about it?
			instances, err := ParseCSVToInstances("../examples/datasets/chim.csv", true)
			So(err, ShouldBeNil)

			Convey("It parses the data correctly, assigning the correct types to attributes", func() {
				attributes := instances.AllAttributes()
				So(attributes[0].GetType(), ShouldEqual, Float64Type)
				So(attributes[1].GetType(), ShouldEqual, CategoricalType)
			})
		})
	})
}
