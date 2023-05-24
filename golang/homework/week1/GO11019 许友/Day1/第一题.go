package main

import "fmt"

func Nine() {
	for i := 1; i < 10; i++ {

		for j := 1; j < i+1; j++ {
			// 输出时，将i和j的顺序调整一下，结果会更美观。
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()

	}
}

func main() {
	Nine()
}
