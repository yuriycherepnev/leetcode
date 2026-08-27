/*
вернуть индекс уникального символа или -1
*/
package main

import "fmt"

func main() {
	str := "leetcode"
	char := firstUniqChar(str)
	fmt.Println(char)
}

/* решение через map */
func firstUniqChar(s string) int {
	chars := make(map[rune]int)

	for _, value := range s {
		chars[value]++
	}

	for index, value := range s {
		if chars[value] == 1 {
			return index
		}
	}
	return -1
}

/* решение через слайс */
/*
func firstUniqChar(s string) int {
	chars := make([]int, 128)

	for _, ch := range s {
		chars[ch]++
	}

	for i, ch := range s {
		if chars[ch] == 1 {
			return i
		}
	}
	return -1
}
*/
