package trees

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"math/rand"
)

// RandomTreeRuleGenerator is used to generate decision rules for Random Trees
type RandomTreeRuleGenerator struct {
	Attributes   int
	internalRule InformationGainRuleGenerator
}

// GenerateSplitRule returns the best attribute out of those randomly chosen
// which maximises Information Gain
func (r *RandomTreeRuleGenerator) GenerateSplitRule(f base.FixedDataGrid) *DecisionTreeRule {

	var consideredAttributes []base.Attribute

	// First step is to generate the random attributes that we'll consider
	allAttributes := base.AttributeDifferenceReferences(f.AllAttributes(), f.AllClassAttributes())
	maximumAttribute := len(allAttributes)

	attrCounter := 0
	for {
		if len(consideredAttributes) >= r.Attributes {
			break
		}
		selectedAttrIndex := rand.Intn(maximumAttribute)
		selectedAttribute := allAttributes[selectedAttrIndex]
		matched := false
		for _, a := range consideredAttributes {
			if a.Equals(selectedAttribute) {
				matched = true
				break
			}
		}
		if matched {
			continue
		}
		consideredAttributes = append(consideredAttributes, selectedAttribute)
		attrCounter++
	}

	return r.internalRule.GetSplitRuleFromSelection(consideredAttributes, f)
}

// RandomTree builds a decision tree by considering a fixed number
// of randomly-chosen attributes at each node
type RandomTree struct {
	base.BaseClassifier
	Root *DecisionTreeNode
	Rule *RandomTreeRuleGenerator
}

// NewRandomTree returns a new RandomTree which considers attrs randomly
// chosen attributes at each node.
func NewRandomTree(attrs int) *RandomTree {
	return &RandomTree{
		base.BaseClassifier{},
		nil,
		&RandomTreeRuleGenerator{
			attrs,
			InformationGainRuleGenerator{},
		},
	}
}

// Fit builds a RandomTree suitable for prediction
func (rt *RandomTree) Fit(from base.FixedDataGrid) error {
	rt.Root = InferID3Tree(from, rt.Rule)
	return nil
}

// Predict returns a set of Instances containing predictions
func (rt *RandomTree) Predict(from base.FixedDataGrid) (base.FixedDataGrid, error) {
	return rt.Root.Predict(from)
}

// String returns a human-readable representation of this structure
func (rt *RandomTree) String() string {
	return fmt.Sprintf("RandomTree(%s)", rt.Root)
}

// Prune removes nodes from the tree which are detrimental
// to determining the accuracy of the test set (with)
func (rt *RandomTree) Prune(with base.FixedDataGrid) {
	rt.Root.Prune(with)
}

// Save outputs this model to a file
func (rt *RandomTree) Save(filePath string) error {
	writer, err := base.CreateSerializedClassifierStub(filePath, rt.GetMetadata())
	if err != nil {
		return err
	}
	defer func() {
		writer.Close()
	}()
	return rt.SaveWithPrefix(writer, "")
}

// SaveWithPrefix outputs this model to a file with a prefix.
func (rt *RandomTree) SaveWithPrefix(writer *base.ClassifierSerializer, prefix string) error {
	return rt.Root.SaveWithPrefix(writer, prefix)
}

// Load retrieves this model from a file
func (rt *RandomTree) Load(filePath string) error {
	reader, err := base.ReadSerializedClassifierStub(filePath)
	if err != nil {
		return err
	}
	return rt.LoadWithPrefix(reader, "")
}

// LoadWithPrefix retrives this random tree from disk with a given prefix.
func (rt *RandomTree) LoadWithPrefix(reader *base.ClassifierDeserializer, prefix string) error {
	rt.Root = &DecisionTreeNode{}
	return rt.Root.LoadWithPrefix(reader, prefix)
}

// GetMetadata returns required serialization metadata
func (rt *RandomTree) GetMetadata() base.ClassifierMetadataV1 {
	return base.ClassifierMetadataV1{
		FormatVersion:      1,
		ClassifierName:     "KNN",
		ClassifierVersion:  "1.0",
		ClassifierMetadata: nil,
	}
}
