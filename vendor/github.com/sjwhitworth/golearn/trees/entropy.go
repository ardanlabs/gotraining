package trees

import (
	"github.com/sjwhitworth/golearn/base"
	"math"
	"sort"
)

//
// Information gain rule generator
//

// InformationGainRuleGenerator generates DecisionTreeRules which
// maximize information gain at each node.
type InformationGainRuleGenerator struct {
}

// GenerateSplitRule returns a DecisionTreeNode based on a non-class Attribute
// which maximises the information gain.
//
// IMPORTANT: passing a base.Instances with no Attributes other than the class
// variable will panic()
func (r *InformationGainRuleGenerator) GenerateSplitRule(f base.FixedDataGrid) *DecisionTreeRule {

	attrs := f.AllAttributes()
	classAttrs := f.AllClassAttributes()
	candidates := base.AttributeDifferenceReferences(attrs, classAttrs)

	return r.GetSplitRuleFromSelection(candidates, f)
}

// GetSplitRuleFromSelection returns a DecisionTreeRule which maximises
// the information gain amongst the considered Attributes.
//
// IMPORTANT: passing a zero-length consideredAttributes parameter will panic()
func (r *InformationGainRuleGenerator) GetSplitRuleFromSelection(consideredAttributes []base.Attribute, f base.FixedDataGrid) *DecisionTreeRule {

	var selectedAttribute base.Attribute

	// Parameter check
	if len(consideredAttributes) == 0 {
		panic("More Attributes should be considered")
	}

	// Next step is to compute the information gain at this node
	// for each randomly chosen attribute, and pick the one
	// which maximises it
	maxGain := math.Inf(-1)
	selectedVal := math.Inf(1)

	// Compute the base entropy
	classDist := base.GetClassDistribution(f)
	baseEntropy := getBaseEntropy(classDist)

	// Compute the information gain for each attribute
	for _, s := range consideredAttributes {
		var informationGain float64
		var splitVal float64
		if fAttr, ok := s.(*base.FloatAttribute); ok {
			var attributeEntropy float64
			attributeEntropy, splitVal = getNumericAttributeEntropy(f, fAttr)
			informationGain = baseEntropy - attributeEntropy
		} else {
			proposedClassDist := base.GetClassDistributionAfterSplit(f, s)
			localEntropy := getSplitEntropy(proposedClassDist)
			informationGain = baseEntropy - localEntropy
		}

		if informationGain > maxGain {
			maxGain = informationGain
			selectedAttribute = s
			selectedVal = splitVal
		}
	}

	// Pick the one which maximises IG
	return &DecisionTreeRule{selectedAttribute, selectedVal}
}

//
// Entropy functions
//

type numericSplitRef struct {
	val   float64
	class string
}

type splitVec []numericSplitRef

func (a splitVec) Len() int           { return len(a) }
func (a splitVec) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a splitVec) Less(i, j int) bool { return a[i].val < a[j].val }

func getNumericAttributeEntropy(f base.FixedDataGrid, attr *base.FloatAttribute) (float64, float64) {

	// Resolve Attribute
	attrSpec, err := f.GetAttribute(attr)
	if err != nil {
		panic(err)
	}

	// Build sortable vector
	_, rows := f.Size()
	refs := make([]numericSplitRef, rows)
	f.MapOverRows([]base.AttributeSpec{attrSpec}, func(val [][]byte, row int) (bool, error) {
		cls := base.GetClass(f, row)
		v := base.UnpackBytesToFloat(val[0])
		refs[row] = numericSplitRef{v, cls}
		return true, nil
	})

	// Sort
	sort.Sort(splitVec(refs))

	generateCandidateSplitDistribution := func(val float64) map[string]map[string]int {
		presplit := make(map[string]int)
		postplit := make(map[string]int)
		for _, i := range refs {
			if i.val < val {
				presplit[i.class]++
			} else {
				postplit[i.class]++
			}
		}
		ret := make(map[string]map[string]int)
		ret["0"] = presplit
		ret["1"] = postplit
		return ret
	}

	minSplitEntropy := math.Inf(1)
	minSplitVal := math.Inf(1)
	// Consider each possible function
	for i := 0; i < len(refs)-1; i++ {
		val := refs[i].val + refs[i+1].val
		val /= 2
		splitDist := generateCandidateSplitDistribution(val)
		splitEntropy := getSplitEntropy(splitDist)
		if splitEntropy < minSplitEntropy {
			minSplitEntropy = splitEntropy
			minSplitVal = val
		}
	}

	return minSplitEntropy, minSplitVal
}

// getSplitEntropy determines the entropy of the target
// class distribution after splitting on an base.Attribute
func getSplitEntropy(s map[string]map[string]int) float64 {
	ret := 0.0
	count := 0
	for a := range s {
		for c := range s[a] {
			count += s[a][c]
		}
	}
	for a := range s {
		total := 0.0
		for c := range s[a] {
			total += float64(s[a][c])
		}
		for c := range s[a] {
			ret -= float64(s[a][c]) / float64(count) * math.Log(float64(s[a][c])/float64(count)) / math.Log(2)
		}
		ret += total / float64(count) * math.Log(total/float64(count)) / math.Log(2)
	}
	return ret
}

// getBaseEntropy determines the entropy of the target
// class distribution before splitting on an base.Attribute
func getBaseEntropy(s map[string]int) float64 {
	ret := 0.0
	count := 0
	for k := range s {
		count += s[k]
	}
	for k := range s {
		ret -= float64(s[k]) / float64(count) * math.Log(float64(s[k])/float64(count)) / math.Log(2)
	}
	return ret
}
