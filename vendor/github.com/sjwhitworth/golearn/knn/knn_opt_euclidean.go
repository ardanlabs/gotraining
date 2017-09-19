package knn

// #include "knn.h"
import "C"

import (
	"github.com/sjwhitworth/golearn/base"
	"sort"
	"unsafe"
)

type dist _Ctype_struct_dist

type distanceRecs []_Ctype_struct_dist

func (d distanceRecs) Len() int           { return len(d) }
func (d distanceRecs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d distanceRecs) Less(i, j int) bool { return d[i].dist < d[j].dist }

func (KNN *KNNClassifier) optimisedEuclideanPredict(d *base.DenseInstances) base.FixedDataGrid {

	// Create return vector
	ret := base.GeneratePredictionVector(d)
	// Type-assert training data
	tr := KNN.TrainingData.(*base.DenseInstances)
	// Enumeration of AttributeGroups
	agPos := make(map[string]int)
	agTrain := tr.AllAttributeGroups()
	agPred := d.AllAttributeGroups()
	classAttrs := tr.AllClassAttributes()
	counter := 0
	for ag := range agTrain {
		// Detect whether the AttributeGroup has any classes in it
		attrs := agTrain[ag].Attributes()
		//matched := false
		if len(base.AttributeIntersect(classAttrs, attrs)) == 0 {
			agPos[ag] = counter
		}
		counter++
	}
	// Pointers to the start of each prediction row
	rowPointers := make([]*C.double, len(agPred))
	trainPointers := make([]*C.double, len(agPred))
	rowSizes := make([]int, len(agPred))
	for ag := range agPred {
		if ap, ok := agPos[ag]; ok {

			rowPointers[ap] = (*C.double)(unsafe.Pointer(&(agPred[ag].Storage()[0])))
			trainPointers[ap] = (*C.double)(unsafe.Pointer(&(agTrain[ag].Storage()[0])))
			rowSizes[ap] = agPred[ag].RowSizeInBytes() / 8
		}
	}
	_, predRows := d.Size()
	_, trainRows := tr.Size()
	// Crete the distance vector
	distanceVec := distanceRecs(make([]_Ctype_struct_dist, trainRows))
	// Additional datastructures
	voteVec := make([]int, KNN.NearestNeighbours)
	maxMap := make(map[string]int)

	for row := 0; row < predRows; row++ {
		for i := 0; i < trainRows; i++ {
			distanceVec[i].dist = 0
		}
		for ag := range agPred {
			if ap, ok := agPos[ag]; ok {
				C.euclidean_distance(
					&(distanceVec[0]),
					C.int(trainRows),
					C.int(len(agPred[ag].Attributes())),
					C.int(row),
					trainPointers[ap],
					rowPointers[ap],
				)
			}
		}
		sort.Sort(distanceVec)
		votes := distanceVec[:KNN.NearestNeighbours]
		for i, v := range votes {
			voteVec[i] = int(v.p)
		}
		maxClass := KNN.vote(maxMap, voteVec)
		base.SetClass(ret, row, maxClass)
	}
	return ret
}
