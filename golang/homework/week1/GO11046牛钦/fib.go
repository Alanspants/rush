package main

import (
	"fmt"
)

// 1.返回第n个斐波那契数 递归
func fibonacciNumbersRecursion(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciNumbersRecursion(n-1) + fibonacciNumbersRecursion(n-2)
}

// 2.返回第n个斐波那契数 迭代
func fibonacciNumbersIteration(n int) int {
	f0 := 0
	f1 := 1
	if n == 0 {
		return f0
	}
	if n == 1 || n == 2 {
		return f1
	}
	// define f(n-2) as fn_2,f(n-1) as fn_1
	fn_2 := f0
	fn_1 := f1
	i := 2
	for i <= n {
		// f := fn_1
		// fn_1 = fn_2
		// fn_2 = f + fn_2
		fn_1, fn_2 = fn_2, fn_1+fn_2

		i++
	}
	return fn_2
}

func printFibonacciValues(n int, f func(int) int) {

	for i := 0; i <= n; i++ {
		temp := f(i)

		if temp > 100 {
			break
		}
		// TODO print fun name
		fmt.Printf("斐波那契数第%d个数:%d\n", i, temp)

	}
}
func main() {
	var n = 100
	printFibonacciValues(n, fibonacciNumbersRecursion)
	printFibonacciValues(n, fibonacciNumbersIteration)

}
