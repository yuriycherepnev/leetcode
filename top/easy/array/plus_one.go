/*
You are given a large integer represented as an integer array digits,
where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.
*/
package main

import "fmt"

func main() {
	digits := []int{9, 9, 9, 9}
	result := plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
