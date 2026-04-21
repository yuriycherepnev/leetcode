package main

func main() {
	str := "leetcode"
	firstUniqChar(str)
}

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
