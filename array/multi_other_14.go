/*
вернуть массив где на каждой i-ой позиции будет произведение всех элементов кроме i-го
[1, 2, 3] => [23, 13, 1*2] => [6, 3, 2]
*/
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	result := multiOther(numbers)
	fmt.Println(result)
}

func multiOther(nums []int) []int {
	result := make([]int, len(nums))

	left := 1

	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
