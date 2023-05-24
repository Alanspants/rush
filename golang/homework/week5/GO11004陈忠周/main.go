package main

import (
	"fmt"
)

// 1.求n的阶乘。至少使用递归函数完成一次
// 1.1 循环法
func Factorialv1(n int) int {
	if n < 1 {
		panic("n不能小于1")
	}
	p := 1
	for i := 2; i <= n; i++ {
		p = p * i
	}
	return p
}

// 1.2 递归法 优美的公式便于理解
func Factorialv2(n int) int {
	if n < 1 {
		panic("n不能小于1")
	} else if n == 1 {
		return 1
	}
	return Factorialv2(n-1) * n
}

// 递归法 v2 loop
func Factorialv3(n int, p int) int {
	if n < 1 {
		panic("n不能小于1")
	} else if n == 1 {
		return p
	}
	return Factorialv3(n-1, p*n)
}

// 编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
func PrintTriangLev1(n int) {
	widh := 40
	var line string
	for i := 1; i <= n; i++ {
		line = ""
		for j := i; j >= 1; j-- {
			// fmt.Print(j, " ")
			line += fmt.Sprint(j, " ")
		}
		fmt.Printf("%[1]*s\n", widh, line)
		// fmt.Println(line)
	}
}

func PrintTriangLev2(n int) {
	last := ""
	for i := n; i >= 1; i++ {
		last += fmt.Sprintf("%d ", i)
	}
	widh := len(last)
	var line string
	for i := 1; i <= n; i++ {
		line = ""
		for j := i; j >= 1; j-- {
			line += fmt.Sprint(j, " ")
		}
		fmt.Printf("%[1]*s\n", widh, line)
	}
}

func PrintTriangLev3(n int) {
	last := ""
	for i := n; i >= 1; i-- {
		if i == 1 {
			last += fmt.Sprint(i)
		} else {
			last += fmt.Sprintf("%d ", i)
		}
	}
	// fmt.Println(last)
	width := len(last)
	for i := width - 1; i >= 0; i-- {
		if last[i] == 32 {
			fmt.Printf("%*s\n", width, last[i:])
		}
	}
	fmt.Println(last)

}

func main() {
	// fmt.Println(Factorialv1(6))
	// fmt.Println(Factorialv2(6))
	// fmt.Println(Factorialv3(6, 1))
	// fmt.Println(Factorialv3(6, 1))
	// PrintTriangLev1(12)
	// PrintTriangLev2(12)
	PrintTriangLev3(12) //最优解

}
