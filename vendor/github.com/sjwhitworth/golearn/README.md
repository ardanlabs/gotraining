GoLearn
=======

<img src="http://talks.golang.org/2013/advconc/gopherhat.jpg" width=125><br>
[![GoDoc](https://godoc.org/github.com/sjwhitworth/golearn?status.png)](https://godoc.org/github.com/sjwhitworth/golearn)
[![Build Status](https://travis-ci.org/sjwhitworth/golearn.png?branch=master)](https://travis-ci.org/sjwhitworth/golearn)<br>

[![Support via Gittip](https://rawgithub.com/twolfson/gittip-badge/0.2.0/dist/gittip.png)](https://www.gittip.com/sjwhitworth/)

GoLearn is a 'batteries included' machine learning library for Go. **Simplicity**, paired with customisability, is the goal.
We are in active development, and would love comments from users out in the wild. Drop us a line on Twitter.

twitter: [@golearn_ml](http://www.twitter.com/golearn_ml)

Install
=======

See [here](https://github.com/sjwhitworth/golearn/wiki/Installation) for installation instructions.

Getting Started
=======

Data are loaded in as Instances. You can then perform matrix like operations on them, and pass them to estimators.
GoLearn implements the scikit-learn interface of Fit/Predict, so you can easily swap out estimators for trial and error.
GoLearn also includes helper functions for data, like cross validation, and train and test splitting.

```go
package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	rawData, err := base.ParseCSVToInstances("datasets/iris.csv", false)
	if err != nil {
		panic(err)
	}

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
```

```
Iris-virginica	28	2	  56	0.9333	0.9333  0.9333
Iris-setosa	    29	0	  59	1.0000  1.0000	1.0000
Iris-versicolor	27	2	  57	0.9310	0.9310  0.9310
Overall accuracy: 0.9545
```

Examples
========

GoLearn comes with practical examples. Dive in and see what is going on.

```bash
cd $GOPATH/src/github.com/sjwhitworth/golearn/examples/knnclassifier
go run knnclassifier_iris.go
```
```bash
cd $GOPATH/src/github.com/sjwhitworth/golearn/examples/instances
go run instances.go
```
```bash
cd $GOPATH/src/github.com/sjwhitworth/golearn/examples/trees
go run trees.go
```

Docs
====

 * [English](https://github.com/sjwhitworth/golearn/wiki)
 * [中文文档(简体)](doc/zh_CN/Home.md)
 * [中文文档(繁体)](doc/zh_TW/Home.md)

Join the team
=============

Please send me a mail at stephenjameswhitworth@gmail.com
