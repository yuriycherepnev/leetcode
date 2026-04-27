package main

import "fmt"

func main() {
	nums := []int{64, 45, 45, 56, 78, 34, 25, 12, 22, 11}

	fmt.Println(nums)

	shakerSort(nums)

	fmt.Println(nums)
}

func shakerSort(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
		for j := right; j > left; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		left++
	}

	//for left > right {
	//	for i := 0; i < right; i++ {
	//		if arr[i] < arr[i+1] {
	//			arr[i], arr[i+1] = arr[i+1], arr[i]
	//		}
	//	}
	//}
}
