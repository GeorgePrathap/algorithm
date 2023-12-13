package main

import (
	"log"
)

func main() {
	var (
		numbers = []int{2, 3, 5, 6}
		target  = 100
	)

	log.Println("Before in inserted : ", numbers)
	numbers = shiftNumberInSortedArray(numbers, target)
	log.Println("After in inserted : ", numbers)
}

func shiftNumber(numbers []int, target int) []int {
	for index, value := range numbers {

		if index == 0 && value > target {
			return append([]int{target}, numbers...)
		}

		if index == len(numbers)-1 && target > value {
			return append(numbers, target)
		}

		nextIndex := index + 1
		if nextIndex > len(numbers)-1 {
			nextIndex = len(numbers) - 1
		}

		if target > value && target < numbers[nextIndex] {
			shiftNumber := []int{}

			shiftNumber = append(shiftNumber, numbers[0:index+1]...)
			shiftNumber = append(shiftNumber, target)
			shiftNumber = append(shiftNumber, numbers[nextIndex:]...)
			return shiftNumber
		}
	}

	return []int{}
}

/*

## shiftNumber
## Time Complexity: O(n)
## Space Complexity: O(n)

*/

func shiftNumberInSortedArray(numbers []int, target int) []int {
	var i = len(numbers) - 1

	// instead of the below loop we can use the binary search
	for i >= 0 {
		if numbers[i] < target {
			break
		}
		i--
	}

	log.Println(i)

	numbers = append(numbers, 0)
	for j := len(numbers) - 2; j > i; j-- {
		numbers[j+1] = numbers[j]
	}

	numbers[i+1] = target

	return numbers
}

/*

## shiftNumber
## Time Complexity: O(n)
## Space Complexity: O(1)

*/
