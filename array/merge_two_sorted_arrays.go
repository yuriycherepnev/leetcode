package main

import "fmt"

func main() {
	one := []int{1, 2, 3, 4, 5, 6}
	two := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	merged := merge(one, two)

	fmt.Println(merged)
}

func merge(nums1, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}

	return result
}
