package main

import (
	"fmt"
)

func main() {
	var (
		numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		target  int   = 3
	)

	// result := BinarySearch(numbers, target)

	result := BinarySearchRecursion(numbers, target, 0, len(numbers)-1)

	if result == -1 {
		fmt.Printf("%v Number not found\n", target)
	} else {
		fmt.Printf("%v Number at %v", target, result)
	}
}

func BinarySearch(numbers []int, target int) int {
	var (
		left  int = 0
		right int = len(numbers) - 1
	)

	for right >= left {

		mid := (left + right) / 2

		if numbers[mid] == target {
			return mid
		}

		if numbers[mid] > target {
			right = mid - 1
			continue
		}

		left = mid + 1
	}

	return -1
}

func BinarySearchRecursion(numbers []int, target, left, right int) int {

	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if numbers[mid] == target {
		return mid
	}

	if numbers[mid] > target {
		return BinarySearchRecursion(numbers, target, left, mid-1)
	}

	return BinarySearchRecursion(numbers, target, mid+1, right)
}

/*

## sudo code

Start

	Get the numbers array and target value from the user

	Assign the left = 0 and right = len(number) - 1

	While right >= left

		mid = (right + left) / 2 (or) left + (right - left) / 2

		If (number[mid] == target)
			return mid
		If End

		If number[mid] > target
			right = mid - 1
			continue
		If End

		left = mid + 1

End

Time Complexity: O(log N)

Space Complexity: O(1)

Note:

	Imagine you have a dataset of one billion numbers. Linear search, in the worst case, could require a billion steps to find what you need – a slow and exhaustive journey.

	here comes binary search. It's like a superpower. Instead of a billion steps, it needs just around 32 steps max to find your target! That's lightning-fast and incredibly efficient.

*/
