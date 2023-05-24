package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 问题一： 求阶乘
// 1、递归(递推公式)
func factorial1(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial1(n-1)
	}

}

// 2、循环
func factorial2(n int) int {
	res := 1
	if n == 0 {
		return res
	} else {
		for ; n >= 1; n-- {
			res *= n
		}
		return res
	}
}

// 循环改递归
// NOTE 通过传入额外的参数，在每次压栈前进行计算，最后逐层返回。性能接近循环方法，但也有stack over flow的问题
func factorial3(n, res int) int {
	if n == 0 {
		return res
	} else {
		res *= n
		return factorial3(n-1, res)
	}
}

// 问题二：编写一个函数，接受一个正整数参数n，打印上三角的递增数列
func printTriangle(n uint) {
	// 思路：每一行的总长度是一致的，只是当前行数之前的数字，全部用空格代替打印而已
	for i := 1; i <= int(n); i++ {
		for j := int(n); j > 0; j-- {
			if j > i { //不打印大于当前行号的数字
				fmt.Print(strings.Repeat(" ", len(strconv.Itoa(j))+1)) //将数字转化成string，打印其宽度相同数量的空格，并加上数字之间间隔的空格
			} else {
				fmt.Printf("%d ", j)
			}

		}
		fmt.Println()
	}

}

func main() {
	n := 5
	fmt.Printf("factorial1(%d)=%d\n", n, factorial1(n))
	fmt.Printf("factorial2(%d)=%d\n", n, factorial2(n))
	fmt.Printf("factorial3(%d)=%d\n", n, factorial3(n, 1))
	printTriangle(uint(n))

	n = 20
	fmt.Printf("factorial1(%d)=%d\n", n, factorial1(n))
	fmt.Printf("factorial2(%d)=%d\n", n, factorial2(n))
	fmt.Printf("factorial3(%d)=%d\n", n, factorial3(n, 1))
	printTriangle(uint(n))

}
