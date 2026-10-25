/*
перевернуть строку
*/
package main

import "fmt"

func main() {
	str := "Hello"
	reverse := reverseString(str)
	fmt.Println(reverse)
}

func reverseString(s string) string {
	sRunes := []rune(s)
	left := 0
	right := len(s) - 1

	for left < right {
		sRunes[left], sRunes[right] = sRunes[right], sRunes[left]
		left++
		right--
	}

	return string(sRunes)
}
