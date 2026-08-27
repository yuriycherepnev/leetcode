package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	result := isAnagram(s, t)

	fmt.Println(result)
}

func isAnagram(s, t string) bool {
	sRunes := []rune(s)
	tRunes := []rune(t)

	if len(sRunes) != len(tRunes) {
		return false
	}

	count := make(map[rune]int)

	for i := range sRunes {
		count[sRunes[i]]++
		count[tRunes[i]]--
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}
