package heap_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/heap"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestNew validates the New functionality.
func TestNew(t *testing.T) {
	t.Log("Given the need to test New functionality.")
	{
		t.Logf("\tTest 0:\tWhen creating a new heap with invalid capacity.")
		{
			var cap int
			_, err := heap.New(cap)
			if err == nil {
				t.Fatalf("\t%sTest 0:\tShould not be able to create a heap for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%sTest 0:\tShould not be able to create a heap for %d items.", succeed, cap)

			cap = -1
			_, err = heap.New(cap)
			if err == nil {
				t.Fatalf("\t%sTest 0:\tShould not be able to create a heap for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%sTest 0:\tShould not be able to create a heap for %d items.", succeed, cap)
		}
	}
}

// TestStore validates the Store functionality.
func TestStore(t *testing.T) {
	t.Log("Given the need to test Store functionality.")
	{
		const maxSize = 10
		t.Logf("\tTest 0:\tWhen we add %d items to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%sTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%sTest 0:\tShould be able to create a heap for %d items.", succeed, maxSize)

			for i := 0; i < maxSize; i++ {
				if err := h.Store(heap.Data{Value: i}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to store data %d in the heap : %v", failed, i, err)
				}
			}
			if h.Size() != maxSize {
				t.Logf("\t%sTest 0:\tShould be able to store %d items.", failed, maxSize)
				t.Fatalf("\tTest 0:\tGot %d, Expected %d.", h.Size(), maxSize)
			}

			t.Logf("\t%sTest 0:\tShould be able to store %d items.", succeed, maxSize)
		}
	}
}

// TestExtract validates extract functionality.
func TestExtract(t *testing.T) {
	t.Log("Given the need to test Extract functionality.")
	{
		const maxSize = 10
		t.Logf("\tTest 0:\tWhen we add %d items to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%sTest 0:\tShould be able to create a heap for %d items.", succeed, maxSize)

			for i := 0; i < maxSize; i++ {
				if err := h.Store(heap.Data{Value: i}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to store data %d in heap : %v", failed, i, err)
				}
				data, err := h.Extract()
				if err != nil {
					t.Fatalf("\t%sTest 0:\tShould be able to extract data %d from heap : %v", failed, i, err)
				}
				if data.Value != i {
					t.Fatalf("\t%sTest 0:\tShould return minimum value from heap.", failed)
				}
			}
			t.Logf("\t%sTest 0:\tShould return minimum value from heap.", succeed)
		}
	}
}

// TestRemove validates remove functionality.
func TestRemove(t *testing.T) {
	t.Log("Given the need to test Remove functionality.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a heap for %d items.", succeed, maxSize)

			d := heap.Data{
				Value: 50,
			}

			if err := h.Store(d); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in heap : %v", failed, &d, err)
			}

			if err := h.Remove(d); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to remove data %v from heap : %v", failed, &d, err)
			}

			if h.Size() != 0 {
				t.Fatalf("\t%s\tTest 0:\tShould decrease size after removing the data %v from heap : %v", failed, &d, err)
			}

			if data, err := h.Extract(); err == nil && data == nil {
				t.Fatalf("\t%s\tTest 0:\tShould have empty heap after removing the data %v from it.", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould have empty heap after removing the data from it.", succeed)
		}
	}
}

// TestGetRoot validates get root functionality.
func TestGetRoot(t *testing.T) {
	t.Log("Given the need to test GetRoot functionality.")
	{
		const maxSize = 2
		t.Logf("\tTest 0:\tWhen we add %d item to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a heap for %d items.", succeed, maxSize)

			d1 := heap.Data{
				Value: 100,
			}
			d2 := heap.Data{
				Value: 150,
			}

			if err := h.Store(d1); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in heap : %v", failed, &d1, err)
			}

			if err := h.Store(d2); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in heap : %v", failed, &d2, err)
			}

			root, err := h.GetRoot()
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to get root value from heap : %v", failed, err)
			}

			if root.Value != d1.Value {
				t.Fatalf("\t%s\tTest 0:\tShould be able to get appropriate root value from heap : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to get appropriate root value from heap.", succeed)

			if h.Size() != maxSize {
				t.Fatalf("\t%s\tTest 0:\tShould not change the size value: %v of the heap after get root value.", failed, h.Size())
			}
			t.Logf("\t%s\tTest 0:\tShould not change the size value of the heap after get root value.", succeed)
		}
	}
}

// TestGetRootEmpty validates get root functionality when heap is empty.
func TestGetRootEmpty(t *testing.T) {
	t.Log("Given the need to test GetRoot functionality when heap is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a heap.", succeed)

			if _, err := h.GetRoot(); err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty heap error when trying to get root from empty heap.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty heap error when trying to get root from empty heap.", succeed)
		}
	}
}

// TestRemoveEmpty validates remove functionality when heap is empty.
func TestRemoveEmpty(t *testing.T) {
	t.Log("Given the need to test Remove functionality when heap is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a heap.", succeed)

			if err := h.Remove(heap.Data{Value: 1}); err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty heap error when trying to remove from empty heap.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty heap error when trying to remove from empty heap.", succeed)
		}
	}
}

// TestExtractEmpty validates extract functionality when heap is empty.
func TestExtractEmpty(t *testing.T) {
	t.Log("Given the need to test Extract functionality when heap is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the heap.", maxSize)
		{
			h, err := heap.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a heap for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a heap.", succeed)

			if data, err := h.Extract(); err == nil  && data != nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty heap error when trying to extract from empty heap.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty heap error when trying to extract from empty heap.", succeed)
		}
	}
}