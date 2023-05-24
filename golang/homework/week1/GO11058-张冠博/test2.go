package main

import (
	"math/rand"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r10 := rand.New(src)
	var sum int
	var product int = 1

	for i := 1; i <= 20; i++ {

		a := r10.Intn(100)
		if a == 0 {
			i--
			continue
		}
		if i%2 == 0 {
			println(a)
			product = product * a
		} else {
			println(a)
			sum = sum + a
		}
	}
	println(sum, product)
}
