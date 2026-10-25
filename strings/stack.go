package main

import "fmt"

func main() {
	str := "leet**cod*e"
	result := removeStars(str)
	fmt.Println(result)
}

func removeStars(s string) string {
	sRunes := []rune(s)
	stack := make([]rune, 0, len(sRunes))

	for _, value := range sRunes {
		if value == '*' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, value)
	}

	return string(stack)
}
