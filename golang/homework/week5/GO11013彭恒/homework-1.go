package main

import "fmt"

// 循环
// 参数n是求阶乘的数字
// 返回n阶乘的乘积
func factorial1(n int) int {
	if n <= 0 {
		return 0
	}
	product := 1
	for i := n; i > 0; i-- {
		product *= i
	}
	return product
}

// 递归
// 参数n是求阶乘的数字
// 返回n阶乘的乘积
func factorial2(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return n * factorial2(n-1)
}

// 循环改递归
// 参数n是求阶乘的数字，参数product是累乘乘积，传固定值1
// 返回n阶乘的乘积
func factorial3(n int, product int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return product * 2
	}
	product *= n // 4 * 1,4  3 * 4,12  2 * 12,24  当n==2的时候，没有计算这一步 根据条件直接返回 所以product要乘以2
	return factorial3(n-1, product)
}

func main() {
	fmt.Println(factorial1(5))
	fmt.Println(factorial2(5))
	fmt.Println(factorial3(5, 1))
}
