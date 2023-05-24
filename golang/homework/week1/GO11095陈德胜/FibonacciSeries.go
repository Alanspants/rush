package main

import "fmt"

//斐波那契数列例子 1 1 3 5 8 13
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var fiblist []int
	//打印100以内的斐波那契数列
	for i := 1; i <= 100; i++ {
		fiblist = append(fiblist, fib(i))
		// fmt.Println(fib(i))
	}
	fmt.Printf("斐波那契数列为: %v\n", fiblist)
}
