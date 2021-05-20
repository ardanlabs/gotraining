package main

import (
	"fmt"

	bst "github.com/ardanlabs/gotraining/topics/go/algorithms/data/tree/binary"
)

func main() {
	values := []int{40, 5, 80, 2, 25, 65, 98}

	var tree bst.Tree
	for _, value := range values {
		tree.Insert(value)
	}
	bst.PrettyPrint(&tree)

	pre := bst.PreOrder(&tree)
	fmt.Println("Pre-order :", pre)

	in := bst.InOrder(&tree)
	fmt.Println("In-order  :", in)

	post := bst.PostOrder(&tree)
	fmt.Println("Post-order:", post)
}
