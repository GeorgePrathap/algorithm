package main

import (
	"fmt"
)

func main() {
	var (
		numbers []int = []int{1, 2, 3, 4, 5, 6}
		target  int   = 2
	)

	// var result = LinearSearch(numbers, target)
	var result = LinearSearchRecursion(numbers, target, 0)
	if result == -1 {
		fmt.Printf("%v Number not found\n", target)
	} else {
		fmt.Printf("%v Number at %v Index\n", target, result)
	}
}

func LinearSearch(numbers []int, target int) int {

	for index, number := range numbers {
		if number == target {
			return index
		}
	}

	return -1
}

func LinearSearchRecursion(numbers []int, target, index int) int {
	if len(numbers) == index {
		return -1
	}

	if numbers[index] == target {
		return index
	}

	return LinearSearchRecursion(numbers, target, index+1)
}

/*

## sudo code

Start

	Get the numbers arrays and target value from the user

	For i to numbers
		If numbers[i] == target
			return target
		If End
	For End

	return -1

End

----

Time Complexity: O(n)

Space Complexity: O(1)

**/
