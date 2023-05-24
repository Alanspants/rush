package main

import (
	"fmt"
)

func main() {
	number := 0
	// 尽量不要在循环过程中动态修改变量，可能会导致意料之外的问题
	for i := 1; i < 100; i = i + number {
		if i == 1 {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("%d ", i)
		number = i - number
	}
}
