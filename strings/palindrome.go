package main

import "fmt"

func main() {
	str := "ВeussueВ"
	result := isPalindrome(str)
	fmt.Println(result)
}

func isPalindrome(s string) bool {
	sRunes := []rune(s)
	left := 0
	right := len(sRunes) - 1

	for left < right {
		if sRunes[left] != sRunes[right] {
			return false
		}
		left++
		right--
	}

	return true
}
