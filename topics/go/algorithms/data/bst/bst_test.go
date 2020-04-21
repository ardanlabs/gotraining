package bst_test

import (
	bst "github.com/ardanlabs/gotraining/topics/go/algorithms/data/bst"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMax(t *testing.T) {

	t.Log("Given the need to test the bst functionality")
	{
		testID := 0
		b := bst.New()
		t.Logf("\tTest %d:\tWhen getting max.", testID)
		{
			b.Insert(5)

			got, _ := b.Max()
			if got != 5 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct max", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot: %d, Expected: %d", testID, got, 5)
			}
		}
		testID = 1
		t.Logf("\tTest %d: \tWhen getting max after adding new int.", testID)
		{
			b.Insert(25)

			got, _ := b.Max()
			if got != 25 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct max", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot: %d, Expected: %d", testID, got, 5)
			}
		}
	}
}

func TestMin(t *testing.T) {

	t.Log("Given the need to test the bst functionality")
	{
		testID := 0
		b := bst.New()
		t.Logf("\tTest %d:\tWhen getting min.", testID)
		{
			b.Insert(50)
			b.Insert(30)

			got, _ := b.Min()
			if got != 30 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct min", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot: %d, Expected: %d", testID, got, 30)
			}
		}
		testID = 1
		t.Logf("\tTest %d: \tWhen getting max after adding new int.", testID)
		{
			b.Insert(25)

			got, _ := b.Min()
			if got != 25 {
				t.Errorf("\t%s\tTest %d:\tShould have the correct min", failed, testID)
				t.Fatalf("\t\tTest %d:\tGot: %d, Expected: %d", testID, got, 25)
			}
		}
	}
}
