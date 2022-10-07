package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("This is a random number: ", rand.Intn(10))
	fmt.Println("Sum: ", add(12, 13))
}
