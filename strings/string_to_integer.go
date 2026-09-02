package main

import (
	"fmt"
)

func main() {
	str := "123456"
	result := myAtoi(str)
	fmt.Println(result)
}

func myAtoi(s string) int {
	resultNumber := 0

	for _, value := range s {
		if value >= 48 && value <= 57 {
			digit := int(value - '0')
			resultNumber = resultNumber*10 + digit
		}
	}
	return resultNumber
}
