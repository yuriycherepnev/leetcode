package main

import (
	"fmt"
)

func main() {
	str := "фывйцу"
	runes := []rune(str)

	fmt.Println(len(runes))

	count := 0
	for range str {
		count++
	}
	fmt.Println(count) // 6

	for _, value := range str {
		if 64 < value && value < 91 {
			value += 32
		}

	}
}

func isPalindrome(str string) {
	strLength := 0
	for range str {
		strLength++
	}
	for i, j := 0, strLength; i < j; i, j = i+1, j-1 {

	}

}
