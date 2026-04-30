package main

import "fmt"

func main() {
	//digits := []int{1, 2, 3, 4}

	//plusOne(digits)

	number := 12345

	remainder := number % 10000

	number /= 10000

	fmt.Println(number)
	fmt.Println(remainder)
}

func plusOne(digits []int) int {
	length := len(digits)
	increase := 1
	number := 0

	for i := length - 1; i >= 0; i-- {
		number += digits[i] * increase
		increase *= 10
	}

	fmt.Println(increase)

	fmt.Println(number)

	return number
}
