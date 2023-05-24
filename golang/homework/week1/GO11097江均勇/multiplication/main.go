package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			tmp := i * j
			// 注意一下i和j的顺序
			fmt.Printf("%d*%d=%d ", i, j, tmp)
			if tmp < 10 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
