package list_test

import (
	"fmt"
	"testing"

	list "github.com/ardanlabs/gotraining/topics/go/algorithms/data/list/student"
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
				t.Logf("\t%s\tShould be able to add %d nodes.", failed, nodes)
				t.Fatalf("\t\tGot %d, Expected %d.", l.Count, nodes)
			}
			t.Logf("\t%s\tShould be able to add %d nodes.", succeed, nodes)

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.Operate(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate on the list.", succeed)

			if nodeData != orgNodeData {
				t.Logf("\t%s\tShould be able to traverse over %d nodes in order.", failed, nodes)
				t.Fatalf("\t\tGot %s, Expected %s.", nodeData, orgNodeData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d nodes in order.", succeed, nodes)
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
				t.Logf("\t%s\tShould be able to add %d nodes.", failed, nodes)
				t.Fatalf("\t\tGot %d, Expected %d.", l.Count, nodes)
			}
			t.Logf("\t%s\tShould be able to add %d nodes.", succeed, nodes)

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.OperateReverse(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate on the list.", succeed)

			if nodeData != orgNodeData {
				t.Logf("\t%s\tShould be able to traverse over %d nodes in reverse order.", failed, nodes)
				t.Fatalf("\t\tGot %s, Expected %s.", nodeData, orgNodeData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d nodes in reverse order.", succeed, nodes)
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
				t.Fatalf("\t%s\tShould be able to call Find with no error : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to call Find with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tShould be able to find %q : %v", failed, data, err)
				t.Fatalf("\t\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tShould be able to find %q.", succeed, data)
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
				t.Fatalf("\t%s\tShould be able to call FindReverse with no error : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to call FindReverse with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tShould be able to find %q : %v", failed, data, err)
				t.Fatalf("\t\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tShould be able to find %q.", succeed, data)
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
				t.Fatalf("\t%s\tShould be able to call Remove with no error : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to call Remove with no error.", succeed)

			if n.Data != data {
				t.Logf("\t%s\tShould be able to remove %q : %v", failed, data, err)
				t.Fatalf("\t\tGot %s, Expected %s.", n.Data, data)
			}
			t.Logf("\t%s\tShould be able to remove %q.", succeed, data)

			n, err = l.Find(data)
			if err == nil {
				t.Fatalf("\t%s\tShould not be able to call Find without an error.", failed)
			}
			t.Logf("\t%s\tShould not be able to call Find without an error.", succeed)
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
				t.Logf("\t%s\tShould be able to add %d nodes.", failed, len(orgNodeData))
				t.Fatalf("\t\tGot %d, Expected %d.", l.Count, len(orgNodeData))
			}
			t.Logf("\t%s\tShould be able to add %d nodes.", succeed, len(orgNodeData))

			var nodeData string
			f := func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.Operate(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate on the list : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate on the list.", succeed)

			sortedNodeData := "applegrapekiwimangoplum"
			if sortedNodeData != nodeData {
				t.Logf("\t%s\tShould be able to traverse over %d nodes in sort order.", failed, len(orgNodeData))
				t.Fatalf("\t\tGot %s, Expected %s.", nodeData, sortedNodeData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d nodes in sort order.", succeed, len(orgNodeData))

			nodeData = ""
			f = func(n *list.Node) error {
				nodeData += n.Data
				return nil
			}
			if err := l.OperateReverse(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate reverse on the list : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate reverse on the list.", succeed)

			sortedNodeData = "plummangokiwigrapeapple"
			if sortedNodeData != nodeData {
				t.Logf("\t%s\tShould be able to traverse over %d nodes in reverse sort order.", failed, len(orgNodeData))
				t.Fatalf("\t\tGot %s, Expected %s.", nodeData, sortedNodeData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d nodes in reverse sort order.", succeed, len(orgNodeData))
		}
	}
}
