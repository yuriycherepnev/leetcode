/*
найти число которое встречается только 1 раз
*/
package main

import "fmt"

func main() {
	sNums := []int{2, 1, 4, 1, 2}
	single := singleNumber(sNums)
	fmt.Println(single)
}

func singleNumber(nums []int) int {
	numbers := make(map[int]struct{})
	result := 0

	for _, num := range nums {
		_, exist := numbers[num]
		if !exist {
			numbers[num] = struct{}{}
		} else {
			delete(numbers, num)
		}
	}

	for key := range numbers {
		result = key
	}
	return result
}

// оба варианта решения верны
// решение через XOR
func singleNumberXOR(nums []int) int {
	number := 0

	for _, value := range nums {
		number = number ^ value
	}

	return number
}
