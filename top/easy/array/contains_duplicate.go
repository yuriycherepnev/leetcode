/*
Дан целочисленный массив nums.
Верните true, если какое-либо значение встречается в массиве как минимум дважды,
и false, если все элементы различны.

Input: nums = [1,2,3,1]
Output: true
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}

	duplicates := containsDuplicate(nums)

	fmt.Println(duplicates)
}

func containsDuplicate(nums []int) bool {
	keyNumbers := make(map[int]bool)

	for _, value := range nums {
		if keyNumbers[value] {
			return true
		}
		keyNumbers[value] = true
	}
	return false
}
