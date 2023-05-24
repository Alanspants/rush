package main

import "fmt"

func main() {
	var result string

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			var k int
			k = j * i
			if i > 2 && i < 5 && j > 2 && j < 4 {
				result = fmt.Sprintf(" %v*%v=%v ", j, i, k)

			} else {
				result = fmt.Sprintf("%v*%v=%v ", j, i, k)
			}

			fmt.Print(result)

		}
		fmt.Printf("\n")
	}

}

// 批改意见
// 1. 后面数字的对齐除了你的方法，还可以尝试一下制表符\t
