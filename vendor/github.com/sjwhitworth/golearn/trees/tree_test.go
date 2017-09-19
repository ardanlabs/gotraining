package trees

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func verifyTreeClassification(trainData, testData base.FixedDataGrid) {
	rand.Seed(44414515)
	Convey("Using InferID3Tree to create the tree and do the fitting", func() {
		Convey("Using a RandomTreeRule", func() {
			randomTreeRuleGenerator := new(RandomTreeRuleGenerator)
			randomTreeRuleGenerator.Attributes = 2
			root := InferID3Tree(trainData, randomTreeRuleGenerator)

			Convey("Predicting with the tree", func() {
				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.5)
				})
			})
		})

		Convey("Using a InformationGainRule", func() {
			informationGainRuleGenerator := new(InformationGainRuleGenerator)
			root := InferID3Tree(trainData, informationGainRuleGenerator)

			Convey("Predicting with the tree", func() {
				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.5)
				})
			})
		})
		Convey("Using a GiniCoefficientRuleGenerator", func() {
			gRuleGen := new(GiniCoefficientRuleGenerator)
			root := InferID3Tree(trainData, gRuleGen)
			Convey("Predicting with the tree", func() {
				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.5)
				})
			})
		})
		Convey("Using a InformationGainRatioRuleGenerator", func() {
			gRuleGen := new(InformationGainRatioRuleGenerator)
			root := InferID3Tree(trainData, gRuleGen)
			Convey("Predicting with the tree", func() {
				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.5)
				})
			})
		})

	})

	Convey("Using NewRandomTree to create the tree", func() {
		root := NewRandomTree(2)

		Convey("Fitting with the tree", func() {
			err := root.Fit(trainData)
			So(err, ShouldBeNil)

			Convey("Predicting with the tree, *without* pruning first", func() {
				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.5)
				})
			})

			Convey("Predicting with the tree, pruning first", func() {
				root.Prune(testData)

				predictions, err := root.Predict(testData)
				So(err, ShouldBeNil)

				confusionMatrix, err := evaluation.GetConfusionMatrix(testData, predictions)
				So(err, ShouldBeNil)

				Convey("Predictions should be somewhat accurate", func() {
					So(evaluation.GetAccuracy(confusionMatrix), ShouldBeGreaterThan, 0.4)
				})
			})
		})
	})
}

func TestRandomTreeClassificationAfterDiscretisation(t *testing.T) {
	Convey("Predictions on filtered data with a Random Tree", t, func() {
		instances, err := base.ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldBeNil)

		trainData, testData := base.InstancesTrainTestSplit(instances, 0.6)

		filter := filters.NewChiMergeFilter(instances, 0.9)
		for _, a := range base.NonClassFloatAttributes(instances) {
			filter.AddAttribute(a)
		}
		filter.Train()
		filteredTrainData := base.NewLazilyFilteredInstances(trainData, filter)
		filteredTestData := base.NewLazilyFilteredInstances(testData, filter)
		verifyTreeClassification(filteredTrainData, filteredTestData)
	})
}

func TestRandomTreeClassificationWithoutDiscretisation(t *testing.T) {
	Convey("Predictions on filtered data with a Random Tree", t, func() {
		instances, err := base.ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		So(err, ShouldBeNil)

		trainData, testData := base.InstancesTrainTestSplit(instances, 0.6)

		verifyTreeClassification(trainData, testData)
	})
}

func TestPRIVATEgetSplitEntropy(t *testing.T) {
	outlook := make(map[string]map[string]int)
	outlook["sunny"] = make(map[string]int)
	outlook["overcast"] = make(map[string]int)
	outlook["rain"] = make(map[string]int)
	outlook["sunny"]["play"] = 2
	outlook["sunny"]["noplay"] = 3
	outlook["overcast"]["play"] = 4
	outlook["rain"]["play"] = 3
	outlook["rain"]["noplay"] = 2

	Convey("Should calculate split entropy accurately", t, func() {
		So(getSplitEntropy(outlook), ShouldAlmostEqual, 0.694, 0.001)
	})
}

func TestID3Inference(t *testing.T) {
	Convey("Producing a decision tree with ID3 inference on a dataset", t, func() {
		instances, err := base.ParseCSVToInstances("../examples/datasets/tennis.csv", true)
		So(err, ShouldBeNil)

		Convey("Using InferID3Tree to create the tree and do the fitting", func() {
			rule := new(InformationGainRuleGenerator)
			root := InferID3Tree(instances, rule)

			itBuildsTheCorrectDecisionTree(root)
		})

		Convey("Using NewID3DecisionTree to build the tree and fitting explicitly", func() {
			tree := NewID3DecisionTree(0.0)
			tree.Fit(instances)
			root := tree.Root

			itBuildsTheCorrectDecisionTree(root)
		})
	})
}

func TestPRIVATEgetNumericAttributeEntropy(t *testing.T) {
	Convey("Checking a particular split...", t, func() {
		instances, err := base.ParseCSVToInstances("../examples/datasets/c45-numeric.csv", true)
		So(err, ShouldBeNil)
		Convey("Fetching the right Attribute", func() {
			attr := base.GetAttributeByName(instances, "Attribute2")
			So(attr, ShouldNotEqual, nil)
			Convey("Finding the threshold...", func() {
				_, threshold := getNumericAttributeEntropy(instances, attr.(*base.FloatAttribute))
				So(threshold, ShouldAlmostEqual, 82.5)
			})
		})
	})
}

func itBuildsTheCorrectDecisionTree(root *DecisionTreeNode) {
	Convey("The root should be 'outlook'", func() {
		So(root.SplitRule.SplitAttr.GetName(), ShouldEqual, "outlook")
	})

	sunny := root.Children["sunny"]
	overcast := root.Children["overcast"]
	rainy := root.Children["rainy"]

	Convey("After the 'sunny' node, the decision should split on 'humidity'", func() {
		So(sunny.SplitRule.SplitAttr.GetName(), ShouldEqual, "humidity")
	})
	Convey("After the 'rainy' node, the decision should split on 'windy'", func() {
		So(rainy.SplitRule.SplitAttr.GetName(), ShouldEqual, "windy")
	})
	Convey("There should be no splits after the 'overcast' node", func() {
		So(overcast.SplitRule.SplitAttr, ShouldBeNil)
	})

	highHumidity := sunny.Children["high"]
	normalHumidity := sunny.Children["normal"]
	windy := rainy.Children["true"]
	notWindy := rainy.Children["false"]

	Convey("The leaf nodes should be classified 'yes' or 'no' accurately", func() {
		So(highHumidity.Class, ShouldEqual, "no")
		So(normalHumidity.Class, ShouldEqual, "yes")
		So(windy.Class, ShouldEqual, "no")
		So(notWindy.Class, ShouldEqual, "yes")
		So(overcast.Class, ShouldEqual, "yes")
	})
}
