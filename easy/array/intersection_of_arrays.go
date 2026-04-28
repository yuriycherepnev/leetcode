/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays
and you may return the result in any order.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{5, 2, 8, 1, 9, 3, 8}
	nums2 := []int{4, 6, 8, 8, 9, 10, 1000, 1002, 1004, 2000, 2}

	result := mapIntersect(nums1, nums2)

	fmt.Println(result)
}

/*
решение через мапы
*/
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

/*
решение через сортировку слайсов
*/
func sortIntersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	intSlice := make([]int, 0)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			intSlice = append(intSlice, nums1[i])
			i++
			j++
		} else {
			if nums1[i] > nums2[j] {
				j++
			} else {
				i++
			}
		}
	}

	return intSlice
}
