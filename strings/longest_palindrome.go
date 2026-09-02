package main

import "fmt"

func main() {
	str := "ghghart"
	l := longestPalindrome(str)
	fmt.Println(l)
}

func longestPalindrome(s string) int {
	charsMap := make(map[rune]int)
	oneChar := false

	for _, v := range s {
		charsMap[v]++
	}
	l := 0
	for _, v := range charsMap {
		if v%2 == 0 {
			l += v
		} else {
			l += v - 1
			if !oneChar {
				oneChar = true
				l += 1
			}
		}
	}
	return l
}
