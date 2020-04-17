package bst_test

import (
	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/bst"
	"testing"
)

var tree bst.BST

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

func TestInsert(t *testing.T) {
	createTree(&tree)
	tree.Insert(20)
}

func TestMax(t *testing.T) {
	if tree.Max() != 20 {
		t.Errorf("Max should be 10 but, got %d\n", tree.Max())
	}
}
