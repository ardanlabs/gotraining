// Package knn implements a K Nearest Neighbors object, capable of both classification
// and regression. It accepts data in the form of a slice of float64s, which are then reshaped
// into a X by Y matrix.
package knn

import (
	"errors"
	"fmt"

	"github.com/gonum/matrix"
	"gonum.org/v1/gonum/mat"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/kdtree"
	"github.com/sjwhitworth/golearn/metrics/pairwise"
	"github.com/sjwhitworth/golearn/utilities"
)

// A KNNClassifier consists of a data matrix, associated labels in the same order as the matrix, searching algorithm, and a distance function.
// The accepted distance functions at this time are 'euclidean', 'manhattan', and 'cosine'.
// The accepted searching algorithm here are 'linear', and 'kdtree'.
// Optimisations only occur when things are identically group into identical
// AttributeGroups, which don't include the class variable, in the same order.
// Using weighted KNN when Weighted set to be true (default: false).
type KNNClassifier struct {
	base.BaseEstimator
	TrainingData       base.FixedDataGrid
	DistanceFunc       string
	Algorithm          string
	NearestNeighbours  int
	AllowOptimisations bool
	Weighted           bool
}

// NewKnnClassifier returns a new classifier
func NewKnnClassifier(distfunc, algorithm string, neighbours int) *KNNClassifier {
	KNN := KNNClassifier{}
	KNN.DistanceFunc = distfunc
	KNN.Algorithm = algorithm
	KNN.NearestNeighbours = neighbours
	KNN.Weighted = false
	KNN.AllowOptimisations = true
	return &KNN
}

// Fit stores the training data for later
func (KNN *KNNClassifier) Fit(trainingData base.FixedDataGrid) error {
	KNN.TrainingData = trainingData
	return nil
}

func (KNN *KNNClassifier) canUseOptimisations(what base.FixedDataGrid) bool {
	// Check that the two have exactly the same layout
	if !base.CheckStrictlyCompatible(what, KNN.TrainingData) {
		return false
	}
	// Check that the two are DenseInstances
	whatd, ok1 := what.(*base.DenseInstances)
	_, ok2 := KNN.TrainingData.(*base.DenseInstances)
	if !ok1 || !ok2 {
		return false
	}
	// Check that no Class Attributes are mixed in with the data
	classAttrs := whatd.AllClassAttributes()
	normalAttrs := base.NonClassAttributes(whatd)
	// Retrieve all the AGs
	ags := whatd.AllAttributeGroups()
	classAttrGroups := make([]base.AttributeGroup, 0)
	for agName := range ags {
		ag := ags[agName]
		attrs := ag.Attributes()
		matched := false
		for _, a := range attrs {
			for _, c := range classAttrs {
				if a.Equals(c) {
					matched = true
				}
			}
		}
		if matched {
			classAttrGroups = append(classAttrGroups, ag)
		}
	}
	for _, cag := range classAttrGroups {
		attrs := cag.Attributes()
		common := base.AttributeIntersect(normalAttrs, attrs)
		if len(common) != 0 {
			return false
		}
	}

	// Check that all of the Attributes are numeric
	for _, a := range normalAttrs {
		if _, ok := a.(*base.FloatAttribute); !ok {
			return false
		}
	}
	// If that's fine, return true
	return true
}

// Predict returns a classification for the vector, based on a vector input, using the KNN algorithm.
func (KNN *KNNClassifier) Predict(what base.FixedDataGrid) (base.FixedDataGrid, error) {
	// Check what distance function we are using
	var distanceFunc pairwise.PairwiseDistanceFunc
	switch KNN.DistanceFunc {
	case "euclidean":
		distanceFunc = pairwise.NewEuclidean()
	case "manhattan":
		distanceFunc = pairwise.NewManhattan()
	case "cosine":
		distanceFunc = pairwise.NewCosine()
	default:
		return nil, errors.New("unsupported distance function")
	}

	// Check what searching algorith, we are using
	if KNN.Algorithm != "linear" && KNN.Algorithm != "kdtree" {
		return nil, errors.New("unsupported searching algorithm")
	}

	// Check Compatibility
	allAttrs := base.CheckCompatible(what, KNN.TrainingData)
	if allAttrs == nil {
		// Don't have the same Attributes
		return nil, errors.New("attributes not compatible")
	}

	// Use optimised version if permitted
	if KNN.Algorithm == "linear" && KNN.AllowOptimisations {
		if KNN.DistanceFunc == "euclidean" {
			if KNN.canUseOptimisations(what) {
				return KNN.optimisedEuclideanPredict(what.(*base.DenseInstances)), nil
			}
		}
	}
	fmt.Println("Optimisations are switched off")

	// Remove the Attributes which aren't numeric
	allNumericAttrs := make([]base.Attribute, 0)
	for _, a := range allAttrs {
		if fAttr, ok := a.(*base.FloatAttribute); ok {
			allNumericAttrs = append(allNumericAttrs, fAttr)
		}
	}

	// If every Attribute is a FloatAttribute, then we remove the last one
	// because that is the Attribute we are trying to predict.
	if len(allNumericAttrs) == len(allAttrs) {
		allNumericAttrs = allNumericAttrs[:len(allNumericAttrs)-1]
	}

	// Generate return vector
	ret := base.GeneratePredictionVector(what)

	// Resolve Attribute specifications for both
	whatAttrSpecs := base.ResolveAttributes(what, allNumericAttrs)
	trainAttrSpecs := base.ResolveAttributes(KNN.TrainingData, allNumericAttrs)

	// Reserve storage for most the most similar items
	distances := make(map[int]float64)

	// Reserve storage for voting map
	maxmapInt := make(map[string]int)
	maxmapFloat := make(map[string]float64)

	// Reserve storage for row computations
	trainRowBuf := make([]float64, len(allNumericAttrs))
	predRowBuf := make([]float64, len(allNumericAttrs))

	_, maxRow := what.Size()
	curRow := 0

	// build kdtree if algorithm is 'kdtree'
	kd := kdtree.New()
	srcRowNoMap := make([]int, 0)
	if KNN.Algorithm == "kdtree" {
		buildData := make([][]float64, 0)
		KNN.TrainingData.MapOverRows(trainAttrSpecs, func(trainRow [][]byte, srcRowNo int) (bool, error) {
			oneData := make([]float64, len(allNumericAttrs))
			// Read the float values out
			for i, _ := range allNumericAttrs {
				oneData[i] = base.UnpackBytesToFloat(trainRow[i])
			}
			srcRowNoMap = append(srcRowNoMap, srcRowNo)
			buildData = append(buildData, oneData)
			return true, nil
		})

		err := kd.Build(buildData)
		if err != nil {
			return nil, err
		}
	}

	// Iterate over all outer rows
	what.MapOverRows(whatAttrSpecs, func(predRow [][]byte, predRowNo int) (bool, error) {

		if (curRow%1) == 0 && curRow > 0 {
			fmt.Printf("KNN: %.2f %% done\r", float64(curRow)*100.0/float64(maxRow))
		}
		curRow++

		// Read the float values out
		for i, _ := range allNumericAttrs {
			predRowBuf[i] = base.UnpackBytesToFloat(predRow[i])
		}

		predMat := utilities.FloatsToMatrix(predRowBuf)

		switch KNN.Algorithm {
		case "linear":
			// Find the closest match in the training data
			KNN.TrainingData.MapOverRows(trainAttrSpecs, func(trainRow [][]byte, srcRowNo int) (bool, error) {
				// Read the float values out
				for i, _ := range allNumericAttrs {
					trainRowBuf[i] = base.UnpackBytesToFloat(trainRow[i])
				}

				// Compute the distance
				trainMat := utilities.FloatsToMatrix(trainRowBuf)
				distances[srcRowNo] = distanceFunc.Distance(predMat, trainMat)
				return true, nil
			})

			sorted := utilities.SortIntMap(distances)
			values := sorted[:KNN.NearestNeighbours]

			length := make([]float64, KNN.NearestNeighbours)
			for k, v := range values {
				length[k] = distances[v]
			}

			var maxClass string
			if KNN.Weighted {
				maxClass = KNN.weightedVote(maxmapFloat, values, length)
			} else {
				maxClass = KNN.vote(maxmapInt, values)
			}
			base.SetClass(ret, predRowNo, maxClass)

		case "kdtree":
			// search kdtree
			values, length, err := kd.Search(KNN.NearestNeighbours, distanceFunc, predRowBuf)
			if err != nil {
				return false, err
			}

			// map values to srcRowNo
			for k, v := range values {
				values[k] = srcRowNoMap[v]
			}

			var maxClass string
			if KNN.Weighted {
				maxClass = KNN.weightedVote(maxmapFloat, values, length)
			} else {
				maxClass = KNN.vote(maxmapInt, values)
			}
			base.SetClass(ret, predRowNo, maxClass)
		}

		return true, nil

	})

	return ret, nil
}

func (KNN *KNNClassifier) String() string {
	return fmt.Sprintf("KNNClassifier(%s, %d)", KNN.DistanceFunc, KNN.NearestNeighbours)
}

func (KNN *KNNClassifier) vote(maxmap map[string]int, values []int) string {
	// Reset maxMap
	for a := range maxmap {
		maxmap[a] = 0
	}

	// Refresh maxMap
	for _, elem := range values {
		label := base.GetClass(KNN.TrainingData, elem)
		if _, ok := maxmap[label]; ok {
			maxmap[label]++
		} else {
			maxmap[label] = 1
		}
	}

	// Sort the maxMap
	var maxClass string
	maxVal := -1
	for a := range maxmap {
		if maxmap[a] > maxVal {
			maxVal = maxmap[a]
			maxClass = a
		}
	}
	return maxClass
}

func (KNN *KNNClassifier) weightedVote(maxmap map[string]float64, values []int, length []float64) string {
	// Reset maxMap
	for a := range maxmap {
		maxmap[a] = 0
	}

	// Refresh maxMap
	for k, elem := range values {
		label := base.GetClass(KNN.TrainingData, elem)
		if _, ok := maxmap[label]; ok {
			maxmap[label] += (1 / length[k])
		} else {
			maxmap[label] = (1 / length[k])
		}
	}

	// Sort the maxMap
	var maxClass string
	maxVal := -1.0
	for a := range maxmap {
		if maxmap[a] > maxVal {
			maxVal = maxmap[a]
			maxClass = a
		}
	}
	return maxClass
}

// GetMetadata returns required serialization information for this classifier
func (KNN *KNNClassifier) GetMetadata() base.ClassifierMetadataV1 {

	classifierParams := make(map[string]interface{})
	classifierParams["distance_func"] = KNN.DistanceFunc
	classifierParams["algorithm"] = KNN.Algorithm
	classifierParams["neighbours"] = KNN.NearestNeighbours
	classifierParams["weighted"] = KNN.Weighted
	classifierParams["allow_optimizations"] = KNN.AllowOptimisations

	return base.ClassifierMetadataV1{
		FormatVersion:      1,
		ClassifierName:     "KNN",
		ClassifierVersion:  "1.0",
		ClassifierMetadata: classifierParams,
	}
}

// Save outputs a given KNN classifier.
func (KNN *KNNClassifier) Save(filePath string) error {
	writer, err := base.CreateSerializedClassifierStub(filePath, KNN.GetMetadata())
	if err != nil {
		return err
	}
	fmt.Printf("writer: %v", writer)
	return KNN.SaveWithPrefix(writer, "")
}

// SaveWithPrefix outputs KNN as part of another file.
func (KNN *KNNClassifier) SaveWithPrefix(writer *base.ClassifierSerializer, prefix string) error {
	err := writer.WriteInstancesForKey(writer.Prefix(prefix, "TrainingInstances"), KNN.TrainingData, true)
	if err != nil {
		return err
	}
	err = writer.Close()
	return err
}

// Load reloads a given KNN classifier when it's the only thing in the output file.
func (KNN *KNNClassifier) Load(filePath string) error {
	reader, err := base.ReadSerializedClassifierStub(filePath)
	if err != nil {
		return err
	}

	return KNN.LoadWithPrefix(reader, "")
}

// LoadWithPrefix reloads a given KNN classifier when it's part of another file.
func (KNN *KNNClassifier) LoadWithPrefix(reader *base.ClassifierDeserializer, prefix string) error {

	clsMetadata, err := reader.ReadMetadataAtPrefix(prefix)
	if err != nil {
		return err
	}

	if clsMetadata.ClassifierName != "KNN" {
		return fmt.Errorf("This file doesn't contain a KNN classifier")
	}
	if clsMetadata.ClassifierVersion != "1.0" {
		return fmt.Errorf("Can't understand this file format")
	}

	metadata := clsMetadata.ClassifierMetadata
	KNN.DistanceFunc = metadata["distance_func"].(string)
	KNN.Algorithm = metadata["algorithm"].(string)
	//KNN.NearestNeighbours = metadata["neighbours"].(int)
	KNN.Weighted = metadata["weighted"].(bool)
	KNN.AllowOptimisations = metadata["allow_optimizations"].(bool)

	// 101 on why JSON is a bad serialization format
	floatNeighbours := metadata["neighbours"].(float64)
	KNN.NearestNeighbours = int(floatNeighbours)

	KNN.TrainingData, err = reader.GetInstancesForKey(reader.Prefix(prefix, "TrainingInstances"))

	return err
}

// ReloadKNNClassifier reloads a KNNClassifier when it's the only thing in an output file.
func ReloadKNNClassifier(filePath string) (*KNNClassifier, error) {
	stub := &KNNClassifier{}
	err := stub.Load(filePath)
	if err != nil {
		return nil, err
	}
	return stub, nil
}

// A KNNRegressor consists of a data matrix, associated result variables in the same order as the matrix, and a name.
type KNNRegressor struct {
	base.BaseEstimator
	Values       []float64
	DistanceFunc string
}

// NewKnnRegressor mints a new classifier.
func NewKnnRegressor(distfunc string) *KNNRegressor {
	KNN := KNNRegressor{}
	KNN.DistanceFunc = distfunc
	return &KNN
}

func (KNN *KNNRegressor) Fit(values []float64, numbers []float64, rows int, cols int) {
	if rows != len(values) {
		panic(matrix.ErrShape)
	}

	KNN.Data = mat.NewDense(rows, cols, numbers)
	KNN.Values = values
}

func (KNN *KNNRegressor) Predict(vector *mat.Dense, K int) float64 {
	// Get the number of rows
	rows, _ := KNN.Data.Dims()
	rownumbers := make(map[int]float64)
	labels := make([]float64, 0)

	// Check what distance function we are using
	var distanceFunc pairwise.PairwiseDistanceFunc
	switch KNN.DistanceFunc {
	case "euclidean":
		distanceFunc = pairwise.NewEuclidean()
	case "manhattan":
		distanceFunc = pairwise.NewManhattan()
	default:
		panic("unsupported distance function")
	}

	for i := 0; i < rows; i++ {
		row := KNN.Data.RowView(i)
		distance := distanceFunc.Distance(utilities.VectorToMatrix(row), vector)
		rownumbers[i] = distance
	}

	sorted := utilities.SortIntMap(rownumbers)
	values := sorted[:K]

	var sum float64
	for _, elem := range values {
		value := KNN.Values[elem]
		labels = append(labels, value)
		sum += value
	}

	average := sum / float64(K)
	return average
}
