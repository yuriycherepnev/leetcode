/*
найти вхождение строк
*/
package main

import "fmt"

func main() {
	haystack := "mississipаpi"
	needle := "issipа"
	number := strStr(haystack, needle)
	fmt.Println(number)
}

func strStr(haystack string, needle string) int {
	needLen := len(needle)
	for i := 0; i <= len(haystack)-needLen; i++ {
		if haystack[i:i+needLen] == needle {
			return i
		}
	}
	return -1
}
