package main

import (
	"fmt"
)

// 1、求n的阶乘。至少使用递归函数完成一次。
// 第一种方式：递归方式
func factorialv1(n int) int {
	if n < 0 {
		panic("传入的参数为0")
	} else if n == 1 {
		return 1
	}
	return n * factorialv1(n-1)
}

// 1、求n的阶乘。至少使用递归函数完成一次。
// 第二种方式：循环方式
func factorialv2(n int) int {
	if n < 0 {
		panic("传入的参数小于0")
	} else if n <= 1 {
		return 1
	}
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

// 1、求n的阶乘。至少使用递归函数完成一次。
// 第三种方式：循环改调用
func factorialv3(n, p int) int {
	if n < 0 {
		panic("传入的参数小于0")
	} else if n == 1 {
		return p
	}
	return factorialv3(n-1, p*n)
}

// 2、编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
func upper_triangular(n int) {
	var k int = 1
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			if n-j >= 9 {
				fmt.Printf("%-3v", " ")
			} else {
				fmt.Printf("%-2v", " ")
			}
		}
		for k = i; k > 0; k-- {
			if k >= 10 {
				fmt.Printf("%-3d", k)
			} else if k < 10 {
				fmt.Printf("%-2d", k)
			}
		}
		fmt.Println()
	}
}

func main() {
	//第一题
	//var num int
	//fmt.Println("请输入整数:")
	//fmt.Scan(&num)
	//ret := factorial(num)
	//fmt.Println(ret)
	//fmt.Println(factorialv1(4))

	//fmt.Println(factorialv2(6))
	//fmt.Println(factorialv3(5, 1))

	//第二题
	upper_triangular(12)

}
