/*

	This package implements decision trees.

	ID3DecisionTree:
		Builds a decision tree using the ID3 algorithm
			by picking the Attribute which maximises
			Information Gain at each node.

		Attributes must be CategoricalAttributes at
			present, so discretise beforehand (see
			filters)

	RandomTree:
		Builds a decision tree using the ID3 algorithm
			by picking the Attribute amongst those
			randomly selected that maximises Information
			Gain

		Attributes must be CategoricalAttributes at
			present, so discretise beforehand (see
			filters)

*/

package trees
