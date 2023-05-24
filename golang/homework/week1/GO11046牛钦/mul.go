package main

import "fmt"

// 生成 9*9乘法表
func multiplicationTable(n, m int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d ", j, i, i*j)
		}
		fmt.Println()
	}

}
func main() {
	var (
		n = 9
		m = 9
	)
	multiplicationTable(n, m)

}
