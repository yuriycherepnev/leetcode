/*
Дан массив nums и число k.
Найти непрерывный подмассив длины k, у которого среднее арифметическое максимально.
*/

package main

import (
	"fmt"
)

func main() {
	wNums := []int{1, 12, -5, -6, 50, 3}
	result := slideWindow(wNums, 3)
	fmt.Println(result)
}

func slideWindow(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
