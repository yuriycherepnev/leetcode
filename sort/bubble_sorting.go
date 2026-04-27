package main

func main() {
	nums := []int{64, 45, 45, 56, 78, 34, 25, 12, 22, 11}
	bubbleSort(nums)
}

func bubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	n := len(arr) - 1
	swapped := false

	for n > 0 {
		for i := 0; i < n; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n--
	}
}
