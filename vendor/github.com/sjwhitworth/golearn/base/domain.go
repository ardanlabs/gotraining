// Package base provides base interfaces for GoLearn objects to implement.
// It also provides a raw base for those objects.
package base

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"

	"gonum.org/v1/gonum/mat"
)

// An Estimator is object that can ingest some data and train on it.
type Estimator interface {
	Fit()
}

// A Predictor is an object that provides predictions.
type Predictor interface {
	Predict()
}

// A Model is a supervised learning object, that is
// possible of scoring accuracy against a test set.
type Model interface {
	Score()
}

type BaseEstimator struct {
	Data *mat.Dense
}

// SaveEstimatorToGob serialises an estimator to a provided filepath, in gob format.
// See http://golang.org/pkg/encoding/gob for further details.
func SaveEstimatorToGob(path string, e *Estimator) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	err := enc.Encode(e)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path, b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
