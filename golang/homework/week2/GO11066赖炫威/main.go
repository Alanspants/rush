package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 九九乘法表
func Multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%v*%v=%-2d", j, i, j*i)
			} else {
				fmt.Printf("%v*%v=%-3d", j, i, j*i)
			}
		}
		fmt.Println()
	}
}

// 打印20个100以内的随机数，奇数个求和，偶数个求积
func Int20() {
	var random string
	var (
		sum            int
		multiplication int = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		rt := r.Intn(100)
		srt := fmt.Sprintf("%d ", rt)
		if i&1 > 0 {
			sum += rt
		} else {
			if rt != 0 {
				multiplication *= rt
			}
		}
		random += srt
	}
	fmt.Printf("所有随机数为: %s\n", random)
	fmt.Printf("随机数及数个和为: %d\n", sum)
	fmt.Printf("随机数偶数个积为: %d\n", multiplication)
}

// 斐波那契数列
func Fibonacci() {
	a := 1
	b := 1
	for i := 0; i < 19; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}

func main() {
	Multiplication()
	Int20()
	Fibonacci()
}
