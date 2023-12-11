package main

import (
	"log"
)

func main() {
	var (
		numbers []int = []int{1, 2, 3, 4, 5, 6, 8}
		length        = len(numbers)
		sum           = 0
	)

	log.Println("length :", length)

	sum = (length * (length + 1) / 2)

	for i := 0; i < length-1; i++ {
		sum -= numbers[i]
	}

	log.Printf("Missing Number is %v\n", sum)
}

/*

## sudo code

	Start

		Calculate the sum of the first N natural numbers as sumtotal= N*(N+1)/2.

		Traverse the array from start to end.

			Find the sum of all the array elements.

		Print the missing number as SumTotal – sum of array

		Note: In order to avoid integer overflow, pick one number from the range [1, N] and subtract a number from the given array (don’t subtract the same number twice). This way there won’t be any integer overflow.

*/
