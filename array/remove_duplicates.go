/*
удалить дубликаты из массива сохранив порядок
*/
// input [0,0,1,1,1,2,2,3,3,4]
// output [0 1 2 3 4 2 3 3 4] or [0 1 2 3 4 _ _ _ _]
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 2, 2, 3, 3, 4}
	index := removeDuplicates(nums)
	fmt.Println(nums)
	cutNums := nums[:index:index]
	fmt.Println(len(cutNums), cap(cutNums))
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
