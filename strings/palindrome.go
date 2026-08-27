package main

import "fmt"

func main() {
	str := "ZeussueZ"

	result := isPalindrome(str)

	fmt.Println(result)
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
