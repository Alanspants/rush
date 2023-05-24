package main

import (
	"fmt"
)

//美丽公式
func fac1(n int) int {
	//fac(n)= n * (n-1)!
	if n == 0 || n == 1 {
		return 1
	}
	return n * fac1(n-1)
}

//循环
func fac2(n int) int {
	//fac2(n) = 1 * 2 *3 *...*n
	//fac2(3)= 1 * 2 * 3
	if n == 0 || n == 1 {
		return 1
	}
	a := 1
	for i := 1; i < n+1; i++ {
		a *= i
	}
	return a
}

//循环改调用
func fac3(n, a int) int {
	if n < 2 {
		return a
	}
	return fac3(n-1, a*n) //4 ,3 1*4=4; 3, 2 4*3=12; 2, 1 12*2=24
}

//倒三角循环实现
func n1(n int) {
	//2 1
	//  1
	a := n
	b := 0
	for i := 1; i <= n; n-- {
		if n < 9 {
			b++
			fmt.Printf("%[1]*s", 3*(a-n)-b, "")
		}
		if n > 8 && n < a {
			fmt.Printf("%[1]*s", 3*(a-n), "")
		}
		for j := n; j >= i; j-- {
			fmt.Printf("%d ", j)
		}

		println()
	}
}

//正三角循环实现
func n2(n int) {
	//  1
	//2 1
	b := 9
	for i := 1; i <= n; i++ {
		if i > 8 {
			fmt.Printf("%[1]*s", 3*(n-i), "")
		}
		if i < 9 {
			b--
			fmt.Printf("%[1]*s", 3*(n-i)-b, "")
		}
		for j := i; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		println()
	}
}

func main() {
	//阶乘三种实现
	println(fac1(4))
	println(fac2(4))
	println(fac3(4, 1))

	//倒三角
	n1(12)
	//正三角
	n2(12)
}

// 批改意见
// 1. 最后的换行函数应该是fmt.Println()
// 2. 代码记得写注释，方便自己后面查看
