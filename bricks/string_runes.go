package main

import "fmt"

func main() {
	str := "rfrfrf"
	runes := []rune(str)
	str2 := string(runes)

	fmt.Println(runes)
	fmt.Println(str2)
}
