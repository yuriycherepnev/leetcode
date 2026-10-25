/*
Найти K самых часто встречающихся элементов
*/
package main

import "fmt"

func main() {
	fNums := []int{1, 1, 6, 6, 6, 6, 8, 8, 8, 8, 2, 2, 3, 3, 3}
	k := 3
	top := topKFrequent(fNums, k)
	fmt.Println(top)
}

// bucket sort
func topKFrequent(nums []int, k int) []int {
	mapNums := make(map[int]int)

	for _, number := range nums {
		mapNums[number]++
	}

	frequentNums := make([][]int, len(nums))

	for number, count := range mapNums {
		frequentNums[count] = append(frequentNums[count], number)
	}

	result := make([]int, 0, k)

	for i := len(frequentNums) - 1; i >= 0; i-- {
		for _, value := range frequentNums[i] {
			result = append(result, value)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
