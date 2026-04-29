/*
Given an integer array nums sorted in non-decreasing order,
remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.

Consider the number of unique elements in nums to be k.
After removing duplicates, return the number of unique elements k.

The first k elements of nums should contain the unique numbers in sorted order.
The remaining elements beyond index k - 1 can be ignored.
*/
// input [0,0,1,1,1,2,2,3,3,4]
// output [0 1 2 3 4 2 3 3 4] or [0 1 2 3 4 _ _ _ _]
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	currentIndex := 0

	for _, value := range nums {
		if value != nums[currentIndex] {
			currentIndex++
			nums[currentIndex] = value
		}
	}

	return currentIndex + 1
}
