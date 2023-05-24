package main

import "fmt"

//1、求n的阶乘。至少使用递归函数完成一次。
//2、编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐

func main() {
	work1()
	work1_2(10)
	printN(15)
}

func printN(n int) {
	var tmpStr string
	for i := 1; i <= n; i++ {
		tmpStr = fmt.Sprintf("%d %s", i, tmpStr)
	}
	maxLen := len(tmpStr)
	tmpStr = ""
	for i := 1; i <= n; i++ {
		tmpStr = fmt.Sprintf("%d %s", i, tmpStr)
		fmt.Printf(fmt.Sprintf("%%%ds\n", maxLen), tmpStr)
	}

}

func work1_2(n int) {
	fmt.Printf(" %d的阶乘是%d  \n", n, f1(n))
}

func f1(n int) (total int) {

	if n > 1 {
		total = n * f1(n-1)
	} else {
		total = 1
	}
	return
}

func work1() {
	n := 4
	total := 1
	if n > 0 {
		for i := n; i > 1; i-- {
			total = total * i
		}
	}
	fmt.Printf(" %d的阶乘是%d  \n", n, total)

}

// 批改意见
// 1. 代码中记得写注释，方便自己以后复习回头看
