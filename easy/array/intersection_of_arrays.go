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
	nums2 := []int{2, 4, 6, 8, 8, 9, 10, 1000, 1002, 1004, 2000}

	result := intersect(nums1, nums2)

	fmt.Println(result)
}

func intersect(nums1 []int, nums2 []int) []int {
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
