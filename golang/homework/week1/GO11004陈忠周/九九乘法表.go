package main

import "fmt"

// 打印九九乘法表，如果可以要求间隔均匀
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-4d", j, i, i*j)
		}
		fmt.Println()
	}
}

// 输出时，把i和j的顺序调整一下结果会更美观。
