package binarysearch

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// generateSortList create a sorted list of number.
func generateSortList(value int) (list []int) {
	if value <= 0 {
		return list
	}

	for i := 1; i < value; i++ {
		list = append(list, i)
	}

	return list
}

// TestBinarySearch a function to test binarySearchIterative and binarySearchRecursive algorithms.
func TestBinarySearch(t *testing.T) {
	// generate a list of numbers between 1 and 10000.
	listData := generateSortList(10000)

	data := []struct {
		list   []int
		find   int
		expect int
	}{
		{listData, 50, 49},
		{listData, 159, 158},
		{listData, 1, 0},
		{listData, 500, 499},
		{listData, 9999, 9998},
		{generateSortList(-1), -999, -1},
	}

	t.Run("Binary Search Iterative", func(t *testing.T) {
		t.Log("Start the testing binary search in iterative way.")
		{
			for i := range data {

				index, _ := binarySearchIterative(data[i].list, data[i].find)

				if index != data[i].expect {
					t.Fatalf("\t%s\tExpected %d, but got %d", failed, data[i].expect, index)
				}
				t.Logf("\t%s\tEverything is looks fine, test %d", succeed, i)
			}
		}
	})

	t.Run("Binary Search Recursive", func(t *testing.T) {
		t.Log("Start the testing binary search in recursive way.")
		{
			for i := range data {
				index, _ := binarySearchRecursive(data[i].list, data[i].find, 0, len(data[i].list)-1)

				if index != data[i].expect {
					t.Fatalf("\t%s\tExpected %d, but got %d", failed, data[i].expect, index)
				}
				t.Logf("\t%s\tEverything is looks fine, test %d", succeed, i)
			}
		}
	})

}
