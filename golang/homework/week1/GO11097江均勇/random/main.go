package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [20]int
	var sum int = 0
	var mul int64 = 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		arr[i] = rand.Intn(99) + 1
		fmt.Printf("%d ", arr[i])
		if i%2 == 1 {
			sum += arr[i]
		} else {
			mul *= int64(arr[i])
		}
	}
	fmt.Printf("\n")
	fmt.Printf("The sum is %d\n", sum)
	fmt.Printf("The mul is %d\n", mul)
}
