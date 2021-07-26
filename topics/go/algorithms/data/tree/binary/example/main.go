package main

import (
	"fmt"

	bst "github.com/ardanlabs/gotraining/topics/go/algorithms/data/tree/binary"
)

func main() {
	values := []bst.Data{
		{Key: 65, Name: "Bill"},
		{Key: 45, Name: "Ale"},
		{Key: 35, Name: "Joan"},
		{Key: 75, Name: "Hanna"},
		{Key: 85, Name: "John"},
		{Key: 78, Name: "Steph"},
		{Key: 95, Name: "Sally"},
	}

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

	fmt.Print("\n")
	d35, err := tree.Find(35)
	if err != nil {
		fmt.Println("ERROR: Unable to find 35")
		return
	}
	fmt.Println("found:", d35)

	d78, err := tree.Find(78)
	if err != nil {
		fmt.Println("ERROR: Unable to find 78")
		return
	}
	fmt.Println("found:", d78)

	d3, err := tree.Find(3)
	if err == nil {
		fmt.Println("ERROR: found 3", d3)
		return
	}
	fmt.Println("not-found: 3")

	fmt.Print("\n")
	tree.Delete(75)
	bst.PrettyPrint(tree)

	tree.Delete(85)
	bst.PrettyPrint(tree)
}
