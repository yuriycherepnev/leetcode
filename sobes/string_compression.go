package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compress("AAAABBBCCDAA")) // A4B3C2DA2
	fmt.Println(compress("ABC"))          // ABC
	fmt.Println(compress("AAAA"))         // A4
	fmt.Println(compress("AABCC"))        // A2BC2
}

func compress(s string) string {
	var builder strings.Builder
	sRunes := []rune(s)
	buffer := make([]rune, 0, len(sRunes))

	for i := 0; i < len(sRunes); i++ {
		if i > 0 && sRunes[i] != sRunes[i-1] {
			builder.WriteString(string(sRunes[i-1]))
			if len(buffer) > 1 {
				builder.WriteString(strconv.Itoa(len(buffer)))
			}
			buffer = buffer[:0]
		}
		buffer = append(buffer, sRunes[i])
	}

	if len(buffer) > 0 {
		builder.WriteString(string(buffer[0]))
		if len(buffer) > 1 {
			builder.WriteString(strconv.Itoa(len(buffer)))
		}
		buffer = buffer[:0]
	}

	return builder.String()
}
