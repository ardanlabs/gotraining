package stack_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/stack"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestPush vaidates the Push functionality.
func TestPush(t *testing.T) {
	t.Log("Given the need to test Push functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen pushing %d items", items)
		{
			var s stack.Stack

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData = name + orgData
				s.Push(&stack.Data{Name: name})
			}

			if s.Count() != items {
				t.Logf("\t%s\tShould be able to push %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", s.Count(), items)
			}
			t.Logf("\t%s\tShould be able to push %d items.", succeed, items)

			var data string
			f := func(d *stack.Data) error {
				data += d.Name
				return nil
			}
			if err := s.Operate(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate on the stack : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate on the stack.", succeed)

			if data != orgData {
				t.Logf("\t%s\tShould be able to traverse over %d items in FILO order.", failed, items)
				t.Fatalf("\t\tGot %s, Expected %s.", data, orgData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d items in FILO order.", succeed, items)
		}
	}
}

// TestPop vaidates the Pop functionality.
func TestPop(t *testing.T) {
	t.Log("Given the need to test Pop functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen popping %d items", items)
		{
			var s stack.Stack

			if _, err := s.Pop(); err == nil {
				t.Fatalf("\t%s\tShould not be able to pop an empty stack : %v", failed, err)
			}

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData = name + orgData
				s.Push(&stack.Data{Name: name})
			}

			if s.Count() != items {
				t.Logf("\t%s\tShould be able to push %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", s.Count(), items)
			}
			t.Logf("\t%s\tShould be able to push %d items.", succeed, items)

			var popData string
			for i := 0; i < items; i++ {
				data, err := s.Pop()
				if err != nil {
					t.Logf("\t%s\tShould be able to pop an item.", failed)
				}
				popData += data.Name
			}

			if s.Count() != 0 {
				t.Logf("\t%s\tShould be able to pop all %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected 0.", s.Count())
			}
			t.Logf("\t%s\tShould be able to pop all %d items.", succeed, items)

			if popData != orgData {
				t.Logf("\t%s\tShould be able to pop %d items in FILO order.", failed, items)
				t.Fatalf("\t\tGot %s, Expected %s.", popData, orgData)
			}
			t.Logf("\t%s\tShould be able to pop %d items in FILO order.", succeed, items)
		}
	}
}

// TestPeek vaidates the Peek functionality.
func TestPeek(t *testing.T) {
	t.Log("Given the need to test Peek functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen peeking %d items", items)
		{
			s := stack.Make(5)

			if _, err := s.Peek(0); err == nil {
				t.Fatalf("\t%s\tShould not be able to peek an empty stack : %s", failed, err)
			}

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData = name + orgData
				s.Push(&stack.Data{Name: name})
			}

			if s.Count() != items {
				t.Logf("\t%s\tShould be able to push %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", s.Count(), items)
			}
			t.Logf("\t%s\tShould be able to push %d items.", succeed, items)

			var popData string
			for i := 0; i < items; i++ {
				data, err := s.Peek(i)
				if err != nil {
					t.Logf("\t%s\tShould be able to peek an item.", failed)
				}
				popData += data.Name
			}

			if s.Count() != items {
				t.Logf("\t%s\tShould be able to pop all %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", s.Count(), items)
			}
			t.Logf("\t%s\tShould be able to pop all %d items.", succeed, items)

			if popData != orgData {
				t.Logf("\t%s\tShould be able to peek %d items in FILO order.", failed, items)
				t.Fatalf("\t\tGot %s, Expected %s.", popData, orgData)
			}
			t.Logf("\t%s\tShould be able to peek %d items in FILO order.", succeed, items)
		}
	}
}
