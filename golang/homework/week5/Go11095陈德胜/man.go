package main

import "fmt"

//求n的阶乘。至少使用递归函数完成一次
func factorial1(n int) int {
	switch {
	case n < 0:
		panic("n is not negative")
	case n == 1:
		return 1
	}
	return factorial1(n-1) * n
}

func factorial2(n int) int {
	switch {
	case n < 0:
		panic("n is not negative")
	case n == 1:
		return 1
	}
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}

//编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
func upperHalfAngle(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if j > i {
				if j >= 10 {
					fmt.Printf("%3v", " ")
				} else {
					fmt.Printf("%2v", " ")
				}
			} else {
				if j >= 10 {
					fmt.Printf("%3d", j)
				} else {
					fmt.Printf("%2d", j)
				}
			}
		}
		fmt.Println()
	}
}
func main() {
	//求n的阶乘。至少使用递归函数完成一次
	for i := 1; i < 11; i++ {
		fmt.Println(factorial1(i))
		fmt.Println(factorial2(i))
	}

	//编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
	upperHalfAngle(12)
}
