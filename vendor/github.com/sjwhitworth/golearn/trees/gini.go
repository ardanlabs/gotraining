package trees

import (
	"github.com/sjwhitworth/golearn/base"
	"math"
)

//
// Gini-coefficient rule generator
//

// GiniCoefficientRuleGenerator generates DecisionTreeRules which minimize
// the Geni impurity coefficient at each node.
type GiniCoefficientRuleGenerator struct {
}

// GenerateSplitRule returns the non-class Attribute-based DecisionTreeRule
// which maximises the information gain.
//
// IMPORTANT: passing a base.Instances with no Attributes other than the class
// variable will panic()
func (g *GiniCoefficientRuleGenerator) GenerateSplitRule(f base.FixedDataGrid) *DecisionTreeRule {

	attrs := f.AllAttributes()
	classAttrs := f.AllClassAttributes()
	candidates := base.AttributeDifferenceReferences(attrs, classAttrs)

	return g.GetSplitRuleFromSelection(candidates, f)
}

// GetSplitRuleFromSelection returns the DecisionTreeRule which maximises
// the information gain amongst consideredAttributes
//
// IMPORTANT: passing a zero-length consideredAttributes parameter will panic()
func (g *GiniCoefficientRuleGenerator) GetSplitRuleFromSelection(consideredAttributes []base.Attribute, f base.FixedDataGrid) *DecisionTreeRule {

	var selectedAttribute base.Attribute
	var selectedVal float64

	// Parameter check
	if len(consideredAttributes) == 0 {
		panic("More Attributes should be considered")
	}

	// Minimize the averagge Gini index
	minGini := math.Inf(1)
	for _, s := range consideredAttributes {
		var proposedDist map[string]map[string]int
		var splitVal float64
		if fAttr, ok := s.(*base.FloatAttribute); ok {
			_, splitVal = getNumericAttributeEntropy(f, fAttr)
			proposedDist = base.GetClassDistributionAfterThreshold(f, fAttr, splitVal)
		} else {
			proposedDist = base.GetClassDistributionAfterSplit(f, s)
		}
		avgGini := computeAverageGiniIndex(proposedDist)
		if avgGini < minGini {
			minGini = avgGini
			selectedAttribute = s
			selectedVal = splitVal
		}
	}

	return &DecisionTreeRule{selectedAttribute, selectedVal}
}

//
// Utility functions
//

// computeGini computes the Gini impurity measure
func computeGini(s map[string]int) float64 {
	// Compute probability map
	p := make(map[string]float64)
	for i := range s {
		if p[i] == 0 {
			continue
		}
		p[i] = 1.0 / float64(p[i])
	}
	// Compute overall sum
	sum := 0.0
	for i := range p {
		sum += p[i] * p[i]
	}

	return 1.0 - sum
}

// computeGiniImpurity computes the average Gini index of a
// proposed split
func computeAverageGiniIndex(s map[string]map[string]int) float64 {

	// Figure out the total number of things in this map
	total := 0
	for i := range s {
		for j := range s[i] {
			total += s[i][j]
		}
	}

	sum := 0.0
	for i := range s {
		subtotal := 0.0
		for j := range s[i] {
			subtotal += float64(s[i][j])
		}
		cf := subtotal / float64(total)
		cf *= computeGini(s[i])
		sum += cf
	}
	return sum
}
