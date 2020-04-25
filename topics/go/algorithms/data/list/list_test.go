/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package list

	// Node represents the data being stored.
	type Node struct {
		Data string
		next *Node
		prev *Node
	}

	// List represents a list of nodes.
	type List struct {
		Count int
		first *Node
		last  *Node
	}

	// Add places a new node at the end of the list.
	func (l *List) Add(data string) *Node

	// AddFront places a new node at the front of the list.
	func (l *List) AddFront(data string) *Node

	// Find traverses the list looking for the specified data.
	func (l *List) Find(data string) (*Node, error)

	// FindReverse traverses the list in the opposite direction
	// looking for the specified data.
	func (l *List) FindReverse(data string) (*Node, error)

	// Remove traverses the list looking for the specified data
	// and if found, removes the node from the list.
	func (l *List) Remove(data string) (*Node, error)

	// Operate accepts a function that takes a node and calls
	// the specified function for every node found.
	func (l *List) Operate(f func(n *Node) error) error

	// OperateReverse accepts a function that takes a node and
	// calls the specified function for every node found.
	func (l *List) OperateReverse(f func(n *Node) error) error

	// AddSort adds a node based on lexical ordering.
	func (l *List) AddSort(data string) *Node
*/

package list_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/list"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestAdd validates the Add functionality.
func TestAdd(t *testing.T) {
	t.Log("Given the need to test Add functionality.")
	{
		const nodes = 5
		t.Logf("\tTest 0:\tWhen adding %d nodes", nodes)
		{
			var l list.List

			var orgNodeData string
			for i := 0; i < nodes; i++ {
				data := fmt.Sprintf("Node%d", i)
				orgNodeData += data
				l.Add(data)
			}

			if l.Count != nodes {
				t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", failed, nodes)
				t.Fatalf("\t\tTest 0:\tGot %d, Expected %d.", l.Count, nodes)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", succeed, nodes)

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.Operate(f); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to operate on the list.", succeed)

			if nodeData != orgNodeData {
				t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in order.", failed, nodes)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", nodeData, orgNodeData)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in order.", succeed, nodes)
		}
	}
}

// TestAddFront validates the AddFront functionality.
func TestAddFront(t *testing.T) {
	t.Log("Given the need to test AddFront functionality.")
	{
		const nodes = 5
		t.Logf("\tTest 0:\tWhen adding %d nodes", nodes)
		{
			var l list.List

			var orgNodeData string
			for i := 0; i < nodes; i++ {
				data := fmt.Sprintf("Node%d", i)
				orgNodeData += data
				l.AddFront(data)
			}

			if l.Count != nodes {
				t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", failed, nodes)
				t.Fatalf("\t\tTest 0:\tGot %d, Expected %d.", l.Count, nodes)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", succeed, nodes)

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.OperateReverse(f); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to operate on the list.", succeed)

			if nodeData != orgNodeData {
				t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in reverse order.", failed, nodes)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", nodeData, orgNodeData)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in reverse order.", succeed, nodes)
		}
	}
}

// TestFind validates the Find functionality.
func TestFind(t *testing.T) {
	t.Log("Given the need to test Find functionality.")
	{
		const nodes = 5
		t.Logf("\tTest 0:\tWhen adding %d nodes", nodes)
		{
			var l list.List

			var orgNodeData string
			for i := 0; i < nodes; i++ {
				data := fmt.Sprintf("Node%d", i)
				orgNodeData = data + orgNodeData
				l.AddFront(data)
			}

			data := "Node3"
			n, err := l.Find(data)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to call Find with no error : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to call Find with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tTest 0:\tShould be able to find %q : %v", failed, data, err)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to find %q.", succeed, data)
		}
	}
}

// TestFindReverse validates the FindReverse functionality.
func TestFindReverse(t *testing.T) {
	t.Log("Given the need to test FindReverse functionality.")
	{
		const nodes = 5
		t.Logf("\tTest 0:\tWhen adding %d nodes", nodes)
		{
			var l list.List

			var orgNodeData string
			for i := 0; i < nodes; i++ {
				data := fmt.Sprintf("Node%d", i)
				orgNodeData = data + orgNodeData
				l.AddFront(data)
			}

			data := "Node3"
			n, err := l.FindReverse(data)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to call FindReverse with no error : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to call FindReverse with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tTest 0:\tShould be able to find %q : %v", failed, data, err)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to find %q.", succeed, data)
		}
	}
}

// TestRemove validates the Remove functionality.
func TestRemove(t *testing.T) {
	t.Log("Given the need to test Remove functionality.")
	{
		const nodes = 5
		t.Logf("\tTest 0:\tWhen adding %d nodes", nodes)
		{
			var l list.List

			var orgNodeData string
			for i := 0; i < nodes; i++ {
				data := fmt.Sprintf("Node%d", i)
				orgNodeData = data + orgNodeData
				l.AddFront(data)
			}

			data := "Node3"
			n, err := l.Remove(data)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to call Remove with no error : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to call Remove with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tTest 0:\tShould be able to remove %q : %v", failed, data, err)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to remove %q.", succeed, data)

			n, err = l.Find(data)
			if err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould not be able to call Find without an error.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould not be able to call Find without an error.", succeed)
		}
	}
}

// TestAddSort validates the AddSort functionality.
func TestAddSort(t *testing.T) {
	t.Log("Given the need to test AddSort functionality.")
	{
		orgNodeData := []string{"grape", "apple", "plum", "mango", "kiwi"}
		t.Logf("\tTest 0:\tWhen adding %d nodes", len(orgNodeData))
		{
			var l list.List

			for _, data := range orgNodeData {
				l.AddSort(data)
			}

			if l.Count != len(orgNodeData) {
				t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", failed, len(orgNodeData))
				t.Fatalf("\t\tTest 0:\tGot %d, Expected %d.", l.Count, len(orgNodeData))
			}
			t.Logf("\t%s\tTest 0:\tShould be able to add %d nodes.", succeed, len(orgNodeData))

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.Operate(f); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to operate on the list.", succeed)

			sortedNodeData := "applegrapekiwimangoplum"
			if sortedNodeData != nodeData {
				t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in sort order.", failed, len(orgNodeData))
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", nodeData, sortedNodeData)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in sort order.", succeed, len(orgNodeData))

			nodeData = ""
			f = func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.OperateReverse(f); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to operate reverse on the list : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to operate reverse on the list.", succeed)

			sortedNodeData = "plummangokiwigrapeapple"
			if sortedNodeData != nodeData {
				t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in reverse sort order.", failed, len(orgNodeData))
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", nodeData, sortedNodeData)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d nodes in reverse sort order.", succeed, len(orgNodeData))
		}
	}
}
