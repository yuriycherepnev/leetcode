/*
перевернуть число
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
