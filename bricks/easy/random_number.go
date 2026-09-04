package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(100) 0 <= n < 100
// для диапазона 50 - 100
// rand.Intn(51) + 50

func main() {
	number := random()
	fmt.Println(number)
}

func random() int {
	number := rand.Intn(101) + 50
	return number
}
