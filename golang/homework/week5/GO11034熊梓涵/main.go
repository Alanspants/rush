package main

import "fmt"

// 循环实现
func n1(n int) int {
	switch {
	case n < 0:
		panic("<0")
	case n == 0:
		return 0
	case n == 1:
		return 1
	}
	result := 1
	for {
		result *= n
		n--
		if n == 1 {
			break
		}
	}
	return result
}

// 循环改调用
func n2(n, result int) int { // result初始给1
	if n == 1 {
		return result // n=1时，返回此时的乘积result
	}
	return n2(n-1, result*n) // 每调用一次，就n-1，result就乘一次n
}

// 递归实现
func n3(n int) int {
	switch {
	case n < 0:
		panic("<0")
	case n == 0:
		return 0
	case n == 1:
		return 1
	}
	return n * n3(n-1)
}

func Triangle(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if i < j { // <上， >下
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
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i := 1; i < 11; i++ {
		fmt.Print(n1(i), " ")
	}
	fmt.Println()
	for i := 1; i < 11; i++ {
		fmt.Print(n2(i, 1), " ")
	}
	fmt.Println()
	for i := 1; i < 11; i++ {
		fmt.Print(n3(i), " ")
	}
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	Triangle(15)
}

// 批改意见
// 实现逻辑都不错，结果也都是对的
// 代码中记得写注释，防止后面自己回头看的时候看不懂
