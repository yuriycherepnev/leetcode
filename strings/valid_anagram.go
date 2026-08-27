/*
is valid anagram
*/
package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	result := isValidAnagram(s, t)

	fmt.Println(result)
}

/* решение через map */
/* работает с любыми символами unicode */
func isValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	stringMap := make(map[rune]int)
	for _, v := range s {
		stringMap[v]++
	}
	sLength := len(stringMap)
	for _, v := range t {
		stringMap[v]--
		if stringMap[v] == 0 {
			sLength--
		}
	}
	return sLength == 0
}

/* решение через slice и ascii */
/* не будет работать с кириллицей */
/*

func isValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	strings := make([]int, 26)

	for i := 0; i < len(s); i++ {
		strings[s[i]-'a']++
		strings[t[i]-'a']--
	}

	for _, value := range strings {
		if value != 0 {
			return false
		}
	}
	return true
}
*/
