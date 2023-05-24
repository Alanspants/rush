package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	r100 := rand.New(rand.NewSource(time.Now().UnixNano()))
	var c = 0
	var d = 1

	for i := 1; i < 21; i++ {
		a := r100.Intn(99) + 1 

		fmt.Printf("%d,%d\n", i, a)
		if i&1 == 1 {

			c += a

			fmt.Println("第单数个累加求和:", c)
		} else {
			d *= a
			fmt.Println("第偶数个累乘求积:", d)
		}

	}

}
