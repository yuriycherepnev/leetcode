package main

import "fmt"

func main() {
	number := 900009
	result := isPalindrome(number)
	fmt.Println(result)
}

// сравнение всего числа
func isPalindrome(x int) bool {
	reversed := 0

	for x > reversed {
		remain := x % 10
		reversed = reversed*10 + remain
		x /= 10
	}

	return reversed == x
}

// сравнение половины числа
func isPalindromeTwo(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}
