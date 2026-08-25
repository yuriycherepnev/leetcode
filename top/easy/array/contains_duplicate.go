/*
Верните true, если есть дубли
и false, если нет.

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
	keyNumbers := make(map[int]struct{})

	for _, value := range nums {
		_, exist := keyNumbers[value]
		if exist {
			return true
		}
		keyNumbers[value] = struct{}{}
	}
	return false
}
