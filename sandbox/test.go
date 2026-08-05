package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 2, 5)

	add(sl)

	fmt.Println(sl)
}

func add(sl []int) {
	sl[1] = 1
	sl = append(sl, []int{1, 2, 3}...)
}
