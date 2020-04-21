package bst_test

import (
	"fmt"
	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/bst"
	"testing"
)

var tree bst.BST

const succeed = "\u2713"
const failed = "\u2717"

func createTree(bst *bst.BST) {
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(9)
}

func TestMaxEmptyTree(t *testing.T) {

	// create empty tree.
	var emptyTree bst.BST

	maxTests := []struct {
		name     string
		input    bst.BST
		expected int
		err      error
	}{
		{"test empty tree", emptyTree, 0, fmt.Errorf("root node is nil")},
	}

	for _, tt := range maxTests {
		_, err := tt.input.Max()
		if err != tt.err {
			t.Logf("\t%s\tShould be able return error %s", failed, err)
			t.Fatalf("\t\tGot %v, Expected %v.", err, tt.err)
		}
		t.Logf("\t%s\tShould be able return error %s", succeed, err)
	}
}

func TestMax(t *testing.T) {

	// create tree with nodes.
	createTree(&tree)

	maxTests := []struct {
		name     string
		input    bst.BST
		expected int
		err      error
	}{
		{"non-empty tree", tree, 10, nil},
	}

	for _, tt := range maxTests {
		got, err := tt.input.Max()
		if got != tt.expected && err == nil {
			t.Logf("\t%s\tShould be able return max int: %d", failed, got)
			t.Fatalf("\t\tGot %v, Expected %v.", got, tt.expected)
		}
		t.Logf("\t%s\tShould be able return max int: %d", succeed, got)
	}
}
