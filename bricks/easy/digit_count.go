package main

import "fmt"

func main() {
	result := digitCount(6789)
	fmt.Println(result)
}

func digitCount(number int) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}
