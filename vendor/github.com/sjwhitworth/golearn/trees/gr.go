package trees

import (
	"github.com/sjwhitworth/golearn/base"
	"math"
)

//
// Information Gatio Ratio generator
//

// InformationGainRatioRuleGenerator generates DecisionTreeRules which
// maximise the InformationGain at each node.
type InformationGainRatioRuleGenerator struct {
}

// GenerateSplitRule returns a DecisionTreeRule which maximises information
// gain ratio considering every available Attribute.
//
// IMPORTANT: passing a base.Instances with no Attributes other than the class
// variable will panic()
func (r *InformationGainRatioRuleGenerator) GenerateSplitRule(f base.FixedDataGrid) *DecisionTreeRule {

	attrs := f.AllAttributes()
	classAttrs := f.AllClassAttributes()
	candidates := base.AttributeDifferenceReferences(attrs, classAttrs)

	return r.GetSplitRuleFromSelection(candidates, f)
}

// GetSplitRuleFromSelection returns the DecisionRule which maximizes information gain,
// considering only a subset of Attributes.
//
// IMPORTANT: passing a zero-length consideredAttributes parameter will panic()
func (r *InformationGainRatioRuleGenerator) GetSplitRuleFromSelection(consideredAttributes []base.Attribute, f base.FixedDataGrid) *DecisionTreeRule {

	var selectedAttribute base.Attribute
	var selectedVal float64

	// Parameter check
	if len(consideredAttributes) == 0 {
		panic("More Attributes should be considered")
	}

	// Next step is to compute the information gain at this node
	// for each randomly chosen attribute, and pick the one
	// which maximises it
	maxRatio := math.Inf(-1)

	// Compute the base entropy
	classDist := base.GetClassDistribution(f)
	baseEntropy := getBaseEntropy(classDist)

	// Compute the information gain for each attribute
	for _, s := range consideredAttributes {
		var informationGain float64
		var localEntropy float64
		var splitVal float64
		if fAttr, ok := s.(*base.FloatAttribute); ok {
			localEntropy, splitVal = getNumericAttributeEntropy(f, fAttr)
		} else {
			proposedClassDist := base.GetClassDistributionAfterSplit(f, s)
			localEntropy = getSplitEntropy(proposedClassDist)
		}
		informationGain = baseEntropy - localEntropy
		informationGainRatio := informationGain / localEntropy
		if informationGainRatio > maxRatio {
			maxRatio = informationGainRatio
			selectedAttribute = s
			selectedVal = splitVal
		}
	}

	// Pick the one which maximises IG
	return &DecisionTreeRule{selectedAttribute, selectedVal}
}
