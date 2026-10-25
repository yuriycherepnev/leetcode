/*
массив цифр это число,
нужно увеличить это число на 1 и вернуть
*/
package main

import "fmt"

func main() {
	digits := []int{9, 9, 7, 9}
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

// если нужна конвертация к инту
func convertToInt(digits []int) int {
	number := 0
	for _, v := range digits {
		number = number*10 + v
	}
	return number
}
