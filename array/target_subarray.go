/*
Необходимо найти минимальную длину подмассива,
сумма элементов которого больше или равна `target`.
*/

package main

import (
	"fmt"
)

func main() {
	target := 7
	tNums := []int{2, 3, 1, 2, 4, 3}
	result := targetSubArray(target, tNums)
	fmt.Println(result)
}

func targetSubArray(target int, nums []int) int {
	left := 0
	sum := 0
	minLength := len(nums)

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			length := right - left + 1
			if length < minLength {
				minLength = length
			}
			sum -= nums[left]
			left++
		}
	}

	return minLength
}
