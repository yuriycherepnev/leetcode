package main

func longestSubstring(s string) int {
	lastSeen := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if pos, ok := lastSeen[ch]; ok && pos >= left {
			left = pos + 1
		}

		lastSeen[ch] = right

		length := right - left + 1

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
