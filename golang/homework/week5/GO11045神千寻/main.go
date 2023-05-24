package main

import (
	"fmt"
)

// 1、求n的阶乘。至少使用递归函数完成一次。
// 循环版
func FactorialCirculation(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	// var product uint64 = 1
	for i := n - 1; i > 1; i-- {
		n *= i
	}
	return n
}

// 循环改调用
var p uint64 = 1

func FactorialRecursion1(n uint64) uint64 {
	p *= n
	if n == 0 {
		return 1
	} else if n == 1 { //边界条件
		return p
	}
	return FactorialRecursion1(n - 1)
}

// 公式递归
func FactorialRecursion2(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	return n * FactorialRecursion2(n-1) //阶乘公式
}

// 2、编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
var width int

// Deprecated
func UpperTriangularMatrix(n int) {
	var line string
	width := 30
	for i := 1; i <= n; i++ {
		line = ""
		for j := i; j >= 1; j-- { //时间复杂度：O(n²)
			line += fmt.Sprint(j) + " " //规模随着n增大而变大，大量拼接字符串，重复计算效率极其低下
		}
		fmt.Printf("%*s \n", width, line)
	}
}

// NOTE: 上三角最优版
func UpperTriangularMatrix2(n int) {
	ceilLine := ""
	for i := n; i > 0; i-- {
		if i == 1 {
			ceilLine += fmt.Sprint(i)
		} else {
			ceilLine += fmt.Sprintf("%d ", i)
		}
	}
	fmt.Println(ceilLine)
	length := len(ceilLine)
	for i := length - 1; i > 0; i-- { //i必须从length-1开始
		if ceilLine[i] == '\x20' { //若 i=length，ceilLine[i]索引超界
			fmt.Printf("%-*s\n", length, ceilLine[:i])
		}
	}
}

// Deprecated
func LowerTriangularMatrix(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- { //时间复杂度：O(n²),效率低下
			if j > 9 {
				width = 3
			} else {
				width = 2
			}
			if j > i { //对角线上半部分区域
				fmt.Printf("%-*s", width, "")
			} else { //下半角(等于i包含对角线数字)
				fmt.Printf("%-*d", width, j) //j小于i时，打印数字j，j==i时，是对角线上的数字算在下半部分，若算在上半角，则打印结果最左侧和最顶层多一列一行空格
			}
		}
		fmt.Println()
	}
}

// NOTE: 作业最优版（下三角）
func LowerTriangularMatrix2(n int) {
	floorLine := ""
	for i := n; i >= 1; i-- {
		if i == 1 {
			floorLine += fmt.Sprint(i)
		} else {
			floorLine += fmt.Sprintf("%d ", i)
		}
	}
	width := len(floorLine)
	for i := width - 1; i > 0; i-- {
		if floorLine[i] == 0x20 {
			fmt.Printf("%*s\n", width, floorLine[i+1:])
		}
	}
	fmt.Println(floorLine)
}
func main() {
	//第一个作业开始：
	fmt.Println(FactorialCirculation(10)) //3628800 39916800 479001600 6227020800 87178291200
	fmt.Println("------------------------阶乘循环改递归调用--------------------")
	fmt.Println(FactorialRecursion1(13)) //17的阶乘:355687428096000
	/* str := "abcdefgh"
	length := len(str)
	for i := length; i > 0; i-- {
		fmt.Printf("类型：%[2]T，No.%[1]d:%[2]s\n", i, str[:i])
	} */
	fmt.Println("=========================阶乘公式递归 开始============================")
	fmt.Println(FactorialRecursion2(14))
	fmt.Println("+++++++++++++++++++++++++++++++The second HomeWork begin++++++++++++++++++++++++++++++")
	UpperTriangularMatrix(12)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	LowerTriangularMatrix2(12)
	/* π := 3.1415926
	fmt.Printf("%T 【%30[1]f】", π) */
	test()
}

var count = 1

func test() {
	fmt.Println("enter test function")
	count = 0                                   //没有:=短格式声明，就是全局变量count向内穿透；反之，就是用的test函数中的局部变量
	defer func() { fmt.Println("1：", count) }() //匿名函数没有入参，内部的打印变量采用就近原则
	count++                                     //1
	defer fmt.Println("2：", count)
	defer foo(count) //有入参count符合就近原则，看是全局变量还是局部变量
	count++          //2
	defer fmt.Println("4：", count)
	defer bar() //没有传参,始终用的全局变量count
	count++     //3
	fmt.Println("exit test")
}
func foo(count int) {
	fmt.Println("3：", count)
}
func bar() {
	fmt.Println("5：", count)
}
