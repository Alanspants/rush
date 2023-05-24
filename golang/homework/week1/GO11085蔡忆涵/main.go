package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	run()
}
func homework_1() {
	i := 1
	for i < 10 {
		j := 1
		for j <= i {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
			j++
		}
		i++
		fmt.Println()
	}
}
func homework_2() {

	i := 0
	t := rand.New(rand.NewSource(time.Now().UnixMilli()))
	even := 0
	odd := 0
	for i < 20 {
		num := t.Intn(100) + 1
		if num&1 == 0 {
			even += num
		} else {
			odd += num
		}
		i++
	}
	fmt.Println("偶数和:", even, "奇数和:", odd)
}
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n >= 2 {
		return fib(n-1) + fib(n-2)
	} else {
		return 0
	}
}
func run() {
	homework_1()
	homework_2()
	fmt.Println(fib(100))
}

// 批改意见
// 1. 对奇偶数的判断有问题
// 2. 要求返回100以内的斐波那契数列，不是第100个，第100个计算量太大，会把电脑卡死。
