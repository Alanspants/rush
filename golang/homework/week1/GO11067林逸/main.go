package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("第一题：")
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			a := i * j
			fmt.Printf("%d*%d=%d\t", i, j, a)
		}
		fmt.Println("\r")
	}

	fmt.Println("\n第二题：")
	var sum int = 0
	var mul int = 1
	// 函数有问题，注意编译器提示
	r1 := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for a := 1; a <= 20; a++ {
		w := r1.Intn(100)
		fmt.Printf("%d ", w)
		if a%2 == 0 {
			mul *= w
		} else {
			sum += w
		}
	}
	fmt.Printf("\n%d %d\n", sum, mul)

	fmt.Println("\n第三题:")
	fmt.Println("1")
	for a, b := 0, 1; ; {
		c := a + b
		a = b
		b = c
		if c > 100 {
			break
		} else {
			fmt.Println(c)
			continue
		}
	}

}
