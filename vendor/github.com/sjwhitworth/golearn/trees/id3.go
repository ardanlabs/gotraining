package trees

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"sort"
)

// NodeType determines whether a DecisionTreeNode is a leaf or not.
type NodeType int

const (
	// LeafNode means there are no children
	LeafNode NodeType = 1
	// RuleNode means we should look at the next attribute value
	RuleNode NodeType = 2
)

// RuleGenerator implementations analyse instances and determine
// the best value to split on.
type RuleGenerator interface {
	GenerateSplitRule(base.FixedDataGrid) *DecisionTreeRule
}

// DecisionTreeRule represents the "decision" in "decision tree".
type DecisionTreeRule struct {
	SplitAttr base.Attribute `json:"split_attribute"`
	SplitVal  float64        `json:"split_val"`
}

func (d *DecisionTreeRule) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})
	if d.SplitAttr == nil {
		ret["split_attribute"] = "unknown"
		ret["split_val"] = 0
		return json.Marshal(ret)
	}
	marshaledSplitAttrRaw, err := d.SplitAttr.MarshalJSON()
	if err != nil {
		return nil, err
	}
	marshaledSplitAttr := make(map[string]interface{})
	err = json.Unmarshal(marshaledSplitAttrRaw, &marshaledSplitAttr)
	if err != nil {
		panic(err)
	}
	ret["split_attribute"] = marshaledSplitAttr
	ret["split_val"] = d.SplitVal
	return json.Marshal(ret)
}

func (d *DecisionTreeRule) unmarshalJSON(data []byte) error {

	var jsonMap map[string]interface{}
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		return err
	}
	if splitVal, ok := jsonMap["split_val"]; ok {
		d.SplitVal = splitVal.(float64)
	}
	split := jsonMap["split_attribute"]
	splitBytes, err := json.Marshal(split)
	if err != nil {
		panic(err)
	}
	if string(splitBytes) != "\"unknown\"" {
		d.SplitAttr, err = base.DeserializeAttribute(splitBytes)
		if err != nil {
			return err
		}
		if d.SplitAttr == nil {
			panic("Should not be nil")
			return fmt.Errorf("base.DeserializeAttribute returned nil")
		}
	} else {
		d.SplitAttr = nil
	}
	return nil
}

func (d *DecisionTreeRule) UnmarshalJSON(data []byte) error {
	ret := d.unmarshalJSON(data)
	return ret
}

// String prints a human-readable summary of this thing.
func (d *DecisionTreeRule) String() string {

	if d.SplitAttr == nil {
		return fmt.Sprintf("INVALID:DecisionTreeRule(SplitAttr is nil)")
	}

	if _, ok := d.SplitAttr.(*base.FloatAttribute); ok {
		return fmt.Sprintf("DecisionTreeRule(%s <= %f)", d.SplitAttr.GetName(), d.SplitVal)
	}
	return fmt.Sprintf("DecisionTreeRule(%s)", d.SplitAttr.GetName())
}

// DecisionTreeNode represents a given portion of a decision tree.
type DecisionTreeNode struct {
	Type      NodeType                     `json:"node_type"`
	Children  map[string]*DecisionTreeNode `json:"children"`
	ClassDist map[string]int               `json:"class_dist"`
	Class     string                       `json:"class_string"`
	ClassAttr base.Attribute               `json:"-"`
	SplitRule *DecisionTreeRule            `json:"decision_tree_rule"`
}

func getClassAttr(from base.FixedDataGrid) base.Attribute {
	allClassAttrs := from.AllClassAttributes()
	return allClassAttrs[0]
}

// Save sends the classification tree to an output file
func (d *DecisionTreeNode) Save(filePath string) error {
	metadata := base.ClassifierMetadataV1{
		FormatVersion:     1,
		ClassifierName:    "test",
		ClassifierVersion: "1",
	}
	serializer, err := base.CreateSerializedClassifierStub(filePath, metadata)
	if err != nil {
		return err
	}
	err = d.SaveWithPrefix(serializer, "")
	if err != nil {
		return err
	}
	return serializer.Close()
}

func (d *DecisionTreeNode) SaveWithPrefix(writer *base.ClassifierSerializer, prefix string) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = writer.WriteBytesForKey(writer.Prefix(prefix, "tree"), b)
	if err != nil {
		return err
	}

	err = writer.WriteAttributeForKey(writer.Prefix(prefix, "treeClassAttr"), d.ClassAttr)
	if err != nil {
		return err
	}

	return nil
}

// Load reads from the classifier from an output file
func (d *DecisionTreeNode) Load(filePath string) error {

	reader, err := base.ReadSerializedClassifierStub(filePath)
	if err != nil {
		return err
	}

	return d.LoadWithPrefix(reader, "")
}

// LoadWithPrefix reads from the classifier from part of another model
func (d *DecisionTreeNode) LoadWithPrefix(reader *base.ClassifierDeserializer, prefix string) error {

	b, err := reader.GetBytesForKey(reader.Prefix(prefix, "tree"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, d)
	if err != nil {
		return err
	}

	a, err := reader.GetAttributeForKey(reader.Prefix(prefix, "treeClassAttr"))
	if err != nil {
		return err
	}

	d.ClassAttr = a

	return nil
}

// InferID3Tree builds a decision tree using a RuleGenerator
// from a set of Instances (implements the ID3 algorithm)
func InferID3Tree(from base.FixedDataGrid, with RuleGenerator) *DecisionTreeNode {
	// Count the number of classes at this node
	classes := base.GetClassDistribution(from)
	classAttr := getClassAttr(from)
	// If there's only one class, return a DecisionTreeLeaf with
	// the only class available
	if len(classes) == 1 {
		maxClass := ""
		for i := range classes {
			maxClass = i
		}
		ret := &DecisionTreeNode{
			LeafNode,
			nil,
			classes,
			maxClass,
			classAttr,
			&DecisionTreeRule{nil, 0.0},
		}
		return ret
	}

	// Only have the class attribute
	maxVal := 0
	maxClass := ""
	for i := range classes {
		if classes[i] > maxVal {
			maxClass = i
			maxVal = classes[i]
		}
	}

	// If there are no more Attributes left to split on,
	// return a DecisionTreeLeaf with the majority class
	cols, _ := from.Size()
	if cols == 2 {
		ret := &DecisionTreeNode{
			LeafNode,
			nil,
			classes,
			maxClass,
			classAttr,
			&DecisionTreeRule{nil, 0.0},
		}
		return ret
	}

	// Generate a return structure
	ret := &DecisionTreeNode{
		RuleNode,
		nil,
		classes,
		maxClass,
		classAttr,
		nil,
	}

	// Generate the splitting rule
	splitRule := with.GenerateSplitRule(from)
	if splitRule == nil || splitRule.SplitAttr == nil {
		// Can't determine, just return what we have
		return ret
	}

	// Split the attributes based on this attribute's value
	var splitInstances map[string]base.FixedDataGrid
	if _, ok := splitRule.SplitAttr.(*base.FloatAttribute); ok {
		splitInstances = base.DecomposeOnNumericAttributeThreshold(from,
			splitRule.SplitAttr, splitRule.SplitVal)
	} else {
		splitInstances = base.DecomposeOnAttributeValues(from, splitRule.SplitAttr)
	}
	// Create new children from these attributes
	ret.Children = make(map[string]*DecisionTreeNode)
	for k := range splitInstances {
		newInstances := splitInstances[k]
		ret.Children[k] = InferID3Tree(newInstances, with)
	}
	ret.SplitRule = splitRule
	return ret
}

// getNestedString returns the contents of node d
// prefixed by level number of tags (also prints children)
func (d *DecisionTreeNode) getNestedString(level int) string {
	buf := bytes.NewBuffer(nil)
	tmp := bytes.NewBuffer(nil)
	for i := 0; i < level; i++ {
		tmp.WriteString("\t")
	}
	buf.WriteString(tmp.String())
	if d.Children == nil {
		buf.WriteString(fmt.Sprintf("Leaf(%s)", d.Class))
	} else {
		var keys []string
		buf.WriteString(fmt.Sprintf("Rule(%s)", d.SplitRule))
		for k := range d.Children {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			buf.WriteString("\n")
			buf.WriteString(tmp.String())
			buf.WriteString("\t")
			buf.WriteString(k)
			buf.WriteString("\n")
			buf.WriteString(d.Children[k].getNestedString(level + 1))
		}
	}
	return buf.String()
}

// String returns a human-readable representation of a given node
// and it's children
func (d *DecisionTreeNode) String() string {
	return d.getNestedString(0)
}

// computeAccuracy is a helper method for Prune()
func computeAccuracy(predictions base.FixedDataGrid, from base.FixedDataGrid) float64 {
	cf, _ := evaluation.GetConfusionMatrix(from, predictions)
	return evaluation.GetAccuracy(cf)
}

// Prune eliminates branches which hurt accuracy
func (d *DecisionTreeNode) Prune(using base.FixedDataGrid) {
	// If you're a leaf, you're already pruned
	if d.Children == nil {
		return
	}
	if d.SplitRule == nil {
		return
	}

	// Recursively prune children of this node
	sub := base.DecomposeOnAttributeValues(using, d.SplitRule.SplitAttr)
	for k := range d.Children {
		if sub[k] == nil {
			continue
		}
		subH, subV := sub[k].Size()
		if subH == 0 || subV == 0 {
			continue
		}
		d.Children[k].Prune(sub[k])
	}

	// Get a baseline accuracy
	predictions, _ := d.Predict(using)
	baselineAccuracy := computeAccuracy(predictions, using)

	// Speculatively remove the children and re-evaluate
	tmpChildren := d.Children
	d.Children = nil

	predictions, _ = d.Predict(using)
	newAccuracy := computeAccuracy(predictions, using)

	// Keep the children removed if better, else restore
	if newAccuracy < baselineAccuracy {
		d.Children = tmpChildren
	}
}

// Predict outputs a base.Instances containing predictions from this tree
func (d *DecisionTreeNode) Predict(what base.FixedDataGrid) (base.FixedDataGrid, error) {
	predictions := base.GeneratePredictionVector(what)
	classAttr := getClassAttr(predictions)
	classAttrSpec, err := predictions.GetAttribute(classAttr)
	if err != nil {
		panic(err)
	}
	predAttrs := base.AttributeDifferenceReferences(what.AllAttributes(), predictions.AllClassAttributes())
	predAttrSpecs := base.ResolveAttributes(what, predAttrs)
	what.MapOverRows(predAttrSpecs, func(row [][]byte, rowNo int) (bool, error) {
		cur := d
		for {
			if cur.Children == nil {
				predictions.Set(classAttrSpec, rowNo, classAttr.GetSysValFromString(cur.Class))
				break
			} else {
				splitVal := cur.SplitRule.SplitVal
				at := cur.SplitRule.SplitAttr
				ats, err := what.GetAttribute(at)
				if err != nil {
					//predictions.Set(classAttrSpec, rowNo, classAttr.GetSysValFromString(cur.Class))
					//break
					panic(err)
				}

				var classVar string
				if _, ok := ats.GetAttribute().(*base.FloatAttribute); ok {
					// If it's a numeric Attribute (e.g. FloatAttribute) check that
					// the value of the current node is greater than the old one
					classVal := base.UnpackBytesToFloat(what.Get(ats, rowNo))
					if classVal > splitVal {
						classVar = "1"
					} else {
						classVar = "0"
					}
				} else {
					classVar = ats.GetAttribute().GetStringFromSysVal(what.Get(ats, rowNo))
				}
				if next, ok := cur.Children[classVar]; ok {
					cur = next
				} else {
					// Suspicious of this
					var bestChild string
					for c := range cur.Children {
						bestChild = c
						if c > classVar {
							break
						}
					}
					cur = cur.Children[bestChild]
				}
			}
		}
		return true, nil
	})
	return predictions, nil
}

type ClassProba struct {
	Probability float64
	ClassValue  string
}

type ClassesProba []ClassProba

func (o ClassesProba) Len() int {
	return len(o)
}
func (o ClassesProba) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
func (o ClassesProba) Less(i, j int) bool {
	return o[i].Probability > o[j].Probability
}

// Predict class probabilities of the input samples what, returns a sorted array (by probability) of classes, and another array representing it's probabilities
func (t *ID3DecisionTree) PredictProba(what base.FixedDataGrid) (ClassesProba, error) {
	d := t.Root
	predictions := base.GeneratePredictionVector(what)
	predAttrs := base.AttributeDifferenceReferences(what.AllAttributes(), predictions.AllClassAttributes())
	predAttrSpecs := base.ResolveAttributes(what, predAttrs)

	_, rowCount := what.Size()
	if rowCount > 1 {
		panic("PredictProba supports only 1 row predictions")
	}
	var results ClassesProba
	what.MapOverRows(predAttrSpecs, func(row [][]byte, rowNo int) (bool, error) {
		cur := d
		for {
			if cur.Children == nil {
				totalDist := 0
				for _, dist := range cur.ClassDist {
					totalDist += dist
				}
				for class, dist := range cur.ClassDist {
					classProba := ClassProba{ClassValue: class, Probability: float64(float64(dist) / float64(totalDist))}
					results = append(results, classProba)
				}
				sort.Sort(results)
				break
			} else {
				splitVal := cur.SplitRule.SplitVal
				at := cur.SplitRule.SplitAttr
				ats, err := what.GetAttribute(at)
				if err != nil {
					//predictions.Set(classAttrSpec, rowNo, classAttr.GetSysValFromString(cur.Class))
					//break
					panic(err)
				}

				var classVar string
				if _, ok := ats.GetAttribute().(*base.FloatAttribute); ok {
					// If it's a numeric Attribute (e.g. FloatAttribute) check that
					// the value of the current node is greater than the old one
					classVal := base.UnpackBytesToFloat(what.Get(ats, rowNo))
					if classVal > splitVal {
						classVar = "1"
					} else {
						classVar = "0"
					}
				} else {
					classVar = ats.GetAttribute().GetStringFromSysVal(what.Get(ats, rowNo))
				}
				if next, ok := cur.Children[classVar]; ok {
					cur = next
				} else {
					// Suspicious of this
					var bestChild string
					for c := range cur.Children {
						bestChild = c
						if c > classVar {
							break
						}
					}
					cur = cur.Children[bestChild]
				}
			}
		}
		return true, nil
	})
	return results, nil
}

//
// ID3 Tree type
//

// ID3DecisionTree represents an ID3-based decision tree
// using the Information Gain metric to select which attributes
// to split on at each node.
type ID3DecisionTree struct {
	base.BaseClassifier
	Root       *DecisionTreeNode
	PruneSplit float64
	Rule       RuleGenerator
}

// NewID3DecisionTree returns a new ID3DecisionTree with the specified test-prune
// ratio and InformationGain as the rule generator.
// If the ratio is less than 0.001, the tree isn't pruned.
func NewID3DecisionTree(prune float64) *ID3DecisionTree {
	return &ID3DecisionTree{
		base.BaseClassifier{},
		nil,
		prune,
		new(InformationGainRuleGenerator),
	}
}

// NewID3DecisionTreeFromRule returns a new ID3DecisionTree with the specified test-prun
// ratio and the given rule gnereator.
func NewID3DecisionTreeFromRule(prune float64, rule RuleGenerator) *ID3DecisionTree {
	return &ID3DecisionTree{
		base.BaseClassifier{},
		nil,
		prune,
		rule,
	}
}

// Fit builds the ID3 decision tree
func (t *ID3DecisionTree) Fit(on base.FixedDataGrid) error {
	if t.PruneSplit > 0.001 {
		trainData, testData := base.InstancesTrainTestSplit(on, t.PruneSplit)
		t.Root = InferID3Tree(trainData, t.Rule)
		t.Root.Prune(testData)
	} else {
		t.Root = InferID3Tree(on, t.Rule)
	}
	return nil
}

// Predict outputs predictions from the ID3 decision tree
func (t *ID3DecisionTree) Predict(what base.FixedDataGrid) (base.FixedDataGrid, error) {
	return t.Root.Predict(what)
}

// String returns a human-readable version of this ID3 tree
func (t *ID3DecisionTree) String() string {
	return fmt.Sprintf("ID3DecisionTree(%s\n)", t.Root)
}

func (t *ID3DecisionTree) GetMetadata() base.ClassifierMetadataV1 {
	return base.ClassifierMetadataV1{
		FormatVersion:      1,
		ClassifierName:     "ID3",
		ClassifierVersion:  "1.0",
		ClassifierMetadata: nil,
	}
}

func (t *ID3DecisionTree) Save(filePath string) error {
	writer, err := base.CreateSerializedClassifierStub(filePath, t.GetMetadata())
	if err != nil {
		return err
	}
	fmt.Printf("writer: %v", writer)
	return t.SaveWithPrefix(writer, "")
}

func (t *ID3DecisionTree) SaveWithPrefix(writer *base.ClassifierSerializer, prefix string) error {
	return t.Root.SaveWithPrefix(writer, prefix)
}

func (t *ID3DecisionTree) Load(filePath string) error {
	reader, err := base.ReadSerializedClassifierStub(filePath)
	if err != nil {
		return err
	}
	return t.LoadWithPrefix(reader, "")
}

func (t *ID3DecisionTree) LoadWithPrefix(reader *base.ClassifierDeserializer, prefix string) error {
	t.Root = &DecisionTreeNode{}
	return t.Root.LoadWithPrefix(reader, "")
}
