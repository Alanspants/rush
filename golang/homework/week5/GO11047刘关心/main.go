package main

import (
	"fmt"
)

// "~~~~~~~~~~~~第一题~~~~~~~~~~~~~~~"
// 求n的阶乘。至少使用递归函数完成一次。
func factorial1(n int) int {
	a := 1
	for i := 1; i <= n; i++ {
		// fmt.Println(i, a)
		a = a * i
		// fmt.Println(i, a)

	}
	return a
}

func factorial2(n, a, b int) int {
	if n < 1 {
		return a

	}
	// fmt.Println(n, a, b)
	return factorial2(n-1, b*a, (b + 1))

}

// "~~~~~~~~~~~~第二题~~~~~~~~~~~~~~~"
// 编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
/*
上三角
	1
  2 1
3 2 1
*/

func UpperTriangle(n int) {
	for i := 1; i < n; i++ {
		// fmt.Println(i)
		for j := n; j > 0; j-- {
			if j > i {
				fmt.Printf("%v", "   ")
			} else {
				fmt.Printf("%-3d", j)
			}

		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("~~~~~~~~~~~~第一题~~~~~~~~~~~~~~~")
	value := factorial1(10)
	fmt.Println(value)

	c := factorial2(10, 1, 1)
	fmt.Println(c)

	fmt.Println("~~~~~~~~~~~~第二题~~~~~~~~~~~~~~~")
	UpperTriangle(15)

}
