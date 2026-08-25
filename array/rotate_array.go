/*
Дан целочисленный массив nums; необходимо циклически сдвинуть
его вправо на k позиций, где k — неотрицательное число.
*/
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotateArray(nums, 9)
	fmt.Println(nums)
}

func rotateArray(nums []int, k int) {
	length := len(nums)
	k = k % length

	secondPart := nums[:length-k]
	firstPart := nums[length-k:]

	firstPart = append(firstPart, secondPart...)

	for index, value := range firstPart {
		nums[index] = value
	}
}

// эффективное по памяти решение
func reverseArray(nums []int, k int) {
	length := len(nums)
	k = k % length

	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
