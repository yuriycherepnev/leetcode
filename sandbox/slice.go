package main

import "fmt"

func main() {
	sl := []int{1, 2, 3}
	change(sl)
	fmt.Println(sl)
}

func change(sl []int) {
	sl[0] = 10
}
