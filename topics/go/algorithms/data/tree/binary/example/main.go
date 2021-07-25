package main

import (
	"fmt"

	bst "github.com/ardanlabs/gotraining/topics/go/algorithms/data/tree/binary"
)

func main() {
	values := []int{65, 45, 35, 75, 85, 78, 95}

	var tree bst.Tree
	for _, value := range values {
		tree.Insert(value)
	}

	bst.PrettyPrint(tree)
	pre := tree.PreOrder()
	fmt.Println("Pre-order :", pre)
	in := tree.InOrder()
	fmt.Println("In-order  :", in)
	post := tree.PostOrder()
	fmt.Println("Post-order:", post)
}
