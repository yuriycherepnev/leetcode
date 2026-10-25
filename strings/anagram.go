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

	mapRunes := make(map[rune]int)

	for i := range sRunes {
		mapRunes[sRunes[i]]++
		mapRunes[tRunes[i]]--
	}

	for _, value := range mapRunes {
		if value != 0 {
			return false
		}
	}

	return true
}
