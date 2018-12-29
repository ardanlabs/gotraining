package queue_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/queue"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestNew vaidates the New functionality.
func TestNew(t *testing.T) {
	t.Log("Given the need to test New functionality.")
	{
		t.Logf("\tTest 0:\tWhen creating a new queue with invalid capacity.")
		{
			var cap int
			_, err := queue.New(cap)
			if err == nil {
				t.Fatalf("\t%s\tShould not be able to create a queue for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%s\tShould not be able to create a queue for %d items.", succeed, cap)

			cap = -1
			_, err = queue.New(cap)
			if err == nil {
				t.Fatalf("\t%s\tShould not be able to create a queue for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%s\tShould not be able to create a queue for %d items.", succeed, cap)
		}
	}
}

// TestEnqueue vaidates the Enqueue functionality.
func TestEnqueue(t *testing.T) {
	t.Log("Given the need to test Enqueue functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen enqueing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tShould be able to create a queue for %d items.", succeed, items)

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData += name
				if err := q.Enqueue(&queue.Data{Name: name}); err != nil {
					t.Fatalf("\t%s\tShould be able to enqueue item %d in the queue : %v", failed, i, err)
				}
			}

			if q.Count != items {
				t.Logf("\t%s\tShould be able to enqueue %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tShould be able to enqueue %d items.", succeed, items)

			var data string
			f := func(d *queue.Data) error {
				data += d.Name
				return nil
			}
			if err := q.Operate(f); err != nil {
				t.Fatalf("\t%s\tShould be able to operate on the queue : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to operate on the queue.", succeed)

			if data != orgData {
				t.Logf("\t%s\tShould be able to traverse over %d items in FIFO order.", failed, items)
				t.Fatalf("\t\tGot %s, Expected %s.", data, orgData)
			}
			t.Logf("\t%s\tShould be able to traverse over %d items in FIFO order.", succeed, items)
		}
	}
}

// TestDequeue vaidates the Dequeue functionality.
func TestDequeue(t *testing.T) {
	t.Log("Given the need to test Dequeue functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen dequeing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tShould be able to create a queue for %d items.", succeed, items)

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData += name
				if err := q.Enqueue(&queue.Data{Name: name}); err != nil {
					t.Fatalf("\t%s\tShould be able to enqueue item %d in the queue : %v", failed, i+1, err)
				}
			}

			if q.Count != items {
				t.Logf("\t%s\tShould be able to enqueue %d items.", failed, items)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tShould be able to enqueue %d items.", succeed, items)

			var data string
			for i := 0; i < items; i++ {
				d, err := q.Dequeue()
				if err != nil {
					t.Fatalf("\t%s\tShould be able to dequeue an item from the queue : %d, %v", failed, i+1, err)
				}
				data += d.Name
			}

			if data != orgData {
				t.Logf("\t%s\tShould be able to dequeue over %d items in FIFO order.", failed, items)
				t.Fatalf("\t\tGot %s, Expected %s.", data, orgData)
			}
			t.Logf("\t%s\tShould be able to dequeue over %d items in FIFO order.", succeed, items)
		}
	}
}

// TestEnqueueFull vaidates the Enqueue functionality when full.
func TestEnqueueFull(t *testing.T) {
	t.Log("Given the need to test Enqueue functionality for being full.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen enqueing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tShould be able to create a queue for %d items.", succeed, items)

			for i := 0; i < items; i++ {
				if err := q.Enqueue(&queue.Data{Name: "test"}); err != nil {
					t.Fatalf("\t%s\tShould be able to enqueue item %d in the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to enqueue %d items in the queue.", succeed, items)

			if q.Count != items {
				t.Logf("\t%s\tShould be able to see queue is full.", failed)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tShould be able to see queue is full.", succeed)

			if err := q.Enqueue(&queue.Data{Name: "test"}); err == nil {
				t.Fatalf("\t%s\tShould not be able to enqueue another item in the queue.", failed)
			}
			t.Logf("\t%s\tShould not be able to enqueue another item in the queue.", succeed)

			for i := 0; i < items-1; i++ {
				if _, err := q.Dequeue(); err != nil {
					t.Fatalf("\t%s\tShould be able to dequeue item %d from the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to dequeue %d items from the queue.", succeed, items-1)

			for i := 0; i < items-1; i++ {
				if err := q.Enqueue(&queue.Data{Name: "test"}); err != nil {
					t.Fatalf("\t%s\tShould be able to enqueue item %d back in the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to enqueue %d items back in the queue.", succeed, items-1)

			if q.Count != items {
				t.Logf("\t%s\tShould be able to see queue is full.", failed)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tShould be able to see queue is full.", succeed)

			if err := q.Enqueue(&queue.Data{Name: "test"}); err == nil {
				t.Fatalf("\t%s\tShould not be able to enqueue another item in the queue.", failed)
			}
			t.Logf("\t%s\tShould not be able to enqueue another item in the queue.", succeed)
		}
	}
}

// TestDequeueEmpty vaidates the Dequeue functionality when empty.
func TestDequeueEmpty(t *testing.T) {
	t.Log("Given the need to test Dequeue functionality for being empty.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen enqueing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tShould be able to create a queue for %d items.", succeed, items)

			for i := 0; i < items; i++ {
				if err := q.Enqueue(&queue.Data{Name: "test"}); err != nil {
					t.Fatalf("\t%s\tShould be able to enqueue item %d in the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to enqueue %d items in the queue.", succeed, items)

			if q.Count != items {
				t.Logf("\t%s\tShould be able to see queue is full.", failed)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tShould be able to see queue is full.", succeed)

			for i := 0; i < items-1; i++ {
				if _, err := q.Dequeue(); err != nil {
					t.Fatalf("\t%s\tShould be able to dequeue item %d from the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to dequeue %d items from the queue.", succeed, items-1)

			if q.Count != 1 {
				t.Logf("\t%s\tShould be able to see queue has 1 item.", failed)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, 1)
			}
			t.Logf("\t%s\tShould be able to see queue has 1 item.", succeed)

			if err := q.Enqueue(&queue.Data{Name: "test"}); err != nil {
				t.Fatalf("\t%s\tShould be able to enqueue another item in the queue.", failed)
			}
			t.Logf("\t%s\tShould be able to enqueue another item in the queue.", succeed)

			for i := 0; i < items-3; i++ {
				if _, err := q.Dequeue(); err != nil {
					t.Fatalf("\t%s\tShould be able to dequeue item %d from in the queue : %v", failed, i+1, err)
				}
			}
			t.Logf("\t%s\tShould be able to dequeue %d items from the queue.", succeed, items-3)

			if q.Count != 0 {
				t.Logf("\t%s\tShould be able to see queue is empty.", failed)
				t.Fatalf("\t\tGot %d, Expected %d.", q.Count, 0)
			}
			t.Logf("\t%s\tShould be able to see queue is empty.", succeed)

			if _, err := q.Dequeue(); err == nil {
				t.Fatalf("\t%s\tShould not be able to dequeue another item from the queue.", failed)
			}
			t.Logf("\t%s\tShould not be able to dequeue another item from the queue.", succeed)
		}
	}
}
