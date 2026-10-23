/*
Верните true, если есть дубли
и false, если нет.

Input: nums = [1,2,3,1]
Output: true
*/

package main

import "fmt"

func main() {
	dNums := []int{1, 2, 3, 1}

	duplicates := containsDuplicate(dNums)

	fmt.Println(duplicates)
}

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]struct{})
	for _, v := range nums {
		_, exist := duplicates[v]
		if exist {
			return true
		}
		duplicates[v] = struct{}{}
	}
	return false
}
