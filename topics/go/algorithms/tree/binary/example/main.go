package main

import (
	"fmt"

	bst "github.com/ardanlabs/gotraining/topics/go/algorithms/tree/binary"
)

func main() {
	values := []int{40, 5, 80, 2, 25, 65, 98}

	var tree bst.Tree
	for _, value := range values {
		tree.Insert(value)
	}
	bst.PrettyPrint(&tree)

	in := bst.InOrder(&tree)
	fmt.Println("In-order  :", in)
	pre := bst.PreOrder(&tree)
	fmt.Println("Pre-order :", pre)
	post := bst.PostOrder(&tree)
	fmt.Println("Post-order:", post)
}
