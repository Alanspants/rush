package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			// 注意输出的数字顺序
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Print("\n")
	}

}
