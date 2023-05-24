package main

import (
	"fmt"
	"math/rand"
)

func homework1() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v x %v = %v\t", i, j, i*j)
			if j == i {
				fmt.Printf("\n")
			}
		}
	}
}

func homework2() {
	a := rand.New(rand.NewSource(100))
	var b int64 = 0
	var c int64 = 1
	for i := 0; i < 20; i++ {
		d := a.Int63n(100)
		if d == int64(0) {
			d = d + 1
		}
		fmt.Printf("%+v\n", d)
		if i%2 == 0 {
			c = c * d
		} else {
			b = b + d
		}
	}
	fmt.Printf("奇数位累加求和为：%v\n偶数位累乘求积为：%v\n", b, c)
}

func homework3() {
	n := 1
	m := 1

	fmt.Printf("%v %v", n, m)
	for s := 2; s < 100; s = n + m {
		n = m
		m = s
		fmt.Printf(" %v", s)
	}
}

func main() {
	fmt.Println("开始打印99乘法表")
	homework1()
	fmt.Println("开始打印随机数")
	homework2()
	fmt.Println("开始打印斐波纳契数列")
	homework3()
}
