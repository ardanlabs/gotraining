package binarysearch

import "fmt"

// 	binarySearchIterative it takes the list of the sorted numbers and check it with `iterative` process
//	to find the value and return the index of array or return the error if the value not found
//	the worst case of this algorithm is O(logn)
//	the best case of this algorithm is O(1)
//	this algorithm should be sorted
func binarySearchIterative(List []int, Target int) (int, error) {

	var low int           // 0 is the first index of array
	high := len(List) - 1 // to get last index of array

	// 	Check if low index is smaller or equal with high index
	// 	than continue the for loop
	for low <= high {

		// 	find the middle index of array
		mid := (low + high) / 2

		// 	check is List[mid] value equal with Target value than return the middle value
		if List[mid] == Target {
			return mid, nil
		}

		// 	If List[mid] is bigger than Target
		//	the high value should change to middle value minus by one
		if List[mid] > Target {
			high = mid - 1
		}

		// 	If List[mid] is smaller than Target
		//	the low value should change to middle value plus by one
		if List[mid] < Target {
			low = mid + 1
		}
	}

	// 	If the value not found in the array,
	//	I returned -1 and an error to show user value not found
	return -1, fmt.Errorf("sorry value not found")
}

// 	binarySearchRecursive it takes the list of the sorted numbers and check it with `recursive` process
//	to find the value and return the index of array or return the error if the value not found
func binarySearchRecursive(List []int, Target int, low int, high int) (int, error) {

	// 	find the middle index of array
	mid := (low + high) / 2

	// 	If low is greater than high it means the value not found,
	//	I returned -1 and an error to show user value not found
	for low > high {
		return -1, fmt.Errorf("sorry value not found")
	}

	// 	check is List[mid] value equal with Target value than return the middle value
	if List[mid] == Target {
		return mid, nil
	}

	// 	If List[mid] is bigger than Target
	//	I returned the binarySearchRecursive function with changing the high value to mid minus by one
	if List[mid] > Target {
		return binarySearchRecursive(List, Target, low, mid-1)
	}

	// 	If List[mid] is bigger than Target
	//	I returned the binarySearchRecursive function with changing the low value to mid plus by one
	if List[mid] < Target {
		return binarySearchRecursive(List, Target, mid+1, high)
	}

	// return -1 if user not found
	return -1, fmt.Errorf("sorry value not found")
}
