package main

import (
	"fmt"
)

// loop
func factorialv1(n int) int {
	if n < 1 {
		panic("n 不能小于 1")
	}

	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}

	return p
}

// recursion v1
func factorialv2(n int) int {
	if n < 1 {
		panic("n 不能小于 1")
	} else if n == 1 {
		return 1
	}

	return factorialv2(n-1) * n
}

// recursion v2
func factorialv3(n, p int) int {
	if n < 1 {
		panic("n 不能小于 1")
	} else if n == 1 {
		return p
	}
	return factorialv3(n-1, n*p)
}

func main() {
	// 6 的阶乘为 720
	fmt.Println(factorialv1(6))
	fmt.Println(factorialv2(6))
	fmt.Println(factorialv3(6, 1))
}
