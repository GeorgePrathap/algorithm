package main

import (
	"fmt"
)

func main() {
	var (
		numbers              []int = []int{1, 2, 77, 99, 50, 101, 200, 500}
		first, second, third       = numbers[0], 0, 0
	)

	for i := 1; i < len(numbers); i++ {

		if numbers[i] > first {
			third = second
			second = first
			first = numbers[i]
			continue
		}

		if numbers[i] > second {
			third = second
			second = numbers[i]
			continue
		}

		if numbers[i] > third {
			third = numbers[i]
		}
	}

	fmt.Printf("First : %v, Second : %v, Third : %v\n", first, second, third)
}

/*

## sudo code

	Get the numbers array and target from user

	Initialize first = numbers[0], second = 0, third = 0

	For i = 1 to len(Numbers) {

		if numbers[i] > first {
			third = second
			second = first
			first = numbers[i]
			continue
		}

		if numbers[i] > second {
			third = second
			second = numbers[i]
			continue
		}

		if (numbers[i] > third) {
			third = first
		}
	}

## Time complexity : O(n)

## Space complexity : O(1)

*/
