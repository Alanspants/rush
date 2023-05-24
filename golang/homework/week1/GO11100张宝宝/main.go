package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 九九乘法表
func nineMultiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, i*j)
			}
		}
		fmt.Println()
	}
}

// 生成随机数，第单数个累加求和，第偶数个累乘求积
func calculationNumber() {
	var (
		a        = 0
		b uint64 = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		j := r.Intn(99) + 1
		fmt.Println(j) // 打印二十个随机数字并不含0
		if j&1 == 0 {
			a += j
		} else {
			b *= uint64(j)
		}
	}
	fmt.Println(a, b)
}

// 打印100以内的斐波那契数列
func fibonacciSequence() {
	var (
		a = 0
		b = 1
	)
	for {
		fmt.Printf("%d,", a)
		a, b = b, a+b
		if a > 100 {
			break
		}
	}
}

func main() {
	nineMultiplication()
	calculationNumber()
	fibonacciSequence()
}

// 批改意见
// 1. 第二题要求的是奇数个和偶数个，不是奇数和偶数
