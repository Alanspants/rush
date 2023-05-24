package main

import (
	"fmt"
	"strings"
)

// FactorialByN 求n的阶乘。至少使用递归函数完成一次
// 6*5*4*3*2*1
func FactorialByN(n int) int {
	if n < 0 {
		panic("负数不能阶乘！")
	}
	// 0!=1
	if n == 1 || n == 0 {
		return 1
	}
	return n * FactorialByN(n-1)
}

// PositiveTriangle 编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
// 正三角
func PositiveTriangle(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%v", strings.Join(make([]string, n-i+1), "\t"))
		for j := i; j > 0; j-- {
			fmt.Printf("%v\t", j)
		}
		fmt.Println()
	}
}

func PositiveTriangleV2(n int) {
	var last string
	for i := n; i > 0; i-- {
		if i == 1 {
			last += fmt.Sprintf("%d", i)
		} else {
			last += fmt.Sprintf("%d ", i)
		}
	}
	//fmt.Println(last)
	width := len(last)
	for i := width - 1; i >= 0; i-- {
		if last[i] == 32 {
			fmt.Printf("%*s\n", width, last[i+1:])
		}
	}
	fmt.Println(last)
}

func main() {
	//s := FactorialByN(3)
	//fmt.Println(s)
	//PositiveTriangle(10)
	PositiveTriangleV2(12)
}
