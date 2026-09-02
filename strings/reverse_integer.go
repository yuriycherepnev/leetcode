/*
перевернуть число
*/
package main

import (
	"fmt"
)

func main() {
	number := 123456789
	newNumber := reverseInteger(number)
	fmt.Println(newNumber)
}

func reverseInteger(x int) int {
	reversed := 0
	for x != 0 {
		remainder := x % 10
		x /= 10
		reversed = reversed*10 + remainder
	}
	return reversed
}

/*
проверка на диапазон

const MaxInt = 2147483647
if reversed > MaxInt || reversed < -MaxInt {
			return 0
		}
*/
