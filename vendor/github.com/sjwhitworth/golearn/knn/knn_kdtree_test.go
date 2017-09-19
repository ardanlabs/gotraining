package knn

import (
	"testing"

	"github.com/sjwhitworth/golearn/base"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKnnClassifierWithoutOptimisationsWithKdtree(t *testing.T) {
	Convey("Given labels, a classifier and data", t, func() {
		trainingData, err := base.ParseCSVToInstances("knn_train_1.csv", false)
		So(err, ShouldBeNil)

		testingData, err := base.ParseCSVToInstances("knn_test_1.csv", false)
		So(err, ShouldBeNil)

		cls := NewKnnClassifier("euclidean", "kdtree", 2)
		cls.AllowOptimisations = false
		cls.Fit(trainingData)
		predictions, err := cls.Predict(testingData)
		So(err, ShouldBeNil)
		So(predictions, ShouldNotEqual, nil)

		Convey("When predicting the label for our first vector", func() {
			result := base.GetClass(predictions, 0)
			Convey("The result should be 'blue", func() {
				So(result, ShouldEqual, "blue")
			})
		})

		Convey("When predicting the label for our second vector", func() {
			result2 := base.GetClass(predictions, 1)
			Convey("The result should be 'red", func() {
				So(result2, ShouldEqual, "red")
			})
		})
	})
}

func TestKnnClassifierWithTemplatedInstances1WithKdtree(t *testing.T) {
	Convey("Given two basically identical files...", t, func() {
		trainingData, err := base.ParseCSVToInstances("knn_train_2.csv", true)
		So(err, ShouldBeNil)
		testingData, err := base.ParseCSVToTemplatedInstances("knn_test_2.csv", true, trainingData)
		So(err, ShouldBeNil)

		cls := NewKnnClassifier("euclidean", "kdtree", 2)
		cls.Fit(trainingData)
		predictions, err := cls.Predict(testingData)
		So(err, ShouldBeNil)
		So(predictions, ShouldNotBeNil)
	})
}

func TestKnnClassifierWithTemplatedInstances1SubsetWithKdtree(t *testing.T) {
	Convey("Given two basically identical files...", t, func() {
		trainingData, err := base.ParseCSVToInstances("knn_train_2.csv", true)
		So(err, ShouldBeNil)
		testingData, err := base.ParseCSVToTemplatedInstances("knn_test_2_subset.csv", true, trainingData)
		So(err, ShouldBeNil)

		cls := NewKnnClassifier("euclidean", "kdtree", 2)
		cls.Fit(trainingData)
		predictions, err := cls.Predict(testingData)
		So(err, ShouldBeNil)
		So(predictions, ShouldNotBeNil)
	})
}

func TestKnnClassifierImplementsClassifierWithKdtree(t *testing.T) {
	cls := NewKnnClassifier("euclidean", "kdtree", 2)
	var c base.Classifier = cls
	if len(c.String()) < 1 {
		t.Fail()
	}
}
