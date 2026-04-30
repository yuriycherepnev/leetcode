package main

import "fmt"

func main() {
	digits := []int{0, 1, 0, 3, 12}

	moveZeroes(digits)
}

func moveZeroes(nums []int) {
	zeroIndex := 0
	notZeroIndex := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroIndex = i
		}
		if nums[i] != 0 {
			notZeroIndex = i
		}
		fmt.Println(zeroIndex, notZeroIndex)

	}
}
