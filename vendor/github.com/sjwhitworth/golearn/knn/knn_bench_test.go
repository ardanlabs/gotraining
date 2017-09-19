package knn

import (
	"fmt"
	"testing"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
)

func readMnist() (*base.DenseInstances, *base.DenseInstances) {
	// Create the class Attribute
	classAttrs := make(map[int]base.Attribute)
	classAttrs[0] = base.NewCategoricalAttribute()
	classAttrs[0].SetName("label")
	// Setup the class Attribute to be in its own group
	classAttrGroups := make(map[string]string)
	classAttrGroups["label"] = "ClassGroup"
	// The rest can go in a default group
	attrGroups := make(map[string]string)

	inst1, err := base.ParseCSVToInstancesWithAttributeGroups(
		"../examples/datasets/mnist_train.csv",
		attrGroups,
		classAttrGroups,
		classAttrs,
		true,
	)
	if err != nil {
		panic(err)
	}
	inst2, err := base.ParseCSVToTemplatedInstances(
		"../examples/datasets/mnist_test.csv",
		true,
		inst1,
	)
	if err != nil {
		panic(err)
	}
	return inst1, inst2
}

func BenchmarkKNNWithOpts(b *testing.B) {
	// Load
	train, test := readMnist()
	cls := NewKnnClassifier("euclidean", "linear", 1)
	cls.AllowOptimisations = true
	cls.Fit(train)
	predictions, err := cls.Predict(test)
	if err != nil {
		b.Error(err)
	}
	c, err := evaluation.GetConfusionMatrix(test, predictions)
	if err != nil {
		panic(err)
	}
	fmt.Println(evaluation.GetSummary(c))
	fmt.Println(evaluation.GetAccuracy(c))
}

func BenchmarkKNNWithNoOpts(b *testing.B) {
	// Load
	train, test := readMnist()
	cls := NewKnnClassifier("euclidean", "linear", 1)
	cls.AllowOptimisations = false
	cls.Fit(train)
	predictions, err := cls.Predict(test)
	if err != nil {
		b.Error(err)
	}
	c, err := evaluation.GetConfusionMatrix(test, predictions)
	if err != nil {
		panic(err)
	}
	fmt.Println(evaluation.GetSummary(c))
	fmt.Println(evaluation.GetAccuracy(c))
}
