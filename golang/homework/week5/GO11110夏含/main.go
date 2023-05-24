package main

import (
	"fmt"
	"strconv"
	"strings"
)

func nFactorial(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * nFactorial(n-1)
}

func triangle(n int, m int) {
	if n <= 0 {
		panic("n应该为正整数")
	}
	if n == 1 {
		fmt.Printf("上三角\n")
		return
	}
	triangle(n-1, m)
	for j := m - 1; j > n-1; j-- {
		fmt.Printf(strings.Repeat(" ", len(strconv.Itoa(j))))
		fmt.Printf(" ")
	}
	for i := n - 1; i > 0; i-- {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")
}

func main() {
	//jc := nFactorial(20) // 未考虑超界的情况
	//fmt.Println(jc)
	triangle(30, 30)
}
