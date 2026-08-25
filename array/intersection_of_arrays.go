/*
найти пересечение массивов
*/
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{5, 2, 8, 1, 9, 3, 8}
	nums2 := []int{4, 6, 8, 8, 9, 10, 1000, 1002, 1004, 2000, 2}

	result := mapIntersect(nums1, nums2)

	fmt.Println(result)
}

func mapIntersect(nums1 []int, nums2 []int) []int {
	mapNums := make(map[int]int)

	for _, v := range nums1 {
		mapNums[v]++
	}

	result := make([]int, 0, len(nums1))
	for _, v := range nums2 {
		if mapNums[v] > 0 {
			mapNums[v]--
			result = append(result, v)
		}
	}
	return result
}
