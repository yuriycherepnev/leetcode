package main

import "fmt"

func main() {
	str := "Zeus was deified, saw Suez."

	result := isPalindrome(str)

	fmt.Println(result)
}

func isPalindrome(s string) bool {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; {
		if 'A' <= runes[i] && runes[i] <= 'Z' {
			runes[i] += 32
		}
		if 'A' <= runes[j] && runes[j] <= 'Z' {
			runes[j] += 32
		}

		if !((runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9')) {
			i++
			continue
		}
		if !((runes[j] >= 'a' && runes[j] <= 'z') || (runes[j] >= '0' && runes[j] <= '9')) {
			j--
			continue
		}
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
