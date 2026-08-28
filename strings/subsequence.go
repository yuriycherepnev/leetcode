/*
s1 = "abcde"
s2 = "ace"
→ true

s1 = "abcde"
s2 = "aec"
→ false
*/

package main

func isSubsequence(s1, s2 string) bool {
	i := 0

	for j := 0; j < len(s1) && i < len(s2); j++ {
		if s1[j] == s2[i] {
			i++
		}
	}

	return i == len(s2)
}
