package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 1, 2}
	single := singleNumber(nums)
	fmt.Println(single)
}

// оба варианта решения верны
// решение через XOR
func singleNumber(nums []int) int {
	number := 0

	for _, value := range nums {
		number = number ^ value
	}

	return number
}

// решение через map
/*
func singleNumber(nums []int) int {
	numbers := make(map[int]struct{})
	result := 0

	for _, num := range nums {
		_, ok := numbers[num]
		if !ok {
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
*/
