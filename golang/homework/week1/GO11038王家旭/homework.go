package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 第一题
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		for k := 1; k <= i; k++ {
			fmt.Printf("%d*%d=%2d ", k, i, k*i)
			fmt.Print(" ")
		}
		fmt.Println()
	}

	// 第二题
	num := rand.New(rand.NewSource(time.Now().UnixNano()))
	var (
		enenNum = 1
		oddNum  int
	)
	for i := 1; i < 21; i++ {
		if i%2 != 1 {
			enenNum *= num.Intn(100)
		} else {
			oddNum += num.Intn(100)
		}

	}
	fmt.Println(enenNum, oddNum)

	// 第三题

	fmt.Println(0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89)

}
