package main

import (
	"fmt"
	"sort"
)

var (
	words = []string{
		"banana",
		"apple",
		"orange",
		"grape",
	}
	nums = []int{8, 6, 34, 6, 2, 3, 9}
)

func main() {
	sort.Strings(words)
	fmt.Println(words)
	sort.Ints(nums)
	fmt.Println(nums)
}
