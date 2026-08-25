/*
Нули переместить в конец,
сохраняя порядок ненулевых элементов
*/

package main

import "fmt"

func main() {
	digits := []int{2, 1, 0, 3, 0, 12, 9, 67}
	move := moveZeroes(digits)
	fmt.Println(move)
}

func moveZeroes(nums []int) []int {
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++
		}
	}

	return nums
}
