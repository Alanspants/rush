package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := a + b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for c < 100 {
		a = b
		b = c
		c = a + b
		if c < 100 {
			fmt.Println(c)
		}
	}
}

// 批改意见
// 记得写注释，方便自己以后看代码
