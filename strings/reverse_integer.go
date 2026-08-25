/*
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
package main

import (
	"fmt"
)

const MaxInt = 2147483647

func main() {
	number := -123456789
	newNumber := reverseInteger(number)
	fmt.Println(newNumber)
}

func reverseInteger(x int) int {
	reversed := 0
	for x != 0 {
		remainder := x % 10
		x /= 10
		reversed = reversed*10 + remainder
		if reversed > MaxInt || reversed < -MaxInt {
			return 0
		}
	}
	return reversed
}
