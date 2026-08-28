/*
найти вхождение строк
*/
package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "issip"

	number := strStr(haystack, needle)

	fmt.Println(number)
}

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
