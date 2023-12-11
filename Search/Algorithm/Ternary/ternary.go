package main

import (
	"fmt"
)

func main() {
	var (
		numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		target  int   = 3
	)

	// result := TernarySearch(numbers, target)
	result := TernarySearchRecursion(numbers, target, 0, len(numbers)-1)
	if result == -1 {
		fmt.Printf("%v Number not found\n", target)
	} else {
		fmt.Printf("%v Number at %v\n", target, result)
	}
}

func TernarySearch(numbers []int, target int) int {
	var (
		left, right = 0, len(numbers) - 1
	)

	for right >= left {
		middleIndexOne := left + (right-left)/3
		middleIndexTwo := right - (right-left)/3

		fmt.Printf("%v : %v\n", left, right)
		fmt.Printf("%v : %v\n", middleIndexOne, middleIndexTwo)
		fmt.Println("")

		if numbers[middleIndexOne] == target {
			return middleIndexOne
		}

		if numbers[middleIndexTwo] == target {
			return middleIndexTwo
		}

		if numbers[middleIndexOne] > target {
			right = middleIndexOne - 1
			continue
		}

		if numbers[middleIndexTwo] < target {
			left = middleIndexTwo + 1
			continue
		}

		left = middleIndexOne + 1

		right = middleIndexTwo - 1
	}

	return -1
}

func TernarySearchRecursion(numbers []int, target, left, right int) int {
	if left > right {
		return -1
	}

	midIndexOne := left + (right-left)/3
	midIndexTwo := right - (right-left)/3

	if numbers[midIndexOne] == target {
		return midIndexOne
	}

	if numbers[midIndexTwo] == target {
		return midIndexTwo
	}

	if numbers[midIndexOne] > target {
		return TernarySearchRecursion(numbers, target, left, midIndexOne-1)
	}

	if numbers[midIndexTwo] < target {
		return TernarySearchRecursion(numbers, target, midIndexTwo+1, right)
	}

	return TernarySearchRecursion(numbers, target, midIndexOne+1, midIndexTwo-1)
}

/*

## sudo code

Start

	Get the numbers arrays and target value from the user

	Set the left as 0 and right = len(numbers) - 1

	For right >= left

		Calculate midOne and MidTwo

		MidOne = left + (right - left) / 3

		MidTwo = right - (right - left) / 3

		If numbers[midOne] == target {
			return midOne
		}

		If numbers[midTwo] == target {
			return midOne
		}

		If numbers[midOne] > target {
			right = midOne - 1
			continue
		}

		If numbers[midTwo] < target {
			right = midTwo + 1
			continue
		}

		left = midOne + 1

		right midTwo - 1


	For End

	return -1

End

----

Time Complexity: O(log3n)

Space Complexity: O(1)

Note: The time complexity of the binary search is more than  the ternary search but it does not mean that ternary search is better. In reality, the number of comparisons in ternary search much more which makes it slower than binary search.

**/
