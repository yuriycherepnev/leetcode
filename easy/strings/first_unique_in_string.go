/*
Given a string s, find the first non-repeating character in it and return its index.
If it does not exist, return -1.
*/
package main

func main() {
	str := "leetcode"
	firstUniqChar(str)
}

/* решение через слайс */
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

/* решение через map */
/*
func firstUniqChar(s string) int {
    chars := make(map[int32]int)

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
*/
