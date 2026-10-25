/*
вернуть 2 числа сумма которых равна target
*/

package main

import "fmt"

func main() {
	sNums := []int{3, 1, 4, 6, 2}
	sum := twoSum(sNums, 7)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)

	for _, val := range nums {
		_, exist := numbers[val]
		remainder := target - val
		if exist {
			return []int{val, numbers[val]}
		}
		numbers[remainder] = val
	}
	return nil
}
