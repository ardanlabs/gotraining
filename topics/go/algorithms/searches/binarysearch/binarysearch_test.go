package binarysearch

import (
	"testing"
)

//	generateSortList create a sort list numbers
func generateSortList(value int) (list []int) {
	if value <= 0 {
		return list
	}

	for i := 1; i < value; i++ {
		list = append(list, i)
	}

	return list
}

// 	TestBinarySearch a function to test binarySearchIterative and binarySearchRecursive algorithms
func TestBinarySearch(t *testing.T) {
	// generate list of 10000 numbers
	listData := generateSortList(10000)

	// 	list of data to check
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
		for i := range data {
			index, _ := binarySearchIterative(data[i].list, data[i].find)

			if index != data[i].expect {
				t.Errorf("expected %d, but got %d", data[i].expect, index)
			}
		}
	})

	t.Run("Binary Search Recursive", func(t *testing.T) {
		for i := range data {
			index, _ := binarySearchRecursive(data[i].list, data[i].find, 0, len(data[i].list)-1)

			if index != data[i].expect {
				t.Errorf("expected %d, but got %d", data[i].expect, index)
			}
		}
	})

}
