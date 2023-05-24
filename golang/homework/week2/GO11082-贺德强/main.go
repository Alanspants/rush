package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// the first homework
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, j*i)
		}
		fmt.Println()
	}

	sum := 0
	multi := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		number := r.Intn(100) + 1

		if i%2 == 0 {
			sum += number
		} else {
			multi *= number
		}
	}
	fmt.Printf("和：%d,乘积：%d \n", sum, multi)

	// 打印100以内的斐波那契数列
	a, b := 0, 1
	for {
		c := a + b
		if c > 100 {
			break
		}
		// 不要漏掉开始的元素
		fmt.Printf("%d ", c)
		a, b = b, c
	}

}

// 批改意见
// 1. 斐波那契数列少了1个值
