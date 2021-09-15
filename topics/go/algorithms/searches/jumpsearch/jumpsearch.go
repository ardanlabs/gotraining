package jumpsearch

import (
	"math"
)
// jumpSearch is like binarysearch this algorithms need a sorted list
// the worst cast time complexity of this algorithm is O(√n)
// the best cast of this algorithm is O(1)
func jumpSearch(list []int, target int) int {
	var index int // index of array start from 0
	jump := int(math.Sqrt(float64(len(list)))) // to calculate jump we take sqrt of the length of array

	// if array is empty I returned -1
	if len(list) <= 0 {
		return -1
	}

	// if index is smaller or equal with the length of array continue the loop
	for index <= len(list) - 1{

		// if the value of list and target is equal,
		// the index of array will be returned
		if list[index] == target{
			return index
		}

		// list value is greater than target I'll break the for loop
		if list[index] > target{
			break
		}

		// add jump value to index
		index = index + jump
	}

	// at this part if first loop is break we continue
	// step back as linear search to see between the last index and previous index there is any value or not


	previous := index - jump // previous is for previous index
	// 	if the index is greater the length of array
	//	we set the index value to the length of array
	if index > len(list) -1 {
		index = len(list) - 1 // set length of array to the index
	}

	for list[index] >= target{
		// check if target is found or not
		if list[index] == target{
			return index
		}

		// if the index is equal to the previous jump index so it means value not found
		if index == previous{
			return -1 // if value not found return -1
		}

		// minus index value by one
		index--
	}


	// if value not found return -1
	return -1
}
