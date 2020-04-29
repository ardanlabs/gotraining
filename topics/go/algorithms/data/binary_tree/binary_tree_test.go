package binary_tree_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/binary_tree"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestNew validates the New functionality.
func TestNew(t *testing.T) {
	t.Log("Given the need to test New functionality.")
	{
		t.Logf("\tTest 0:\tWhen creating a new binary_tree with invalid capacity.")
		{
			var cap int
			_, err := binary_tree.New(cap)
			if err == nil {
				t.Fatalf("\t%sTest 0:\tShould not be able to create a binary_tree for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%sTest 0:\tShould not be able to create a binary_tree for %d items.", succeed, cap)

			cap = -1
			_, err = binary_tree.New(cap)
			if err == nil {
				t.Fatalf("\t%sTest 0:\tShould not be able to create a binary_tree for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%sTest 0:\tShould not be able to create a binary_tree for %d items.", succeed, cap)
		}
	}
}

// TestStore validates the Store functionality.
func TestStore(t *testing.T) {
	t.Log("Given the need to test Store functionality.")
	{
		const maxSize = 10
		t.Logf("\tTest 0:\tWhen we add %d items to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%sTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%sTest 0:\tShould be able to create a binary_tree for %d items.", succeed, maxSize)

			for i := 0; i < maxSize; i++ {
				if err := b.Store(&binary_tree.Data{Value: i}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to store data %d in the binary_tree : %v", failed, i, err)
				}
			}
			if b.Size() != maxSize {
				t.Logf("\t%sTest 0:\tShould be able to store %d items.", failed, maxSize)
				t.Fatalf("\tTest 0:\tGot %d, Expected %d.", b.Size(), maxSize)
			}

			t.Logf("\t%s\tTest 0:\tShould be able to store %d items.", succeed, maxSize)
		}
	}
}

// TestExtract validates extract functionality.
func TestExtract(t *testing.T) {
	t.Log("Given the need to test Extract functionality.")
	{
		const maxSize = 10
		t.Logf("\tTest 0:\tWhen we add %d items to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%sTest 0:\tShould be able to create a binary_tree for %d items.", succeed, maxSize)

			for i := 0; i < maxSize; i++ {
				if err := b.Store(&binary_tree.Data{Value: i}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to store data %d in binary_tree : %v", failed, i, err)
				}
				data, err := b.Extract()
				if err != nil {
					t.Fatalf("\t%sTest 0:\tShould be able to extract data %d from binary_tree : %v", failed, i, err)
				}
				if data.Value != i {
					t.Fatalf("\t%sTest 0:\tShould return minimum value from binary_tree.", failed)
				}
			}
			t.Logf("\t%s\tTest 0:\tShould return minimum value from binary_tree.", succeed)
		}
	}
}

// TestRemove validates remove functionality.
func TestRemove(t *testing.T) {
	t.Log("Given the need to test Remove functionality.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items.", succeed, maxSize)

			d := binary_tree.Data{
				Value: 50,
			}

			if err := b.Store(&d); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in binary_tree : %v", failed, &d, err)
			}

			if err := b.Remove(&d); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to remove data %v from binary_tree : %v", failed, &d, err)
			}

			if b.Size() != 0 {
				t.Fatalf("\t%s\tTest 0:\tShould decrease size after removing the data %v from binary_tree : %v", failed, &d, err)
			}

			if data, err := b.Extract(); err == nil && data == nil {
				t.Fatalf("\t%s\tTest 0:\tShould have empty binary_tree after removing the data %v from it.", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould have empty binary_tree after removing the data from it.", succeed)
		}
	}
}

// TestGetRoot validates get root functionality.
func TestGetRoot(t *testing.T) {
	t.Log("Given the need to test GetRoot functionality.")
	{
		const maxSize = 2
		t.Logf("\tTest 0:\tWhen we add %d item to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items.", succeed, maxSize)

			d1 := binary_tree.Data{
				Value: 100,
			}
			d2 := binary_tree.Data{
				Value: 150,
			}

			if err := b.Store(&d1); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in binary_tree : %v", failed, &d1, err)
			}

			if err := b.Store(&d2); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to store data %v in binary_tree : %v", failed, &d2, err)
			}

			root, err := b.GetRoot()
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to get root value from binary_tree : %v", failed, err)
			}

			if root.Value != d1.Value {
				t.Fatalf("\t%s\tTest 0:\tShould be able to get appropriate root value from binary_tree : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to get appropriate root value from binary_tree.", succeed)

			if b.Size() != maxSize {
				t.Fatalf("\t%s\tTest 0:\tShould not change the size value: %v of the binary_tree after get root value.", failed, b.Size())
			}
			t.Logf("\t%s\tTest 0:\tShould not change the size value of the binary_tree after get root value.", succeed)
		}
	}
}

// TestGetRootEmpty validates get root functionality when binary_tree is empty.
func TestGetRootEmpty(t *testing.T) {
	t.Log("Given the need to test GetRoot functionality when binary_tree is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a binary_tree.", succeed)

			if _, err := b.GetRoot(); err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to get root from empty binary_tree.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to get root from empty binary_tree.", succeed)
		}
	}
}

// TestRemoveEmpty validates remove functionality when binary_tree is empty.
func TestRemoveEmpty(t *testing.T) {
	t.Log("Given the need to test Remove functionality when binary_tree is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a binary_tree.", succeed)

			if err := b.Remove(&binary_tree.Data{Value: 1}); err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to remove from empty binary_tree.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to remove from empty binary_tree.", succeed)
		}
	}
}

// TestExtractEmpty validates extract functionality when binary_tree is empty.
func TestExtractEmpty(t *testing.T) {
	t.Log("Given the need to test Extract functionality when binary_tree is empty.")
	{
		const maxSize = 1
		t.Logf("\tTest 0:\tWhen we add %d item to store in the binary_tree.", maxSize)
		{
			b, err := binary_tree.New(maxSize)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a binary_tree for %d items : %v", failed, maxSize, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a binary_tree.", succeed)

			if data, err := b.Extract(); err == nil  && data != nil {
				t.Fatalf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to extract from empty binary_tree.", failed)
			}
			t.Logf("\t%s\tTest 0:\tShould get empty binary_tree error when trying to extract from empty binary_tree.", succeed)
		}
	}
}