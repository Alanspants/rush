package main

import "fmt"

func main() {
	i := 1
	for {
		if fib(i) < 100 {
			fmt.Println(fib(i))
		} else {
			break
		}
		i++
	}
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return (fib(n-2) + fib(n-1))
}
