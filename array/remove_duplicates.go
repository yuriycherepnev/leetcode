/*
удалить дубликаты из массива сохранив порядок
*/
// input [0,0,1,1,1,2,2,3,3,4]
// output [0 1 2 3 4 2 3 3 4] or [0 1 2 3 4 _ _ _ _]
package main

import "fmt"

func main() {
	rNums := []int{0, 1, 2, 2, 2, 2, 3, 3, 4}
	index := removeDuplicates(rNums)
	cutNums := rNums[:index:index]
	fmt.Println(cutNums)
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
