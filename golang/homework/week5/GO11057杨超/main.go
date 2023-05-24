package main

import "fmt"

// 1、求n的阶乘。至少使用递归函数完成一次。
//for循环阶乘

func work1v1(n int) int {

	if n < 1 {
		panic("阶乘应该大于0")
	}
	p := 1
	for i := 2; i <= n; i++ {
		p = p * i
	}
	return p
}

func work1v2(n int) int {
	sum := 1
	switch {
	case n <= 0:
		panic("阶乘应该大于0")
		// case n == 1:
		// 	return 1
	}
	for i := 2; i < n; i++ {
		sum = sum * i
	}
	return sum
}

func work1v3(n int, p int) int {
	switch {
	case n <= 0:
		fmt.Println(n)
		panic("不能小于等于0")

	case n == 1:
		return p
	}

	return work1v3(n-1, p*n) //workv3(2, 1*3) workv3(1, 1*3)

}

//
// func work2v4(n int) int {
// 	switch {
// 	case n <= 0:
// 		fmt.Println(n)
// 		panic("不能小于等于0")

// 	case n == 1:
// 		return 2
// 	}

// 	return work2(n-1) * n

// }

// // 2、编写一个函数，接受一个参数n，n为正整数。要求数字必须对齐
// func work2(n int) {
// 	var space int

// 	for i := 0; i < n; i++ {
// 		for j := n; j >= 1; j-- {
// 			if i > j {
// 				if j >= 10 {
// 					space = 3

// 				} else {
// 					space = 2
// 				}
// 			}
// 		}
// 		fmt.Println(space)
// 	}
// }

//for循环小于是++ ,大于是--
func work2v2(n int) {
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

//func a() {
//	var width int
//	for i := 1; i < 10; i++ {
//		for j := 1; j < i; j++ {
//			if j == 1 {
//				width = 2
//			} else {
//				width = 3
//			}
//			fmt.Printf("%d*%d=%-[4]*[3]d", j, i, i*j, width)
//		}
//		fmt.Println()
//	}
//}

func main() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~循环版阶乘各版本调用")
	fmt.Println("for循环阶乘", work1v1(6))
	fmt.Println("switch判断阶乘", work1v2(6))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~递归版阶乘版本调用")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~循环改递归版阶乘调用")
	fmt.Println(work1v3(3, 1))
	// fmt.Println(work2v2(12))
	//// work1(5)
	// work2(13)
	work2v2(13)
	// a()
}
