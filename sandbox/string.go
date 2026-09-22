package main

import "fmt"

func main() {
	str := "раз два"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for _, value := range str {
		fmt.Println(value)
	}
}

func str3() {
	x := []string{"1", "2", "3"}

	fmt.Println(x[2:])
	fmt.Println(x[3:])
	fmt.Println(x[4:])
}
