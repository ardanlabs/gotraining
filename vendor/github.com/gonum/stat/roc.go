// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat

import "sort"

// ROC returns paired false positive rate (FPR) and true positive rate
// (TPR) values corresponding to n cutoffs spanning the relative
// (or receiver) operator characteristic (ROC) curve obtained when y is
// treated as a binary classifier for classes with weights.
//
// Cutoffs are equally spaced from eps less than the minimum value of y
// to the maximum value of y, including both endpoints meaning that the
// resulting ROC curve will always begin at (0,0) and end at (1,1).
//
// The input y must be sorted, and SortWeightedLabeled can be used in
// order to sort y together with classes and weights.
//
// For a given cutoff value, observations corresponding to entries in y
// greater than the cutoff value are classified as false, while those
// below (or equal to) the cutoff value are classified as true. These
// assigned class labels are compared with the true values in the classes
// slice and used to calculate the FPR and TPR.
//
// If weights is nil, all weights are treated as 1.
//
// When n is zero all possible cutoffs are calculated, resulting
// in fpr and tpr having length one greater than the number of unique
// values in y. When n is greater than one fpr and tpr will be returned
// with length n. ROC will panic if n is equal to one or less than 0.
//
// More details about ROC curves are available at
// https://en.wikipedia.org/wiki/Receiver_operating_characteristic
func ROC(n int, y []float64, classes []bool, weights []float64) (tpr, fpr []float64) {
	if len(y) != len(classes) {
		panic("stat: slice length mismatch")
	}
	if weights != nil && len(y) != len(weights) {
		panic("stat: slice length mismatch")
	}
	if !sort.Float64sAreSorted(y) {
		panic("stat: input must be sorted")
	}

	var incWidth, tol float64
	if n == 0 {
		if len(y) == 0 {
			return nil, nil
		}
		tpr = make([]float64, len(y)+1)
		fpr = make([]float64, len(y)+1)
	} else {
		if n < 2 {
			panic("stat: cannot calculate fewer than 2 points on a ROC curve")
		}
		if len(y) == 0 {
			return nil, nil
		}
		tpr = make([]float64, n)
		fpr = make([]float64, n)
		incWidth = (y[len(y)-1] - y[0]) / float64(n-1)
		tol = y[0] + incWidth
		if incWidth == 0 {
			tpr[n-1] = 1
			fpr[n-1] = 1
			return
		}
	}

	var bin int = 1 // the initial bin is known to have 0 fpr and 0 tpr
	var nPos, nNeg float64
	for i, u := range classes {
		var posWeight, negWeight float64 = 0, 1
		if weights != nil {
			negWeight = weights[i]
		}
		if u {
			posWeight, negWeight = negWeight, posWeight
		}
		nPos += posWeight
		nNeg += negWeight
		tpr[bin] += posWeight
		fpr[bin] += negWeight

		// Assess if the bin needs to be updated. If n is zero,
		// the bin is always updated, unless consecutive y values
		// are equal. Otherwise, the bin must be updated until it
		// matches the next y value (skipping empty bins).
		if n == 0 {
			if i != (len(y)-1) && y[i] != y[i+1] {
				bin++
				tpr[bin] = tpr[bin-1]
				fpr[bin] = fpr[bin-1]
			}
		} else {
			for i != (len(y)-1) && y[i+1] > tol {
				tol += incWidth
				bin++
				tpr[bin] = tpr[bin-1]
				fpr[bin] = fpr[bin-1]
			}
		}
	}
	if n == 0 {
		tpr = tpr[:(bin + 1)]
		fpr = fpr[:(bin + 1)]
	}

	invNeg := 1 / nNeg
	invPos := 1 / nPos
	for i := range tpr {
		tpr[i] *= invPos
		fpr[i] *= invNeg
	}
	tpr[len(tpr)-1] = 1
	fpr[len(fpr)-1] = 1

	return tpr, fpr
}
