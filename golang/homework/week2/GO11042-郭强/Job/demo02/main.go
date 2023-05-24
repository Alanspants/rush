package main

import (
	"fmt"
)

//斐波那契数列
func Feibo(n int) int {
	if n < 2 {
		return n
	}
	return Feibo(n-2) + Feibo(n-1)
}
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t", Feibo(i))
	}
}

// 批改意见
// 斐波那契数列从1开始，不是从0开始
