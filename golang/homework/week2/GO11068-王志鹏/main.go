package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 打印斐波那契
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for x := 1; x <= i; x++ {
			fmt.Printf("%d*%d=%d ", i, x, i*x)
		}
		fmt.Println()
	}

	var i int
	for i = 0; i < 20; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	// 输出100以内的随机数，打印单数和、偶数之积
	sum := 0
	multi := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("20个随机数如下:  \n")
	for i := 0; i < 20; i++ {
		res := r.Intn(100)
		if i%2 == 0 {
			sum += res
		} else {
			multi *= res
		}
		//如果随机数为0,则重新循环一次
		if res == 0 {
			fmt.Printf("\n")
			i--
			continue
		}
		fmt.Printf("%d ", res)
	}
	fmt.Printf("\n")
	fmt.Printf("单数和为 %d\n", sum)
	fmt.Printf("偶数之积为 %d\n", multi)
}

// 批改意见：
// 1. 斐波那契数列是从0开始，而且要求是100以内
