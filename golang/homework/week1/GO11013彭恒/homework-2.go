package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var oddNumberSum int = 0
	var evenNumberProduct int = 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < 21; i++ {
		number := r.Intn(100)
		if number == 0 {
			i = i - 1
			continue
		} else if i == 20 {
			fmt.Printf("%d:%d\n", i, number)
		} else {
			fmt.Printf("%d:%d ", i, number)
		}
		if i&1 == 0 {
			evenNumberProduct = evenNumberProduct * number
		} else {
			oddNumberSum = oddNumberSum + number
		}
	}
	fmt.Printf("sum of odd numbers: %d\n", oddNumberSum)
	fmt.Printf("product of even numbers: %d\n", evenNumberProduct)
}
